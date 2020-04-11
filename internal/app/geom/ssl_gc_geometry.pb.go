// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ssl_gc_geometry.proto

package geom

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

// A vector with two dimensions
type Vector2 struct {
	X                    *float32 `protobuf:"fixed32,1,req,name=x" json:"x,omitempty"`
	Y                    *float32 `protobuf:"fixed32,2,req,name=y" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vector2) Reset()         { *m = Vector2{} }
func (m *Vector2) String() string { return proto.CompactTextString(m) }
func (*Vector2) ProtoMessage()    {}
func (*Vector2) Descriptor() ([]byte, []int) {
	return fileDescriptor_760314ee40a5c6e3, []int{0}
}

func (m *Vector2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vector2.Unmarshal(m, b)
}
func (m *Vector2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vector2.Marshal(b, m, deterministic)
}
func (m *Vector2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vector2.Merge(m, src)
}
func (m *Vector2) XXX_Size() int {
	return xxx_messageInfo_Vector2.Size(m)
}
func (m *Vector2) XXX_DiscardUnknown() {
	xxx_messageInfo_Vector2.DiscardUnknown(m)
}

var xxx_messageInfo_Vector2 proto.InternalMessageInfo

func (m *Vector2) GetX() float32 {
	if m != nil && m.X != nil {
		return *m.X
	}
	return 0
}

func (m *Vector2) GetY() float32 {
	if m != nil && m.Y != nil {
		return *m.Y
	}
	return 0
}

// A vector with three dimensions
type Vector3 struct {
	X                    *float32 `protobuf:"fixed32,1,req,name=x" json:"x,omitempty"`
	Y                    *float32 `protobuf:"fixed32,2,req,name=y" json:"y,omitempty"`
	Z                    *float32 `protobuf:"fixed32,3,req,name=z" json:"z,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vector3) Reset()         { *m = Vector3{} }
func (m *Vector3) String() string { return proto.CompactTextString(m) }
func (*Vector3) ProtoMessage()    {}
func (*Vector3) Descriptor() ([]byte, []int) {
	return fileDescriptor_760314ee40a5c6e3, []int{1}
}

func (m *Vector3) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vector3.Unmarshal(m, b)
}
func (m *Vector3) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vector3.Marshal(b, m, deterministic)
}
func (m *Vector3) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vector3.Merge(m, src)
}
func (m *Vector3) XXX_Size() int {
	return xxx_messageInfo_Vector3.Size(m)
}
func (m *Vector3) XXX_DiscardUnknown() {
	xxx_messageInfo_Vector3.DiscardUnknown(m)
}

var xxx_messageInfo_Vector3 proto.InternalMessageInfo

func (m *Vector3) GetX() float32 {
	if m != nil && m.X != nil {
		return *m.X
	}
	return 0
}

func (m *Vector3) GetY() float32 {
	if m != nil && m.Y != nil {
		return *m.Y
	}
	return 0
}

func (m *Vector3) GetZ() float32 {
	if m != nil && m.Z != nil {
		return *m.Z
	}
	return 0
}

func init() {
	proto.RegisterType((*Vector2)(nil), "Vector2")
	proto.RegisterType((*Vector3)(nil), "Vector3")
}

func init() {
	proto.RegisterFile("ssl_gc_geometry.proto", fileDescriptor_760314ee40a5c6e3)
}

var fileDescriptor_760314ee40a5c6e3 = []byte{
	// 157 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x2e, 0xce, 0x89,
	0x4f, 0x4f, 0x8e, 0x4f, 0x4f, 0xcd, 0xcf, 0x4d, 0x2d, 0x29, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x57, 0x52, 0xe5, 0x62, 0x0f, 0x4b, 0x4d, 0x2e, 0xc9, 0x2f, 0x32, 0x12, 0xe2, 0xe1, 0x62,
	0xac, 0x90, 0x60, 0x54, 0x60, 0xd2, 0x60, 0x0a, 0x62, 0xac, 0x00, 0xf1, 0x2a, 0x25, 0x98, 0x20,
	0xbc, 0x4a, 0x25, 0x63, 0x98, 0x32, 0x63, 0x7c, 0xca, 0x40, 0xbc, 0x2a, 0x09, 0x66, 0x08, 0xaf,
	0xca, 0xc9, 0x2e, 0xca, 0x26, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x3f,
	0x28, 0x3f, 0x29, 0xdf, 0xb9, 0xb4, 0x40, 0x37, 0x38, 0xd8, 0x47, 0xbf, 0xb8, 0x38, 0x47, 0x37,
	0x3d, 0x31, 0x37, 0x55, 0x37, 0x39, 0x3f, 0xaf, 0xa4, 0x28, 0x3f, 0x27, 0x27, 0xb5, 0x48, 0x3f,
	0x33, 0xaf, 0x24, 0xb5, 0x28, 0x2f, 0x31, 0x47, 0x3f, 0xb1, 0xa0, 0x40, 0x1f, 0xe4, 0x4a, 0x40,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x54, 0x5b, 0x46, 0x0c, 0xb3, 0x00, 0x00, 0x00,
}
