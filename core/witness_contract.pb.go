// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.21.12
// source: protocol/core/contract/witness_contract.proto

package core

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

type WitnessCreateContract struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerAddress []byte `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	Url          []byte `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *WitnessCreateContract) Reset() {
	*x = WitnessCreateContract{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_core_contract_witness_contract_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WitnessCreateContract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WitnessCreateContract) ProtoMessage() {}

func (x *WitnessCreateContract) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_core_contract_witness_contract_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WitnessCreateContract.ProtoReflect.Descriptor instead.
func (*WitnessCreateContract) Descriptor() ([]byte, []int) {
	return file_protocol_core_contract_witness_contract_proto_rawDescGZIP(), []int{0}
}

func (x *WitnessCreateContract) GetOwnerAddress() []byte {
	if x != nil {
		return x.OwnerAddress
	}
	return nil
}

func (x *WitnessCreateContract) GetUrl() []byte {
	if x != nil {
		return x.Url
	}
	return nil
}

type WitnessUpdateContract struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerAddress []byte `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	UpdateUrl    []byte `protobuf:"bytes,12,opt,name=update_url,json=updateUrl,proto3" json:"update_url,omitempty"`
}

func (x *WitnessUpdateContract) Reset() {
	*x = WitnessUpdateContract{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_core_contract_witness_contract_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WitnessUpdateContract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WitnessUpdateContract) ProtoMessage() {}

func (x *WitnessUpdateContract) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_core_contract_witness_contract_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WitnessUpdateContract.ProtoReflect.Descriptor instead.
func (*WitnessUpdateContract) Descriptor() ([]byte, []int) {
	return file_protocol_core_contract_witness_contract_proto_rawDescGZIP(), []int{1}
}

func (x *WitnessUpdateContract) GetOwnerAddress() []byte {
	if x != nil {
		return x.OwnerAddress
	}
	return nil
}

func (x *WitnessUpdateContract) GetUpdateUrl() []byte {
	if x != nil {
		return x.UpdateUrl
	}
	return nil
}

type VoteWitnessContract struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerAddress []byte                      `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	Votes        []*VoteWitnessContract_Vote `protobuf:"bytes,2,rep,name=votes,proto3" json:"votes,omitempty"`
	Support      bool                        `protobuf:"varint,3,opt,name=support,proto3" json:"support,omitempty"`
}

func (x *VoteWitnessContract) Reset() {
	*x = VoteWitnessContract{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_core_contract_witness_contract_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoteWitnessContract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoteWitnessContract) ProtoMessage() {}

func (x *VoteWitnessContract) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_core_contract_witness_contract_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoteWitnessContract.ProtoReflect.Descriptor instead.
func (*VoteWitnessContract) Descriptor() ([]byte, []int) {
	return file_protocol_core_contract_witness_contract_proto_rawDescGZIP(), []int{2}
}

func (x *VoteWitnessContract) GetOwnerAddress() []byte {
	if x != nil {
		return x.OwnerAddress
	}
	return nil
}

func (x *VoteWitnessContract) GetVotes() []*VoteWitnessContract_Vote {
	if x != nil {
		return x.Votes
	}
	return nil
}

func (x *VoteWitnessContract) GetSupport() bool {
	if x != nil {
		return x.Support
	}
	return false
}

type VoteWitnessContract_Vote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VoteAddress []byte `protobuf:"bytes,1,opt,name=vote_address,json=voteAddress,proto3" json:"vote_address,omitempty"`
	VoteCount   int64  `protobuf:"varint,2,opt,name=vote_count,json=voteCount,proto3" json:"vote_count,omitempty"`
}

func (x *VoteWitnessContract_Vote) Reset() {
	*x = VoteWitnessContract_Vote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_core_contract_witness_contract_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoteWitnessContract_Vote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoteWitnessContract_Vote) ProtoMessage() {}

func (x *VoteWitnessContract_Vote) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_core_contract_witness_contract_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoteWitnessContract_Vote.ProtoReflect.Descriptor instead.
func (*VoteWitnessContract_Vote) Descriptor() ([]byte, []int) {
	return file_protocol_core_contract_witness_contract_proto_rawDescGZIP(), []int{2, 0}
}

func (x *VoteWitnessContract_Vote) GetVoteAddress() []byte {
	if x != nil {
		return x.VoteAddress
	}
	return nil
}

func (x *VoteWitnessContract_Vote) GetVoteCount() int64 {
	if x != nil {
		return x.VoteCount
	}
	return 0
}

var File_protocol_core_contract_witness_contract_proto protoreflect.FileDescriptor

var file_protocol_core_contract_witness_contract_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2f, 0x77, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73,
	0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22, 0x4e, 0x0a, 0x15, 0x57, 0x69, 0x74,
	0x6e, 0x65, 0x73, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x5b, 0x0a, 0x15, 0x57, 0x69, 0x74,
	0x6e, 0x65, 0x73, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x55, 0x72, 0x6c, 0x22, 0xd8, 0x01, 0x0a, 0x13, 0x56, 0x6f, 0x74, 0x65, 0x57,
	0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x23,
	0x0a, 0x0d, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x38, 0x0a, 0x05, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x56, 0x6f,
	0x74, 0x65, 0x57, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x2e, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x05, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x1a, 0x48, 0x0a, 0x04, 0x56, 0x6f, 0x74, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x76, 0x6f, 0x74, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x76, 0x6f, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x6f, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x76, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x42, 0x3b, 0x0a, 0x18, 0x6f, 0x72, 0x67, 0x2e, 0x74, 0x72, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5a, 0x1f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x31, 0x6d, 0x70, 0x69, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2d, 0x74, 0x72, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protocol_core_contract_witness_contract_proto_rawDescOnce sync.Once
	file_protocol_core_contract_witness_contract_proto_rawDescData = file_protocol_core_contract_witness_contract_proto_rawDesc
)

func file_protocol_core_contract_witness_contract_proto_rawDescGZIP() []byte {
	file_protocol_core_contract_witness_contract_proto_rawDescOnce.Do(func() {
		file_protocol_core_contract_witness_contract_proto_rawDescData = protoimpl.X.CompressGZIP(file_protocol_core_contract_witness_contract_proto_rawDescData)
	})
	return file_protocol_core_contract_witness_contract_proto_rawDescData
}

var file_protocol_core_contract_witness_contract_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protocol_core_contract_witness_contract_proto_goTypes = []interface{}{
	(*WitnessCreateContract)(nil),    // 0: protocol.WitnessCreateContract
	(*WitnessUpdateContract)(nil),    // 1: protocol.WitnessUpdateContract
	(*VoteWitnessContract)(nil),      // 2: protocol.VoteWitnessContract
	(*VoteWitnessContract_Vote)(nil), // 3: protocol.VoteWitnessContract.Vote
}
var file_protocol_core_contract_witness_contract_proto_depIdxs = []int32{
	3, // 0: protocol.VoteWitnessContract.votes:type_name -> protocol.VoteWitnessContract.Vote
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protocol_core_contract_witness_contract_proto_init() }
func file_protocol_core_contract_witness_contract_proto_init() {
	if File_protocol_core_contract_witness_contract_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protocol_core_contract_witness_contract_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WitnessCreateContract); i {
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
		file_protocol_core_contract_witness_contract_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WitnessUpdateContract); i {
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
		file_protocol_core_contract_witness_contract_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoteWitnessContract); i {
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
		file_protocol_core_contract_witness_contract_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoteWitnessContract_Vote); i {
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
			RawDescriptor: file_protocol_core_contract_witness_contract_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protocol_core_contract_witness_contract_proto_goTypes,
		DependencyIndexes: file_protocol_core_contract_witness_contract_proto_depIdxs,
		MessageInfos:      file_protocol_core_contract_witness_contract_proto_msgTypes,
	}.Build()
	File_protocol_core_contract_witness_contract_proto = out.File
	file_protocol_core_contract_witness_contract_proto_rawDesc = nil
	file_protocol_core_contract_witness_contract_proto_goTypes = nil
	file_protocol_core_contract_witness_contract_proto_depIdxs = nil
}
