// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: pb/user/user.proto

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

type Preq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID int32 `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *Preq) Reset() {
	*x = Preq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_user_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Preq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Preq) ProtoMessage() {}

func (x *Preq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_user_user_proto_msgTypes[0]
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
	return file_pb_user_user_proto_rawDescGZIP(), []int{0}
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
		mi := &file_pb_user_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pres) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pres) ProtoMessage() {}

func (x *Pres) ProtoReflect() protoreflect.Message {
	mi := &file_pb_user_user_proto_msgTypes[1]
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
	return file_pb_user_user_proto_rawDescGZIP(), []int{1}
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

var File_pb_user_user_proto protoreflect.FileDescriptor

var file_pb_user_user_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x62, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x1e, 0x0a, 0x04, 0x50, 0x72,
	0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x4a, 0x0a, 0x04, 0x50, 0x72,
	0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x44, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x46, 0x72, 0x65, 0x65,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x50, 0x61, 0x79, 0x70, 0x61, 0x6c, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x73, 0x12, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x65, 0x71, 0x1a, 0x0a,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x65, 0x73, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09,
	0x2e, 0x2f, 0x70, 0x62, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_pb_user_user_proto_rawDescOnce sync.Once
	file_pb_user_user_proto_rawDescData = file_pb_user_user_proto_rawDesc
)

func file_pb_user_user_proto_rawDescGZIP() []byte {
	file_pb_user_user_proto_rawDescOnce.Do(func() {
		file_pb_user_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_user_user_proto_rawDescData)
	})
	return file_pb_user_user_proto_rawDescData
}

var file_pb_user_user_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pb_user_user_proto_goTypes = []interface{}{
	(*Preq)(nil), // 0: user.Preq
	(*Pres)(nil), // 1: user.Pres
}
var file_pb_user_user_proto_depIdxs = []int32{
	0, // 0: user.UserService.GetFreelancerPaypalEmails:input_type -> user.Preq
	1, // 1: user.UserService.GetFreelancerPaypalEmails:output_type -> user.Pres
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_user_user_proto_init() }
func file_pb_user_user_proto_init() {
	if File_pb_user_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_user_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_pb_user_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_user_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_user_user_proto_goTypes,
		DependencyIndexes: file_pb_user_user_proto_depIdxs,
		MessageInfos:      file_pb_user_user_proto_msgTypes,
	}.Build()
	File_pb_user_user_proto = out.File
	file_pb_user_user_proto_rawDesc = nil
	file_pb_user_user_proto_goTypes = nil
	file_pb_user_user_proto_depIdxs = nil
}
