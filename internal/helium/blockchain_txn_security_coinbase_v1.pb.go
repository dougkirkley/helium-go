// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: blockchain_txn_security_coinbase_v1.proto

package helium

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

type BlockchainTxnSecurityCoinbaseV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payee  []byte `protobuf:"bytes,1,opt,name=payee,proto3" json:"payee,omitempty"`
	Amount uint64 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *BlockchainTxnSecurityCoinbaseV1) Reset() {
	*x = BlockchainTxnSecurityCoinbaseV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_txn_security_coinbase_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockchainTxnSecurityCoinbaseV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockchainTxnSecurityCoinbaseV1) ProtoMessage() {}

func (x *BlockchainTxnSecurityCoinbaseV1) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_txn_security_coinbase_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockchainTxnSecurityCoinbaseV1.ProtoReflect.Descriptor instead.
func (*BlockchainTxnSecurityCoinbaseV1) Descriptor() ([]byte, []int) {
	return file_blockchain_txn_security_coinbase_v1_proto_rawDescGZIP(), []int{0}
}

func (x *BlockchainTxnSecurityCoinbaseV1) GetPayee() []byte {
	if x != nil {
		return x.Payee
	}
	return nil
}

func (x *BlockchainTxnSecurityCoinbaseV1) GetAmount() uint64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

var File_blockchain_txn_security_coinbase_v1_proto protoreflect.FileDescriptor

var file_blockchain_txn_security_coinbase_v1_proto_rawDesc = []byte{
	0x0a, 0x29, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x78, 0x6e,
	0x5f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61,
	0x73, 0x65, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x68, 0x65, 0x6c,
	0x69, 0x75, 0x6d, 0x22, 0x53, 0x0a, 0x23, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x5f, 0x74, 0x78, 0x6e, 0x5f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x63,
	0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x76, 0x31, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61,
	0x79, 0x65, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x70, 0x61, 0x79, 0x65, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x75, 0x67, 0x6b, 0x69, 0x72, 0x6b, 0x6c,
	0x65, 0x79, 0x2f, 0x68, 0x65, 0x6c, 0x69, 0x75, 0x6d, 0x2d, 0x67, 0x6f, 0x2f, 0x68, 0x65, 0x6c,
	0x69, 0x75, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blockchain_txn_security_coinbase_v1_proto_rawDescOnce sync.Once
	file_blockchain_txn_security_coinbase_v1_proto_rawDescData = file_blockchain_txn_security_coinbase_v1_proto_rawDesc
)

func file_blockchain_txn_security_coinbase_v1_proto_rawDescGZIP() []byte {
	file_blockchain_txn_security_coinbase_v1_proto_rawDescOnce.Do(func() {
		file_blockchain_txn_security_coinbase_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_blockchain_txn_security_coinbase_v1_proto_rawDescData)
	})
	return file_blockchain_txn_security_coinbase_v1_proto_rawDescData
}

var file_blockchain_txn_security_coinbase_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_blockchain_txn_security_coinbase_v1_proto_goTypes = []interface{}{
	(*BlockchainTxnSecurityCoinbaseV1)(nil), // 0: helium.blockchain_txn_security_coinbase_v1
}
var file_blockchain_txn_security_coinbase_v1_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_blockchain_txn_security_coinbase_v1_proto_init() }
func file_blockchain_txn_security_coinbase_v1_proto_init() {
	if File_blockchain_txn_security_coinbase_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_blockchain_txn_security_coinbase_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockchainTxnSecurityCoinbaseV1); i {
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
			RawDescriptor: file_blockchain_txn_security_coinbase_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_blockchain_txn_security_coinbase_v1_proto_goTypes,
		DependencyIndexes: file_blockchain_txn_security_coinbase_v1_proto_depIdxs,
		MessageInfos:      file_blockchain_txn_security_coinbase_v1_proto_msgTypes,
	}.Build()
	File_blockchain_txn_security_coinbase_v1_proto = out.File
	file_blockchain_txn_security_coinbase_v1_proto_rawDesc = nil
	file_blockchain_txn_security_coinbase_v1_proto_goTypes = nil
	file_blockchain_txn_security_coinbase_v1_proto_depIdxs = nil
}
