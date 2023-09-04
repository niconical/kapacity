//
//Copyright 2023 The Kapacity Authors.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.3
// source: provider.proto

package api

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

type QueryLatestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query *Query `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
}

func (x *QueryLatestRequest) Reset() {
	*x = QueryLatestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provider_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryLatestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryLatestRequest) ProtoMessage() {}

func (x *QueryLatestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_provider_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryLatestRequest.ProtoReflect.Descriptor instead.
func (*QueryLatestRequest) Descriptor() ([]byte, []int) {
	return file_provider_proto_rawDescGZIP(), []int{0}
}

func (x *QueryLatestRequest) GetQuery() *Query {
	if x != nil {
		return x.Query
	}
	return nil
}

type QueryLatestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Samples []*Sample `protobuf:"bytes,1,rep,name=samples,proto3" json:"samples,omitempty"`
}

func (x *QueryLatestResponse) Reset() {
	*x = QueryLatestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provider_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryLatestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryLatestResponse) ProtoMessage() {}

func (x *QueryLatestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_provider_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryLatestResponse.ProtoReflect.Descriptor instead.
func (*QueryLatestResponse) Descriptor() ([]byte, []int) {
	return file_provider_proto_rawDescGZIP(), []int{1}
}

func (x *QueryLatestResponse) GetSamples() []*Sample {
	if x != nil {
		return x.Samples
	}
	return nil
}

type QueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query *Query                 `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Start *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=start,proto3" json:"start,omitempty"`
	End   *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=end,proto3" json:"end,omitempty"`
	Step  *durationpb.Duration   `protobuf:"bytes,4,opt,name=step,proto3" json:"step,omitempty"`
}

func (x *QueryRequest) Reset() {
	*x = QueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provider_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRequest) ProtoMessage() {}

func (x *QueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_provider_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRequest.ProtoReflect.Descriptor instead.
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return file_provider_proto_rawDescGZIP(), []int{2}
}

func (x *QueryRequest) GetQuery() *Query {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *QueryRequest) GetStart() *timestamppb.Timestamp {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *QueryRequest) GetEnd() *timestamppb.Timestamp {
	if x != nil {
		return x.End
	}
	return nil
}

func (x *QueryRequest) GetStep() *durationpb.Duration {
	if x != nil {
		return x.Step
	}
	return nil
}

type QueryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Series []*Series `protobuf:"bytes,1,rep,name=series,proto3" json:"series,omitempty"`
}

func (x *QueryResponse) Reset() {
	*x = QueryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provider_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryResponse) ProtoMessage() {}

func (x *QueryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_provider_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryResponse.ProtoReflect.Descriptor instead.
func (*QueryResponse) Descriptor() ([]byte, []int) {
	return file_provider_proto_rawDescGZIP(), []int{3}
}

func (x *QueryResponse) GetSeries() []*Series {
	if x != nil {
		return x.Series
	}
	return nil
}

var File_provider_proto protoreflect.FileDescriptor

var file_provider_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x17, 0x69, 0x6f, 0x2e, 0x6b, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x73, 0x74, 0x61,
	0x63, 0x6b, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x1a, 0x0c, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4a, 0x0a, 0x12, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34,
	0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x69, 0x6f, 0x2e, 0x6b, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x73, 0x74, 0x61, 0x63, 0x6b,
	0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x22, 0x50, 0x0a, 0x13, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4c, 0x61, 0x74,
	0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x07, 0x73,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x69,
	0x6f, 0x2e, 0x6b, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x07, 0x73,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x22, 0xd3, 0x01, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x69, 0x6f, 0x2e, 0x6b, 0x61, 0x70, 0x61,
	0x63, 0x69, 0x74, 0x79, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x30, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12,
	0x2c, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x2d, 0x0a,
	0x04, 0x73, 0x74, 0x65, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x73, 0x74, 0x65, 0x70, 0x22, 0x48, 0x0a, 0x0d,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a,
	0x06, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x69, 0x6f, 0x2e, 0x6b, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x73, 0x74, 0x61, 0x63, 0x6b,
	0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x52, 0x06,
	0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x32, 0xd3, 0x01, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x68, 0x0a, 0x0b, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x2e, 0x69, 0x6f, 0x2e, 0x6b,
	0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x69, 0x6f, 0x2e, 0x6b, 0x61, 0x70, 0x61,
	0x63, 0x69, 0x74, 0x79, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x25, 0x2e,
	0x69, 0x6f, 0x2e, 0x6b, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x73, 0x74, 0x61, 0x63, 0x6b,
	0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x69, 0x6f, 0x2e, 0x6b, 0x61, 0x70, 0x61, 0x63, 0x69,
	0x74, 0x79, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x38, 0x5a, 0x36,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x72, 0x61, 0x61, 0x73,
	0x2d, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x6b, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_provider_proto_rawDescOnce sync.Once
	file_provider_proto_rawDescData = file_provider_proto_rawDesc
)

func file_provider_proto_rawDescGZIP() []byte {
	file_provider_proto_rawDescOnce.Do(func() {
		file_provider_proto_rawDescData = protoimpl.X.CompressGZIP(file_provider_proto_rawDescData)
	})
	return file_provider_proto_rawDescData
}

var file_provider_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_provider_proto_goTypes = []interface{}{
	(*QueryLatestRequest)(nil),    // 0: io.kapacitystack.metric.QueryLatestRequest
	(*QueryLatestResponse)(nil),   // 1: io.kapacitystack.metric.QueryLatestResponse
	(*QueryRequest)(nil),          // 2: io.kapacitystack.metric.QueryRequest
	(*QueryResponse)(nil),         // 3: io.kapacitystack.metric.QueryResponse
	(*Query)(nil),                 // 4: io.kapacitystack.metric.Query
	(*Sample)(nil),                // 5: io.kapacitystack.metric.Sample
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),   // 7: google.protobuf.Duration
	(*Series)(nil),                // 8: io.kapacitystack.metric.Series
}
var file_provider_proto_depIdxs = []int32{
	4, // 0: io.kapacitystack.metric.QueryLatestRequest.query:type_name -> io.kapacitystack.metric.Query
	5, // 1: io.kapacitystack.metric.QueryLatestResponse.samples:type_name -> io.kapacitystack.metric.Sample
	4, // 2: io.kapacitystack.metric.QueryRequest.query:type_name -> io.kapacitystack.metric.Query
	6, // 3: io.kapacitystack.metric.QueryRequest.start:type_name -> google.protobuf.Timestamp
	6, // 4: io.kapacitystack.metric.QueryRequest.end:type_name -> google.protobuf.Timestamp
	7, // 5: io.kapacitystack.metric.QueryRequest.step:type_name -> google.protobuf.Duration
	8, // 6: io.kapacitystack.metric.QueryResponse.series:type_name -> io.kapacitystack.metric.Series
	0, // 7: io.kapacitystack.metric.ProviderService.QueryLatest:input_type -> io.kapacitystack.metric.QueryLatestRequest
	2, // 8: io.kapacitystack.metric.ProviderService.Query:input_type -> io.kapacitystack.metric.QueryRequest
	1, // 9: io.kapacitystack.metric.ProviderService.QueryLatest:output_type -> io.kapacitystack.metric.QueryLatestResponse
	3, // 10: io.kapacitystack.metric.ProviderService.Query:output_type -> io.kapacitystack.metric.QueryResponse
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_provider_proto_init() }
func file_provider_proto_init() {
	if File_provider_proto != nil {
		return
	}
	file_metric_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_provider_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryLatestRequest); i {
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
		file_provider_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryLatestResponse); i {
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
		file_provider_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRequest); i {
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
		file_provider_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_provider_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_provider_proto_goTypes,
		DependencyIndexes: file_provider_proto_depIdxs,
		MessageInfos:      file_provider_proto_msgTypes,
	}.Build()
	File_provider_proto = out.File
	file_provider_proto_rawDesc = nil
	file_provider_proto_goTypes = nil
	file_provider_proto_depIdxs = nil
}
