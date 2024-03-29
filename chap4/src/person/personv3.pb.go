// Code generated by protoc-gen-go. DO NOT EDIT.
// source: personv3.proto

package person

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

type Person struct {
	Name                 *Person_Name    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email                []*Person_Email `protobuf:"bytes,2,rep,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_945d10db9c163165, []int{0}
}

func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetName() *Person_Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *Person) GetEmail() []*Person_Email {
	if m != nil {
		return m.Email
	}
	return nil
}

type Person_Name struct {
	Family               string   `protobuf:"bytes,1,opt,name=family,proto3" json:"family,omitempty"`
	Personal             string   `protobuf:"bytes,2,opt,name=personal,proto3" json:"personal,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Person_Name) Reset()         { *m = Person_Name{} }
func (m *Person_Name) String() string { return proto.CompactTextString(m) }
func (*Person_Name) ProtoMessage()    {}
func (*Person_Name) Descriptor() ([]byte, []int) {
	return fileDescriptor_945d10db9c163165, []int{0, 0}
}

func (m *Person_Name) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person_Name.Unmarshal(m, b)
}
func (m *Person_Name) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person_Name.Marshal(b, m, deterministic)
}
func (m *Person_Name) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person_Name.Merge(m, src)
}
func (m *Person_Name) XXX_Size() int {
	return xxx_messageInfo_Person_Name.Size(m)
}
func (m *Person_Name) XXX_DiscardUnknown() {
	xxx_messageInfo_Person_Name.DiscardUnknown(m)
}

var xxx_messageInfo_Person_Name proto.InternalMessageInfo

func (m *Person_Name) GetFamily() string {
	if m != nil {
		return m.Family
	}
	return ""
}

func (m *Person_Name) GetPersonal() string {
	if m != nil {
		return m.Personal
	}
	return ""
}

type Person_Email struct {
	Kind                 string   `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Person_Email) Reset()         { *m = Person_Email{} }
func (m *Person_Email) String() string { return proto.CompactTextString(m) }
func (*Person_Email) ProtoMessage()    {}
func (*Person_Email) Descriptor() ([]byte, []int) {
	return fileDescriptor_945d10db9c163165, []int{0, 1}
}

func (m *Person_Email) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person_Email.Unmarshal(m, b)
}
func (m *Person_Email) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person_Email.Marshal(b, m, deterministic)
}
func (m *Person_Email) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person_Email.Merge(m, src)
}
func (m *Person_Email) XXX_Size() int {
	return xxx_messageInfo_Person_Email.Size(m)
}
func (m *Person_Email) XXX_DiscardUnknown() {
	xxx_messageInfo_Person_Email.DiscardUnknown(m)
}

var xxx_messageInfo_Person_Email proto.InternalMessageInfo

func (m *Person_Email) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *Person_Email) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*Person)(nil), "person.Person")
	proto.RegisterType((*Person_Name)(nil), "person.Person.Name")
	proto.RegisterType((*Person_Email)(nil), "person.Person.Email")
}

func init() { proto.RegisterFile("personv3.proto", fileDescriptor_945d10db9c163165) }

var fileDescriptor_945d10db9c163165 = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x48, 0x2d, 0x2a,
	0xce, 0xcf, 0x2b, 0x33, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0x95, 0x2e,
	0x30, 0x72, 0xb1, 0x05, 0x80, 0x99, 0x42, 0xea, 0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x8c,
	0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0xc2, 0x7a, 0x10, 0x15, 0x7a, 0x10, 0x59, 0x3d, 0xbf, 0xc4, 0xdc,
	0xd4, 0x20, 0xb0, 0x02, 0x21, 0x2d, 0x2e, 0xd6, 0xd4, 0xdc, 0xc4, 0xcc, 0x1c, 0x09, 0x26, 0x05,
	0x66, 0x0d, 0x6e, 0x23, 0x11, 0x34, 0x95, 0xae, 0x20, 0xb9, 0x20, 0x88, 0x12, 0x29, 0x2b, 0x2e,
	0x16, 0x90, 0x4e, 0x21, 0x31, 0x2e, 0xb6, 0xb4, 0xc4, 0xdc, 0xcc, 0x9c, 0x4a, 0xb0, 0xf1, 0x9c,
	0x41, 0x50, 0x9e, 0x90, 0x14, 0x17, 0x07, 0x44, 0x77, 0x22, 0xc8, 0x38, 0x90, 0x0c, 0x9c, 0x2f,
	0x65, 0xca, 0xc5, 0x0a, 0x36, 0x4b, 0x48, 0x88, 0x8b, 0x25, 0x3b, 0x33, 0x2f, 0x05, 0xaa, 0x15,
	0xcc, 0x16, 0x92, 0xe0, 0x62, 0x4f, 0x4c, 0x49, 0x29, 0x4a, 0x2d, 0x2e, 0x86, 0xea, 0x83, 0x71,
	0x93, 0xd8, 0xc0, 0x3e, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x7e, 0x8e, 0x7f, 0x36, 0xf3,
	0x00, 0x00, 0x00,
}
