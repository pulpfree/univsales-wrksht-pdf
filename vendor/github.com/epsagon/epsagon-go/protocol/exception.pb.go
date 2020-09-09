// Code generated by protoc-gen-go. DO NOT EDIT.
// source: exception.proto

package protocol

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Exception struct {
	Type                 string            `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Message              string            `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Traceback            string            `protobuf:"bytes,3,opt,name=traceback,proto3" json:"traceback,omitempty"`
	Time                 float64           `protobuf:"fixed64,4,opt,name=time,proto3" json:"time,omitempty"`
	AdditionalData       map[string]string `protobuf:"bytes,5,rep,name=additional_data,json=additionalData,proto3" json:"additional_data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Exception) Reset()         { *m = Exception{} }
func (m *Exception) String() string { return proto.CompactTextString(m) }
func (*Exception) ProtoMessage()    {}
func (*Exception) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb6dc2e8aaa76c47, []int{0}
}

func (m *Exception) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Exception.Unmarshal(m, b)
}
func (m *Exception) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Exception.Marshal(b, m, deterministic)
}
func (m *Exception) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Exception.Merge(m, src)
}
func (m *Exception) XXX_Size() int {
	return xxx_messageInfo_Exception.Size(m)
}
func (m *Exception) XXX_DiscardUnknown() {
	xxx_messageInfo_Exception.DiscardUnknown(m)
}

var xxx_messageInfo_Exception proto.InternalMessageInfo

func (m *Exception) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Exception) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Exception) GetTraceback() string {
	if m != nil {
		return m.Traceback
	}
	return ""
}

func (m *Exception) GetTime() float64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *Exception) GetAdditionalData() map[string]string {
	if m != nil {
		return m.AdditionalData
	}
	return nil
}

func init() {
	proto.RegisterType((*Exception)(nil), "protocol.Exception")
	proto.RegisterMapType((map[string]string)(nil), "protocol.Exception.AdditionalDataEntry")
}

func init() { proto.RegisterFile("exception.proto", fileDescriptor_fb6dc2e8aaa76c47) }

var fileDescriptor_fb6dc2e8aaa76c47 = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0xad, 0x48, 0x4e,
	0x2d, 0x28, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0xc9,
	0xf9, 0x39, 0x4a, 0x0d, 0x4c, 0x5c, 0x9c, 0xae, 0x30, 0x59, 0x21, 0x21, 0x2e, 0x96, 0x92, 0xca,
	0x82, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0x48, 0x82, 0x8b, 0x3d, 0x37,
	0xb5, 0xb8, 0x38, 0x31, 0x3d, 0x55, 0x82, 0x09, 0x2c, 0x0c, 0xe3, 0x0a, 0xc9, 0x70, 0x71, 0x96,
	0x14, 0x25, 0x26, 0xa7, 0x26, 0x25, 0x26, 0x67, 0x4b, 0x30, 0x83, 0xe5, 0x10, 0x02, 0x60, 0xb3,
	0x32, 0x73, 0x53, 0x25, 0x58, 0x14, 0x18, 0x35, 0x18, 0x83, 0xc0, 0x6c, 0xa1, 0x00, 0x2e, 0xfe,
	0xc4, 0x94, 0x94, 0x4c, 0x90, 0x5d, 0x89, 0x39, 0xf1, 0x29, 0x89, 0x25, 0x89, 0x12, 0xac, 0x0a,
	0xcc, 0x1a, 0xdc, 0x46, 0xea, 0x7a, 0x30, 0x17, 0xe9, 0xc1, 0x5d, 0xa3, 0xe7, 0x08, 0x57, 0xea,
	0x92, 0x58, 0x92, 0xe8, 0x9a, 0x57, 0x52, 0x54, 0x19, 0xc4, 0x97, 0x88, 0x22, 0x28, 0xe5, 0xc8,
	0x25, 0x8c, 0x45, 0x99, 0x90, 0x00, 0x17, 0x73, 0x76, 0x6a, 0x25, 0xd4, 0x1f, 0x20, 0xa6, 0x90,
	0x08, 0x17, 0x6b, 0x59, 0x62, 0x4e, 0x29, 0xcc, 0x13, 0x10, 0x8e, 0x15, 0x93, 0x05, 0x63, 0x12,
	0x1b, 0xd8, 0x6a, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbb, 0x2e, 0x17, 0x93, 0x26, 0x01,
	0x00, 0x00,
}
