// Code generated by protoc-gen-go. DO NOT EDIT.
// source: request.proto

package request // import "github.com/RTradeLtd/grpc/pay/request"

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

// ethAddress common.Address, paymentMethod uint8, paymentNumber, chargeAmountInWei *big.Int
type SignRequest struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Method               string   `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	Number               string   `protobuf:"bytes,3,opt,name=number,proto3" json:"number,omitempty"`
	ChargeAmount         string   `protobuf:"bytes,4,opt,name=chargeAmount,proto3" json:"chargeAmount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignRequest) Reset()         { *m = SignRequest{} }
func (m *SignRequest) String() string { return proto.CompactTextString(m) }
func (*SignRequest) ProtoMessage()    {}
func (*SignRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_request_76e48e6daa10fc3f, []int{0}
}
func (m *SignRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignRequest.Unmarshal(m, b)
}
func (m *SignRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignRequest.Marshal(b, m, deterministic)
}
func (dst *SignRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignRequest.Merge(dst, src)
}
func (m *SignRequest) XXX_Size() int {
	return xxx_messageInfo_SignRequest.Size(m)
}
func (m *SignRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignRequest proto.InternalMessageInfo

func (m *SignRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *SignRequest) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *SignRequest) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *SignRequest) GetChargeAmount() string {
	if m != nil {
		return m.ChargeAmount
	}
	return ""
}

func init() {
	proto.RegisterType((*SignRequest)(nil), "request.SignRequest")
}

func init() { proto.RegisterFile("request.proto", fileDescriptor_request_76e48e6daa10fc3f) }

var fileDescriptor_request_76e48e6daa10fc3f = []byte{
	// 166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4a, 0x2d, 0x2c,
	0x4d, 0x2d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0xaa, 0xb9,
	0xb8, 0x83, 0x33, 0xd3, 0xf3, 0x82, 0x20, 0x5c, 0x21, 0x09, 0x2e, 0xf6, 0xc4, 0x94, 0x94, 0xa2,
	0xd4, 0xe2, 0x62, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x18, 0x57, 0x48, 0x8c, 0x8b, 0x2d,
	0x37, 0xb5, 0x24, 0x23, 0x3f, 0x45, 0x82, 0x09, 0x2c, 0x01, 0xe5, 0x81, 0xc4, 0xf3, 0x4a, 0x73,
	0x93, 0x52, 0x8b, 0x24, 0x98, 0x21, 0xe2, 0x10, 0x9e, 0x90, 0x12, 0x17, 0x4f, 0x72, 0x46, 0x62,
	0x51, 0x7a, 0xaa, 0x63, 0x6e, 0x7e, 0x69, 0x5e, 0x89, 0x04, 0x0b, 0x58, 0x16, 0x45, 0xcc, 0x49,
	0x3d, 0x4a, 0x35, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x3f, 0x28, 0xa4,
	0x28, 0x31, 0x25, 0xd5, 0xa7, 0x24, 0x45, 0x3f, 0xbd, 0xa8, 0x20, 0x59, 0xbf, 0x20, 0xb1, 0x52,
	0x1f, 0xea, 0xca, 0x24, 0x36, 0xb0, 0xab, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x54, 0x0b,
	0x9c, 0x1e, 0xc6, 0x00, 0x00, 0x00,
}
