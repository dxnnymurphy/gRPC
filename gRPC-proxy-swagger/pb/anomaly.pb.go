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

	Topic string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
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

func (x *Metric) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

type Metric1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=endTime,proto3" json:"endTime,omitempty"`
}

func (x *Metric1) Reset() {
	*x = Metric1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anomaly_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metric1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metric1) ProtoMessage() {}

func (x *Metric1) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Metric1.ProtoReflect.Descriptor instead.
func (*Metric1) Descriptor() ([]byte, []int) {
	return file_anomaly_proto_rawDescGZIP(), []int{1}
}

func (x *Metric1) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *Metric1) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
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
		mi := &file_anomaly_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnomalyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnomalyRequest) ProtoMessage() {}

func (x *AnomalyRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use AnomalyRequest.ProtoReflect.Descriptor instead.
func (*AnomalyRequest) Descriptor() ([]byte, []int) {
	return file_anomaly_proto_rawDescGZIP(), []int{2}
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
		mi := &file_anomaly_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnomalyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnomalyResponse) ProtoMessage() {}

func (x *AnomalyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_anomaly_proto_msgTypes[3]
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
	return file_anomaly_proto_rawDescGZIP(), []int{3}
}

func (x *AnomalyResponse) GetResponse() []string {
	if x != nil {
		return x.Response
	}
	return nil
}

type ModelTrainRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metrics []*Metric1 `protobuf:"bytes,1,rep,name=metrics,proto3" json:"metrics,omitempty"`
}

func (x *ModelTrainRequest) Reset() {
	*x = ModelTrainRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anomaly_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelTrainRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelTrainRequest) ProtoMessage() {}

func (x *ModelTrainRequest) ProtoReflect() protoreflect.Message {
	mi := &file_anomaly_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelTrainRequest.ProtoReflect.Descriptor instead.
func (*ModelTrainRequest) Descriptor() ([]byte, []int) {
	return file_anomaly_proto_rawDescGZIP(), []int{4}
}

func (x *ModelTrainRequest) GetMetrics() []*Metric1 {
	if x != nil {
		return x.Metrics
	}
	return nil
}

type ModelTrainResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response []string `protobuf:"bytes,1,rep,name=response,proto3" json:"response,omitempty"`
}

func (x *ModelTrainResponse) Reset() {
	*x = ModelTrainResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anomaly_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelTrainResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelTrainResponse) ProtoMessage() {}

func (x *ModelTrainResponse) ProtoReflect() protoreflect.Message {
	mi := &file_anomaly_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelTrainResponse.ProtoReflect.Descriptor instead.
func (*ModelTrainResponse) Descriptor() ([]byte, []int) {
	return file_anomaly_proto_rawDescGZIP(), []int{5}
}

func (x *ModelTrainResponse) GetResponse() []string {
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
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1e, 0x0a, 0x06, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x70, 0x69, 0x63, 0x22, 0x79, 0x0a, 0x07, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x31, 0x12,
	0x38, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x65, 0x6e, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22,
	0x36, 0x0a, 0x0e, 0x41, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x24, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x07,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x22, 0x2d, 0x0a, 0x0f, 0x41, 0x6e, 0x6f, 0x6d, 0x61,
	0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3a, 0x0a, 0x11, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x54,
	0x72, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x07, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70,
	0x62, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x31, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x22, 0x30, 0x0a, 0x12, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x54, 0x72, 0x61, 0x69, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x32, 0x82, 0x01, 0x0a, 0x10, 0x41, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79,
	0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x07, 0x50, 0x72, 0x65,
	0x64, 0x69, 0x63, 0x74, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x6e, 0x6f, 0x6d, 0x61, 0x6c,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x6e,
	0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x38, 0x0a, 0x05, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x20, 0x5a, 0x1e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x78, 0x6e, 0x6e, 0x79, 0x6d, 0x75, 0x72,
	0x70, 0x68, 0x79, 0x2f, 0x67, 0x52, 0x50, 0x43, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
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

var file_anomaly_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_anomaly_proto_goTypes = []interface{}{
	(*Metric)(nil),                // 0: pb.Metric
	(*Metric1)(nil),               // 1: pb.Metric1
	(*AnomalyRequest)(nil),        // 2: pb.AnomalyRequest
	(*AnomalyResponse)(nil),       // 3: pb.AnomalyResponse
	(*ModelTrainRequest)(nil),     // 4: pb.ModelTrainRequest
	(*ModelTrainResponse)(nil),    // 5: pb.ModelTrainResponse
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_anomaly_proto_depIdxs = []int32{
	6, // 0: pb.Metric1.startTime:type_name -> google.protobuf.Timestamp
	6, // 1: pb.Metric1.endTime:type_name -> google.protobuf.Timestamp
	0, // 2: pb.AnomalyRequest.metrics:type_name -> pb.Metric
	1, // 3: pb.ModelTrainRequest.metrics:type_name -> pb.Metric1
	2, // 4: pb.AnomalyDetection.Predict:input_type -> pb.AnomalyRequest
	4, // 5: pb.AnomalyDetection.Train:input_type -> pb.ModelTrainRequest
	3, // 6: pb.AnomalyDetection.Predict:output_type -> pb.AnomalyResponse
	5, // 7: pb.AnomalyDetection.Train:output_type -> pb.ModelTrainResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
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
			switch v := v.(*Metric1); i {
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
		file_anomaly_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_anomaly_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelTrainRequest); i {
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
		file_anomaly_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelTrainResponse); i {
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
			NumMessages:   6,
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