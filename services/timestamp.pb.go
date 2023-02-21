// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: timestamp.proto

package services

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// *
// An exact date and time. This is the same data structure as the protobuf Timestamp.proto (see the
// comments in https://github.com/google/protobuf/blob/master/src/google/protobuf/timestamp.proto)
type Timestamp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// Number of complete seconds since the start of the epoch
	Seconds int64 `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
	// *
	// Number of nanoseconds since the start of the last second
	Nanos int32 `protobuf:"varint,2,opt,name=nanos,proto3" json:"nanos,omitempty"`
}

func (x *Timestamp) Reset() {
	*x = Timestamp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_timestamp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Timestamp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Timestamp) ProtoMessage() {}

func (x *Timestamp) ProtoReflect() protoreflect.Message {
	mi := &file_timestamp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Timestamp.ProtoReflect.Descriptor instead.
func (*Timestamp) Descriptor() ([]byte, []int) {
	return file_timestamp_proto_rawDescGZIP(), []int{0}
}

func (x *Timestamp) GetSeconds() int64 {
	if x != nil {
		return x.Seconds
	}
	return 0
}

func (x *Timestamp) GetNanos() int32 {
	if x != nil {
		return x.Nanos
	}
	return 0
}

// *
// An exact date and time,  with a resolution of one second (no nanoseconds).
type TimestampSeconds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// Number of complete seconds since the start of the epoch
	Seconds int64 `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
}

func (x *TimestampSeconds) Reset() {
	*x = TimestampSeconds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_timestamp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimestampSeconds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimestampSeconds) ProtoMessage() {}

func (x *TimestampSeconds) ProtoReflect() protoreflect.Message {
	mi := &file_timestamp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimestampSeconds.ProtoReflect.Descriptor instead.
func (*TimestampSeconds) Descriptor() ([]byte, []int) {
	return file_timestamp_proto_rawDescGZIP(), []int{1}
}

func (x *TimestampSeconds) GetSeconds() int64 {
	if x != nil {
		return x.Seconds
	}
	return 0
}

var File_timestamp_proto protoreflect.FileDescriptor

var file_timestamp_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3b, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6e, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x6e, 0x61, 0x6e, 0x6f, 0x73, 0x22, 0x2c, 0x0a, 0x10, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x73, 0x42, 0x26, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x64, 0x65, 0x72,
	0x61, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6a, 0x61, 0x76, 0x61, 0x50, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_timestamp_proto_rawDescOnce sync.Once
	file_timestamp_proto_rawDescData = file_timestamp_proto_rawDesc
)

func file_timestamp_proto_rawDescGZIP() []byte {
	file_timestamp_proto_rawDescOnce.Do(func() {
		file_timestamp_proto_rawDescData = protoimpl.X.CompressGZIP(file_timestamp_proto_rawDescData)
	})
	return file_timestamp_proto_rawDescData
}

var file_timestamp_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_timestamp_proto_goTypes = []interface{}{
	(*Timestamp)(nil),        // 0: proto.Timestamp
	(*TimestampSeconds)(nil), // 1: proto.TimestampSeconds
}
var file_timestamp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_timestamp_proto_init() }
func file_timestamp_proto_init() {
	if File_timestamp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_timestamp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Timestamp); i {
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
		file_timestamp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimestampSeconds); i {
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
			RawDescriptor: file_timestamp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_timestamp_proto_goTypes,
		DependencyIndexes: file_timestamp_proto_depIdxs,
		MessageInfos:      file_timestamp_proto_msgTypes,
	}.Build()
	File_timestamp_proto = out.File
	file_timestamp_proto_rawDesc = nil
	file_timestamp_proto_goTypes = nil
	file_timestamp_proto_depIdxs = nil
}
