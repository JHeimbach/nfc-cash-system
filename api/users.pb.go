// Code generated by protoc-gen-go. DO NOT EDIT.
// source: users.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type AuthenticateResponse_TokenType int32

const (
	AuthenticateResponse_BEARER AuthenticateResponse_TokenType = 0
)

var AuthenticateResponse_TokenType_name = map[int32]string{
	0: "BEARER",
}

var AuthenticateResponse_TokenType_value = map[string]int32{
	"BEARER": 0,
}

func (x AuthenticateResponse_TokenType) String() string {
	return proto.EnumName(AuthenticateResponse_TokenType_name, int32(x))
}

func (AuthenticateResponse_TokenType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{0, 0}
}

type AuthenticateResponse struct {
	TokenType            AuthenticateResponse_TokenType `protobuf:"varint,1,opt,name=token_type,json=tokenType,proto3,enum=api.AuthenticateResponse_TokenType" json:"token_type,omitempty"`
	AccessToken          string                         `protobuf:"bytes,2,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken         string                         `protobuf:"bytes,3,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	ExpiresIn            int64                          `protobuf:"varint,4,opt,name=expires_in,json=expiresIn,proto3" json:"expires_in,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *AuthenticateResponse) Reset()         { *m = AuthenticateResponse{} }
func (m *AuthenticateResponse) String() string { return proto.CompactTextString(m) }
func (*AuthenticateResponse) ProtoMessage()    {}
func (*AuthenticateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{0}
}

func (m *AuthenticateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticateResponse.Unmarshal(m, b)
}
func (m *AuthenticateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticateResponse.Marshal(b, m, deterministic)
}
func (m *AuthenticateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticateResponse.Merge(m, src)
}
func (m *AuthenticateResponse) XXX_Size() int {
	return xxx_messageInfo_AuthenticateResponse.Size(m)
}
func (m *AuthenticateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticateResponse proto.InternalMessageInfo

func (m *AuthenticateResponse) GetTokenType() AuthenticateResponse_TokenType {
	if m != nil {
		return m.TokenType
	}
	return AuthenticateResponse_BEARER
}

func (m *AuthenticateResponse) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *AuthenticateResponse) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

func (m *AuthenticateResponse) GetExpiresIn() int64 {
	if m != nil {
		return m.ExpiresIn
	}
	return 0
}

type User struct {
	Id                   int32                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email                string               `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Created              *timestamp.Timestamp `protobuf:"bytes,4,opt,name=created,proto3" json:"created,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{1}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetCreated() *timestamp.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func init() {
	proto.RegisterEnum("api.AuthenticateResponse_TokenType", AuthenticateResponse_TokenType_name, AuthenticateResponse_TokenType_value)
	proto.RegisterType((*AuthenticateResponse)(nil), "api.AuthenticateResponse")
	proto.RegisterType((*User)(nil), "api.User")
}

func init() { proto.RegisterFile("users.proto", fileDescriptor_030765f334c86cea) }

var fileDescriptor_030765f334c86cea = []byte{
	// 531 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0xcf, 0x8f, 0xd2, 0x40,
	0x14, 0xde, 0xc2, 0xee, 0x1a, 0x1e, 0xbb, 0x80, 0x13, 0xa2, 0x58, 0x7f, 0xd5, 0xae, 0x26, 0x84,
	0x48, 0x1b, 0xd1, 0x93, 0x1a, 0x93, 0x12, 0x39, 0x68, 0xf4, 0x52, 0x31, 0x9e, 0x0c, 0x19, 0xca,
	0xa3, 0x8c, 0xc2, 0x4c, 0xd3, 0x99, 0xe2, 0x72, 0xf5, 0x4f, 0xc0, 0x7f, 0xc1, 0x9b, 0x7f, 0x8e,
	0x37, 0x13, 0x6f, 0xfe, 0x21, 0xa6, 0xd3, 0x61, 0x45, 0x71, 0x3d, 0x78, 0xea, 0xf4, 0xcd, 0xf7,
	0xbe, 0xef, 0x9b, 0xef, 0x3d, 0xa8, 0x66, 0x12, 0x53, 0xe9, 0x25, 0xa9, 0x50, 0x82, 0x94, 0x69,
	0xc2, 0xec, 0x9b, 0xb1, 0x10, 0xf1, 0x1c, 0x7d, 0x5d, 0x1a, 0x67, 0x53, 0x5f, 0xb1, 0x05, 0x4a,
	0x45, 0x17, 0x49, 0x81, 0xb2, 0xaf, 0x19, 0x00, 0x4d, 0x98, 0x4f, 0x39, 0x17, 0x8a, 0x2a, 0x26,
	0xb8, 0xe1, 0xb0, 0xaf, 0xfe, 0xd9, 0x8e, 0x8b, 0x44, 0xad, 0xcc, 0xe5, 0x5d, 0xfd, 0x89, 0xba,
	0x31, 0xf2, 0xae, 0xfc, 0x40, 0xe3, 0x18, 0x53, 0x5f, 0x24, 0xba, 0x7d, 0x97, 0xca, 0xfd, 0x66,
	0x41, 0x33, 0xc8, 0xd4, 0x0c, 0xb9, 0x62, 0x11, 0x55, 0x18, 0xa2, 0x4c, 0x04, 0x97, 0x48, 0xfa,
	0x00, 0x4a, 0xbc, 0x47, 0x3e, 0x52, 0xab, 0x04, 0x5b, 0x96, 0x63, 0xb5, 0x6b, 0xbd, 0x13, 0x8f,
	0x26, 0xcc, 0xfb, 0x1b, 0xdc, 0x1b, 0xe6, 0xd8, 0xe1, 0x2a, 0xc1, 0xb0, 0xa2, 0x36, 0x47, 0x72,
	0x0b, 0x8e, 0x68, 0x14, 0xa1, 0x94, 0x23, 0x5d, 0x6b, 0x95, 0x1c, 0xab, 0x5d, 0x09, 0xab, 0x45,
	0x4d, 0x77, 0x90, 0x13, 0x38, 0x4e, 0x71, 0x9a, 0xa2, 0x9c, 0x19, 0x4c, 0x59, 0x63, 0x8e, 0x4c,
	0xb1, 0x00, 0x5d, 0x07, 0xc0, 0xd3, 0x84, 0xa5, 0x28, 0x47, 0x8c, 0xb7, 0xf6, 0x1d, 0xab, 0x5d,
	0x0e, 0x2b, 0xa6, 0xf2, 0x8c, 0xbb, 0x97, 0xa1, 0x72, 0x26, 0x4f, 0x00, 0x0e, 0xfb, 0x83, 0x20,
	0x1c, 0x84, 0x8d, 0x3d, 0x77, 0x09, 0xfb, 0xaf, 0x25, 0xa6, 0xa4, 0x06, 0x25, 0x36, 0xd1, 0x6f,
	0x38, 0x08, 0x4b, 0x6c, 0x42, 0x08, 0xec, 0x73, 0xba, 0x40, 0xe3, 0x47, 0x9f, 0x49, 0x13, 0x0e,
	0x70, 0x41, 0xd9, 0xdc, 0x18, 0x28, 0x7e, 0xc8, 0x03, 0xb8, 0x10, 0xa5, 0x48, 0x15, 0x4e, 0xb4,
	0x6c, 0xb5, 0x67, 0x7b, 0x45, 0xf6, 0xde, 0x26, 0x7b, 0x6f, 0xb8, 0x19, 0x5d, 0xb8, 0x81, 0xf6,
	0xbe, 0x97, 0xa1, 0x9a, 0x0b, 0xbf, 0xc2, 0x74, 0xc9, 0x22, 0x24, 0x9f, 0x2d, 0x68, 0x6c, 0xa7,
	0xa6, 0x4d, 0x5d, 0xda, 0x61, 0x1a, 0xe4, 0x53, 0xb4, 0xaf, 0x9c, 0x1b, 0xb2, 0xfb, 0x76, 0x1d,
	0x3c, 0xb5, 0xef, 0x84, 0xa8, 0xb2, 0x94, 0x4b, 0xe7, 0xf9, 0x9b, 0xa1, 0xa3, 0xdf, 0x2d, 0x9d,
	0xa9, 0x48, 0x1d, 0xfa, 0xab, 0x83, 0x09, 0xde, 0x81, 0x17, 0x22, 0x66, 0xdc, 0xc9, 0xa5, 0xc6,
	0x75, 0x38, 0x86, 0x4a, 0x9f, 0x4a, 0x16, 0xe5, 0xb4, 0x64, 0xef, 0xe3, 0xd7, 0x1f, 0x9f, 0x4a,
	0x0d, 0x52, 0xf3, 0x97, 0xf7, 0xfc, 0x7c, 0x37, 0xfd, 0x79, 0x8e, 0x25, 0x2b, 0xc8, 0x9b, 0x44,
	0xa6, 0xfe, 0xe9, 0xef, 0x9c, 0xba, 0xfb, 0x68, 0x1d, 0xdc, 0xe8, 0x54, 0x0b, 0x82, 0x2d, 0x59,
	0xed, 0x70, 0x4b, 0xb6, 0xe9, 0xd6, 0xb7, 0x65, 0x45, 0xa6, 0x1e, 0x5a, 0x1d, 0xf2, 0xc5, 0x82,
	0xa3, 0x70, 0x7b, 0xe4, 0xff, 0x91, 0xce, 0x6c, 0x1d, 0xbc, 0xec, 0x34, 0x0d, 0x8b, 0x13, 0xe8,
	0x25, 0x2b, 0x02, 0xda, 0x71, 0xf2, 0xee, 0x36, 0xd4, 0x4f, 0xbb, 0x66, 0xc7, 0xba, 0x7a, 0xf1,
	0xc8, 0x45, 0xbb, 0xfe, 0xf8, 0xb7, 0x55, 0x7c, 0xa2, 0xfd, 0x12, 0xd2, 0x38, 0xf3, 0x6b, 0xae,
	0xc7, 0x87, 0xda, 0xd4, 0xfd, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xdf, 0x92, 0x0a, 0x6a, 0xd9,
	0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	AuthenticateUser(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*AuthenticateResponse, error)
	LogoutUser(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
	RefreshToken(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*AuthenticateResponse, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) AuthenticateUser(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*AuthenticateResponse, error) {
	out := new(AuthenticateResponse)
	err := c.cc.Invoke(ctx, "/api.UserService/AuthenticateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) LogoutUser(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.UserService/LogoutUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) RefreshToken(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*AuthenticateResponse, error) {
	out := new(AuthenticateResponse)
	err := c.cc.Invoke(ctx, "/api.UserService/RefreshToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	AuthenticateUser(context.Context, *empty.Empty) (*AuthenticateResponse, error)
	LogoutUser(context.Context, *empty.Empty) (*empty.Empty, error)
	RefreshToken(context.Context, *empty.Empty) (*AuthenticateResponse, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) AuthenticateUser(ctx context.Context, req *empty.Empty) (*AuthenticateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthenticateUser not implemented")
}
func (*UnimplementedUserServiceServer) LogoutUser(ctx context.Context, req *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogoutUser not implemented")
}
func (*UnimplementedUserServiceServer) RefreshToken(ctx context.Context, req *empty.Empty) (*AuthenticateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_AuthenticateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AuthenticateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserService/AuthenticateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AuthenticateUser(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_LogoutUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).LogoutUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserService/LogoutUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).LogoutUser(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_RefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).RefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserService/RefreshToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).RefreshToken(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AuthenticateUser",
			Handler:    _UserService_AuthenticateUser_Handler,
		},
		{
			MethodName: "LogoutUser",
			Handler:    _UserService_LogoutUser_Handler,
		},
		{
			MethodName: "RefreshToken",
			Handler:    _UserService_RefreshToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "users.proto",
}
