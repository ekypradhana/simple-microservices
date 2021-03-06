// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/transaction/service.proto

/*
Package transaction is a generated protocol buffer package.

It is generated from these files:
	rpc/transaction/service.proto

It has these top-level messages:
	TransactionModel
	ResultInfo
*/
package transaction

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TransactionModel struct {
	CustomerName string `protobuf:"bytes,1,opt,name=customer_name,json=customerName" json:"customer_name,omitempty"`
	ProductId    string `protobuf:"bytes,2,opt,name=product_id,json=productId" json:"product_id,omitempty"`
	OrderAmount  int32  `protobuf:"varint,3,opt,name=order_amount,json=orderAmount" json:"order_amount,omitempty"`
}

func (m *TransactionModel) Reset()                    { *m = TransactionModel{} }
func (m *TransactionModel) String() string            { return proto.CompactTextString(m) }
func (*TransactionModel) ProtoMessage()               {}
func (*TransactionModel) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TransactionModel) GetCustomerName() string {
	if m != nil {
		return m.CustomerName
	}
	return ""
}

func (m *TransactionModel) GetProductId() string {
	if m != nil {
		return m.ProductId
	}
	return ""
}

func (m *TransactionModel) GetOrderAmount() int32 {
	if m != nil {
		return m.OrderAmount
	}
	return 0
}

type ResultInfo struct {
	Status  string `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *ResultInfo) Reset()                    { *m = ResultInfo{} }
func (m *ResultInfo) String() string            { return proto.CompactTextString(m) }
func (*ResultInfo) ProtoMessage()               {}
func (*ResultInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ResultInfo) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *ResultInfo) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*TransactionModel)(nil), "simple.microservices.transaction.TransactionModel")
	proto.RegisterType((*ResultInfo)(nil), "simple.microservices.transaction.ResultInfo")
}

func init() { proto.RegisterFile("rpc/transaction/service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0x59, 0xa5, 0x95, 0xce, 0xb6, 0x20, 0x39, 0xc8, 0x22, 0x14, 0xd6, 0x7a, 0xe9, 0x41,
	0x52, 0xa8, 0x77, 0x41, 0x6f, 0x3d, 0xa8, 0xb0, 0x78, 0xf2, 0xb2, 0xc4, 0x64, 0x94, 0xc0, 0x26,
	0xb3, 0x64, 0x12, 0xd1, 0x7f, 0x2f, 0x2e, 0x29, 0x5d, 0x7a, 0xe9, 0xf1, 0x7d, 0x33, 0xc9, 0x7b,
	0x6f, 0x60, 0x19, 0x7a, 0xbd, 0x89, 0x41, 0x79, 0x56, 0x3a, 0x5a, 0xf2, 0x1b, 0xc6, 0xf0, 0x6d,
	0x35, 0xca, 0x3e, 0x50, 0x24, 0x51, 0xb3, 0x75, 0x7d, 0x87, 0xd2, 0x59, 0x1d, 0x28, 0x8f, 0x58,
	0x8e, 0xf6, 0x57, 0xbf, 0x70, 0xf9, 0x76, 0x90, 0xcf, 0x64, 0xb0, 0x13, 0xb7, 0xb0, 0xd0, 0x89,
	0x23, 0x39, 0x0c, 0xad, 0x57, 0x0e, 0xab, 0xa2, 0x2e, 0xd6, 0xb3, 0x66, 0xbe, 0x87, 0x2f, 0xca,
	0xa1, 0x58, 0x02, 0xf4, 0x81, 0x4c, 0xd2, 0xb1, 0xb5, 0xa6, 0x3a, 0x1b, 0x36, 0x66, 0x99, 0xec,
	0x8c, 0xb8, 0x81, 0x39, 0x05, 0x83, 0xa1, 0x55, 0x8e, 0x92, 0x8f, 0xd5, 0x79, 0x5d, 0xac, 0x27,
	0x4d, 0x39, 0xb0, 0xc7, 0x01, 0xad, 0x1e, 0x00, 0x1a, 0xe4, 0xd4, 0xc5, 0x9d, 0xff, 0x24, 0x71,
	0x05, 0x53, 0x8e, 0x2a, 0x26, 0xce, 0x6e, 0x59, 0x89, 0x0a, 0x2e, 0x1c, 0x32, 0xab, 0x2f, 0xcc,
	0x26, 0x7b, 0xb9, 0xfd, 0x81, 0x72, 0x14, 0x5d, 0x58, 0x98, 0xbc, 0xfe, 0xff, 0x2e, 0xb6, 0xf2,
	0x54, 0x6b, 0x79, 0x5c, 0xf9, 0xfa, 0xee, 0xf4, 0x9b, 0x43, 0xd6, 0xa7, 0xc5, 0x7b, 0x39, 0x9a,
	0x7c, 0x4c, 0x87, 0x63, 0xdf, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x75, 0x88, 0x63, 0x72, 0x8d,
	0x01, 0x00, 0x00,
}
