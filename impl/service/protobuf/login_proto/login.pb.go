// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.0
// source: login/login.proto

package login_proto

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

type CMD int32

const (
	CMD_def            CMD = 0
	CMD_LOGIN_MSG      CMD = 327681
	CMD_LOGIN_KICK_MSG CMD = 327682
)

// Enum value maps for CMD.
var (
	CMD_name = map[int32]string{
		0:      "def",
		327681: "LOGIN_MSG",
		327682: "LOGIN_KICK_MSG",
	}
	CMD_value = map[string]int32{
		"def":            0,
		"LOGIN_MSG":      327681,
		"LOGIN_KICK_MSG": 327682,
	}
)

func (x CMD) Enum() *CMD {
	p := new(CMD)
	*p = x
	return p
}

func (x CMD) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CMD) Descriptor() protoreflect.EnumDescriptor {
	return file_login_login_proto_enumTypes[0].Descriptor()
}

func (CMD) Type() protoreflect.EnumType {
	return &file_login_login_proto_enumTypes[0]
}

func (x CMD) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CMD.Descriptor instead.
func (CMD) EnumDescriptor() ([]byte, []int) {
	return file_login_login_proto_rawDescGZIP(), []int{0}
}

type LoginMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip   *string `protobuf:"bytes,1,opt,name=ip,proto3,oneof" json:"ip,omitempty"`
	Port *uint32 `protobuf:"varint,2,opt,name=port,proto3,oneof" json:"port,omitempty"`
}

func (x *LoginMsg) Reset() {
	*x = LoginMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_login_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginMsg) ProtoMessage() {}

func (x *LoginMsg) ProtoReflect() protoreflect.Message {
	mi := &file_login_login_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginMsg.ProtoReflect.Descriptor instead.
func (*LoginMsg) Descriptor() ([]byte, []int) {
	return file_login_login_proto_rawDescGZIP(), []int{0}
}

func (x *LoginMsg) GetIp() string {
	if x != nil && x.Ip != nil {
		return *x.Ip
	}
	return ""
}

func (x *LoginMsg) GetPort() uint32 {
	if x != nil && x.Port != nil {
		return *x.Port
	}
	return 0
}

type LoginMsgRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Platform *uint32 `protobuf:"varint,1,opt,name=platform,proto3,oneof" json:"platform,omitempty"` //平台号
	Account  *string `protobuf:"bytes,2,opt,name=account,proto3,oneof" json:"account,omitempty"`    //帐号
	Session  *string `protobuf:"bytes,3,opt,name=session,proto3,oneof" json:"session,omitempty"`    //登录验证时使用的session
}

func (x *LoginMsgRes) Reset() {
	*x = LoginMsgRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_login_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginMsgRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginMsgRes) ProtoMessage() {}

func (x *LoginMsgRes) ProtoReflect() protoreflect.Message {
	mi := &file_login_login_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginMsgRes.ProtoReflect.Descriptor instead.
func (*LoginMsgRes) Descriptor() ([]byte, []int) {
	return file_login_login_proto_rawDescGZIP(), []int{1}
}

func (x *LoginMsgRes) GetPlatform() uint32 {
	if x != nil && x.Platform != nil {
		return *x.Platform
	}
	return 0
}

func (x *LoginMsgRes) GetAccount() string {
	if x != nil && x.Account != nil {
		return *x.Account
	}
	return ""
}

func (x *LoginMsgRes) GetSession() string {
	if x != nil && x.Session != nil {
		return *x.Session
	}
	return ""
}

type LoginKickMsgRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Platform *uint32 `protobuf:"varint,1,opt,name=platform,proto3,oneof" json:"platform,omitempty"` //平台号
	Account  *string `protobuf:"bytes,2,opt,name=account,proto3,oneof" json:"account,omitempty"`    //帐号
}

func (x *LoginKickMsgRes) Reset() {
	*x = LoginKickMsgRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_login_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginKickMsgRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginKickMsgRes) ProtoMessage() {}

func (x *LoginKickMsgRes) ProtoReflect() protoreflect.Message {
	mi := &file_login_login_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginKickMsgRes.ProtoReflect.Descriptor instead.
func (*LoginKickMsgRes) Descriptor() ([]byte, []int) {
	return file_login_login_proto_rawDescGZIP(), []int{2}
}

func (x *LoginKickMsgRes) GetPlatform() uint32 {
	if x != nil && x.Platform != nil {
		return *x.Platform
	}
	return 0
}

func (x *LoginKickMsgRes) GetAccount() string {
	if x != nil && x.Account != nil {
		return *x.Account
	}
	return ""
}

var File_login_login_proto protoreflect.FileDescriptor

var file_login_login_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x49, 0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x6d, 0x73, 0x67,
	0x12, 0x13, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x02,
	0x69, 0x70, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x48, 0x01, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x88, 0x01, 0x01, 0x42, 0x05,
	0x0a, 0x03, 0x5f, 0x69, 0x70, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x93,
	0x01, 0x0a, 0x0d, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x6d, 0x73, 0x67, 0x5f, 0x72, 0x65, 0x73,
	0x12, 0x1f, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x48, 0x00, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x88, 0x01,
	0x01, 0x12, 0x1d, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x01, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01,
	0x12, 0x1d, 0x0a, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x02, 0x52, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42,
	0x0b, 0x0a, 0x09, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x42, 0x0a, 0x0a, 0x08,
	0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x22, 0x6d, 0x0a, 0x12, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x6b, 0x69,
	0x63, 0x6b, 0x5f, 0x6d, 0x73, 0x67, 0x5f, 0x72, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x08, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x08,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x07,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2a, 0x35, 0x0a, 0x03, 0x43, 0x4d, 0x44, 0x12, 0x07, 0x0a, 0x03, 0x64, 0x65,
	0x66, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x09, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x4d, 0x53, 0x47,
	0x10, 0x81, 0x80, 0x14, 0x12, 0x14, 0x0a, 0x0e, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x4b, 0x49,
	0x43, 0x4b, 0x5f, 0x4d, 0x53, 0x47, 0x10, 0x82, 0x80, 0x14, 0x42, 0x21, 0x5a, 0x1f, 0x2e, 0x2e,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_login_login_proto_rawDescOnce sync.Once
	file_login_login_proto_rawDescData = file_login_login_proto_rawDesc
)

func file_login_login_proto_rawDescGZIP() []byte {
	file_login_login_proto_rawDescOnce.Do(func() {
		file_login_login_proto_rawDescData = protoimpl.X.CompressGZIP(file_login_login_proto_rawDescData)
	})
	return file_login_login_proto_rawDescData
}

var file_login_login_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_login_login_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_login_login_proto_goTypes = []interface{}{
	(CMD)(0),                // 0: CMD
	(*LoginMsg)(nil),        // 1: login_msg
	(*LoginMsgRes)(nil),     // 2: login_msg_res
	(*LoginKickMsgRes)(nil), // 3: login_kick_msg_res
}
var file_login_login_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_login_login_proto_init() }
func file_login_login_proto_init() {
	if File_login_login_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_login_login_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginMsg); i {
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
		file_login_login_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginMsgRes); i {
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
		file_login_login_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginKickMsgRes); i {
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
	file_login_login_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_login_login_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_login_login_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_login_login_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_login_login_proto_goTypes,
		DependencyIndexes: file_login_login_proto_depIdxs,
		EnumInfos:         file_login_login_proto_enumTypes,
		MessageInfos:      file_login_login_proto_msgTypes,
	}.Build()
	File_login_login_proto = out.File
	file_login_login_proto_rawDesc = nil
	file_login_login_proto_goTypes = nil
	file_login_login_proto_depIdxs = nil
}
