// Copyright 2022 Dolthub, Inc.
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

// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package serial

import (
	"strconv"

	flatbuffers "github.com/google/flatbuffers/go"
)

type TupleFormat uint16

const (
	TupleFormatUnknown TupleFormat = 0
	TupleFormatV1      TupleFormat = 1
)

var EnumNamesTupleFormat = map[TupleFormat]string{
	TupleFormatUnknown: "Unknown",
	TupleFormatV1:      "V1",
}

var EnumValuesTupleFormat = map[string]TupleFormat{
	"Unknown": TupleFormatUnknown,
	"V1":      TupleFormatV1,
}

func (v TupleFormat) String() string {
	if s, ok := EnumNamesTupleFormat[v]; ok {
		return s
	}
	return "TupleFormat(" + strconv.FormatInt(int64(v), 10) + ")"
}

type Map struct {
	_tab flatbuffers.Table
}

func GetRootAsMap(buf []byte, offset flatbuffers.UOffsetT) *Map {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Map{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsMap(buf []byte, offset flatbuffers.UOffsetT) *Map {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Map{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Map) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Map) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Map) KeyTuples(j int) int8 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetInt8(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *Map) KeyTuplesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Map) MutateKeyTuples(j int, n int8) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateInt8(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func (rcv *Map) KeyOffsets(j int) uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetUint16(a + flatbuffers.UOffsetT(j*2))
	}
	return 0
}

func (rcv *Map) KeyOffsetsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Map) MutateKeyOffsets(j int, n uint16) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateUint16(a+flatbuffers.UOffsetT(j*2), n)
	}
	return false
}

func (rcv *Map) KeyFormat() TupleFormat {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return TupleFormat(rcv._tab.GetUint16(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *Map) MutateKeyFormat(n TupleFormat) bool {
	return rcv._tab.MutateUint16Slot(8, uint16(n))
}

func (rcv *Map) ValueTuples(j int) int8 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetInt8(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *Map) ValueTuplesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Map) MutateValueTuples(j int, n int8) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateInt8(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func (rcv *Map) ValueOffsets(j int) uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetUint16(a + flatbuffers.UOffsetT(j*2))
	}
	return 0
}

func (rcv *Map) ValueOffsetsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Map) MutateValueOffsets(j int, n uint16) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateUint16(a+flatbuffers.UOffsetT(j*2), n)
	}
	return false
}

func (rcv *Map) ValueFormat() TupleFormat {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return TupleFormat(rcv._tab.GetUint16(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *Map) MutateValueFormat(n TupleFormat) bool {
	return rcv._tab.MutateUint16Slot(14, uint16(n))
}

func (rcv *Map) RefArray(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *Map) RefArrayLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Map) RefArrayBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Map) MutateRefArray(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func (rcv *Map) TreeCount() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Map) MutateTreeCount(n uint64) bool {
	return rcv._tab.MutateUint64Slot(18, n)
}

func (rcv *Map) NodeCount() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Map) MutateNodeCount(n uint16) bool {
	return rcv._tab.MutateUint16Slot(20, n)
}

func (rcv *Map) TreeLevel() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Map) MutateTreeLevel(n byte) bool {
	return rcv._tab.MutateByteSlot(22, n)
}

func MapStart(builder *flatbuffers.Builder) {
	builder.StartObject(10)
}
func MapAddKeyTuples(builder *flatbuffers.Builder, keyTuples flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(keyTuples), 0)
}
func MapStartKeyTuplesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func MapAddKeyOffsets(builder *flatbuffers.Builder, keyOffsets flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(keyOffsets), 0)
}
func MapStartKeyOffsetsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(2, numElems, 2)
}
func MapAddKeyFormat(builder *flatbuffers.Builder, keyFormat TupleFormat) {
	builder.PrependUint16Slot(2, uint16(keyFormat), 0)
}
func MapAddValueTuples(builder *flatbuffers.Builder, valueTuples flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(valueTuples), 0)
}
func MapStartValueTuplesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func MapAddValueOffsets(builder *flatbuffers.Builder, valueOffsets flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(valueOffsets), 0)
}
func MapStartValueOffsetsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(2, numElems, 2)
}
func MapAddValueFormat(builder *flatbuffers.Builder, valueFormat TupleFormat) {
	builder.PrependUint16Slot(5, uint16(valueFormat), 0)
}
func MapAddRefArray(builder *flatbuffers.Builder, refArray flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(refArray), 0)
}
func MapStartRefArrayVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func MapAddTreeCount(builder *flatbuffers.Builder, treeCount uint64) {
	builder.PrependUint64Slot(7, treeCount, 0)
}
func MapAddNodeCount(builder *flatbuffers.Builder, nodeCount uint16) {
	builder.PrependUint16Slot(8, nodeCount, 0)
}
func MapAddTreeLevel(builder *flatbuffers.Builder, treeLevel byte) {
	builder.PrependByteSlot(9, treeLevel, 0)
}
func MapEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}

type RefMap struct {
	_tab flatbuffers.Table
}

func GetRootAsRefMap(buf []byte, offset flatbuffers.UOffsetT) *RefMap {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &RefMap{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsRefMap(buf []byte, offset flatbuffers.UOffsetT) *RefMap {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &RefMap{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *RefMap) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *RefMap) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *RefMap) Names(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *RefMap) NamesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *RefMap) RefArray(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *RefMap) RefArrayLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *RefMap) RefArrayBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *RefMap) MutateRefArray(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func (rcv *RefMap) NodeCount() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *RefMap) MutateNodeCount(n uint16) bool {
	return rcv._tab.MutateUint16Slot(8, n)
}

func (rcv *RefMap) TreeCount() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *RefMap) MutateTreeCount(n uint64) bool {
	return rcv._tab.MutateUint64Slot(10, n)
}

func (rcv *RefMap) TreeLevel() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *RefMap) MutateTreeLevel(n byte) bool {
	return rcv._tab.MutateByteSlot(12, n)
}

func RefMapStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func RefMapAddNames(builder *flatbuffers.Builder, names flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(names), 0)
}
func RefMapStartNamesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func RefMapAddRefArray(builder *flatbuffers.Builder, refArray flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(refArray), 0)
}
func RefMapStartRefArrayVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func RefMapAddNodeCount(builder *flatbuffers.Builder, nodeCount uint16) {
	builder.PrependUint16Slot(2, nodeCount, 0)
}
func RefMapAddTreeCount(builder *flatbuffers.Builder, treeCount uint64) {
	builder.PrependUint64Slot(3, treeCount, 0)
}
func RefMapAddTreeLevel(builder *flatbuffers.Builder, treeLevel byte) {
	builder.PrependByteSlot(4, treeLevel, 0)
}
func RefMapEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
