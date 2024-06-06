// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: protos/vote/public_vote.proto

package vote

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

type PublicVote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ElectionId string `protobuf:"bytes,2,opt,name=election_id,json=electionId,proto3" json:"election_id,omitempty"`
	PublicId   string `protobuf:"bytes,3,opt,name=public_id,json=publicId,proto3" json:"public_id,omitempty"`
	CreatedAt  string `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *PublicVote) Reset() {
	*x = PublicVote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_vote_public_vote_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicVote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicVote) ProtoMessage() {}

func (x *PublicVote) ProtoReflect() protoreflect.Message {
	mi := &file_protos_vote_public_vote_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicVote.ProtoReflect.Descriptor instead.
func (*PublicVote) Descriptor() ([]byte, []int) {
	return file_protos_vote_public_vote_proto_rawDescGZIP(), []int{0}
}

func (x *PublicVote) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PublicVote) GetElectionId() string {
	if x != nil {
		return x.ElectionId
	}
	return ""
}

func (x *PublicVote) GetPublicId() string {
	if x != nil {
		return x.PublicId
	}
	return ""
}

func (x *PublicVote) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

// Request for creating a new public vote
type PublicVoteCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ElectionId string `protobuf:"bytes,1,opt,name=election_id,json=electionId,proto3" json:"election_id,omitempty"`
	PublicId   string `protobuf:"bytes,2,opt,name=public_id,json=publicId,proto3" json:"public_id,omitempty"`
}

func (x *PublicVoteCreate) Reset() {
	*x = PublicVoteCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_vote_public_vote_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicVoteCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicVoteCreate) ProtoMessage() {}

func (x *PublicVoteCreate) ProtoReflect() protoreflect.Message {
	mi := &file_protos_vote_public_vote_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicVoteCreate.ProtoReflect.Descriptor instead.
func (*PublicVoteCreate) Descriptor() ([]byte, []int) {
	return file_protos_vote_public_vote_proto_rawDescGZIP(), []int{1}
}

func (x *PublicVoteCreate) GetElectionId() string {
	if x != nil {
		return x.ElectionId
	}
	return ""
}

func (x *PublicVoteCreate) GetPublicId() string {
	if x != nil {
		return x.PublicId
	}
	return ""
}

// Request for getting a public vote by its ID
type PublicVoteById struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PublicVoteById) Reset() {
	*x = PublicVoteById{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_vote_public_vote_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicVoteById) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicVoteById) ProtoMessage() {}

func (x *PublicVoteById) ProtoReflect() protoreflect.Message {
	mi := &file_protos_vote_public_vote_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicVoteById.ProtoReflect.Descriptor instead.
func (*PublicVoteById) Descriptor() ([]byte, []int) {
	return file_protos_vote_public_vote_proto_rawDescGZIP(), []int{2}
}

func (x *PublicVoteById) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// Request for getting all public votes
type GetAllPublicVoteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional filter for election ID
	ElectionId string `protobuf:"bytes,1,opt,name=election_id,json=electionId,proto3" json:"election_id,omitempty"`
	// Optional filter for public ID
	PublicId string `protobuf:"bytes,2,opt,name=public_id,json=publicId,proto3" json:"public_id,omitempty"`
}

func (x *GetAllPublicVoteReq) Reset() {
	*x = GetAllPublicVoteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_vote_public_vote_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllPublicVoteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllPublicVoteReq) ProtoMessage() {}

func (x *GetAllPublicVoteReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_vote_public_vote_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllPublicVoteReq.ProtoReflect.Descriptor instead.
func (*GetAllPublicVoteReq) Descriptor() ([]byte, []int) {
	return file_protos_vote_public_vote_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllPublicVoteReq) GetElectionId() string {
	if x != nil {
		return x.ElectionId
	}
	return ""
}

func (x *GetAllPublicVoteReq) GetPublicId() string {
	if x != nil {
		return x.PublicId
	}
	return ""
}

// Response for getting all public votes
type GetAllPublicVoteRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PublicVotes []*PublicVote `protobuf:"bytes,1,rep,name=public_votes,json=publicVotes,proto3" json:"public_votes,omitempty"`
	Count       int32         `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetAllPublicVoteRes) Reset() {
	*x = GetAllPublicVoteRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_vote_public_vote_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllPublicVoteRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllPublicVoteRes) ProtoMessage() {}

func (x *GetAllPublicVoteRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_vote_public_vote_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllPublicVoteRes.ProtoReflect.Descriptor instead.
func (*GetAllPublicVoteRes) Descriptor() ([]byte, []int) {
	return file_protos_vote_public_vote_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllPublicVoteRes) GetPublicVotes() []*PublicVote {
	if x != nil {
		return x.PublicVotes
	}
	return nil
}

func (x *GetAllPublicVoteRes) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_protos_vote_public_vote_proto protoreflect.FileDescriptor

var file_protos_vote_public_vote_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x6f, 0x74, 0x65, 0x2f, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x5f, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x22, 0x79, 0x0a, 0x0a, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0x50, 0x0a, 0x10, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x56, 0x6f, 0x74, 0x65,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x49, 0x64, 0x22, 0x20, 0x0a, 0x0e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x56, 0x6f,
	0x74, 0x65, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x53, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a,
	0x0b, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x64, 0x22, 0x62, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x56, 0x6f, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x12, 0x35, 0x0a, 0x0c, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x76, 0x6f, 0x74,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x0b, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32,
	0xcc, 0x01, 0x0a, 0x11, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x56, 0x6f, 0x74, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x56,
	0x6f, 0x74, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x56, 0x6f, 0x74, 0x65, 0x22, 0x00, 0x12,
	0x37, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x56, 0x6f, 0x74, 0x65, 0x42, 0x79,
	0x49, 0x64, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x56, 0x6f, 0x74, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a,
	0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x0c,
	0x5a, 0x0a, 0x2f, 0x76, 0x6f, 0x74, 0x65, 0x3b, 0x76, 0x6f, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_vote_public_vote_proto_rawDescOnce sync.Once
	file_protos_vote_public_vote_proto_rawDescData = file_protos_vote_public_vote_proto_rawDesc
)

func file_protos_vote_public_vote_proto_rawDescGZIP() []byte {
	file_protos_vote_public_vote_proto_rawDescOnce.Do(func() {
		file_protos_vote_public_vote_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_vote_public_vote_proto_rawDescData)
	})
	return file_protos_vote_public_vote_proto_rawDescData
}

var file_protos_vote_public_vote_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protos_vote_public_vote_proto_goTypes = []interface{}{
	(*PublicVote)(nil),          // 0: protos.PublicVote
	(*PublicVoteCreate)(nil),    // 1: protos.PublicVoteCreate
	(*PublicVoteById)(nil),      // 2: protos.PublicVoteById
	(*GetAllPublicVoteReq)(nil), // 3: protos.GetAllPublicVoteReq
	(*GetAllPublicVoteRes)(nil), // 4: protos.GetAllPublicVoteRes
}
var file_protos_vote_public_vote_proto_depIdxs = []int32{
	0, // 0: protos.GetAllPublicVoteRes.public_votes:type_name -> protos.PublicVote
	1, // 1: protos.PublicVoteService.Create:input_type -> protos.PublicVoteCreate
	2, // 2: protos.PublicVoteService.GetById:input_type -> protos.PublicVoteById
	3, // 3: protos.PublicVoteService.GetAll:input_type -> protos.GetAllPublicVoteReq
	0, // 4: protos.PublicVoteService.Create:output_type -> protos.PublicVote
	0, // 5: protos.PublicVoteService.GetById:output_type -> protos.PublicVote
	4, // 6: protos.PublicVoteService.GetAll:output_type -> protos.GetAllPublicVoteRes
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_vote_public_vote_proto_init() }
func file_protos_vote_public_vote_proto_init() {
	if File_protos_vote_public_vote_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_vote_public_vote_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicVote); i {
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
		file_protos_vote_public_vote_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicVoteCreate); i {
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
		file_protos_vote_public_vote_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicVoteById); i {
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
		file_protos_vote_public_vote_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllPublicVoteReq); i {
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
		file_protos_vote_public_vote_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllPublicVoteRes); i {
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
			RawDescriptor: file_protos_vote_public_vote_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_vote_public_vote_proto_goTypes,
		DependencyIndexes: file_protos_vote_public_vote_proto_depIdxs,
		MessageInfos:      file_protos_vote_public_vote_proto_msgTypes,
	}.Build()
	File_protos_vote_public_vote_proto = out.File
	file_protos_vote_public_vote_proto_rawDesc = nil
	file_protos_vote_public_vote_proto_goTypes = nil
	file_protos_vote_public_vote_proto_depIdxs = nil
}