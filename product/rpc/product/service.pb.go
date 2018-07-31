// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/product/service.proto

/*
Package product is a generated protocol buffer package.

It is generated from these files:
	rpc/product/service.proto

It has these top-level messages:
	ProductModel
	UpdateStockParam
	ParamString
	ListProductModel
	ResultInfo
*/
package product

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

type ProductModel struct {
	ProductId string `protobuf:"bytes,1,opt,name=product_id,json=productId" json:"product_id,omitempty"`
	Stock     int32  `protobuf:"varint,2,opt,name=stock" json:"stock,omitempty"`
	Name      string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
}

func (m *ProductModel) Reset()                    { *m = ProductModel{} }
func (m *ProductModel) String() string            { return proto.CompactTextString(m) }
func (*ProductModel) ProtoMessage()               {}
func (*ProductModel) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ProductModel) GetProductId() string {
	if m != nil {
		return m.ProductId
	}
	return ""
}

func (m *ProductModel) GetStock() int32 {
	if m != nil {
		return m.Stock
	}
	return 0
}

func (m *ProductModel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type UpdateStockParam struct {
	ProductId string `protobuf:"bytes,1,opt,name=product_id,json=productId" json:"product_id,omitempty"`
	Amount    int32  `protobuf:"varint,2,opt,name=amount" json:"amount,omitempty"`
}

func (m *UpdateStockParam) Reset()                    { *m = UpdateStockParam{} }
func (m *UpdateStockParam) String() string            { return proto.CompactTextString(m) }
func (*UpdateStockParam) ProtoMessage()               {}
func (*UpdateStockParam) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UpdateStockParam) GetProductId() string {
	if m != nil {
		return m.ProductId
	}
	return ""
}

func (m *UpdateStockParam) GetAmount() int32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type ParamString struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *ParamString) Reset()                    { *m = ParamString{} }
func (m *ParamString) String() string            { return proto.CompactTextString(m) }
func (*ParamString) ProtoMessage()               {}
func (*ParamString) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ParamString) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ListProductModel struct {
	Result []*ProductModel `protobuf:"bytes,1,rep,name=result" json:"result,omitempty"`
}

func (m *ListProductModel) Reset()                    { *m = ListProductModel{} }
func (m *ListProductModel) String() string            { return proto.CompactTextString(m) }
func (*ListProductModel) ProtoMessage()               {}
func (*ListProductModel) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ListProductModel) GetResult() []*ProductModel {
	if m != nil {
		return m.Result
	}
	return nil
}

type ResultInfo struct {
	Status  string `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *ResultInfo) Reset()                    { *m = ResultInfo{} }
func (m *ResultInfo) String() string            { return proto.CompactTextString(m) }
func (*ResultInfo) ProtoMessage()               {}
func (*ResultInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

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
	proto.RegisterType((*ProductModel)(nil), "simple.microservices.product.ProductModel")
	proto.RegisterType((*UpdateStockParam)(nil), "simple.microservices.product.UpdateStockParam")
	proto.RegisterType((*ParamString)(nil), "simple.microservices.product.ParamString")
	proto.RegisterType((*ListProductModel)(nil), "simple.microservices.product.ListProductModel")
	proto.RegisterType((*ResultInfo)(nil), "simple.microservices.product.ResultInfo")
}

func init() { proto.RegisterFile("rpc/product/service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x5f, 0x4b, 0xfb, 0x30,
	0x14, 0xa5, 0xeb, 0xd6, 0xd1, 0xbb, 0x1f, 0x3f, 0x46, 0x10, 0xa9, 0xc3, 0xc1, 0xe8, 0x53, 0xf5,
	0xa1, 0x83, 0xf9, 0xee, 0xc3, 0x10, 0x64, 0xa0, 0x30, 0x3a, 0x54, 0x10, 0x45, 0x63, 0x7b, 0x1d,
	0xc1, 0xb6, 0x29, 0xc9, 0xed, 0xbe, 0x87, 0xdf, 0x58, 0x96, 0x45, 0x9d, 0x0a, 0xfb, 0xf3, 0x76,
	0x4f, 0x72, 0xce, 0xb9, 0x27, 0x97, 0x1b, 0x38, 0x52, 0x55, 0x3a, 0xac, 0x94, 0xcc, 0xea, 0x94,
	0x86, 0x1a, 0xd5, 0x42, 0xa4, 0x18, 0x57, 0x4a, 0x92, 0x64, 0xc7, 0x5a, 0x14, 0x55, 0x8e, 0x71,
	0x21, 0x52, 0x25, 0xed, 0x95, 0x8e, 0x2d, 0x37, 0xbc, 0x83, 0x7f, 0xd3, 0x55, 0x79, 0x2d, 0x33,
	0xcc, 0x59, 0x1f, 0xc0, 0x5e, 0x3d, 0x89, 0x2c, 0x70, 0x06, 0x4e, 0xe4, 0x27, 0xbe, 0x3d, 0x99,
	0x64, 0xec, 0x00, 0x5a, 0x9a, 0x64, 0xfa, 0x16, 0x34, 0x06, 0x4e, 0xd4, 0x4a, 0x56, 0x80, 0x31,
	0x68, 0x96, 0xbc, 0xc0, 0xc0, 0x35, 0x74, 0x53, 0x87, 0x13, 0xe8, 0xde, 0x54, 0x19, 0x27, 0x9c,
	0x2d, 0x29, 0x53, 0xae, 0x78, 0xb1, 0xcd, 0xfc, 0x10, 0x3c, 0x5e, 0xc8, 0xba, 0x24, 0xeb, 0x6e,
	0x51, 0xd8, 0x87, 0x8e, 0xd1, 0xcf, 0x48, 0x89, 0x72, 0xce, 0xfe, 0x43, 0xe3, 0x4b, 0xdd, 0x10,
	0x59, 0x78, 0x0b, 0xdd, 0x2b, 0xa1, 0xe9, 0xc7, 0x33, 0xc6, 0xe0, 0x29, 0xd4, 0x75, 0x4e, 0x81,
	0x33, 0x70, 0xa3, 0xce, 0xe8, 0x34, 0xde, 0x34, 0x85, 0x78, 0x5d, 0x9b, 0x58, 0x65, 0x78, 0x0e,
	0x90, 0x98, 0x6a, 0x52, 0xbe, 0xca, 0x65, 0x38, 0x4d, 0x9c, 0x6a, 0x6d, 0x3b, 0x5b, 0xc4, 0x02,
	0x68, 0x17, 0xa8, 0x35, 0x9f, 0xa3, 0x49, 0xed, 0x27, 0x9f, 0x70, 0xf4, 0xee, 0x42, 0xdb, 0x1a,
	0xb3, 0x07, 0x68, 0xce, 0xf8, 0x02, 0xd9, 0x1e, 0x39, 0x7a, 0xd1, 0x66, 0xee, 0x5a, 0xb6, 0x67,
	0x70, 0x2f, 0x91, 0xd8, 0xc9, 0x16, 0xf3, 0xef, 0x19, 0xf6, 0xe2, 0xcd, 0xd4, 0x3f, 0xf3, 0x7c,
	0x04, 0xef, 0x02, 0x73, 0x24, 0xdc, 0xa7, 0xc9, 0xee, 0x0f, 0x98, 0x43, 0x67, 0x6d, 0x59, 0xd8,
	0x96, 0x74, 0xbf, 0xf7, 0x6a, 0xf7, 0x46, 0x63, 0xff, 0xbe, 0x6d, 0x4f, 0x5f, 0x3c, 0xf3, 0x3d,
	0xce, 0x3e, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa3, 0x10, 0x75, 0xb3, 0x3b, 0x03, 0x00, 0x00,
}
