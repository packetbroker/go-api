// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: packetbroker/api/v3/contact.proto

package packetbroker

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ContactInfo struct {
	// Person or organization name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Email address.
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	// URL.
	Url                  string   `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContactInfo) Reset()         { *m = ContactInfo{} }
func (m *ContactInfo) String() string { return proto.CompactTextString(m) }
func (*ContactInfo) ProtoMessage()    {}
func (*ContactInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_eea700e68c0af50b, []int{0}
}
func (m *ContactInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContactInfo.Unmarshal(m, b)
}
func (m *ContactInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContactInfo.Marshal(b, m, deterministic)
}
func (m *ContactInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContactInfo.Merge(m, src)
}
func (m *ContactInfo) XXX_Size() int {
	return xxx_messageInfo_ContactInfo.Size(m)
}
func (m *ContactInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ContactInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ContactInfo proto.InternalMessageInfo

func (m *ContactInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ContactInfo) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *ContactInfo) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type ContactInfoValue struct {
	Value                *ContactInfo `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ContactInfoValue) Reset()         { *m = ContactInfoValue{} }
func (m *ContactInfoValue) String() string { return proto.CompactTextString(m) }
func (*ContactInfoValue) ProtoMessage()    {}
func (*ContactInfoValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_eea700e68c0af50b, []int{1}
}
func (m *ContactInfoValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContactInfoValue.Unmarshal(m, b)
}
func (m *ContactInfoValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContactInfoValue.Marshal(b, m, deterministic)
}
func (m *ContactInfoValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContactInfoValue.Merge(m, src)
}
func (m *ContactInfoValue) XXX_Size() int {
	return xxx_messageInfo_ContactInfoValue.Size(m)
}
func (m *ContactInfoValue) XXX_DiscardUnknown() {
	xxx_messageInfo_ContactInfoValue.DiscardUnknown(m)
}

var xxx_messageInfo_ContactInfoValue proto.InternalMessageInfo

func (m *ContactInfoValue) GetValue() *ContactInfo {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*ContactInfo)(nil), "org.packetbroker.v3.ContactInfo")
	proto.RegisterType((*ContactInfoValue)(nil), "org.packetbroker.v3.ContactInfoValue")
}

func init() { proto.RegisterFile("packetbroker/api/v3/contact.proto", fileDescriptor_eea700e68c0af50b) }

var fileDescriptor_eea700e68c0af50b = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2c, 0x48, 0x4c, 0xce,
	0x4e, 0x2d, 0x49, 0x2a, 0xca, 0xcf, 0x4e, 0x2d, 0xd2, 0x4f, 0x2c, 0xc8, 0xd4, 0x2f, 0x33, 0xd6,
	0x4f, 0xce, 0xcf, 0x2b, 0x49, 0x4c, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xce,
	0x2f, 0x4a, 0xd7, 0x43, 0x56, 0xa6, 0x57, 0x66, 0xac, 0xe4, 0xc9, 0xc5, 0xed, 0x0c, 0x51, 0xe5,
	0x99, 0x97, 0x96, 0x2f, 0x24, 0xc4, 0xc5, 0x92, 0x97, 0x98, 0x9b, 0x2a, 0xc1, 0xa8, 0xc0, 0xa8,
	0xc1, 0x19, 0x04, 0x66, 0x0b, 0x89, 0x70, 0xb1, 0xa6, 0xe6, 0x26, 0x66, 0xe6, 0x48, 0x30, 0x81,
	0x05, 0x21, 0x1c, 0x21, 0x01, 0x2e, 0xe6, 0xd2, 0xa2, 0x1c, 0x09, 0x66, 0xb0, 0x18, 0x88, 0xa9,
	0xe4, 0xc5, 0x25, 0x80, 0x64, 0x54, 0x58, 0x62, 0x4e, 0x69, 0xaa, 0x90, 0x19, 0x17, 0x6b, 0x19,
	0x88, 0x01, 0x36, 0x90, 0xdb, 0x48, 0x41, 0x0f, 0x8b, 0x1b, 0xf4, 0x90, 0x74, 0x05, 0x41, 0x94,
	0x3b, 0x69, 0x46, 0xa9, 0xa7, 0xe7, 0xa3, 0x2a, 0xcc, 0x2f, 0x4a, 0x87, 0xfa, 0xcb, 0x1a, 0x59,
	0x3c, 0x89, 0x0d, 0xec, 0x3b, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x79, 0xf1, 0x21, 0x5e,
	0x02, 0x01, 0x00, 0x00,
}
