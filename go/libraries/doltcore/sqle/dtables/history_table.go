// Copyright 2019 Dolthub, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dtables

import (
	"context"
	"io"
	"strings"

	"github.com/dolthub/dolt/go/libraries/doltcore/sqle"
	"github.com/dolthub/go-mysql-server/sql"
	"github.com/dolthub/go-mysql-server/sql/expression"
	"github.com/dolthub/go-mysql-server/sql/transform"

	"github.com/dolthub/dolt/go/libraries/doltcore/doltdb"
	"github.com/dolthub/dolt/go/libraries/doltcore/rowconv"
	"github.com/dolthub/dolt/go/libraries/doltcore/schema"
	"github.com/dolthub/dolt/go/libraries/doltcore/sqle/sqlutil"
	"github.com/dolthub/dolt/go/libraries/utils/set"
	"github.com/dolthub/dolt/go/store/hash"
	"github.com/dolthub/dolt/go/store/types"
)

const (
	// DoltHistoryTablePrefix is the name prefix for each history table

	// CommitHashCol is the name of the column containing the commit hash in the result set
	CommitHashCol = "commit_hash"

	// CommitterCol is the name of the column containing the committer in the result set
	CommitterCol = "committer"

	// CommitDateCol is the name of the column containing the commit date in the result set
	CommitDateCol = "commit_date"
)

var _ sql.Table = (*HistoryTable)(nil)
var _ sql.FilteredTable = (*HistoryTable)(nil)
var _ sql.IndexAddressableTable = (*HistoryTable)(nil)
var _ sql.IndexedTable = (*HistoryTable)(nil)

// HistoryTable is a system table that shows the history of rows over time
type HistoryTable struct {
	doltTable             *sqle.DoltTable
	commitFilters         []sql.Expression
	cmItr                 doltdb.CommitItr
	indexLookup           sql.IndexLookup
	readerCreateFuncCache *ThreadSafeCRFuncCache
	sqlSch                sql.PrimaryKeySchema
	targetSch             schema.Schema
}

func (ht *HistoryTable) GetIndexes(ctx *sql.Context) ([]sql.Index, error) {
	return ht.doltTable.GetIndexes(ctx)
}

func (ht HistoryTable) WithIndexLookup(lookup sql.IndexLookup) sql.Table {
	ht.indexLookup = lookup
	return &ht
}

// NewHistoryTable creates a history table
func NewHistoryTable(ctx *sql.Context, table *sqle.DoltTable) (sql.Table, error) {
	currentSch := table.Schema()

	sch := schema.MustSchemaFromCols(currentSch.GetAllCols().Append(
		schema.NewColumn(CommitHashCol, schema.HistoryCommitHashTag, types.StringKind, false),
		schema.NewColumn(CommitterCol, schema.HistoryCommitterTag, types.StringKind, false),
		schema.NewColumn(CommitDateCol, schema.HistoryCommitDateTag, types.TimestampKind, false),
	))

	if sch.GetAllCols().Size() <= 3 {
		return nil, sql.ErrTableNotFound.New(doltdb.DoltHistoryTablePrefix + table.Name())
	}

	sqlSch, err := sqlutil.FromDoltSchema(doltdb.DoltHistoryTablePrefix+table.Name(), sch)
	if err != nil {
		return nil, err
	}

	return &HistoryTable{
		doltTable:             table,
		sqlSch:                sqlSch,
		targetSch:             sch,
	}, nil
}

// HandledFilters returns the list of filters that will be handled by the table itself
func (ht *HistoryTable) HandledFilters(filters []sql.Expression) []sql.Expression {
	ht.commitFilters = filterFilters(filters, getColumnFilterCheck(commitFilterCols))
	return ht.commitFilters
}

// Filters returns the list of filters that are applied to this table.
func (ht *HistoryTable) Filters() []sql.Expression {
	return ht.commitFilters
}

// WithFilters returns a new sql.Table instance with the filters applied
func (ht HistoryTable) WithFilters(ctx *sql.Context, filters []sql.Expression) sql.Table {
	if ht.commitFilters == nil {
		ht.commitFilters = filterFilters(filters, getColumnFilterCheck(commitFilterCols))
	}

	if len(ht.commitFilters) > 0 {
		commitCheck, err := commitFilterForExprs(ctx, ht.commitFilters)
		if err != nil {
			return sqlutil.NewStaticErrorTable(&ht, err)
		}

		ht.cmItr = doltdb.NewFilteringCommitItr(ht.cmItr, commitCheck)
	}

	return &ht
}

var commitFilterCols = set.NewStrSet([]string{CommitHashCol, CommitDateCol, CommitterCol})

func getColumnFilterCheck(colNameSet *set.StrSet) func(sql.Expression) bool {
	return func(filter sql.Expression) bool {
		isCommitFilter := true
		sql.Inspect(filter, func(e sql.Expression) (cont bool) {
			if e == nil {
				return true
			}

			switch val := e.(type) {
			case *expression.GetField:
				if !colNameSet.Contains(strings.ToLower(val.Name())) {
					isCommitFilter = false
					return false
				}
			}

			return true
		})

		return isCommitFilter
	}
}

func filterFilters(filters []sql.Expression, predicate func(filter sql.Expression) bool) []sql.Expression {
	matching := make([]sql.Expression, 0, len(filters))
	for _, f := range filters {
		if predicate(f) {
			matching = append(matching, f)
		}
	}
	return matching
}

func commitFilterForExprs(ctx *sql.Context, filters []sql.Expression) (doltdb.CommitFilter, error) {
	filters = transformFilters(ctx, filters...)

	return func(ctx context.Context, h hash.Hash, cm *doltdb.Commit) (filterOut bool, err error) {
		meta, err := cm.GetCommitMeta(ctx)

		if err != nil {
			return false, err
		}

		sc := sql.NewContext(ctx)
		r := sql.Row{h.String(), meta.Name, meta.Time()}

		for _, filter := range filters {
			res, err := filter.Eval(sc, r)
			if err != nil {
				return false, err
			}
			b, ok := res.(bool)
			if ok && !b {
				return true, nil
			}
		}

		return false, err
	}, nil
}

func transformFilters(ctx *sql.Context, filters ...sql.Expression) []sql.Expression {
	for i := range filters {
		filters[i], _, _ = transform.Expr(filters[i], func(e sql.Expression) (sql.Expression, transform.TreeIdentity, error) {
			gf, ok := e.(*expression.GetField)
			if !ok {
				return e, transform.SameTree, nil
			}
			switch gf.Name() {
			case CommitHashCol:
				return gf.WithIndex(0), transform.NewTree, nil
			case CommitterCol:
				return gf.WithIndex(1), transform.NewTree, nil
			case CommitDateCol:
				return gf.WithIndex(2), transform.NewTree, nil
			default:
				return gf, transform.SameTree, nil
			}
		})
	}
	return filters
}

func (ht *HistoryTable) WithProjection(colNames []string) sql.Table {
	return ht
}

func (ht *HistoryTable) Projection() []string {
	return []string{}
}

// Name returns the name of the history table
func (ht *HistoryTable) Name() string {
	return doltdb.DoltHistoryTablePrefix + ht.doltTable.Name()
}

// String returns the name of the history table
func (ht *HistoryTable) String() string {
	return doltdb.DoltHistoryTablePrefix + ht.doltTable.Name()
}

// Schema returns the schema for the history table
func (ht *HistoryTable) Schema() sql.Schema {
	return ht.sqlSch.Schema
}

// Partitions returns a PartitionIter which will be used in getting partitions each of which is used to create RowIter.
func (ht *HistoryTable) Partitions(ctx *sql.Context) (sql.PartitionIter, error) {
	return &commitPartitioner{ht.cmItr}, nil
}

// PartitionRows takes a partition and returns a row iterator for that partition
func (ht *HistoryTable) PartitionRows(ctx *sql.Context, part sql.Partition) (sql.RowIter, error) {
	cp := part.(*commitPartition)

	return newRowItrForTableAtCommit(ctx, cp.h, cp.cm, ht.targetSch, ht.doltTable)
}

// commitPartition is a single commit
type commitPartition struct {
	h  hash.Hash
	cm *doltdb.Commit
}

// Key returns the hash of the commit for this partition which is used as the partition key
func (cp *commitPartition) Key() []byte {
	return cp.h[:]
}

// commitPartitioner creates partitions from a CommitItr
type commitPartitioner struct {
	cmItr doltdb.CommitItr
}

// Next returns the next partition and nil, io.EOF when complete
func (cp commitPartitioner) Next(ctx *sql.Context) (sql.Partition, error) {
	h, cm, err := cp.cmItr.Next(ctx)

	if err != nil {
		return nil, err
	}

	return &commitPartition{h, cm}, nil
}

// Close closes the partitioner
func (cp commitPartitioner) Close(*sql.Context) error {
	return nil
}

type rowItrForTableAtCommit struct {
	table           *sqle.DoltTable
	tablePartitions sql.PartitionIter
	currPart        sql.RowIter
	sch             schema.Schema
	rowConverter    *rowconv.RowConverter
	extraVals       map[uint64]types.Value
	empty           bool
}

func newRowItrForTableAtCommit(
		ctx *sql.Context,
		h hash.Hash,
		cm *doltdb.Commit,
		sch schema.Schema,
		table *sqle.DoltTable,
) (*rowItrForTableAtCommit, error) {

	root, err := cm.GetRootValue(ctx)
	if err != nil {
		return nil, err
	}

	meta, err := cm.GetCommitMeta(ctx)
	if err != nil {
		return nil, err
	}

	tbl, _, ok, err := root.GetTableInsensitive(ctx, table.Name())
	if err != nil {
		return nil, err
	}
	if !ok {
		return &rowItrForTableAtCommit{empty: true}, nil
	}

	table = table.LockedToRoot(root)

	// TODO: apply index lookups conditionally based on index presence at this revision

	tblSch, err := tbl.GetSchema(ctx)
	if err != nil {
		return nil, err
	}

	rowConverter, err := rowConvForSchema(ctx, tbl.ValueReadWriter(), sch, tblSch)
	if err != nil {
		return nil, err
	}

	hashCol, hashOK := sch.GetAllCols().GetByName(CommitHashCol)
	dateCol, dateOK := sch.GetAllCols().GetByName(CommitDateCol)
	committerCol, committerOK := sch.GetAllCols().GetByName(CommitterCol)
	if !hashOK || !dateOK || !committerOK {
		panic("Bug: missing meta columns in history table schema")
	}

	tablePartitions, err := table.Partitions(ctx)
	if err != nil {
		return nil, err
	}

	return &rowItrForTableAtCommit{
		sch:          sch,
		tablePartitions: tablePartitions,
		rowConverter: rowConverter,
		extraVals: map[uint64]types.Value{
			hashCol.Tag:      types.String(h.String()),
			dateCol.Tag:      types.Timestamp(meta.Time()),
			committerCol.Tag: types.String(meta.Name),
		},
		empty: false,
	}, nil
}

// Next retrieves the next row. It will return io.EOF if it's the last row. After retrieving the last row, Close
// will be automatically closed.
func (i *rowItrForTableAtCommit) Next(ctx *sql.Context) (sql.Row, error) {
	if i.empty {
		return nil, io.EOF
	}

	if i.currPart == nil {
		nextPart, err := i.tablePartitions.Next(ctx)
		if err != nil {
			return nil, err
		}

		rowIter, err := i.table.PartitionRows(ctx, nextPart)
		if err != nil {
			return nil, err
		}

		i.currPart = rowIter
		return i.Next(ctx)
	}

	r, err := i.currPart.Next(ctx)
	if err == io.EOF {
		i.currPart = nil
		return i.Next(ctx)
	} else if err != nil {
		return nil, err
	}

	// TODO: add in extra columns
	return r, nil
}

// Close the iterator.
func (i *rowItrForTableAtCommit) Close(ctx *sql.Context) error {
	return nil
}

// rowConvForSchema creates a RowConverter for transforming rows with the given schema to the target schema.
func rowConvForSchema(ctx context.Context, vrw types.ValueReadWriter, targetSch schema.Schema, sch schema.Schema) (*rowconv.RowConverter, error) {
	if schema.SchemasAreEqual(sch, schema.EmptySchema) {
		return rowconv.IdentityConverter, nil
	}

	fm, err := rowconv.TagMappingWithNameFallback(sch, targetSch)
	if err != nil {
		return nil, err
	}

	return rowconv.NewRowConverter(ctx, vrw, fm)
}
