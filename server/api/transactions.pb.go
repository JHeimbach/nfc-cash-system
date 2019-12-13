// Code generated by protoc-gen-go. DO NOT EDIT.
// source: transactions.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type ListTransactionRequest struct {
	Paging               *Paging  `protobuf:"bytes,1,opt,name=paging,proto3" json:"paging,omitempty"`
	Order                string   `protobuf:"bytes,2,opt,name=order,proto3" json:"order,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTransactionRequest) Reset()         { *m = ListTransactionRequest{} }
func (m *ListTransactionRequest) String() string { return proto.CompactTextString(m) }
func (*ListTransactionRequest) ProtoMessage()    {}
func (*ListTransactionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b72849cf10e9c77, []int{0}
}

func (m *ListTransactionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTransactionRequest.Unmarshal(m, b)
}
func (m *ListTransactionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTransactionRequest.Marshal(b, m, deterministic)
}
func (m *ListTransactionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTransactionRequest.Merge(m, src)
}
func (m *ListTransactionRequest) XXX_Size() int {
	return xxx_messageInfo_ListTransactionRequest.Size(m)
}
func (m *ListTransactionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTransactionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListTransactionRequest proto.InternalMessageInfo

func (m *ListTransactionRequest) GetPaging() *Paging {
	if m != nil {
		return m.Paging
	}
	return nil
}

func (m *ListTransactionRequest) GetOrder() string {
	if m != nil {
		return m.Order
	}
	return ""
}

type ListTransactionsByAccountRequest struct {
	AccountId            int32    `protobuf:"varint,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Paging               *Paging  `protobuf:"bytes,2,opt,name=paging,proto3" json:"paging,omitempty"`
	Order                string   `protobuf:"bytes,3,opt,name=order,proto3" json:"order,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTransactionsByAccountRequest) Reset()         { *m = ListTransactionsByAccountRequest{} }
func (m *ListTransactionsByAccountRequest) String() string { return proto.CompactTextString(m) }
func (*ListTransactionsByAccountRequest) ProtoMessage()    {}
func (*ListTransactionsByAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b72849cf10e9c77, []int{1}
}

func (m *ListTransactionsByAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTransactionsByAccountRequest.Unmarshal(m, b)
}
func (m *ListTransactionsByAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTransactionsByAccountRequest.Marshal(b, m, deterministic)
}
func (m *ListTransactionsByAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTransactionsByAccountRequest.Merge(m, src)
}
func (m *ListTransactionsByAccountRequest) XXX_Size() int {
	return xxx_messageInfo_ListTransactionsByAccountRequest.Size(m)
}
func (m *ListTransactionsByAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTransactionsByAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListTransactionsByAccountRequest proto.InternalMessageInfo

func (m *ListTransactionsByAccountRequest) GetAccountId() int32 {
	if m != nil {
		return m.AccountId
	}
	return 0
}

func (m *ListTransactionsByAccountRequest) GetPaging() *Paging {
	if m != nil {
		return m.Paging
	}
	return nil
}

func (m *ListTransactionsByAccountRequest) GetOrder() string {
	if m != nil {
		return m.Order
	}
	return ""
}

type GetTransactionRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AccountId            int32    `protobuf:"varint,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTransactionRequest) Reset()         { *m = GetTransactionRequest{} }
func (m *GetTransactionRequest) String() string { return proto.CompactTextString(m) }
func (*GetTransactionRequest) ProtoMessage()    {}
func (*GetTransactionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b72849cf10e9c77, []int{2}
}

func (m *GetTransactionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTransactionRequest.Unmarshal(m, b)
}
func (m *GetTransactionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTransactionRequest.Marshal(b, m, deterministic)
}
func (m *GetTransactionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTransactionRequest.Merge(m, src)
}
func (m *GetTransactionRequest) XXX_Size() int {
	return xxx_messageInfo_GetTransactionRequest.Size(m)
}
func (m *GetTransactionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTransactionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTransactionRequest proto.InternalMessageInfo

func (m *GetTransactionRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GetTransactionRequest) GetAccountId() int32 {
	if m != nil {
		return m.AccountId
	}
	return 0
}

type ListTransactionsResponse struct {
	Transactions         []*Transaction `protobuf:"bytes,1,rep,name=transactions,proto3" json:"transactions,omitempty"`
	TotalCount           int32          `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ListTransactionsResponse) Reset()         { *m = ListTransactionsResponse{} }
func (m *ListTransactionsResponse) String() string { return proto.CompactTextString(m) }
func (*ListTransactionsResponse) ProtoMessage()    {}
func (*ListTransactionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b72849cf10e9c77, []int{3}
}

func (m *ListTransactionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTransactionsResponse.Unmarshal(m, b)
}
func (m *ListTransactionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTransactionsResponse.Marshal(b, m, deterministic)
}
func (m *ListTransactionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTransactionsResponse.Merge(m, src)
}
func (m *ListTransactionsResponse) XXX_Size() int {
	return xxx_messageInfo_ListTransactionsResponse.Size(m)
}
func (m *ListTransactionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTransactionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListTransactionsResponse proto.InternalMessageInfo

func (m *ListTransactionsResponse) GetTransactions() []*Transaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

func (m *ListTransactionsResponse) GetTotalCount() int32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

type Transaction struct {
	Id                   int32                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OldSaldo             float64              `protobuf:"fixed64,2,opt,name=old_saldo,json=oldSaldo,proto3" json:"old_saldo,omitempty"`
	NewSaldo             float64              `protobuf:"fixed64,3,opt,name=new_saldo,json=newSaldo,proto3" json:"new_saldo,omitempty"`
	Amount               float64              `protobuf:"fixed64,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Created              *timestamp.Timestamp `protobuf:"bytes,5,opt,name=created,proto3" json:"created,omitempty"`
	Account              *Account             `protobuf:"bytes,6,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b72849cf10e9c77, []int{4}
}

func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transaction.Unmarshal(m, b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
}
func (m *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(m, src)
}
func (m *Transaction) XXX_Size() int {
	return xxx_messageInfo_Transaction.Size(m)
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Transaction) GetOldSaldo() float64 {
	if m != nil {
		return m.OldSaldo
	}
	return 0
}

func (m *Transaction) GetNewSaldo() float64 {
	if m != nil {
		return m.NewSaldo
	}
	return 0
}

func (m *Transaction) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Transaction) GetCreated() *timestamp.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *Transaction) GetAccount() *Account {
	if m != nil {
		return m.Account
	}
	return nil
}

type CreateTransactionRequest struct {
	OldSaldo             float64  `protobuf:"fixed64,1,opt,name=old_saldo,json=oldSaldo,proto3" json:"old_saldo,omitempty"`
	NewSaldo             float64  `protobuf:"fixed64,2,opt,name=new_saldo,json=newSaldo,proto3" json:"new_saldo,omitempty"`
	Amount               float64  `protobuf:"fixed64,3,opt,name=amount,proto3" json:"amount,omitempty"`
	AccountId            int32    `protobuf:"varint,4,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTransactionRequest) Reset()         { *m = CreateTransactionRequest{} }
func (m *CreateTransactionRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTransactionRequest) ProtoMessage()    {}
func (*CreateTransactionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b72849cf10e9c77, []int{5}
}

func (m *CreateTransactionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTransactionRequest.Unmarshal(m, b)
}
func (m *CreateTransactionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTransactionRequest.Marshal(b, m, deterministic)
}
func (m *CreateTransactionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTransactionRequest.Merge(m, src)
}
func (m *CreateTransactionRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTransactionRequest.Size(m)
}
func (m *CreateTransactionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTransactionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTransactionRequest proto.InternalMessageInfo

func (m *CreateTransactionRequest) GetOldSaldo() float64 {
	if m != nil {
		return m.OldSaldo
	}
	return 0
}

func (m *CreateTransactionRequest) GetNewSaldo() float64 {
	if m != nil {
		return m.NewSaldo
	}
	return 0
}

func (m *CreateTransactionRequest) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *CreateTransactionRequest) GetAccountId() int32 {
	if m != nil {
		return m.AccountId
	}
	return 0
}

func init() {
	proto.RegisterType((*ListTransactionRequest)(nil), "api.ListTransactionRequest")
	proto.RegisterType((*ListTransactionsByAccountRequest)(nil), "api.ListTransactionsByAccountRequest")
	proto.RegisterType((*GetTransactionRequest)(nil), "api.GetTransactionRequest")
	proto.RegisterType((*ListTransactionsResponse)(nil), "api.ListTransactionsResponse")
	proto.RegisterType((*Transaction)(nil), "api.Transaction")
	proto.RegisterType((*CreateTransactionRequest)(nil), "api.CreateTransactionRequest")
}

func init() { proto.RegisterFile("transactions.proto", fileDescriptor_0b72849cf10e9c77) }

var fileDescriptor_0b72849cf10e9c77 = []byte{
	// 710 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x41, 0x4f, 0x13, 0x41,
	0x14, 0xce, 0xb4, 0x50, 0x64, 0x0a, 0xa5, 0x0c, 0x02, 0x75, 0x91, 0x30, 0x59, 0x83, 0x92, 0x0d,
	0x74, 0x23, 0x72, 0xea, 0xad, 0x92, 0x48, 0x4c, 0x38, 0x98, 0x2d, 0x77, 0x9c, 0x76, 0x87, 0x75,
	0x92, 0xed, 0xcc, 0xba, 0x33, 0xa5, 0x31, 0xc8, 0xc5, 0x83, 0x77, 0xeb, 0xd5, 0x93, 0x07, 0xfd,
	0x17, 0xfe, 0x09, 0x7f, 0x81, 0x89, 0x7f, 0xc1, 0xbb, 0xe9, 0xec, 0x16, 0x66, 0xbb, 0x8b, 0xf6,
	0xb4, 0x99, 0xf7, 0xbd, 0x79, 0xdf, 0xf7, 0xcd, 0x7b, 0x6f, 0x21, 0x52, 0x31, 0xe1, 0x92, 0xf4,
	0x14, 0x13, 0x5c, 0x36, 0xa3, 0x58, 0x28, 0x81, 0xca, 0x24, 0x62, 0xd6, 0x72, 0x10, 0x8a, 0x2e,
	0x09, 0xd3, 0x98, 0x55, 0x23, 0xbd, 0x9e, 0x18, 0x70, 0x35, 0x39, 0xef, 0x04, 0x42, 0x04, 0x21,
	0x75, 0xf5, 0xa9, 0x3b, 0xb8, 0x70, 0x15, 0xeb, 0x53, 0xa9, 0x48, 0x3f, 0x4a, 0x13, 0x1e, 0xa6,
	0x09, 0x24, 0x62, 0x2e, 0xe1, 0x5c, 0x28, 0x62, 0x50, 0x58, 0xfb, 0xfa, 0xd3, 0x3b, 0x08, 0x28,
	0x3f, 0x90, 0x43, 0x12, 0x04, 0x34, 0x76, 0x45, 0xa4, 0x33, 0xf2, 0xd9, 0x76, 0x07, 0x6e, 0x9c,
	0x32, 0xa9, 0xce, 0x6e, 0xa5, 0x7a, 0xf4, 0xed, 0x80, 0x4a, 0x85, 0x1e, 0xc1, 0x4a, 0x44, 0x02,
	0xc6, 0x83, 0x06, 0xc0, 0x60, 0xaf, 0x7a, 0x58, 0x6d, 0x92, 0x88, 0x35, 0x5f, 0xe9, 0x90, 0x97,
	0x42, 0xe8, 0x3e, 0x9c, 0x17, 0xb1, 0x4f, 0xe3, 0x46, 0x09, 0x83, 0xbd, 0x45, 0x2f, 0x39, 0xd8,
	0xef, 0x21, 0x9e, 0x2a, 0x2a, 0x9f, 0xbf, 0x6b, 0x27, 0x2e, 0x27, 0xe5, 0xb7, 0x21, 0x4c, 0x7d,
	0x9f, 0x33, 0x5f, 0x53, 0xcc, 0x7b, 0x8b, 0x69, 0xe4, 0xa5, 0x6f, 0xb0, 0x97, 0x66, 0x60, 0x2f,
	0x9b, 0xec, 0x2f, 0xe0, 0xfa, 0x09, 0x2d, 0x72, 0x54, 0x83, 0xa5, 0x1b, 0xaa, 0x12, 0xf3, 0xa7,
	0x24, 0x94, 0xa6, 0x24, 0xd8, 0x1f, 0x01, 0x6c, 0x4c, 0xdb, 0xf0, 0xa8, 0x8c, 0x04, 0x97, 0x14,
	0x1d, 0xc1, 0x25, 0xb3, 0xbd, 0x0d, 0x80, 0xcb, 0x7b, 0xd5, 0xc3, 0xba, 0x56, 0x69, 0x52, 0x67,
	0xb2, 0xd0, 0x0e, 0xac, 0x2a, 0xa1, 0x48, 0x78, 0xae, 0x39, 0x52, 0x4a, 0xa8, 0x43, 0xc7, 0xe3,
	0x48, 0x6b, 0x6d, 0xd4, 0xae, 0xc3, 0x9a, 0xb3, 0x64, 0x72, 0xda, 0xbf, 0x00, 0xac, 0x1a, 0x81,
	0x9c, 0x8f, 0x2d, 0xb8, 0x28, 0x42, 0xff, 0x5c, 0x92, 0xd0, 0x17, 0xba, 0x26, 0xf0, 0xee, 0x89,
	0xd0, 0xef, 0x8c, 0xcf, 0x63, 0x90, 0xd3, 0x61, 0x0a, 0x96, 0x13, 0x90, 0xd3, 0x61, 0x02, 0x6e,
	0xc0, 0x0a, 0xe9, 0x6b, 0x29, 0x73, 0x1a, 0x49, 0x4f, 0xe8, 0x08, 0x2e, 0xf4, 0x62, 0x4a, 0x14,
	0xf5, 0x1b, 0xf3, 0xfa, 0xf9, 0xad, 0x66, 0x32, 0x73, 0xcd, 0xc9, 0x50, 0x36, 0xcf, 0x26, 0x43,
	0xe9, 0x4d, 0x52, 0xd1, 0x63, 0xb8, 0x90, 0xbe, 0x5e, 0xa3, 0xa2, 0x6f, 0x2d, 0xe9, 0xe7, 0x98,
	0x34, 0x7e, 0x02, 0xb6, 0xd0, 0xa8, 0xbd, 0x02, 0x97, 0x1d, 0xd3, 0x93, 0xfd, 0x1d, 0xc0, 0xc6,
	0xb1, 0xae, 0x53, 0xd0, 0xb8, 0x8c, 0x41, 0xf0, 0x2f, 0x83, 0xa5, 0x3b, 0x0d, 0x96, 0x33, 0x06,
	0xb3, 0xad, 0x9f, 0x9b, 0x6a, 0x7d, 0xcb, 0x1a, 0xb5, 0x37, 0xe1, 0xba, 0xb3, 0x66, 0x68, 0xd1,
	0xe2, 0x98, 0xe0, 0x87, 0x5f, 0x2a, 0xd0, 0x8c, 0xcb, 0x0e, 0x8d, 0x2f, 0x59, 0x8f, 0xa2, 0x1f,
	0x00, 0xd6, 0xa7, 0xc7, 0x05, 0x6d, 0xe9, 0x17, 0x28, 0xde, 0x30, 0x6b, 0xbb, 0x08, 0xbc, 0x19,
	0x31, 0x9b, 0x8f, 0xda, 0x1d, 0xab, 0x35, 0x86, 0x25, 0x26, 0x61, 0x88, 0xcd, 0x49, 0xda, 0xc7,
	0x3d, 0xc2, 0x71, 0x97, 0xe2, 0x90, 0xf5, 0x99, 0xa2, 0x3e, 0x1e, 0x32, 0xf5, 0x06, 0x27, 0x6b,
	0x81, 0xd3, 0x6d, 0x77, 0xd6, 0xc7, 0x77, 0x73, 0x57, 0x3f, 0xfc, 0xfc, 0xfd, 0xb9, 0x84, 0x50,
	0xdd, 0xbd, 0x7c, 0xea, 0x66, 0x86, 0xf3, 0x0f, 0x80, 0x0f, 0xee, 0x5c, 0x5b, 0xb4, 0x5b, 0x28,
	0x76, 0x7a, 0xad, 0xff, 0xe7, 0xe9, 0x13, 0x18, 0xb5, 0x89, 0x75, 0x7a, 0x6b, 0xca, 0xcc, 0xc2,
	0x17, 0x22, 0xc6, 0x01, 0xbb, 0xa4, 0x1c, 0xa7, 0xad, 0x98, 0xc9, 0xe6, 0xaa, 0xb6, 0x99, 0xb3,
	0xf8, 0x04, 0xed, 0x8e, 0x2d, 0xa6, 0x95, 0xdc, 0xab, 0xdb, 0x7e, 0x5f, 0x67, 0x7d, 0x7f, 0x03,
	0x70, 0x35, 0x37, 0x7a, 0x28, 0x31, 0x72, 0xd7, 0x48, 0x5a, 0xb9, 0x4d, 0xb7, 0x5f, 0x8f, 0xda,
	0x07, 0xd6, 0x66, 0x72, 0x41, 0x62, 0x4e, 0x87, 0xa6, 0x24, 0x07, 0x25, 0x80, 0x19, 0xd3, 0x2a,
	0x1d, 0x7b, 0x36, 0x95, 0x2d, 0xe0, 0xa0, 0xaf, 0x00, 0xd6, 0xb2, 0x7f, 0x36, 0x64, 0x69, 0x19,
	0x85, 0xbf, 0xbb, 0x02, 0x89, 0xdd, 0xb1, 0x44, 0xcb, 0xa3, 0x6a, 0x10, 0x73, 0x89, 0x25, 0xe3,
	0x41, 0x98, 0x51, 0xe4, 0xac, 0x9c, 0x50, 0x95, 0x93, 0xb8, 0x8f, 0x9c, 0x99, 0x24, 0xba, 0x57,
	0xcc, 0xbf, 0xee, 0x56, 0xf4, 0x1f, 0xe2, 0xd9, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x82, 0x6b,
	0x6c, 0x45, 0xfe, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TransactionsServiceClient is the client API for TransactionsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TransactionsServiceClient interface {
	ListTransactions(ctx context.Context, in *ListTransactionRequest, opts ...grpc.CallOption) (*ListTransactionsResponse, error)
	ListTransactionsByAccount(ctx context.Context, in *ListTransactionsByAccountRequest, opts ...grpc.CallOption) (*ListTransactionsResponse, error)
	CreateTransaction(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*Transaction, error)
	GetTransaction(ctx context.Context, in *GetTransactionRequest, opts ...grpc.CallOption) (*Transaction, error)
}

type transactionsServiceClient struct {
	cc *grpc.ClientConn
}

func NewTransactionsServiceClient(cc *grpc.ClientConn) TransactionsServiceClient {
	return &transactionsServiceClient{cc}
}

func (c *transactionsServiceClient) ListTransactions(ctx context.Context, in *ListTransactionRequest, opts ...grpc.CallOption) (*ListTransactionsResponse, error) {
	out := new(ListTransactionsResponse)
	err := c.cc.Invoke(ctx, "/api.TransactionsService/ListTransactions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionsServiceClient) ListTransactionsByAccount(ctx context.Context, in *ListTransactionsByAccountRequest, opts ...grpc.CallOption) (*ListTransactionsResponse, error) {
	out := new(ListTransactionsResponse)
	err := c.cc.Invoke(ctx, "/api.TransactionsService/ListTransactionsByAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionsServiceClient) CreateTransaction(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*Transaction, error) {
	out := new(Transaction)
	err := c.cc.Invoke(ctx, "/api.TransactionsService/CreateTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionsServiceClient) GetTransaction(ctx context.Context, in *GetTransactionRequest, opts ...grpc.CallOption) (*Transaction, error) {
	out := new(Transaction)
	err := c.cc.Invoke(ctx, "/api.TransactionsService/GetTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionsServiceServer is the server API for TransactionsService service.
type TransactionsServiceServer interface {
	ListTransactions(context.Context, *ListTransactionRequest) (*ListTransactionsResponse, error)
	ListTransactionsByAccount(context.Context, *ListTransactionsByAccountRequest) (*ListTransactionsResponse, error)
	CreateTransaction(context.Context, *CreateTransactionRequest) (*Transaction, error)
	GetTransaction(context.Context, *GetTransactionRequest) (*Transaction, error)
}

// UnimplementedTransactionsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTransactionsServiceServer struct {
}

func (*UnimplementedTransactionsServiceServer) ListTransactions(ctx context.Context, req *ListTransactionRequest) (*ListTransactionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTransactions not implemented")
}
func (*UnimplementedTransactionsServiceServer) ListTransactionsByAccount(ctx context.Context, req *ListTransactionsByAccountRequest) (*ListTransactionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTransactionsByAccount not implemented")
}
func (*UnimplementedTransactionsServiceServer) CreateTransaction(ctx context.Context, req *CreateTransactionRequest) (*Transaction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTransaction not implemented")
}
func (*UnimplementedTransactionsServiceServer) GetTransaction(ctx context.Context, req *GetTransactionRequest) (*Transaction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransaction not implemented")
}

func RegisterTransactionsServiceServer(s *grpc.Server, srv TransactionsServiceServer) {
	s.RegisterService(&_TransactionsService_serviceDesc, srv)
}

func _TransactionsService_ListTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionsServiceServer).ListTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.TransactionsService/ListTransactions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionsServiceServer).ListTransactions(ctx, req.(*ListTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionsService_ListTransactionsByAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTransactionsByAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionsServiceServer).ListTransactionsByAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.TransactionsService/ListTransactionsByAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionsServiceServer).ListTransactionsByAccount(ctx, req.(*ListTransactionsByAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionsService_CreateTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionsServiceServer).CreateTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.TransactionsService/CreateTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionsServiceServer).CreateTransaction(ctx, req.(*CreateTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionsService_GetTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionsServiceServer).GetTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.TransactionsService/GetTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionsServiceServer).GetTransaction(ctx, req.(*GetTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TransactionsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.TransactionsService",
	HandlerType: (*TransactionsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListTransactions",
			Handler:    _TransactionsService_ListTransactions_Handler,
		},
		{
			MethodName: "ListTransactionsByAccount",
			Handler:    _TransactionsService_ListTransactionsByAccount_Handler,
		},
		{
			MethodName: "CreateTransaction",
			Handler:    _TransactionsService_CreateTransaction_Handler,
		},
		{
			MethodName: "GetTransaction",
			Handler:    _TransactionsService_GetTransaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transactions.proto",
}
