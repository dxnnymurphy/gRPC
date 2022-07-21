// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.2
// source: anomaly.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type Metric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time    *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	Topic   string                 `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	Refresh bool                   `protobuf:"varint,3,opt,name=refresh,proto3" json:"refresh,omitempty"`
	Stop    bool                   `protobuf:"varint,4,opt,name=stop,proto3" json:"stop,omitempty"`
}

func (x *Metric) Reset() {
	*x = Metric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anomaly_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metric) ProtoMessage() {}

func (x *Metric) ProtoReflect() protoreflect.Message {
	mi := &file_anomaly_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metric.ProtoReflect.Descriptor instead.
func (*Metric) Descriptor() ([]byte, []int) {
	return file_anomaly_proto_rawDescGZIP(), []int{0}
}

func (x *Metric) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *Metric) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *Metric) GetRefresh() bool {
	if x != nil {
		return x.Refresh
	}
	return false
}

func (x *Metric) GetStop() bool {
	if x != nil {
		return x.Stop
	}
	return false
}

type AnomalyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metrics []*Metric `protobuf:"bytes,1,rep,name=metrics,proto3" json:"metrics,omitempty"`
}

func (x *AnomalyRequest) Reset() {
	*x = AnomalyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anomaly_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnomalyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnomalyRequest) ProtoMessage() {}

func (x *AnomalyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_anomaly_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnomalyRequest.ProtoReflect.Descriptor instead.
func (*AnomalyRequest) Descriptor() ([]byte, []int) {
	return file_anomaly_proto_rawDescGZIP(), []int{1}
}

func (x *AnomalyRequest) GetMetrics() []*Metric {
	if x != nil {
		return x.Metrics
	}
	return nil
}

type AnomalyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response []string `protobuf:"bytes,1,rep,name=response,proto3" json:"response,omitempty"`
}

func (x *AnomalyResponse) Reset() {
	*x = AnomalyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anomaly_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnomalyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnomalyResponse) ProtoMessage() {}

func (x *AnomalyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_anomaly_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnomalyResponse.ProtoReflect.Descriptor instead.
func (*AnomalyResponse) Descriptor() ([]byte, []int) {
	return file_anomaly_proto_rawDescGZIP(), []int{2}
}

func (x *AnomalyResponse) GetResponse() []string {
	if x != nil {
		return x.Response
	}
	return nil
}

var File_anomaly_proto protoreflect.FileDescriptor

var file_anomaly_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7c, 0x0a, 0x06, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x2e,
	0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x70, 0x69, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x74, 0x6f, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x73, 0x74,
	0x6f, 0x70, 0x22, 0x36, 0x0a, 0x0e, 0x41, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x22, 0x2d, 0x0a, 0x0f, 0x41, 0x6e,
	0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x48, 0x0a, 0x10, 0x41, 0x6e, 0x6f,
	0x6d, 0x61, 0x6c, 0x79, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a,
	0x07, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x6e,
	0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70,
	0x62, 0x2e, 0x41, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x20, 0x5a, 0x1e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x64, 0x78, 0x6e, 0x6e, 0x79, 0x6d, 0x75, 0x72, 0x70, 0x68, 0x79, 0x2f, 0x67, 0x52,
	0x50, 0x43, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_anomaly_proto_rawDescOnce sync.Once
	file_anomaly_proto_rawDescData = file_anomaly_proto_rawDesc
)

func file_anomaly_proto_rawDescGZIP() []byte {
	file_anomaly_proto_rawDescOnce.Do(func() {
		file_anomaly_proto_rawDescData = protoimpl.X.CompressGZIP(file_anomaly_proto_rawDescData)
	})
	return file_anomaly_proto_rawDescData
}

var file_anomaly_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_anomaly_proto_goTypes = []interface{}{
	(*Metric)(nil),                // 0: pb.Metric
	(*AnomalyRequest)(nil),        // 1: pb.AnomalyRequest
	(*AnomalyResponse)(nil),       // 2: pb.AnomalyResponse
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_anomaly_proto_depIdxs = []int32{
	3, // 0: pb.Metric.time:type_name -> google.protobuf.Timestamp
	0, // 1: pb.AnomalyRequest.metrics:type_name -> pb.Metric
	1, // 2: pb.AnomalyDetection.Predict:input_type -> pb.AnomalyRequest
	2, // 3: pb.AnomalyDetection.Predict:output_type -> pb.AnomalyResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_anomaly_proto_init() }
func file_anomaly_proto_init() {
	if File_anomaly_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_anomaly_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metric); i {
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
		file_anomaly_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnomalyRequest); i {
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
		file_anomaly_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnomalyResponse); i {
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
			RawDescriptor: file_anomaly_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_anomaly_proto_goTypes,
		DependencyIndexes: file_anomaly_proto_depIdxs,
		MessageInfos:      file_anomaly_proto_msgTypes,
	}.Build()
	File_anomaly_proto = out.File
	file_anomaly_proto_rawDesc = nil
	file_anomaly_proto_goTypes = nil
	file_anomaly_proto_depIdxs = nil
}