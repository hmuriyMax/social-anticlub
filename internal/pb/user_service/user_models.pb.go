// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.25.3
// source: api/user_models.proto

package user_service

import (
	date "google.golang.org/genproto/googleapis/type/date"
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

type UserInfo_Gender int32

const (
	UserInfo_UNSET  UserInfo_Gender = 0
	UserInfo_MALE   UserInfo_Gender = 1
	UserInfo_FEMALE UserInfo_Gender = 2
)

// Enum value maps for UserInfo_Gender.
var (
	UserInfo_Gender_name = map[int32]string{
		0: "UNSET",
		1: "MALE",
		2: "FEMALE",
	}
	UserInfo_Gender_value = map[string]int32{
		"UNSET":  0,
		"MALE":   1,
		"FEMALE": 2,
	}
)

func (x UserInfo_Gender) Enum() *UserInfo_Gender {
	p := new(UserInfo_Gender)
	*p = x
	return p
}

func (x UserInfo_Gender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserInfo_Gender) Descriptor() protoreflect.EnumDescriptor {
	return file_api_user_models_proto_enumTypes[0].Descriptor()
}

func (UserInfo_Gender) Type() protoreflect.EnumType {
	return &file_api_user_models_proto_enumTypes[0]
}

func (x UserInfo_Gender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserInfo_Gender.Descriptor instead.
func (UserInfo_Gender) EnumDescriptor() ([]byte, []int) {
	return file_api_user_models_proto_rawDescGZIP(), []int{2, 0}
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token   string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Expires int64  `protobuf:"varint,2,opt,name=expires,proto3" json:"expires,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_models_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_models_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_api_user_models_proto_rawDescGZIP(), []int{0}
}

func (x *Token) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *Token) GetExpires() int64 {
	if x != nil {
		return x.Expires
	}
	return 0
}

type LoginInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID int64  `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Token  *Token `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *LoginInfo) Reset() {
	*x = LoginInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_models_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginInfo) ProtoMessage() {}

func (x *LoginInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_models_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginInfo.ProtoReflect.Descriptor instead.
func (*LoginInfo) Descriptor() ([]byte, []int) {
	return file_api_user_models_proto_rawDescGZIP(), []int{1}
}

func (x *LoginInfo) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *LoginInfo) GetToken() *Token {
	if x != nil {
		return x.Token
	}
	return nil
}

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Surname  string          `protobuf:"bytes,2,opt,name=surname,proto3" json:"surname,omitempty"`
	Birthday *date.Date      `protobuf:"bytes,3,opt,name=birthday,proto3" json:"birthday,omitempty"`
	Gender   UserInfo_Gender `protobuf:"varint,4,opt,name=gender,proto3,enum=social_anti_club.UserInfo_Gender" json:"gender,omitempty"`
	Hobbies  string          `protobuf:"bytes,5,opt,name=hobbies,proto3" json:"hobbies,omitempty"`
	HomeTown string          `protobuf:"bytes,6,opt,name=home_town,proto3" json:"home_town,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_models_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_models_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_api_user_models_proto_rawDescGZIP(), []int{2}
}

func (x *UserInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserInfo) GetSurname() string {
	if x != nil {
		return x.Surname
	}
	return ""
}

func (x *UserInfo) GetBirthday() *date.Date {
	if x != nil {
		return x.Birthday
	}
	return nil
}

func (x *UserInfo) GetGender() UserInfo_Gender {
	if x != nil {
		return x.Gender
	}
	return UserInfo_UNSET
}

func (x *UserInfo) GetHobbies() string {
	if x != nil {
		return x.Hobbies
	}
	return ""
}

func (x *UserInfo) GetHomeTown() string {
	if x != nil {
		return x.HomeTown
	}
	return ""
}

var File_api_user_models_proto protoreflect.FileDescriptor

var file_api_user_models_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x5f,
	0x61, 0x6e, 0x74, 0x69, 0x5f, 0x63, 0x6c, 0x75, 0x62, 0x1a, 0x16, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x37, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x18, 0x0a, 0x07, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x22, 0x52, 0x0a, 0x09, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12,
	0x2d, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x5f, 0x61, 0x6e, 0x74, 0x69, 0x5f, 0x63, 0x6c, 0x75,
	0x62, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x85,
	0x02, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x08, 0x62, 0x69, 0x72,
	0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x52, 0x08,
	0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x12, 0x39, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x73, 0x6f, 0x63, 0x69, 0x61,
	0x6c, 0x5f, 0x61, 0x6e, 0x74, 0x69, 0x5f, 0x63, 0x6c, 0x75, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x06, 0x67, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x6f, 0x62, 0x62, 0x69, 0x65, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x68, 0x6f, 0x62, 0x62, 0x69, 0x65, 0x73, 0x12, 0x1c, 0x0a,
	0x09, 0x68, 0x6f, 0x6d, 0x65, 0x5f, 0x74, 0x6f, 0x77, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x68, 0x6f, 0x6d, 0x65, 0x5f, 0x74, 0x6f, 0x77, 0x6e, 0x22, 0x29, 0x0a, 0x06, 0x47,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x09, 0x0a, 0x05, 0x55, 0x4e, 0x53, 0x45, 0x54, 0x10, 0x00,
	0x12, 0x08, 0x0a, 0x04, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x45,
	0x4d, 0x41, 0x4c, 0x45, 0x10, 0x02, 0x42, 0x0f, 0x5a, 0x0d, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_user_models_proto_rawDescOnce sync.Once
	file_api_user_models_proto_rawDescData = file_api_user_models_proto_rawDesc
)

func file_api_user_models_proto_rawDescGZIP() []byte {
	file_api_user_models_proto_rawDescOnce.Do(func() {
		file_api_user_models_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_user_models_proto_rawDescData)
	})
	return file_api_user_models_proto_rawDescData
}

var file_api_user_models_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_user_models_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_user_models_proto_goTypes = []interface{}{
	(UserInfo_Gender)(0), // 0: social_anti_club.UserInfo.Gender
	(*Token)(nil),        // 1: social_anti_club.Token
	(*LoginInfo)(nil),    // 2: social_anti_club.LoginInfo
	(*UserInfo)(nil),     // 3: social_anti_club.UserInfo
	(*date.Date)(nil),    // 4: google.type.Date
}
var file_api_user_models_proto_depIdxs = []int32{
	1, // 0: social_anti_club.LoginInfo.token:type_name -> social_anti_club.Token
	4, // 1: social_anti_club.UserInfo.birthday:type_name -> google.type.Date
	0, // 2: social_anti_club.UserInfo.gender:type_name -> social_anti_club.UserInfo.Gender
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_user_models_proto_init() }
func file_api_user_models_proto_init() {
	if File_api_user_models_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_user_models_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
		file_api_user_models_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginInfo); i {
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
		file_api_user_models_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
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
			RawDescriptor: file_api_user_models_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_user_models_proto_goTypes,
		DependencyIndexes: file_api_user_models_proto_depIdxs,
		EnumInfos:         file_api_user_models_proto_enumTypes,
		MessageInfos:      file_api_user_models_proto_msgTypes,
	}.Build()
	File_api_user_models_proto = out.File
	file_api_user_models_proto_rawDesc = nil
	file_api_user_models_proto_goTypes = nil
	file_api_user_models_proto_depIdxs = nil
}
