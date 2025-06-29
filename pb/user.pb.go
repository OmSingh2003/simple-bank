// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: user.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type User struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	Username          string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	FullName          string                 `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Email             string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	PasswordChangedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=password_changed_at,json=passwordChangedAt,proto3" json:"password_changed_at,omitempty"`
	CreatedAt         *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *User) Reset() {
	*x = User{}
	mi := &file_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetPasswordChangedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.PasswordChangedAt
	}
	return nil
}

func (x *User) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type CreateUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	FullName      string                 `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Email         string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Password      string                 `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	mi := &file_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUserRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CreateUserRequest) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *CreateUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type CreateUserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          *User                  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateUserResponse) Reset() {
	*x = CreateUserResponse{}
	mi := &file_user_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserResponse) ProtoMessage() {}

func (x *CreateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserResponse.ProtoReflect.Descriptor instead.
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

func (x *CreateUserResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type LoginUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginUserRequest) Reset() {
	*x = LoginUserRequest{}
	mi := &file_user_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginUserRequest) ProtoMessage() {}

func (x *LoginUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginUserRequest.ProtoReflect.Descriptor instead.
func (*LoginUserRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{3}
}

func (x *LoginUserRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginUserResponse struct {
	state                 protoimpl.MessageState `protogen:"open.v1"`
	User                  *User                  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	AccessToken           string                 `protobuf:"bytes,2,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	AccessTokenExpiresAt  *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=access_token_expires_at,json=accessTokenExpiresAt,proto3" json:"access_token_expires_at,omitempty"`
	RefreshToken          string                 `protobuf:"bytes,4,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	RefreshTokenExpiresAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=refresh_token_expires_at,json=refreshTokenExpiresAt,proto3" json:"refresh_token_expires_at,omitempty"`
	SessionId             string                 `protobuf:"bytes,6,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *LoginUserResponse) Reset() {
	*x = LoginUserResponse{}
	mi := &file_user_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginUserResponse) ProtoMessage() {}

func (x *LoginUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginUserResponse.ProtoReflect.Descriptor instead.
func (*LoginUserResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{4}
}

func (x *LoginUserResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *LoginUserResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *LoginUserResponse) GetAccessTokenExpiresAt() *timestamppb.Timestamp {
	if x != nil {
		return x.AccessTokenExpiresAt
	}
	return nil
}

func (x *LoginUserResponse) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *LoginUserResponse) GetRefreshTokenExpiresAt() *timestamppb.Timestamp {
	if x != nil {
		return x.RefreshTokenExpiresAt
	}
	return nil
}

func (x *LoginUserResponse) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

type UpdateUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	FullName      *string                `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3,oneof" json:"full_name,omitempty"`
	Email         *string                `protobuf:"bytes,3,opt,name=email,proto3,oneof" json:"email,omitempty"`
	Password      *string                `protobuf:"bytes,4,opt,name=password,proto3,oneof" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateUserRequest) Reset() {
	*x = UpdateUserRequest{}
	mi := &file_user_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserRequest) ProtoMessage() {}

func (x *UpdateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateUserRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UpdateUserRequest) GetFullName() string {
	if x != nil && x.FullName != nil {
		return *x.FullName
	}
	return ""
}

func (x *UpdateUserRequest) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

func (x *UpdateUserRequest) GetPassword() string {
	if x != nil && x.Password != nil {
		return *x.Password
	}
	return ""
}

type UpdateUserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          *User                  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateUserResponse) Reset() {
	*x = UpdateUserResponse{}
	mi := &file_user_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserResponse) ProtoMessage() {}

func (x *UpdateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserResponse.ProtoReflect.Descriptor instead.
func (*UpdateUserResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateUserResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

var File_user_proto protoreflect.FileDescriptor

const file_user_proto_rawDesc = "" +
	"\n" +
	"\n" +
	"user.proto\x12\x02pb\x1a\x1fgoogle/protobuf/timestamp.proto\"\xdc\x01\n" +
	"\x04User\x12\x1a\n" +
	"\busername\x18\x01 \x01(\tR\busername\x12\x1b\n" +
	"\tfull_name\x18\x02 \x01(\tR\bfullName\x12\x14\n" +
	"\x05email\x18\x03 \x01(\tR\x05email\x12J\n" +
	"\x13password_changed_at\x18\x04 \x01(\v2\x1a.google.protobuf.TimestampR\x11passwordChangedAt\x129\n" +
	"\n" +
	"created_at\x18\x05 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\"~\n" +
	"\x11CreateUserRequest\x12\x1a\n" +
	"\busername\x18\x01 \x01(\tR\busername\x12\x1b\n" +
	"\tfull_name\x18\x02 \x01(\tR\bfullName\x12\x14\n" +
	"\x05email\x18\x03 \x01(\tR\x05email\x12\x1a\n" +
	"\bpassword\x18\x04 \x01(\tR\bpassword\"2\n" +
	"\x12CreateUserResponse\x12\x1c\n" +
	"\x04user\x18\x01 \x01(\v2\b.pb.UserR\x04user\"J\n" +
	"\x10LoginUserRequest\x12\x1a\n" +
	"\busername\x18\x01 \x01(\tR\busername\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\"\xc0\x02\n" +
	"\x11LoginUserResponse\x12\x1c\n" +
	"\x04user\x18\x01 \x01(\v2\b.pb.UserR\x04user\x12!\n" +
	"\faccess_token\x18\x02 \x01(\tR\vaccessToken\x12Q\n" +
	"\x17access_token_expires_at\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\x14accessTokenExpiresAt\x12#\n" +
	"\rrefresh_token\x18\x04 \x01(\tR\frefreshToken\x12S\n" +
	"\x18refresh_token_expires_at\x18\x05 \x01(\v2\x1a.google.protobuf.TimestampR\x15refreshTokenExpiresAt\x12\x1d\n" +
	"\n" +
	"session_id\x18\x06 \x01(\tR\tsessionId\"\xb2\x01\n" +
	"\x11UpdateUserRequest\x12\x1a\n" +
	"\busername\x18\x01 \x01(\tR\busername\x12 \n" +
	"\tfull_name\x18\x02 \x01(\tH\x00R\bfullName\x88\x01\x01\x12\x19\n" +
	"\x05email\x18\x03 \x01(\tH\x01R\x05email\x88\x01\x01\x12\x1f\n" +
	"\bpassword\x18\x04 \x01(\tH\x02R\bpassword\x88\x01\x01B\f\n" +
	"\n" +
	"_full_nameB\b\n" +
	"\x06_emailB\v\n" +
	"\t_password\"2\n" +
	"\x12UpdateUserResponse\x12\x1c\n" +
	"\x04user\x18\x01 \x01(\v2\b.pb.UserR\x04userB*Z(github.com/OmSingh2003/vaultguard-api/pbb\x06proto3"

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData []byte
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_user_proto_rawDesc), len(file_user_proto_rawDesc)))
	})
	return file_user_proto_rawDescData
}

var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_user_proto_goTypes = []any{
	(*User)(nil),                  // 0: pb.User
	(*CreateUserRequest)(nil),     // 1: pb.CreateUserRequest
	(*CreateUserResponse)(nil),    // 2: pb.CreateUserResponse
	(*LoginUserRequest)(nil),      // 3: pb.LoginUserRequest
	(*LoginUserResponse)(nil),     // 4: pb.LoginUserResponse
	(*UpdateUserRequest)(nil),     // 5: pb.UpdateUserRequest
	(*UpdateUserResponse)(nil),    // 6: pb.UpdateUserResponse
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_user_proto_depIdxs = []int32{
	7, // 0: pb.User.password_changed_at:type_name -> google.protobuf.Timestamp
	7, // 1: pb.User.created_at:type_name -> google.protobuf.Timestamp
	0, // 2: pb.CreateUserResponse.user:type_name -> pb.User
	0, // 3: pb.LoginUserResponse.user:type_name -> pb.User
	7, // 4: pb.LoginUserResponse.access_token_expires_at:type_name -> google.protobuf.Timestamp
	7, // 5: pb.LoginUserResponse.refresh_token_expires_at:type_name -> google.protobuf.Timestamp
	0, // 6: pb.UpdateUserResponse.user:type_name -> pb.User
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	file_user_proto_msgTypes[5].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_user_proto_rawDesc), len(file_user_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}
