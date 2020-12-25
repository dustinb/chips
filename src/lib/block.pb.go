// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.11.4
// source: protos/block.proto

package chips

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

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HashBlock      []byte         `protobuf:"bytes,1,opt,name=hashBlock,proto3" json:"hashBlock,omitempty"`
	HashPrevBlock  []byte         `protobuf:"bytes,2,opt,name=hashPrevBlock,proto3" json:"hashPrevBlock,omitempty"`
	HashMerkleRoot []byte         `protobuf:"bytes,3,opt,name=hashMerkleRoot,proto3" json:"hashMerkleRoot,omitempty"`
	Difficulty     uint32         `protobuf:"varint,4,opt,name=difficulty,proto3" json:"difficulty,omitempty"`
	Transactions   []*Transaction `protobuf:"bytes,5,rep,name=transactions,proto3" json:"transactions,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_block_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_protos_block_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_protos_block_proto_rawDescGZIP(), []int{0}
}

func (x *Block) GetHashBlock() []byte {
	if x != nil {
		return x.HashBlock
	}
	return nil
}

func (x *Block) GetHashPrevBlock() []byte {
	if x != nil {
		return x.HashPrevBlock
	}
	return nil
}

func (x *Block) GetHashMerkleRoot() []byte {
	if x != nil {
		return x.HashMerkleRoot
	}
	return nil
}

func (x *Block) GetDifficulty() uint32 {
	if x != nil {
		return x.Difficulty
	}
	return 0
}

func (x *Block) GetTransactions() []*Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

type Input struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxnHash   []byte `protobuf:"bytes,1,opt,name=txnHash,proto3" json:"txnHash,omitempty"`
	Index     int32  `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	Signature []byte `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *Input) Reset() {
	*x = Input{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_block_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Input) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Input) ProtoMessage() {}

func (x *Input) ProtoReflect() protoreflect.Message {
	mi := &file_protos_block_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Input.ProtoReflect.Descriptor instead.
func (*Input) Descriptor() ([]byte, []int) {
	return file_protos_block_proto_rawDescGZIP(), []int{1}
}

func (x *Input) GetTxnHash() []byte {
	if x != nil {
		return x.TxnHash
	}
	return nil
}

func (x *Input) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Input) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

//*
//Outputs are either unspent UTXO or spent TXO
type Output struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"` // pay to pub key encumbrance
	Amount  uint64 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`  // value in coin
	Spent   bool   `protobuf:"varint,4,opt,name=spent,proto3" json:"spent,omitempty"`
}

func (x *Output) Reset() {
	*x = Output{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_block_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Output) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Output) ProtoMessage() {}

func (x *Output) ProtoReflect() protoreflect.Message {
	mi := &file_protos_block_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Output.ProtoReflect.Descriptor instead.
func (*Output) Descriptor() ([]byte, []int) {
	return file_protos_block_proto_rawDescGZIP(), []int{2}
}

func (x *Output) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Output) GetAmount() uint64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Output) GetSpent() bool {
	if x != nil {
		return x.Spent
	}
	return false
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version int32     `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Hash    []byte    `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	Inputs  []*Input  `protobuf:"bytes,3,rep,name=inputs,proto3" json:"inputs,omitempty"`
	Outputs []*Output `protobuf:"bytes,4,rep,name=outputs,proto3" json:"outputs,omitempty"`
	Message string    `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_block_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_protos_block_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_protos_block_proto_rawDescGZIP(), []int{3}
}

func (x *Transaction) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Transaction) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *Transaction) GetInputs() []*Input {
	if x != nil {
		return x.Inputs
	}
	return nil
}

func (x *Transaction) GetOutputs() []*Output {
	if x != nil {
		return x.Outputs
	}
	return nil
}

func (x *Transaction) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_protos_block_proto protoreflect.FileDescriptor

var file_protos_block_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0xcb, 0x01, 0x0a, 0x05,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x68, 0x61, 0x73, 0x68, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x68, 0x61, 0x73, 0x68, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x12, 0x24, 0x0a, 0x0d, 0x68, 0x61, 0x73, 0x68, 0x50, 0x72, 0x65, 0x76, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x68, 0x61, 0x73, 0x68,
	0x50, 0x72, 0x65, 0x76, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x26, 0x0a, 0x0e, 0x68, 0x61, 0x73,
	0x68, 0x4d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0e, 0x68, 0x61, 0x73, 0x68, 0x4d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x52, 0x6f, 0x6f,
	0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74,
	0x79, 0x12, 0x36, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x55, 0x0a, 0x05, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x78, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x74, 0x78, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x12, 0x14, 0x0a, 0x05,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x22, 0x50, 0x0a, 0x06, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x70, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x73, 0x70, 0x65,
	0x6e, 0x74, 0x22, 0xa4, 0x01, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68,
	0x12, 0x24, 0x0a, 0x06, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x06,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x12, 0x27, 0x0a, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x63,
	0x68, 0x69, 0x70, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_block_proto_rawDescOnce sync.Once
	file_protos_block_proto_rawDescData = file_protos_block_proto_rawDesc
)

func file_protos_block_proto_rawDescGZIP() []byte {
	file_protos_block_proto_rawDescOnce.Do(func() {
		file_protos_block_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_block_proto_rawDescData)
	})
	return file_protos_block_proto_rawDescData
}

var file_protos_block_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protos_block_proto_goTypes = []interface{}{
	(*Block)(nil),       // 0: block.Block
	(*Input)(nil),       // 1: block.Input
	(*Output)(nil),      // 2: block.Output
	(*Transaction)(nil), // 3: block.Transaction
}
var file_protos_block_proto_depIdxs = []int32{
	3, // 0: block.Block.transactions:type_name -> block.Transaction
	1, // 1: block.Transaction.inputs:type_name -> block.Input
	2, // 2: block.Transaction.outputs:type_name -> block.Output
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_protos_block_proto_init() }
func file_protos_block_proto_init() {
	if File_protos_block_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_block_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
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
		file_protos_block_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Input); i {
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
		file_protos_block_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Output); i {
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
		file_protos_block_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
			RawDescriptor: file_protos_block_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_block_proto_goTypes,
		DependencyIndexes: file_protos_block_proto_depIdxs,
		MessageInfos:      file_protos_block_proto_msgTypes,
	}.Build()
	File_protos_block_proto = out.File
	file_protos_block_proto_rawDesc = nil
	file_protos_block_proto_goTypes = nil
	file_protos_block_proto_depIdxs = nil
}
