
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: pb/user/job.proto

package user

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

type CReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *CReq) Reset() {
	*x = CReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_user_job_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CReq) ProtoMessage() {}

func (x *CReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_user_job_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CReq.ProtoReflect.Descriptor instead.
func (*CReq) Descriptor() ([]byte, []int) {
	return file_pb_user_job_proto_rawDescGZIP(), []int{0}
}

func (x *CReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type CRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32  `protobuf:"varint,1,opt,name=Status,proto3" json:"Status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"`
	Exist  bool   `protobuf:"varint,3,opt,name=Exist,proto3" json:"Exist,omitempty"`
}

func (x *CRes) Reset() {
	*x = CRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_user_job_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CRes) ProtoMessage() {}

func (x *CRes) ProtoReflect() protoreflect.Message {
	mi := &file_pb_user_job_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CRes.ProtoReflect.Descriptor instead.
func (*CRes) Descriptor() ([]byte, []int) {
	return file_pb_user_job_proto_rawDescGZIP(), []int{1}
}

func (x *CRes) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CRes) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *CRes) GetExist() bool {
	if x != nil {
		return x.Exist
	}
	return false
}

type Preq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID int32 `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *Preq) Reset() {
	*x = Preq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_user_job_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Preq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Preq) ProtoMessage() {}

func (x *Preq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_user_job_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Preq.ProtoReflect.Descriptor instead.
func (*Preq) Descriptor() ([]byte, []int) {
	return file_pb_user_job_proto_rawDescGZIP(), []int{2}
}

func (x *Preq) GetUserID() int32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

type Pres struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email  string `protobuf:"bytes,1,opt,name=Email,proto3" json:"Email,omitempty"`
	Status int32  `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
}

func (x *Pres) Reset() {
	*x = Pres{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_user_job_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pres) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pres) ProtoMessage() {}

func (x *Pres) ProtoReflect() protoreflect.Message {
	mi := &file_pb_user_job_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pres.ProtoReflect.Descriptor instead.
func (*Pres) Descriptor() ([]byte, []int) {
	return file_pb_user_job_proto_rawDescGZIP(), []int{3}
}

func (x *Pres) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Pres) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Pres) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type Req struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Skill []int64 `protobuf:"varint,1,rep,packed,name=skill,proto3" json:"skill,omitempty"`
}

func (x *Req) Reset() {
	*x = Req{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_user_job_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Req) ProtoMessage() {}

func (x *Req) ProtoReflect() protoreflect.Message {
	mi := &file_pb_user_job_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Req.ProtoReflect.Descriptor instead.
func (*Req) Descriptor() ([]byte, []int) {
	return file_pb_user_job_proto_rawDescGZIP(), []int{4}
}

func (x *Req) GetSkill() []int64 {
	if x != nil {
		return x.Skill
	}
	return nil
}

type Res struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Skill []string `protobuf:"bytes,1,rep,name=skill,proto3" json:"skill,omitempty"`
}

func (x *Res) Reset() {
	*x = Res{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_user_job_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Res) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Res) ProtoMessage() {}

func (x *Res) ProtoReflect() protoreflect.Message {
	mi := &file_pb_user_job_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Res.ProtoReflect.Descriptor instead.
func (*Res) Descriptor() ([]byte, []int) {
	return file_pb_user_job_proto_rawDescGZIP(), []int{5}
}

func (x *Res) GetSkill() []string {
	if x != nil {
		return x.Skill
	}
	return nil
}

var File_pb_user_job_proto protoreflect.FileDescriptor

var file_pb_user_job_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x62, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6a, 0x6f, 0x62, 0x22, 0x1f, 0x0a, 0x04, 0x43, 0x52, 0x65, 0x71,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x4a, 0x0a, 0x04, 0x43, 0x52, 0x65,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x45, 0x78, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05,
	0x45, 0x78, 0x69, 0x73, 0x74, 0x22, 0x1e, 0x0a, 0x04, 0x50, 0x72, 0x65, 0x71, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x4a, 0x0a, 0x04, 0x50, 0x72, 0x65, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x22, 0x1b, 0x0a, 0x03, 0x72, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x6b, 0x69, 0x6c,
	0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x05, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x22, 0x1b,
	0x0a, 0x03, 0x72, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x32, 0x98, 0x01, 0x0a, 0x0a,
	0x4a, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x0d, 0x47, 0x65,
	0x74, 0x4a, 0x6f, 0x62, 0x73, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x12, 0x08, 0x2e, 0x6a, 0x6f,
	0x62, 0x2e, 0x72, 0x65, 0x71, 0x1a, 0x08, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x22,
	0x00, 0x12, 0x32, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x46, 0x72, 0x65, 0x65, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x72, 0x50, 0x61, 0x79, 0x70, 0x61, 0x6c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x09, 0x2e,
	0x6a, 0x6f, 0x62, 0x2e, 0x50, 0x72, 0x65, 0x71, 0x1a, 0x09, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x50,
	0x72, 0x65, 0x73, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x15, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x61,
	0x79, 0x70, 0x61, 0x6c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x65, 0x64, 0x12, 0x09,
	0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x43, 0x52, 0x65, 0x71, 0x1a, 0x09, 0x2e, 0x6a, 0x6f, 0x62, 0x2e,
	0x43, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x70, 0x62, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_user_job_proto_rawDescOnce sync.Once
	file_pb_user_job_proto_rawDescData = file_pb_user_job_proto_rawDesc
)

func file_pb_user_job_proto_rawDescGZIP() []byte {
	file_pb_user_job_proto_rawDescOnce.Do(func() {
		file_pb_user_job_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_user_job_proto_rawDescData)
	})
	return file_pb_user_job_proto_rawDescData
}

var file_pb_user_job_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pb_user_job_proto_goTypes = []interface{}{
	(*CReq)(nil), // 0: job.CReq
	(*CRes)(nil), // 1: job.CRes
	(*Preq)(nil), // 2: job.Preq
	(*Pres)(nil), // 3: job.Pres
	(*Req)(nil),  // 4: job.req
	(*Res)(nil),  // 5: job.res
}
var file_pb_user_job_proto_depIdxs = []int32{
	4, // 0: job.Jobservice.GetJobsSkills:input_type -> job.req
	2, // 1: job.Jobservice.GetFreelancerPaypalEmail:input_type -> job.Preq
	0, // 2: job.Jobservice.CheckPaypalEmailAdded:input_type -> job.CReq
	5, // 3: job.Jobservice.GetJobsSkills:output_type -> job.res
	3, // 4: job.Jobservice.GetFreelancerPaypalEmail:output_type -> job.Pres
	1, // 5: job.Jobservice.CheckPaypalEmailAdded:output_type -> job.CRes
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_user_job_proto_init() }
func file_pb_user_job_proto_init() {
	if File_pb_user_job_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_user_job_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CReq); i {
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
		file_pb_user_job_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CRes); i {
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
		file_pb_user_job_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Preq); i {
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
		file_pb_user_job_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pres); i {
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
		file_pb_user_job_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Req); i {
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
		file_pb_user_job_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Res); i {
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
			RawDescriptor: file_pb_user_job_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_user_job_proto_goTypes,
		DependencyIndexes: file_pb_user_job_proto_depIdxs,
		MessageInfos:      file_pb_user_job_proto_msgTypes,
	}.Build()
	File_pb_user_job_proto = out.File
	file_pb_user_job_proto_rawDesc = nil
	file_pb_user_job_proto_goTypes = nil
	file_pb_user_job_proto_depIdxs = nil
}
