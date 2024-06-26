// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: pb/admin/admin.proto

package admin

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

type LoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginReq) Reset() {
	*x = LoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_admin_admin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReq) ProtoMessage() {}

func (x *LoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_admin_admin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReq.ProtoReflect.Descriptor instead.
func (*LoginReq) Descriptor() ([]byte, []int) {
	return file_pb_admin_admin_proto_rawDescGZIP(), []int{0}
}

func (x *LoginReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *LoginReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error    string `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"`
	Response string `protobuf:"bytes,3,opt,name=Response,proto3" json:"Response,omitempty"`
	Token    string `protobuf:"bytes,4,opt,name=Token,proto3" json:"Token,omitempty"`
}

func (x *LoginRes) Reset() {
	*x = LoginRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_admin_admin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRes) ProtoMessage() {}

func (x *LoginRes) ProtoReflect() protoreflect.Message {
	mi := &file_pb_admin_admin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRes.ProtoReflect.Descriptor instead.
func (*LoginRes) Descriptor() ([]byte, []int) {
	return file_pb_admin_admin_proto_rawDescGZIP(), []int{1}
}

func (x *LoginRes) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *LoginRes) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *LoginRes) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

func (x *LoginRes) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type BlockReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *BlockReq) Reset() {
	*x = BlockReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_admin_admin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockReq) ProtoMessage() {}

func (x *BlockReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_admin_admin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockReq.ProtoReflect.Descriptor instead.
func (*BlockReq) Descriptor() ([]byte, []int) {
	return file_pb_admin_admin_proto_rawDescGZIP(), []int{2}
}

func (x *BlockReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type BlockRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error    string `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"`
	Response string `protobuf:"bytes,3,opt,name=Response,proto3" json:"Response,omitempty"`
}

func (x *BlockRes) Reset() {
	*x = BlockRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_admin_admin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockRes) ProtoMessage() {}

func (x *BlockRes) ProtoReflect() protoreflect.Message {
	mi := &file_pb_admin_admin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockRes.ProtoReflect.Descriptor instead.
func (*BlockRes) Descriptor() ([]byte, []int) {
	return file_pb_admin_admin_proto_rawDescGZIP(), []int{3}
}

func (x *BlockRes) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *BlockRes) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *BlockRes) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type AddSkillReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Skill       string `protobuf:"bytes,1,opt,name=skill,proto3" json:"skill,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *AddSkillReq) Reset() {
	*x = AddSkillReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_admin_admin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddSkillReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddSkillReq) ProtoMessage() {}

func (x *AddSkillReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_admin_admin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddSkillReq.ProtoReflect.Descriptor instead.
func (*AddSkillReq) Descriptor() ([]byte, []int) {
	return file_pb_admin_admin_proto_rawDescGZIP(), []int{4}
}

func (x *AddSkillReq) GetSkill() string {
	if x != nil {
		return x.Skill
	}
	return ""
}

func (x *AddSkillReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type AddSkillRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error    string `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"`
	Response string `protobuf:"bytes,3,opt,name=Response,proto3" json:"Response,omitempty"`
}

func (x *AddSkillRes) Reset() {
	*x = AddSkillRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_admin_admin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddSkillRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddSkillRes) ProtoMessage() {}

func (x *AddSkillRes) ProtoReflect() protoreflect.Message {
	mi := &file_pb_admin_admin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddSkillRes.ProtoReflect.Descriptor instead.
func (*AddSkillRes) Descriptor() ([]byte, []int) {
	return file_pb_admin_admin_proto_rawDescGZIP(), []int{5}
}

func (x *AddSkillRes) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *AddSkillRes) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *AddSkillRes) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

var File_pb_admin_admin_proto protoreflect.FileDescriptor

var file_pb_admin_admin_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x62, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x22, 0x3c, 0x0a,
	0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x6a, 0x0a, 0x08, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x23, 0x0a, 0x08, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x54, 0x0a, 0x08,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x45, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x52, 0x65,
	0x71, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x57, 0x0a, 0x0b, 0x41, 0x64, 0x64,
	0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x32, 0xda, 0x01, 0x0a, 0x0c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x0a, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x12, 0x0f, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x09, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x0f, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x0b, 0x55, 0x6e, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0f, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x08, 0x41, 0x64, 0x64,
	0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x12, 0x12, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x41, 0x64,
	0x64, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x41, 0x64, 0x64, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42,
	0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x70, 0x62, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_admin_admin_proto_rawDescOnce sync.Once
	file_pb_admin_admin_proto_rawDescData = file_pb_admin_admin_proto_rawDesc
)

func file_pb_admin_admin_proto_rawDescGZIP() []byte {
	file_pb_admin_admin_proto_rawDescOnce.Do(func() {
		file_pb_admin_admin_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_admin_admin_proto_rawDescData)
	})
	return file_pb_admin_admin_proto_rawDescData
}

var file_pb_admin_admin_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pb_admin_admin_proto_goTypes = []interface{}{
	(*LoginReq)(nil),    // 0: admin.LoginReq
	(*LoginRes)(nil),    // 1: admin.LoginRes
	(*BlockReq)(nil),    // 2: admin.BlockReq
	(*BlockRes)(nil),    // 3: admin.BlockRes
	(*AddSkillReq)(nil), // 4: admin.AddSkillReq
	(*AddSkillRes)(nil), // 5: admin.AddSkillRes
}
var file_pb_admin_admin_proto_depIdxs = []int32{
	0, // 0: admin.AdminService.AdminLogin:input_type -> admin.LoginReq
	2, // 1: admin.AdminService.BlockUser:input_type -> admin.BlockReq
	2, // 2: admin.AdminService.UnBlockUser:input_type -> admin.BlockReq
	4, // 3: admin.AdminService.AddSkill:input_type -> admin.AddSkillReq
	1, // 4: admin.AdminService.AdminLogin:output_type -> admin.LoginRes
	3, // 5: admin.AdminService.BlockUser:output_type -> admin.BlockRes
	3, // 6: admin.AdminService.UnBlockUser:output_type -> admin.BlockRes
	5, // 7: admin.AdminService.AddSkill:output_type -> admin.AddSkillRes
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_admin_admin_proto_init() }
func file_pb_admin_admin_proto_init() {
	if File_pb_admin_admin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_admin_admin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReq); i {
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
		file_pb_admin_admin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRes); i {
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
		file_pb_admin_admin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockReq); i {
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
		file_pb_admin_admin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockRes); i {
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
		file_pb_admin_admin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddSkillReq); i {
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
		file_pb_admin_admin_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddSkillRes); i {
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
			RawDescriptor: file_pb_admin_admin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_admin_admin_proto_goTypes,
		DependencyIndexes: file_pb_admin_admin_proto_depIdxs,
		MessageInfos:      file_pb_admin_admin_proto_msgTypes,
	}.Build()
	File_pb_admin_admin_proto = out.File
	file_pb_admin_admin_proto_rawDesc = nil
	file_pb_admin_admin_proto_goTypes = nil
	file_pb_admin_admin_proto_depIdxs = nil
}
