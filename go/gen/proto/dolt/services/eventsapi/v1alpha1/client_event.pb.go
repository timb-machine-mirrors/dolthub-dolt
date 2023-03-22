// Copyright 2020 Dolthub, Inc.
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

// WARNING: This file was is automatically generated. DO NOT EDIT BY HAND.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.0
// source: dolt/services/eventsapi/v1alpha1/client_event.proto

package eventsapi

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ClientEventAttribute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    AttributeID `protobuf:"varint,1,opt,name=id,proto3,enum=dolt.services.eventsapi.v1alpha1.AttributeID" json:"id,omitempty"`
	Value string      `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ClientEventAttribute) Reset() {
	*x = ClientEventAttribute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dolt_services_eventsapi_v1alpha1_client_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientEventAttribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientEventAttribute) ProtoMessage() {}

func (x *ClientEventAttribute) ProtoReflect() protoreflect.Message {
	mi := &file_dolt_services_eventsapi_v1alpha1_client_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientEventAttribute.ProtoReflect.Descriptor instead.
func (*ClientEventAttribute) Descriptor() ([]byte, []int) {
	return file_dolt_services_eventsapi_v1alpha1_client_event_proto_rawDescGZIP(), []int{0}
}

func (x *ClientEventAttribute) GetId() AttributeID {
	if x != nil {
		return x.Id
	}
	return AttributeID_ATTRIBUTE_UNSPECIFIED
}

func (x *ClientEventAttribute) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type ClientEventMetric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to MetricOneof:
	//	*ClientEventMetric_Duration
	//	*ClientEventMetric_Count
	MetricOneof isClientEventMetric_MetricOneof `protobuf_oneof:"metric_oneof"`
	MetricId    MetricID                        `protobuf:"varint,100,opt,name=metric_id,json=metricId,proto3,enum=dolt.services.eventsapi.v1alpha1.MetricID" json:"metric_id,omitempty"`
}

func (x *ClientEventMetric) Reset() {
	*x = ClientEventMetric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dolt_services_eventsapi_v1alpha1_client_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientEventMetric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientEventMetric) ProtoMessage() {}

func (x *ClientEventMetric) ProtoReflect() protoreflect.Message {
	mi := &file_dolt_services_eventsapi_v1alpha1_client_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientEventMetric.ProtoReflect.Descriptor instead.
func (*ClientEventMetric) Descriptor() ([]byte, []int) {
	return file_dolt_services_eventsapi_v1alpha1_client_event_proto_rawDescGZIP(), []int{1}
}

func (m *ClientEventMetric) GetMetricOneof() isClientEventMetric_MetricOneof {
	if m != nil {
		return m.MetricOneof
	}
	return nil
}

func (x *ClientEventMetric) GetDuration() *durationpb.Duration {
	if x, ok := x.GetMetricOneof().(*ClientEventMetric_Duration); ok {
		return x.Duration
	}
	return nil
}

func (x *ClientEventMetric) GetCount() int32 {
	if x, ok := x.GetMetricOneof().(*ClientEventMetric_Count); ok {
		return x.Count
	}
	return 0
}

func (x *ClientEventMetric) GetMetricId() MetricID {
	if x != nil {
		return x.MetricId
	}
	return MetricID_METRIC_UNSPECIFIED
}

type isClientEventMetric_MetricOneof interface {
	isClientEventMetric_MetricOneof()
}

type ClientEventMetric_Duration struct {
	Duration *durationpb.Duration `protobuf:"bytes,1,opt,name=duration,proto3,oneof"`
}

type ClientEventMetric_Count struct {
	Count int32 `protobuf:"varint,2,opt,name=count,proto3,oneof"`
}

func (*ClientEventMetric_Duration) isClientEventMetric_MetricOneof() {}

func (*ClientEventMetric_Count) isClientEventMetric_MetricOneof() {}

type ClientEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	StartTime  *timestamppb.Timestamp  `protobuf:"bytes,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime    *timestamppb.Timestamp  `protobuf:"bytes,3,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Type       ClientEventType         `protobuf:"varint,4,opt,name=type,proto3,enum=dolt.services.eventsapi.v1alpha1.ClientEventType" json:"type,omitempty"`
	Attributes []*ClientEventAttribute `protobuf:"bytes,5,rep,name=attributes,proto3" json:"attributes,omitempty"`
	Metrics    []*ClientEventMetric    `protobuf:"bytes,6,rep,name=metrics,proto3" json:"metrics,omitempty"`
}

func (x *ClientEvent) Reset() {
	*x = ClientEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dolt_services_eventsapi_v1alpha1_client_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientEvent) ProtoMessage() {}

func (x *ClientEvent) ProtoReflect() protoreflect.Message {
	mi := &file_dolt_services_eventsapi_v1alpha1_client_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientEvent.ProtoReflect.Descriptor instead.
func (*ClientEvent) Descriptor() ([]byte, []int) {
	return file_dolt_services_eventsapi_v1alpha1_client_event_proto_rawDescGZIP(), []int{2}
}

func (x *ClientEvent) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ClientEvent) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *ClientEvent) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *ClientEvent) GetType() ClientEventType {
	if x != nil {
		return x.Type
	}
	return ClientEventType_TYPE_UNSPECIFIED
}

func (x *ClientEvent) GetAttributes() []*ClientEventAttribute {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *ClientEvent) GetMetrics() []*ClientEventMetric {
	if x != nil {
		return x.Metrics
	}
	return nil
}

type LogEventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MachineId string         `protobuf:"bytes,1,opt,name=machine_id,json=machineId,proto3" json:"machine_id,omitempty"`
	Extra     string         `protobuf:"bytes,2,opt,name=extra,proto3" json:"extra,omitempty"`
	Version   string         `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Platform  Platform       `protobuf:"varint,4,opt,name=platform,proto3,enum=dolt.services.eventsapi.v1alpha1.Platform" json:"platform,omitempty"`
	Events    []*ClientEvent `protobuf:"bytes,5,rep,name=events,proto3" json:"events,omitempty"`
	App       AppID          `protobuf:"varint,6,opt,name=app,proto3,enum=dolt.services.eventsapi.v1alpha1.AppID" json:"app,omitempty"`
}

func (x *LogEventsRequest) Reset() {
	*x = LogEventsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dolt_services_eventsapi_v1alpha1_client_event_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogEventsRequest) ProtoMessage() {}

func (x *LogEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dolt_services_eventsapi_v1alpha1_client_event_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogEventsRequest.ProtoReflect.Descriptor instead.
func (*LogEventsRequest) Descriptor() ([]byte, []int) {
	return file_dolt_services_eventsapi_v1alpha1_client_event_proto_rawDescGZIP(), []int{3}
}

func (x *LogEventsRequest) GetMachineId() string {
	if x != nil {
		return x.MachineId
	}
	return ""
}

func (x *LogEventsRequest) GetExtra() string {
	if x != nil {
		return x.Extra
	}
	return ""
}

func (x *LogEventsRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *LogEventsRequest) GetPlatform() Platform {
	if x != nil {
		return x.Platform
	}
	return Platform_PLATFORM_UNSPECIFIED
}

func (x *LogEventsRequest) GetEvents() []*ClientEvent {
	if x != nil {
		return x.Events
	}
	return nil
}

func (x *LogEventsRequest) GetApp() AppID {
	if x != nil {
		return x.App
	}
	return AppID_APP_ID_UNSPECIFIED
}

type LogEventsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LogEventsResponse) Reset() {
	*x = LogEventsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dolt_services_eventsapi_v1alpha1_client_event_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogEventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogEventsResponse) ProtoMessage() {}

func (x *LogEventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dolt_services_eventsapi_v1alpha1_client_event_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogEventsResponse.ProtoReflect.Descriptor instead.
func (*LogEventsResponse) Descriptor() ([]byte, []int) {
	return file_dolt_services_eventsapi_v1alpha1_client_event_proto_rawDescGZIP(), []int{4}
}

var File_dolt_services_eventsapi_v1alpha1_client_event_proto protoreflect.FileDescriptor

var file_dolt_services_eventsapi_v1alpha1_client_event_proto_rawDesc = []byte{
	0x0a, 0x33, 0x64, 0x6f, 0x6c, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x64, 0x6f, 0x6c, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x36, 0x64, 0x6f, 0x6c, 0x74, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f,
	0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x6b, 0x0a, 0x14, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x12, 0x3d, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x64, 0x6f, 0x6c, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xbd, 0x01,
	0x0a, 0x11, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x12, 0x37, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x48, 0x00, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x47, 0x0a, 0x09, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x69,
	0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x64, 0x6f, 0x6c, 0x74, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x49, 0x44, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x49, 0x64, 0x42, 0x0e, 0x0a,
	0x0c, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x22, 0xfd, 0x02,
	0x0a, 0x0b, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x39, 0x0a,
	0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x45, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x31, 0x2e,
	0x64, 0x6f, 0x6c, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x56, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x64, 0x6f, 0x6c,
	0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x4d,
	0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x33, 0x2e, 0x64, 0x6f, 0x6c, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x22, 0xab, 0x02,
	0x0a, 0x10, 0x4c, 0x6f, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x46, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x64, 0x6f, 0x6c, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x52,
	0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x45, 0x0a, 0x06, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x64, 0x6f, 0x6c, 0x74,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x39, 0x0a, 0x03, 0x61, 0x70, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e,
	0x64, 0x6f, 0x6c, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x41, 0x70, 0x70, 0x49, 0x44, 0x52, 0x03, 0x61, 0x70, 0x70, 0x22, 0x13, 0x0a, 0x11, 0x4c,
	0x6f, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x32, 0x8b, 0x01, 0x0a, 0x13, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x74, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x32, 0x2e, 0x64, 0x6f, 0x6c, 0x74, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x64, 0x6f, 0x6c, 0x74,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x6f, 0x67,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x51,
	0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x6c,
	0x74, 0x68, 0x75, 0x62, 0x2f, 0x64, 0x6f, 0x6c, 0x74, 0x2f, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x6f, 0x6c, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x61, 0x70,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dolt_services_eventsapi_v1alpha1_client_event_proto_rawDescOnce sync.Once
	file_dolt_services_eventsapi_v1alpha1_client_event_proto_rawDescData = file_dolt_services_eventsapi_v1alpha1_client_event_proto_rawDesc
)

func file_dolt_services_eventsapi_v1alpha1_client_event_proto_rawDescGZIP() []byte {
	file_dolt_services_eventsapi_v1alpha1_client_event_proto_rawDescOnce.Do(func() {
		file_dolt_services_eventsapi_v1alpha1_client_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_dolt_services_eventsapi_v1alpha1_client_event_proto_rawDescData)
	})
	return file_dolt_services_eventsapi_v1alpha1_client_event_proto_rawDescData
}

var file_dolt_services_eventsapi_v1alpha1_client_event_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_dolt_services_eventsapi_v1alpha1_client_event_proto_goTypes = []interface{}{
	(*ClientEventAttribute)(nil),  // 0: dolt.services.eventsapi.v1alpha1.ClientEventAttribute
	(*ClientEventMetric)(nil),     // 1: dolt.services.eventsapi.v1alpha1.ClientEventMetric
	(*ClientEvent)(nil),           // 2: dolt.services.eventsapi.v1alpha1.ClientEvent
	(*LogEventsRequest)(nil),      // 3: dolt.services.eventsapi.v1alpha1.LogEventsRequest
	(*LogEventsResponse)(nil),     // 4: dolt.services.eventsapi.v1alpha1.LogEventsResponse
	(AttributeID)(0),              // 5: dolt.services.eventsapi.v1alpha1.AttributeID
	(*durationpb.Duration)(nil),   // 6: google.protobuf.Duration
	(MetricID)(0),                 // 7: dolt.services.eventsapi.v1alpha1.MetricID
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
	(ClientEventType)(0),          // 9: dolt.services.eventsapi.v1alpha1.ClientEventType
	(Platform)(0),                 // 10: dolt.services.eventsapi.v1alpha1.Platform
	(AppID)(0),                    // 11: dolt.services.eventsapi.v1alpha1.AppID
}
var file_dolt_services_eventsapi_v1alpha1_client_event_proto_depIdxs = []int32{
	5,  // 0: dolt.services.eventsapi.v1alpha1.ClientEventAttribute.id:type_name -> dolt.services.eventsapi.v1alpha1.AttributeID
	6,  // 1: dolt.services.eventsapi.v1alpha1.ClientEventMetric.duration:type_name -> google.protobuf.Duration
	7,  // 2: dolt.services.eventsapi.v1alpha1.ClientEventMetric.metric_id:type_name -> dolt.services.eventsapi.v1alpha1.MetricID
	8,  // 3: dolt.services.eventsapi.v1alpha1.ClientEvent.start_time:type_name -> google.protobuf.Timestamp
	8,  // 4: dolt.services.eventsapi.v1alpha1.ClientEvent.end_time:type_name -> google.protobuf.Timestamp
	9,  // 5: dolt.services.eventsapi.v1alpha1.ClientEvent.type:type_name -> dolt.services.eventsapi.v1alpha1.ClientEventType
	0,  // 6: dolt.services.eventsapi.v1alpha1.ClientEvent.attributes:type_name -> dolt.services.eventsapi.v1alpha1.ClientEventAttribute
	1,  // 7: dolt.services.eventsapi.v1alpha1.ClientEvent.metrics:type_name -> dolt.services.eventsapi.v1alpha1.ClientEventMetric
	10, // 8: dolt.services.eventsapi.v1alpha1.LogEventsRequest.platform:type_name -> dolt.services.eventsapi.v1alpha1.Platform
	2,  // 9: dolt.services.eventsapi.v1alpha1.LogEventsRequest.events:type_name -> dolt.services.eventsapi.v1alpha1.ClientEvent
	11, // 10: dolt.services.eventsapi.v1alpha1.LogEventsRequest.app:type_name -> dolt.services.eventsapi.v1alpha1.AppID
	3,  // 11: dolt.services.eventsapi.v1alpha1.ClientEventsService.LogEvents:input_type -> dolt.services.eventsapi.v1alpha1.LogEventsRequest
	4,  // 12: dolt.services.eventsapi.v1alpha1.ClientEventsService.LogEvents:output_type -> dolt.services.eventsapi.v1alpha1.LogEventsResponse
	12, // [12:13] is the sub-list for method output_type
	11, // [11:12] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_dolt_services_eventsapi_v1alpha1_client_event_proto_init() }
func file_dolt_services_eventsapi_v1alpha1_client_event_proto_init() {
	if File_dolt_services_eventsapi_v1alpha1_client_event_proto != nil {
		return
	}
	file_dolt_services_eventsapi_v1alpha1_event_constants_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_dolt_services_eventsapi_v1alpha1_client_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientEventAttribute); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dolt_services_eventsapi_v1alpha1_client_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientEventMetric); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dolt_services_eventsapi_v1alpha1_client_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dolt_services_eventsapi_v1alpha1_client_event_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogEventsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dolt_services_eventsapi_v1alpha1_client_event_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogEventsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_dolt_services_eventsapi_v1alpha1_client_event_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*ClientEventMetric_Duration)(nil),
		(*ClientEventMetric_Count)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_dolt_services_eventsapi_v1alpha1_client_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dolt_services_eventsapi_v1alpha1_client_event_proto_goTypes,
		DependencyIndexes: file_dolt_services_eventsapi_v1alpha1_client_event_proto_depIdxs,
		MessageInfos:      file_dolt_services_eventsapi_v1alpha1_client_event_proto_msgTypes,
	}.Build()
	File_dolt_services_eventsapi_v1alpha1_client_event_proto = out.File
	file_dolt_services_eventsapi_v1alpha1_client_event_proto_rawDesc = nil
	file_dolt_services_eventsapi_v1alpha1_client_event_proto_goTypes = nil
	file_dolt_services_eventsapi_v1alpha1_client_event_proto_depIdxs = nil
}
