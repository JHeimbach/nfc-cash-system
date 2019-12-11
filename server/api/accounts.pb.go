// Code generated by protoc-gen-go. DO NOT EDIT.
// source: accounts.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type AccountListRequest struct {
	GroupFilter          int32    `protobuf:"varint,1,opt,name=group_filter,json=groupFilter,proto3" json:"group_filter,omitempty"`
	Paging               *Paging  `protobuf:"bytes,2,opt,name=paging,proto3" json:"paging,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountListRequest) Reset()         { *m = AccountListRequest{} }
func (m *AccountListRequest) String() string { return proto.CompactTextString(m) }
func (*AccountListRequest) ProtoMessage()    {}
func (*AccountListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1e7723af4c007b7, []int{0}
}

func (m *AccountListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountListRequest.Unmarshal(m, b)
}
func (m *AccountListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountListRequest.Marshal(b, m, deterministic)
}
func (m *AccountListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountListRequest.Merge(m, src)
}
func (m *AccountListRequest) XXX_Size() int {
	return xxx_messageInfo_AccountListRequest.Size(m)
}
func (m *AccountListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AccountListRequest proto.InternalMessageInfo

func (m *AccountListRequest) GetGroupFilter() int32 {
	if m != nil {
		return m.GroupFilter
	}
	return 0
}

func (m *AccountListRequest) GetPaging() *Paging {
	if m != nil {
		return m.Paging
	}
	return nil
}

type Accounts struct {
	Accounts             []*Account `protobuf:"bytes,1,rep,name=accounts,proto3" json:"accounts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Accounts) Reset()         { *m = Accounts{} }
func (m *Accounts) String() string { return proto.CompactTextString(m) }
func (*Accounts) ProtoMessage()    {}
func (*Accounts) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1e7723af4c007b7, []int{1}
}

func (m *Accounts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Accounts.Unmarshal(m, b)
}
func (m *Accounts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Accounts.Marshal(b, m, deterministic)
}
func (m *Accounts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Accounts.Merge(m, src)
}
func (m *Accounts) XXX_Size() int {
	return xxx_messageInfo_Accounts.Size(m)
}
func (m *Accounts) XXX_DiscardUnknown() {
	xxx_messageInfo_Accounts.DiscardUnknown(m)
}

var xxx_messageInfo_Accounts proto.InternalMessageInfo

func (m *Accounts) GetAccounts() []*Account {
	if m != nil {
		return m.Accounts
	}
	return nil
}

type Account struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Saldo                float64  `protobuf:"fixed64,4,opt,name=saldo,proto3" json:"saldo,omitempty"`
	NfcChipId            string   `protobuf:"bytes,5,opt,name=nfc_chip_id,json=nfcChipId,proto3" json:"nfc_chip_id,omitempty"`
	Group                *Group   `protobuf:"bytes,6,opt,name=group,proto3" json:"group,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1e7723af4c007b7, []int{2}
}

func (m *Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account.Unmarshal(m, b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account.Marshal(b, m, deterministic)
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return xxx_messageInfo_Account.Size(m)
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Account) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Account) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Account) GetSaldo() float64 {
	if m != nil {
		return m.Saldo
	}
	return 0
}

func (m *Account) GetNfcChipId() string {
	if m != nil {
		return m.NfcChipId
	}
	return ""
}

func (m *Account) GetGroup() *Group {
	if m != nil {
		return m.Group
	}
	return nil
}

type AccountCreate struct {
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Saldo                float64  `protobuf:"fixed64,4,opt,name=saldo,proto3" json:"saldo,omitempty"`
	NfcChipId            string   `protobuf:"bytes,5,opt,name=nfc_chip_id,json=nfcChipId,proto3" json:"nfc_chip_id,omitempty"`
	GroupId              int32    `protobuf:"varint,6,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountCreate) Reset()         { *m = AccountCreate{} }
func (m *AccountCreate) String() string { return proto.CompactTextString(m) }
func (*AccountCreate) ProtoMessage()    {}
func (*AccountCreate) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1e7723af4c007b7, []int{3}
}

func (m *AccountCreate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountCreate.Unmarshal(m, b)
}
func (m *AccountCreate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountCreate.Marshal(b, m, deterministic)
}
func (m *AccountCreate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountCreate.Merge(m, src)
}
func (m *AccountCreate) XXX_Size() int {
	return xxx_messageInfo_AccountCreate.Size(m)
}
func (m *AccountCreate) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountCreate.DiscardUnknown(m)
}

var xxx_messageInfo_AccountCreate proto.InternalMessageInfo

func (m *AccountCreate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AccountCreate) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *AccountCreate) GetSaldo() float64 {
	if m != nil {
		return m.Saldo
	}
	return 0
}

func (m *AccountCreate) GetNfcChipId() string {
	if m != nil {
		return m.NfcChipId
	}
	return ""
}

func (m *AccountCreate) GetGroupId() int32 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

func init() {
	proto.RegisterType((*AccountListRequest)(nil), "api.AccountListRequest")
	proto.RegisterType((*Accounts)(nil), "api.Accounts")
	proto.RegisterType((*Account)(nil), "api.Account")
	proto.RegisterType((*AccountCreate)(nil), "api.AccountCreate")
}

func init() { proto.RegisterFile("accounts.proto", fileDescriptor_e1e7723af4c007b7) }

var fileDescriptor_e1e7723af4c007b7 = []byte{
	// 751 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0xd1, 0x6e, 0xeb, 0x34,
	0x18, 0x56, 0xda, 0xb5, 0xdb, 0xdc, 0x2e, 0x5b, 0x3d, 0xca, 0xa2, 0x08, 0x84, 0x09, 0x4c, 0xaa,
	0xac, 0x6d, 0x85, 0x0d, 0x71, 0xd1, 0x1b, 0x94, 0x6e, 0x30, 0x8a, 0xa6, 0x09, 0x65, 0xda, 0x1d,
	0xd2, 0xe4, 0xc6, 0x6e, 0x66, 0x91, 0x39, 0x59, 0xec, 0x6e, 0x20, 0xc4, 0x0d, 0xd7, 0x5c, 0x85,
	0x07, 0xe0, 0x12, 0x5e, 0x81, 0xe7, 0xe0, 0x15, 0xce, 0x83, 0x1c, 0xc5, 0x4e, 0xaa, 0x64, 0x3d,
	0xe7, 0xee, 0xe8, 0x5c, 0x45, 0xfe, 0xbe, 0xcf, 0xff, 0xf7, 0xfb, 0x8b, 0x7f, 0x03, 0x9b, 0x84,
	0x61, 0xb2, 0x14, 0x4a, 0x9e, 0xa4, 0x59, 0xa2, 0x12, 0xd8, 0x26, 0x29, 0x77, 0x77, 0xa2, 0x38,
	0x99, 0x93, 0xb8, 0xc4, 0xdc, 0x7e, 0x94, 0x25, 0xcb, 0xb4, 0x5a, 0x7d, 0x14, 0x25, 0x49, 0x14,
	0xb3, 0x31, 0x49, 0xf9, 0x98, 0x08, 0x91, 0x28, 0xa2, 0x78, 0x22, 0x2a, 0xf6, 0x48, 0x7f, 0xc2,
	0xe3, 0x88, 0x89, 0x63, 0xf9, 0x4c, 0xa2, 0x88, 0x65, 0xe3, 0x24, 0xd5, 0x8a, 0x75, 0xb5, 0xf7,
	0x13, 0x80, 0xbe, 0xf1, 0xbf, 0xe2, 0x52, 0x05, 0xec, 0x71, 0xc9, 0xa4, 0x82, 0x9f, 0x02, 0xe3,
	0x78, 0xb7, 0xe0, 0xb1, 0x62, 0x99, 0x63, 0x21, 0x6b, 0xd4, 0x09, 0x7a, 0x1a, 0xfb, 0x4e, 0x43,
	0xf0, 0x33, 0xd0, 0x4d, 0x49, 0xc4, 0x45, 0xe4, 0xb4, 0x90, 0x35, 0xea, 0x9d, 0xf6, 0x4e, 0x48,
	0xca, 0x4f, 0x7e, 0xd4, 0x50, 0x50, 0x52, 0xde, 0xb7, 0x60, 0xab, 0xac, 0x2e, 0xe1, 0x08, 0x6c,
	0x55, 0x27, 0x75, 0x2c, 0xd4, 0x1e, 0xf5, 0x4e, 0xfb, 0x7a, 0x4b, 0x29, 0x08, 0x56, 0xec, 0x64,
	0x37, 0xf7, 0xfb, 0x00, 0xe0, 0xd5, 0x56, 0xef, 0xbf, 0x16, 0xd8, 0x2c, 0x17, 0xf0, 0x13, 0xd0,
	0xe2, 0xd4, 0x34, 0x34, 0x2d, 0x84, 0x18, 0x94, 0x0c, 0x9a, 0x5d, 0x04, 0x2d, 0x4e, 0xe1, 0x21,
	0xd8, 0x10, 0xe4, 0x81, 0xe9, 0xb6, 0xb6, 0xa7, 0x83, 0xdc, 0xb7, 0x71, 0xbf, 0x92, 0x14, 0x44,
	0xa0, 0x69, 0x38, 0x01, 0x3d, 0xca, 0x64, 0x98, 0x71, 0x1d, 0x8d, 0xd3, 0xd6, 0x6a, 0x27, 0xf7,
	0x87, 0x78, 0xbf, 0x52, 0xd7, 0xf8, 0xa0, 0x2e, 0x86, 0x23, 0xd0, 0x91, 0x24, 0xa6, 0x89, 0xb3,
	0x81, 0xac, 0x91, 0x35, 0x85, 0xb9, 0xbf, 0x8b, 0x77, 0xaa, 0x5d, 0x37, 0x05, 0x13, 0x18, 0x41,
	0xe1, 0x22, 0x16, 0xe1, 0x5d, 0x78, 0xcf, 0xd3, 0x3b, 0x4e, 0x9d, 0x8e, 0x76, 0x71, 0x73, 0xff,
	0x00, 0x0f, 0x2b, 0xfd, 0xf5, 0x22, 0x44, 0xe7, 0xf7, 0x3c, 0x45, 0xb7, 0x4b, 0x4e, 0x83, 0x6d,
	0xb1, 0x08, 0x8b, 0xd5, 0x8c, 0xc2, 0xaf, 0x40, 0x47, 0x07, 0xee, 0x74, 0x75, 0xc0, 0x40, 0xa7,
	0x75, 0x59, 0x20, 0x2f, 0x1c, 0x35, 0x16, 0x18, 0xf1, 0xc4, 0xce, 0xfd, 0x1e, 0xd8, 0xc6, 0x55,
	0x5e, 0xde, 0x3f, 0x2d, 0x50, 0x09, 0xcf, 0x33, 0x46, 0x14, 0x7b, 0x1f, 0x01, 0x1d, 0x37, 0x03,
	0x3a, 0xc8, 0xfd, 0x0f, 0x30, 0x5c, 0x05, 0xa4, 0x48, 0xa6, 0xe4, 0xbb, 0x4a, 0xe9, 0x0b, 0xb0,
	0x65, 0xae, 0x2a, 0xa7, 0x3a, 0xa8, 0xce, 0x74, 0x98, 0xfb, 0x10, 0xef, 0x35, 0xc2, 0x29, 0xee,
	0xc6, 0xa6, 0x96, 0xcd, 0xe8, 0xe4, 0xc3, 0xdc, 0xdf, 0x07, 0x03, 0xbc, 0x5b, 0x4f, 0x85, 0x27,
	0xe2, 0xf4, 0xef, 0x0e, 0xb0, 0x4b, 0xec, 0x86, 0x65, 0x4f, 0x3c, 0x64, 0xf0, 0x5f, 0x0b, 0x6c,
	0x14, 0x73, 0x01, 0x0f, 0xea, 0x57, 0xb5, 0x36, 0x29, 0xee, 0x4e, 0x9d, 0x90, 0xde, 0x2f, 0xb9,
	0x4f, 0xdc, 0xab, 0x80, 0xa9, 0x65, 0x26, 0x24, 0x22, 0x71, 0x8c, 0xca, 0x5b, 0x7d, 0x84, 0x42,
	0x22, 0xd0, 0x9c, 0xa1, 0x98, 0x3f, 0x70, 0xc5, 0x28, 0x7a, 0xe6, 0xea, 0x1e, 0xa5, 0x24, 0x62,
	0x14, 0x95, 0xd3, 0x89, 0x88, 0xa0, 0xc8, 0x0c, 0x1c, 0xa3, 0x68, 0xfe, 0x2b, 0xd2, 0x0d, 0xe3,
	0x41, 0x61, 0x57, 0x2f, 0x25, 0xff, 0xf8, 0xff, 0xd5, 0x5f, 0x2d, 0x1b, 0xf6, 0xc7, 0x4f, 0x5f,
	0x8e, 0x2b, 0x0c, 0xfe, 0x0c, 0xba, 0xe5, 0xff, 0x85, 0xf5, 0x96, 0x0c, 0xe6, 0x36, 0x46, 0xcd,
	0xfb, 0x26, 0xf7, 0x0f, 0xdd, 0x7d, 0x43, 0x49, 0x24, 0xd8, 0x73, 0x55, 0x1a, 0xdb, 0x06, 0xac,
	0xd6, 0xda, 0x69, 0xe0, 0x35, 0x9c, 0x26, 0x16, 0x86, 0x12, 0xb4, 0x2f, 0x99, 0x82, 0xb6, 0xae,
	0x3a, 0xa3, 0x55, 0x18, 0x4d, 0x97, 0xeb, 0xdc, 0x3f, 0x73, 0x3f, 0xaf, 0xb2, 0x90, 0x5c, 0x44,
	0xf1, 0xaa, 0xb0, 0x39, 0x7e, 0xc4, 0x9f, 0x98, 0x40, 0x9c, 0xe2, 0xde, 0x25, 0x53, 0x0d, 0x4f,
	0x08, 0xf7, 0x6a, 0x9e, 0xe3, 0xdf, 0x38, 0xfd, 0x1d, 0xfe, 0x69, 0x81, 0xee, 0x6d, 0x4a, 0x8b,
	0x23, 0x36, 0x8c, 0x5e, 0xd8, 0x2e, 0x72, 0xff, 0x7b, 0xf7, 0x6b, 0x23, 0x94, 0x6f, 0xf6, 0x3b,
	0xd2, 0x71, 0x2e, 0x38, 0x8b, 0xa9, 0x44, 0x0f, 0x4b, 0xa9, 0x8a, 0x1f, 0x23, 0x99, 0xa0, 0xd8,
	0x36, 0xfb, 0x1a, 0xbd, 0x0c, 0xdd, 0xb5, 0x5e, 0x8a, 0x0c, 0x1e, 0x41, 0xf7, 0x82, 0xc5, 0x4c,
	0xb1, 0xb5, 0x18, 0xcc, 0x53, 0x78, 0xa3, 0x88, 0x5a, 0x4a, 0xef, 0x87, 0xdc, 0x1f, 0xbb, 0x1f,
	0x1b, 0xe5, 0x5b, 0xda, 0xc1, 0xb6, 0xa1, 0x9b, 0x09, 0xe0, 0x35, 0xd7, 0x79, 0x57, 0xbf, 0xd9,
	0x67, 0xaf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x96, 0xab, 0xad, 0x46, 0x33, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccountServiceClient is the client API for AccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountServiceClient interface {
	List(ctx context.Context, in *AccountListRequest, opts ...grpc.CallOption) (*Accounts, error)
	Create(ctx context.Context, in *AccountCreate, opts ...grpc.CallOption) (*Account, error)
	Get(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Account, error)
	Update(ctx context.Context, in *Account, opts ...grpc.CallOption) (*Account, error)
	Delete(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Status, error)
}

type accountServiceClient struct {
	cc *grpc.ClientConn
}

func NewAccountServiceClient(cc *grpc.ClientConn) AccountServiceClient {
	return &accountServiceClient{cc}
}

func (c *accountServiceClient) List(ctx context.Context, in *AccountListRequest, opts ...grpc.CallOption) (*Accounts, error) {
	out := new(Accounts)
	err := c.cc.Invoke(ctx, "/api.AccountService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) Create(ctx context.Context, in *AccountCreate, opts ...grpc.CallOption) (*Account, error) {
	out := new(Account)
	err := c.cc.Invoke(ctx, "/api.AccountService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) Get(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Account, error) {
	out := new(Account)
	err := c.cc.Invoke(ctx, "/api.AccountService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) Update(ctx context.Context, in *Account, opts ...grpc.CallOption) (*Account, error) {
	out := new(Account)
	err := c.cc.Invoke(ctx, "/api.AccountService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) Delete(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/api.AccountService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServiceServer is the server API for AccountService service.
type AccountServiceServer interface {
	List(context.Context, *AccountListRequest) (*Accounts, error)
	Create(context.Context, *AccountCreate) (*Account, error)
	Get(context.Context, *IdRequest) (*Account, error)
	Update(context.Context, *Account) (*Account, error)
	Delete(context.Context, *IdRequest) (*Status, error)
}

// UnimplementedAccountServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAccountServiceServer struct {
}

func (*UnimplementedAccountServiceServer) List(ctx context.Context, req *AccountListRequest) (*Accounts, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedAccountServiceServer) Create(ctx context.Context, req *AccountCreate) (*Account, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedAccountServiceServer) Get(ctx context.Context, req *IdRequest) (*Account, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedAccountServiceServer) Update(ctx context.Context, req *Account) (*Account, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedAccountServiceServer) Delete(ctx context.Context, req *IdRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterAccountServiceServer(s *grpc.Server, srv AccountServiceServer) {
	s.RegisterService(&_AccountService_serviceDesc, srv)
}

func _AccountService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AccountService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).List(ctx, req.(*AccountListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountCreate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AccountService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).Create(ctx, req.(*AccountCreate))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AccountService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).Get(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Account)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AccountService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).Update(ctx, req.(*Account))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AccountService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).Delete(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccountService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.AccountService",
	HandlerType: (*AccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _AccountService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _AccountService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _AccountService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _AccountService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AccountService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accounts.proto",
}
