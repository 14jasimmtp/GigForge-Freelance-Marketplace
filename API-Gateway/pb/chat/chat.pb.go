// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: pb/chat/chat.proto

package chat

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

type GetChatReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetChatReq) Reset() {
	*x = GetChatReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_chat_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChatReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChatReq) ProtoMessage() {}

func (x *GetChatReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_chat_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChatReq.ProtoReflect.Descriptor instead.
func (*GetChatReq) Descriptor() ([]byte, []int) {
	return file_pb_chat_chat_proto_rawDescGZIP(), []int{0}
}

type GetChatRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetChatRes) Reset() {
	*x = GetChatRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_chat_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChatRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChatRes) ProtoMessage() {}

func (x *GetChatRes) ProtoReflect() protoreflect.Message {
	mi := &file_pb_chat_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChatRes.ProtoReflect.Descriptor instead.
func (*GetChatRes) Descriptor() ([]byte, []int) {
	return file_pb_chat_chat_proto_rawDescGZIP(), []int{1}
}

type SendChatReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendChatReq) Reset() {
	*x = SendChatReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_chat_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendChatReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendChatReq) ProtoMessage() {}

func (x *SendChatReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_chat_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendChatReq.ProtoReflect.Descriptor instead.
func (*SendChatReq) Descriptor() ([]byte, []int) {
	return file_pb_chat_chat_proto_rawDescGZIP(), []int{2}
}

type SendChatRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendChatRes) Reset() {
	*x = SendChatRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_chat_chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendChatRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendChatRes) ProtoMessage() {}

func (x *SendChatRes) ProtoReflect() protoreflect.Message {
	mi := &file_pb_chat_chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendChatRes.ProtoReflect.Descriptor instead.
func (*SendChatRes) Descriptor() ([]byte, []int) {
	return file_pb_chat_chat_proto_rawDescGZIP(), []int{3}
}

var File_pb_chat_chat_proto protoreflect.FileDescriptor

var file_pb_chat_chat_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x62, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x68, 0x61, 0x74, 0x22, 0x0c, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x22, 0x0c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43,
	0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x22, 0x0d, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x68,
	0x61, 0x74, 0x52, 0x65, 0x71, 0x22, 0x0d, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x68, 0x61,
	0x74, 0x52, 0x65, 0x73, 0x32, 0x74, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x73, 0x12,
	0x10, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65,
	0x71, 0x1a, 0x10, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74,
	0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x09, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x68, 0x61,
	0x74, 0x73, 0x12, 0x11, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x68,
	0x61, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x53, 0x65, 0x6e,
	0x64, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f,
	0x70, 0x62, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_chat_chat_proto_rawDescOnce sync.Once
	file_pb_chat_chat_proto_rawDescData = file_pb_chat_chat_proto_rawDesc
)

func file_pb_chat_chat_proto_rawDescGZIP() []byte {
	file_pb_chat_chat_proto_rawDescOnce.Do(func() {
		file_pb_chat_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_chat_chat_proto_rawDescData)
	})
	return file_pb_chat_chat_proto_rawDescData
}

var file_pb_chat_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pb_chat_chat_proto_goTypes = []interface{}{
	(*GetChatReq)(nil),  // 0: chat.GetChatReq
	(*GetChatRes)(nil),  // 1: chat.GetChatRes
	(*SendChatReq)(nil), // 2: chat.SendChatReq
	(*SendChatRes)(nil), // 3: chat.SendChatRes
}
var file_pb_chat_chat_proto_depIdxs = []int32{
	0, // 0: chat.ChatService.GetChats:input_type -> chat.GetChatReq
	2, // 1: chat.ChatService.SendChats:input_type -> chat.SendChatReq
	1, // 2: chat.ChatService.GetChats:output_type -> chat.GetChatRes
	3, // 3: chat.ChatService.SendChats:output_type -> chat.SendChatRes
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_chat_chat_proto_init() }
func file_pb_chat_chat_proto_init() {
	if File_pb_chat_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_chat_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChatReq); i {
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
		file_pb_chat_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChatRes); i {
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
		file_pb_chat_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendChatReq); i {
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
		file_pb_chat_chat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendChatRes); i {
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
			RawDescriptor: file_pb_chat_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_chat_chat_proto_goTypes,
		DependencyIndexes: file_pb_chat_chat_proto_depIdxs,
		MessageInfos:      file_pb_chat_chat_proto_msgTypes,
	}.Build()
	File_pb_chat_chat_proto = out.File
	file_pb_chat_chat_proto_rawDesc = nil
	file_pb_chat_chat_proto_goTypes = nil
	file_pb_chat_chat_proto_depIdxs = nil
}
