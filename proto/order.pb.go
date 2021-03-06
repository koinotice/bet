// Code generated by protoc-gen-go. DO NOT EDIT.
// source: order.proto

package order

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

type CloseRequest struct {
	Roomid               int64    `protobuf:"varint,1,opt,name=roomid,proto3" json:"roomid,omitempty"`
	Orderid              int64    `protobuf:"varint,2,opt,name=orderid,proto3" json:"orderid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CloseRequest) Reset()         { *m = CloseRequest{} }
func (m *CloseRequest) String() string { return proto.CompactTextString(m) }
func (*CloseRequest) ProtoMessage()    {}
func (*CloseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{0}
}

func (m *CloseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloseRequest.Unmarshal(m, b)
}
func (m *CloseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloseRequest.Marshal(b, m, deterministic)
}
func (m *CloseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloseRequest.Merge(m, src)
}
func (m *CloseRequest) XXX_Size() int {
	return xxx_messageInfo_CloseRequest.Size(m)
}
func (m *CloseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CloseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CloseRequest proto.InternalMessageInfo

func (m *CloseRequest) GetRoomid() int64 {
	if m != nil {
		return m.Roomid
	}
	return 0
}

func (m *CloseRequest) GetOrderid() int64 {
	if m != nil {
		return m.Orderid
	}
	return 0
}

type CloseResponse struct {
	Txhash               string   `protobuf:"bytes,2,opt,name=txhash,proto3" json:"txhash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CloseResponse) Reset()         { *m = CloseResponse{} }
func (m *CloseResponse) String() string { return proto.CompactTextString(m) }
func (*CloseResponse) ProtoMessage()    {}
func (*CloseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{1}
}

func (m *CloseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloseResponse.Unmarshal(m, b)
}
func (m *CloseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloseResponse.Marshal(b, m, deterministic)
}
func (m *CloseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloseResponse.Merge(m, src)
}
func (m *CloseResponse) XXX_Size() int {
	return xxx_messageInfo_CloseResponse.Size(m)
}
func (m *CloseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CloseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CloseResponse proto.InternalMessageInfo

func (m *CloseResponse) GetTxhash() string {
	if m != nil {
		return m.Txhash
	}
	return ""
}

func init() {
	proto.RegisterType((*CloseRequest)(nil), "CloseRequest")
	proto.RegisterType((*CloseResponse)(nil), "CloseResponse")
}

func init() { proto.RegisterFile("order.proto", fileDescriptor_cd01338c35d87077) }

var fileDescriptor_cd01338c35d87077 = []byte{
	// 145 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x2f, 0x4a, 0x49,
	0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x72, 0xe0, 0xe2, 0x71, 0xce, 0xc9, 0x2f, 0x4e,
	0x0d, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe3, 0x62, 0x2b, 0xca, 0xcf, 0xcf, 0xcd,
	0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0e, 0x82, 0xf2, 0x84, 0x24, 0xb8, 0xd8, 0xc1, 0xda,
	0x32, 0x53, 0x24, 0x98, 0xc0, 0x12, 0x30, 0xae, 0x92, 0x3a, 0x17, 0x2f, 0xd4, 0x84, 0xe2, 0x82,
	0xfc, 0xbc, 0xe2, 0x54, 0x90, 0x11, 0x25, 0x15, 0x19, 0x89, 0xc5, 0x19, 0x60, 0x95, 0x9c, 0x41,
	0x50, 0x9e, 0x91, 0x09, 0x17, 0x87, 0x53, 0x6a, 0x89, 0x3f, 0x48, 0x9b, 0x90, 0x06, 0x17, 0x2b,
	0x58, 0x93, 0x10, 0xaf, 0x1e, 0xb2, 0xf5, 0x52, 0x7c, 0x7a, 0x28, 0x66, 0x29, 0x31, 0x24, 0xb1,
	0x81, 0xdd, 0x69, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x31, 0x32, 0x53, 0xcf, 0xb6, 0x00, 0x00,
	0x00,
}
