// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.17.3
// source: transaction_get_record.proto

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

//*
// Get the record for a transaction. If the transaction requested a record, then the record lasts
// for one hour, and a state proof is available for it. If the transaction created an account, file,
// or smart contract instance, then the record will contain the ID for what it created. If the
// transaction called a smart contract function, then the record contains the result of that call.
// If the transaction was a cryptocurrency transfer, then the record includes the TransferList which
// gives the details of that transfer. If the transaction didn't return anything that should be in
// the record, then the results field will be set to nothing.
type TransactionGetRecordQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//*
	// Standard info sent from client to node, including the signed payment, and what kind of
	// response is requested (cost, state proof, both, or neither).
	Header *QueryHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	//*
	// The ID of the transaction for which the record is requested.
	TransactionID *TransactionID `protobuf:"bytes,2,opt,name=transactionID,proto3" json:"transactionID,omitempty"`
	//*
	// Whether records of processing duplicate transactions should be returned along with the record
	// of processing the first consensus transaction with the given id whose status was neither
	// <tt>INVALID_NODE_ACCOUNT</tt> nor <tt>INVALID_PAYER_SIGNATURE</tt>; <b>or</b>, if no such
	// record exists, the record of processing the first transaction to reach consensus with the
	// given transaction id..
	IncludeDuplicates bool `protobuf:"varint,3,opt,name=includeDuplicates,proto3" json:"includeDuplicates,omitempty"`
	//*
	// Whether the response should include the records of any child transactions spawned by the
	// top-level transaction with the given transactionID.
	IncludeChildRecords bool `protobuf:"varint,4,opt,name=include_child_records,json=includeChildRecords,proto3" json:"include_child_records,omitempty"`
}

func (x *TransactionGetRecordQuery) Reset() {
	*x = TransactionGetRecordQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_get_record_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionGetRecordQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionGetRecordQuery) ProtoMessage() {}

func (x *TransactionGetRecordQuery) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_get_record_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionGetRecordQuery.ProtoReflect.Descriptor instead.
func (*TransactionGetRecordQuery) Descriptor() ([]byte, []int) {
	return file_transaction_get_record_proto_rawDescGZIP(), []int{0}
}

func (x *TransactionGetRecordQuery) GetHeader() *QueryHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *TransactionGetRecordQuery) GetTransactionID() *TransactionID {
	if x != nil {
		return x.TransactionID
	}
	return nil
}

func (x *TransactionGetRecordQuery) GetIncludeDuplicates() bool {
	if x != nil {
		return x.IncludeDuplicates
	}
	return false
}

func (x *TransactionGetRecordQuery) GetIncludeChildRecords() bool {
	if x != nil {
		return x.IncludeChildRecords
	}
	return false
}

//*
// Response when the client sends the node TransactionGetRecordQuery
type TransactionGetRecordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//*
	// Standard response from node to client, including the requested fields: cost, or state proof,
	// or both, or neither.
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	//*
	// Either the record of processing the first consensus transaction with the given id whose
	// status was neither <tt>INVALID_NODE_ACCOUNT</tt> nor <tt>INVALID_PAYER_SIGNATURE</tt>;
	// <b>or</b>, if no such record exists, the record of processing the first transaction to reach
	// consensus with the given transaction id.
	TransactionRecord *TransactionRecord `protobuf:"bytes,3,opt,name=transactionRecord,proto3" json:"transactionRecord,omitempty"`
	//*
	// The records of processing all consensus transaction with the same id as the distinguished
	// record above, in chronological order.
	DuplicateTransactionRecords []*TransactionRecord `protobuf:"bytes,4,rep,name=duplicateTransactionRecords,proto3" json:"duplicateTransactionRecords,omitempty"`
	//*
	// The records of processing all child transaction spawned by the transaction with the given
	// top-level id, in consensus order. Always empty if the top-level status is UNKNOWN.
	ChildTransactionRecords []*TransactionRecord `protobuf:"bytes,5,rep,name=child_transaction_records,json=childTransactionRecords,proto3" json:"child_transaction_records,omitempty"`
}

func (x *TransactionGetRecordResponse) Reset() {
	*x = TransactionGetRecordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_get_record_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionGetRecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionGetRecordResponse) ProtoMessage() {}

func (x *TransactionGetRecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_get_record_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionGetRecordResponse.ProtoReflect.Descriptor instead.
func (*TransactionGetRecordResponse) Descriptor() ([]byte, []int) {
	return file_transaction_get_record_proto_rawDescGZIP(), []int{1}
}

func (x *TransactionGetRecordResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *TransactionGetRecordResponse) GetTransactionRecord() *TransactionRecord {
	if x != nil {
		return x.TransactionRecord
	}
	return nil
}

func (x *TransactionGetRecordResponse) GetDuplicateTransactionRecords() []*TransactionRecord {
	if x != nil {
		return x.DuplicateTransactionRecords
	}
	return nil
}

func (x *TransactionGetRecordResponse) GetChildTransactionRecords() []*TransactionRecord {
	if x != nil {
		return x.ChildTransactionRecords
	}
	return nil
}

var File_transaction_get_record_proto protoreflect.FileDescriptor

var file_transaction_get_record_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x67, 0x65,
	0x74, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x11, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x12, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe5, 0x01,
	0x0a, 0x19, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x2a, 0x0a, 0x06, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52,
	0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x3a, 0x0a, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x44, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x44, 0x12, 0x2c, 0x0a, 0x11, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x44, 0x75,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11,
	0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x44, 0x75, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x73, 0x12, 0x32, 0x0a, 0x15, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x63, 0x68, 0x69,
	0x6c, 0x64, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x13, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x73, 0x22, 0xc7, 0x02, 0x0a, 0x1c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x46, 0x0a, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x11, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x5a, 0x0a,
	0x1b, 0x64, 0x75, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x1b, 0x64, 0x75,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x54, 0x0a, 0x19, 0x63, 0x68, 0x69,
	0x6c, 0x64, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x17, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x42,
	0x26, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x68, 0x61, 0x73,
	0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x6a, 0x61, 0x76, 0x61, 0x50, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transaction_get_record_proto_rawDescOnce sync.Once
	file_transaction_get_record_proto_rawDescData = file_transaction_get_record_proto_rawDesc
)

func file_transaction_get_record_proto_rawDescGZIP() []byte {
	file_transaction_get_record_proto_rawDescOnce.Do(func() {
		file_transaction_get_record_proto_rawDescData = protoimpl.X.CompressGZIP(file_transaction_get_record_proto_rawDescData)
	})
	return file_transaction_get_record_proto_rawDescData
}

var file_transaction_get_record_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_transaction_get_record_proto_goTypes = []interface{}{
	(*TransactionGetRecordQuery)(nil),    // 0: proto.TransactionGetRecordQuery
	(*TransactionGetRecordResponse)(nil), // 1: proto.TransactionGetRecordResponse
	(*QueryHeader)(nil),                  // 2: proto.QueryHeader
	(*TransactionID)(nil),                // 3: proto.TransactionID
	(*ResponseHeader)(nil),               // 4: proto.ResponseHeader
	(*TransactionRecord)(nil),            // 5: proto.TransactionRecord
}
var file_transaction_get_record_proto_depIdxs = []int32{
	2, // 0: proto.TransactionGetRecordQuery.header:type_name -> proto.QueryHeader
	3, // 1: proto.TransactionGetRecordQuery.transactionID:type_name -> proto.TransactionID
	4, // 2: proto.TransactionGetRecordResponse.header:type_name -> proto.ResponseHeader
	5, // 3: proto.TransactionGetRecordResponse.transactionRecord:type_name -> proto.TransactionRecord
	5, // 4: proto.TransactionGetRecordResponse.duplicateTransactionRecords:type_name -> proto.TransactionRecord
	5, // 5: proto.TransactionGetRecordResponse.child_transaction_records:type_name -> proto.TransactionRecord
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_transaction_get_record_proto_init() }
func file_transaction_get_record_proto_init() {
	if File_transaction_get_record_proto != nil {
		return
	}
	file_transaction_record_proto_init()
	file_basic_types_proto_init()
	file_query_header_proto_init()
	file_response_header_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_transaction_get_record_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionGetRecordQuery); i {
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
		file_transaction_get_record_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionGetRecordResponse); i {
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
			RawDescriptor: file_transaction_get_record_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_transaction_get_record_proto_goTypes,
		DependencyIndexes: file_transaction_get_record_proto_depIdxs,
		MessageInfos:      file_transaction_get_record_proto_msgTypes,
	}.Build()
	File_transaction_get_record_proto = out.File
	file_transaction_get_record_proto_rawDesc = nil
	file_transaction_get_record_proto_goTypes = nil
	file_transaction_get_record_proto_depIdxs = nil
}
