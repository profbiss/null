// Code generated by protoc-gen-go. DO NOT EDIT.
// source: int.proto

package null

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

// String is a nullable string. It supports SQL and JSON serialization.
type Int struct {
	I                    int64    `protobuf:"varint,1,opt,name=i,proto3" json:"i,omitempty"`
	Valid                bool     `protobuf:"varint,2,opt,name=valid,proto3" json:"valid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Int) Reset()         { *m = Int{} }
func (m *Int) String() string { return proto.CompactTextString(m) }
func (*Int) ProtoMessage()    {}
func (*Int) Descriptor() ([]byte, []int) {
	return fileDescriptor_221abb4e6615ef44, []int{0}
}

func (m *Int) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Int.Unmarshal(m, b)
}
func (m *Int) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Int.Marshal(b, m, deterministic)
}
func (m *Int) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Int.Merge(m, src)
}
func (m *Int) XXX_Size() int {
	return xxx_messageInfo_Int.Size(m)
}
func (m *Int) XXX_DiscardUnknown() {
	xxx_messageInfo_Int.DiscardUnknown(m)
}

var xxx_messageInfo_Int proto.InternalMessageInfo

func (m *Int) GetI() int64 {
	if m != nil {
		return m.I
	}
	return 0
}

func (m *Int) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func init() {
	proto.RegisterType((*Int)(nil), "profbiss.null.Int")
}

func init() { proto.RegisterFile("int.proto", fileDescriptor_221abb4e6615ef44) }

var fileDescriptor_221abb4e6615ef44 = []byte{
	// 100 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcc, 0xcc, 0x2b, 0xd1,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2d, 0x28, 0xca, 0x4f, 0x4b, 0xca, 0x2c, 0x2e, 0xd6,
	0xcb, 0x2b, 0xcd, 0xc9, 0x51, 0xd2, 0xe4, 0x62, 0xf6, 0xcc, 0x2b, 0x11, 0xe2, 0xe1, 0x62, 0xcc,
	0x94, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0e, 0x62, 0xcc, 0x14, 0x12, 0xe1, 0x62, 0x2d, 0x4b, 0xcc,
	0xc9, 0x4c, 0x91, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x08, 0x82, 0x70, 0x9c, 0xd8, 0xa2, 0x58, 0x40,
	0x5a, 0x92, 0xd8, 0xc0, 0x06, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x9d, 0x39, 0x7c, 0x76,
	0x55, 0x00, 0x00, 0x00,
}
