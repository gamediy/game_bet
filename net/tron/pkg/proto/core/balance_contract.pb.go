// Code generated by protoc-gen-go. DO NOT EDIT.
// source: core/contract/balance_contract.proto

package core

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type FreezeBalanceContract struct {
	OwnerAddress         []byte       `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	FrozenBalance        int64        `protobuf:"varint,2,opt,name=frozen_balance,json=frozenBalance,proto3" json:"frozen_balance,omitempty"`
	FrozenDuration       int64        `protobuf:"varint,3,opt,name=frozen_duration,json=frozenDuration,proto3" json:"frozen_duration,omitempty"`
	Resource             ResourceCode `protobuf:"varint,10,opt,name=resource,proto3,enum=protocol.ResourceCode" json:"resource,omitempty"`
	ReceiverAddress      []byte       `protobuf:"bytes,15,opt,name=receiver_address,json=receiverAddress,proto3" json:"receiver_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *FreezeBalanceContract) Reset()         { *m = FreezeBalanceContract{} }
func (m *FreezeBalanceContract) String() string { return proto.CompactTextString(m) }
func (*FreezeBalanceContract) ProtoMessage()    {}
func (*FreezeBalanceContract) Descriptor() ([]byte, []int) {
	return fileDescriptor_54daa91fef75922f, []int{0}
}

func (m *FreezeBalanceContract) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FreezeBalanceContract.Unmarshal(m, b)
}
func (m *FreezeBalanceContract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FreezeBalanceContract.Marshal(b, m, deterministic)
}
func (m *FreezeBalanceContract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FreezeBalanceContract.Merge(m, src)
}
func (m *FreezeBalanceContract) XXX_Size() int {
	return xxx_messageInfo_FreezeBalanceContract.Size(m)
}
func (m *FreezeBalanceContract) XXX_DiscardUnknown() {
	xxx_messageInfo_FreezeBalanceContract.DiscardUnknown(m)
}

var xxx_messageInfo_FreezeBalanceContract proto.InternalMessageInfo

func (m *FreezeBalanceContract) GetOwnerAddress() []byte {
	if m != nil {
		return m.OwnerAddress
	}
	return nil
}

func (m *FreezeBalanceContract) GetFrozenBalance() int64 {
	if m != nil {
		return m.FrozenBalance
	}
	return 0
}

func (m *FreezeBalanceContract) GetFrozenDuration() int64 {
	if m != nil {
		return m.FrozenDuration
	}
	return 0
}

func (m *FreezeBalanceContract) GetResource() ResourceCode {
	if m != nil {
		return m.Resource
	}
	return ResourceCode_BANDWIDTH
}

func (m *FreezeBalanceContract) GetReceiverAddress() []byte {
	if m != nil {
		return m.ReceiverAddress
	}
	return nil
}

type UnfreezeBalanceContract struct {
	OwnerAddress         []byte       `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	Resource             ResourceCode `protobuf:"varint,10,opt,name=resource,proto3,enum=protocol.ResourceCode" json:"resource,omitempty"`
	ReceiverAddress      []byte       `protobuf:"bytes,15,opt,name=receiver_address,json=receiverAddress,proto3" json:"receiver_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UnfreezeBalanceContract) Reset()         { *m = UnfreezeBalanceContract{} }
func (m *UnfreezeBalanceContract) String() string { return proto.CompactTextString(m) }
func (*UnfreezeBalanceContract) ProtoMessage()    {}
func (*UnfreezeBalanceContract) Descriptor() ([]byte, []int) {
	return fileDescriptor_54daa91fef75922f, []int{1}
}

func (m *UnfreezeBalanceContract) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnfreezeBalanceContract.Unmarshal(m, b)
}
func (m *UnfreezeBalanceContract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnfreezeBalanceContract.Marshal(b, m, deterministic)
}
func (m *UnfreezeBalanceContract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnfreezeBalanceContract.Merge(m, src)
}
func (m *UnfreezeBalanceContract) XXX_Size() int {
	return xxx_messageInfo_UnfreezeBalanceContract.Size(m)
}
func (m *UnfreezeBalanceContract) XXX_DiscardUnknown() {
	xxx_messageInfo_UnfreezeBalanceContract.DiscardUnknown(m)
}

var xxx_messageInfo_UnfreezeBalanceContract proto.InternalMessageInfo

func (m *UnfreezeBalanceContract) GetOwnerAddress() []byte {
	if m != nil {
		return m.OwnerAddress
	}
	return nil
}

func (m *UnfreezeBalanceContract) GetResource() ResourceCode {
	if m != nil {
		return m.Resource
	}
	return ResourceCode_BANDWIDTH
}

func (m *UnfreezeBalanceContract) GetReceiverAddress() []byte {
	if m != nil {
		return m.ReceiverAddress
	}
	return nil
}

type WithdrawBalanceContract struct {
	OwnerAddress         []byte   `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WithdrawBalanceContract) Reset()         { *m = WithdrawBalanceContract{} }
func (m *WithdrawBalanceContract) String() string { return proto.CompactTextString(m) }
func (*WithdrawBalanceContract) ProtoMessage()    {}
func (*WithdrawBalanceContract) Descriptor() ([]byte, []int) {
	return fileDescriptor_54daa91fef75922f, []int{2}
}

func (m *WithdrawBalanceContract) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WithdrawBalanceContract.Unmarshal(m, b)
}
func (m *WithdrawBalanceContract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WithdrawBalanceContract.Marshal(b, m, deterministic)
}
func (m *WithdrawBalanceContract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WithdrawBalanceContract.Merge(m, src)
}
func (m *WithdrawBalanceContract) XXX_Size() int {
	return xxx_messageInfo_WithdrawBalanceContract.Size(m)
}
func (m *WithdrawBalanceContract) XXX_DiscardUnknown() {
	xxx_messageInfo_WithdrawBalanceContract.DiscardUnknown(m)
}

var xxx_messageInfo_WithdrawBalanceContract proto.InternalMessageInfo

func (m *WithdrawBalanceContract) GetOwnerAddress() []byte {
	if m != nil {
		return m.OwnerAddress
	}
	return nil
}

type TransferContract struct {
	OwnerAddress         []byte   `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	ToAddress            []byte   `protobuf:"bytes,2,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty"`
	Amount               int64    `protobuf:"varint,3,opt,name=order,proto3" json:"order,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransferContract) Reset()         { *m = TransferContract{} }
func (m *TransferContract) String() string { return proto.CompactTextString(m) }
func (*TransferContract) ProtoMessage()    {}
func (*TransferContract) Descriptor() ([]byte, []int) {
	return fileDescriptor_54daa91fef75922f, []int{3}
}

func (m *TransferContract) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferContract.Unmarshal(m, b)
}
func (m *TransferContract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferContract.Marshal(b, m, deterministic)
}
func (m *TransferContract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferContract.Merge(m, src)
}
func (m *TransferContract) XXX_Size() int {
	return xxx_messageInfo_TransferContract.Size(m)
}
func (m *TransferContract) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferContract.DiscardUnknown(m)
}

var xxx_messageInfo_TransferContract proto.InternalMessageInfo

func (m *TransferContract) GetOwnerAddress() []byte {
	if m != nil {
		return m.OwnerAddress
	}
	return nil
}

func (m *TransferContract) GetToAddress() []byte {
	if m != nil {
		return m.ToAddress
	}
	return nil
}

func (m *TransferContract) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func init() {
	proto.RegisterType((*FreezeBalanceContract)(nil), "protocol.FreezeBalanceContract")
	proto.RegisterType((*UnfreezeBalanceContract)(nil), "protocol.UnfreezeBalanceContract")
	proto.RegisterType((*WithdrawBalanceContract)(nil), "protocol.WithdrawBalanceContract")
	proto.RegisterType((*TransferContract)(nil), "protocol.TransferContract")
}

func init() {
	proto.RegisterFile("core/contract/balance_contract.proto", fileDescriptor_54daa91fef75922f)
}

var fileDescriptor_54daa91fef75922f = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0x87, 0x49, 0x0b, 0xa5, 0x0e, 0xfd, 0x47, 0xc0, 0x36, 0x14, 0x84, 0x5a, 0x15, 0xeb, 0x25,
	0x81, 0x7a, 0x56, 0xb0, 0x15, 0x1f, 0x20, 0x28, 0x82, 0x97, 0xb2, 0xd9, 0x4c, 0xd2, 0x40, 0xb3,
	0x23, 0x93, 0x8d, 0x85, 0xbe, 0x8b, 0x4f, 0xe8, 0x4b, 0x48, 0x93, 0xdd, 0xa2, 0x47, 0x0b, 0x9e,
	0x42, 0x7e, 0xf3, 0xcd, 0xf0, 0x0d, 0x3b, 0x70, 0x29, 0x89, 0x31, 0x90, 0xa4, 0x34, 0x0b, 0xa9,
	0x83, 0x48, 0x6c, 0x84, 0x92, 0xb8, 0xb2, 0x81, 0xff, 0xce, 0xa4, 0xc9, 0x6d, 0x57, 0x1f, 0x49,
	0x9b, 0xf1, 0xf8, 0x37, 0x2f, 0x29, 0xcf, 0x49, 0xd5, 0xd4, 0xf4, 0xcb, 0x81, 0xd3, 0x27, 0x46,
	0xdc, 0xe1, 0xa2, 0x1e, 0xb3, 0x34, 0x98, 0x7b, 0x01, 0x5d, 0xda, 0x2a, 0xe4, 0x95, 0x88, 0x63,
	0xc6, 0xa2, 0xf0, 0x9c, 0x89, 0x33, 0xeb, 0x84, 0x9d, 0x2a, 0x7c, 0xa8, 0x33, 0xf7, 0x0a, 0x7a,
	0x09, 0xd3, 0x0e, 0xd5, 0xca, 0x58, 0x78, 0x8d, 0x89, 0x33, 0x6b, 0x86, 0xdd, 0x3a, 0x35, 0x33,
	0xdd, 0x6b, 0xe8, 0x1b, 0x2c, 0x2e, 0x59, 0xe8, 0x8c, 0x94, 0xd7, 0xac, 0x38, 0xd3, 0xfd, 0x68,
	0x52, 0x77, 0x0e, 0x6d, 0xc6, 0x82, 0x4a, 0x96, 0xe8, 0xc1, 0xc4, 0x99, 0xf5, 0xe6, 0x43, 0xdf,
	0xee, 0xe1, 0x87, 0xa6, 0xb2, 0xa4, 0x18, 0xc3, 0x03, 0xe7, 0xde, 0xc0, 0x80, 0x51, 0x62, 0xf6,
	0xf1, 0xc3, 0xb5, 0x5f, 0xb9, 0xf6, 0x6d, 0x6e, 0x74, 0xa7, 0x9f, 0x0e, 0x8c, 0x5e, 0x54, 0x72,
	0xfc, 0xbe, 0xff, 0xec, 0x77, 0x0f, 0xa3, 0xd7, 0x4c, 0xaf, 0x63, 0x16, 0xdb, 0x63, 0xf4, 0xa6,
	0x0a, 0x06, 0xcf, 0x2c, 0x54, 0x91, 0x20, 0xff, 0x6d, 0xaf, 0x33, 0x00, 0x4d, 0x07, 0xa2, 0x51,
	0x11, 0x27, 0x9a, 0x6c, 0x79, 0x08, 0x2d, 0x91, 0x53, 0xa9, 0xb4, 0x79, 0x36, 0xf3, 0xb7, 0xb8,
	0x03, 0x8f, 0x38, 0xf5, 0x35, 0xdb, 0x7b, 0x2a, 0x7c, 0x7b, 0x66, 0x6f, 0xe7, 0x69, 0xa6, 0xd7,
	0x65, 0xe4, 0x4b, 0xca, 0x83, 0x24, 0x2a, 0x28, 0x62, 0xcc, 0x58, 0x04, 0x29, 0xed, 0xe9, 0x60,
	0x7f, 0x92, 0x51, 0xab, 0xea, 0xb9, 0xfd, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xed, 0x9f, 0x5e, 0xa6,
	0xd1, 0x02, 0x00, 0x00,
}
