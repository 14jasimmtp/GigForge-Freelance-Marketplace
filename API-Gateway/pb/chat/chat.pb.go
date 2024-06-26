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

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageID   string `protobuf:"bytes,1,opt,name=MessageID,proto3" json:"MessageID,omitempty"`
	SenderId    string `protobuf:"bytes,2,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	RecipientId string `protobuf:"bytes,3,opt,name=recipient_id,json=recipientId,proto3" json:"recipient_id,omitempty"`
	Content     string `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	Timestamp   string `protobuf:"bytes,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Type        string `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
	Tag         string `protobuf:"bytes,7,opt,name=tag,proto3" json:"tag,omitempty"`
	Status      string `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_chat_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_pb_chat_chat_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetMessageID() string {
	if x != nil {
		return x.MessageID
	}
	return ""
}

func (x *Message) GetSenderId() string {
	if x != nil {
		return x.SenderId
	}
	return ""
}

func (x *Message) GetRecipientId() string {
	if x != nil {
		return x.RecipientId
	}
	return ""
}

func (x *Message) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Message) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *Message) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Message) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *Message) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type GetChatReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SenderId   string `protobuf:"bytes,1,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	RecieverId string `protobuf:"bytes,2,opt,name=reciever_id,json=recieverId,proto3" json:"reciever_id,omitempty"`
}

func (x *GetChatReq) Reset() {
	*x = GetChatReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_chat_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChatReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChatReq) ProtoMessage() {}

func (x *GetChatReq) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetChatReq.ProtoReflect.Descriptor instead.
func (*GetChatReq) Descriptor() ([]byte, []int) {
	return file_pb_chat_chat_proto_rawDescGZIP(), []int{1}
}

func (x *GetChatReq) GetSenderId() string {
	if x != nil {
		return x.SenderId
	}
	return ""
}

func (x *GetChatReq) GetRecieverId() string {
	if x != nil {
		return x.RecieverId
	}
	return ""
}

type GetChatRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chat []*Message `protobuf:"bytes,1,rep,name=Chat,proto3" json:"Chat,omitempty"`
}

func (x *GetChatRes) Reset() {
	*x = GetChatRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_chat_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChatRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChatRes) ProtoMessage() {}

func (x *GetChatRes) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetChatRes.ProtoReflect.Descriptor instead.
func (*GetChatRes) Descriptor() ([]byte, []int) {
	return file_pb_chat_chat_proto_rawDescGZIP(), []int{2}
}

func (x *GetChatRes) GetChat() []*Message {
	if x != nil {
		return x.Chat
	}
	return nil
}

var File_pb_chat_chat_proto protoreflect.FileDescriptor

var file_pb_chat_chat_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x62, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x68, 0x61, 0x74, 0x22, 0xdd, 0x01, 0x0a, 0x07, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x49, 0x44, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74,
	0x61, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x4a, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x69, 0x65, 0x76, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x69,
	0x65, 0x76, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2f, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61,
	0x74, 0x52, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x04, 0x43, 0x68, 0x61, 0x74, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x04, 0x43, 0x68, 0x61, 0x74, 0x32, 0x3f, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61,
	0x74, 0x73, 0x12, 0x10, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61,
	0x74, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x70, 0x62,
	0x2f, 0x63, 0x68, 0x61, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_pb_chat_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pb_chat_chat_proto_goTypes = []interface{}{
	(*Message)(nil),    // 0: chat.Message
	(*GetChatReq)(nil), // 1: chat.GetChatReq
	(*GetChatRes)(nil), // 2: chat.GetChatRes
}
var file_pb_chat_chat_proto_depIdxs = []int32{
	0, // 0: chat.GetChatRes.Chat:type_name -> chat.Message
	1, // 1: chat.ChatService.GetChats:input_type -> chat.GetChatReq
	2, // 2: chat.ChatService.GetChats:output_type -> chat.GetChatRes
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pb_chat_chat_proto_init() }
func file_pb_chat_chat_proto_init() {
	if File_pb_chat_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_chat_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_pb_chat_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_chat_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
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
