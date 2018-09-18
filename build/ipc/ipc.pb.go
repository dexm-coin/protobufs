// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ipc.proto

package ipc

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

type IPCEnvelope_MsgType int32

const (
	IPCEnvelope_GET_BLOCK_NUMBER IPCEnvelope_MsgType = 0
	IPCEnvelope_GET_TS           IPCEnvelope_MsgType = 1
	IPCEnvelope_GET_SENDER       IPCEnvelope_MsgType = 2
	IPCEnvelope_GET_BLOCKHASH    IPCEnvelope_MsgType = 3
	IPCEnvelope_SUICIDE          IPCEnvelope_MsgType = 4
	IPCEnvelope_REVERT           IPCEnvelope_MsgType = 5
	IPCEnvelope_GET              IPCEnvelope_MsgType = 6
	IPCEnvelope_APPEND           IPCEnvelope_MsgType = 7
)

var IPCEnvelope_MsgType_name = map[int32]string{
	0: "GET_BLOCK_NUMBER",
	1: "GET_TS",
	2: "GET_SENDER",
	3: "GET_BLOCKHASH",
	4: "SUICIDE",
	5: "REVERT",
	6: "GET",
	7: "APPEND",
}

var IPCEnvelope_MsgType_value = map[string]int32{
	"GET_BLOCK_NUMBER": 0,
	"GET_TS":           1,
	"GET_SENDER":       2,
	"GET_BLOCKHASH":    3,
	"SUICIDE":          4,
	"REVERT":           5,
	"GET":              6,
	"APPEND":           7,
}

func (x IPCEnvelope_MsgType) String() string {
	return proto.EnumName(IPCEnvelope_MsgType_name, int32(x))
}

func (IPCEnvelope_MsgType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6eb2dac662210585, []int{0, 0}
}

type IPCEnvelope struct {
	Type                 IPCEnvelope_MsgType `protobuf:"varint,1,opt,name=type,proto3,enum=IPCEnvelope_MsgType" json:"type,omitempty"`
	Data                 []byte              `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *IPCEnvelope) Reset()         { *m = IPCEnvelope{} }
func (m *IPCEnvelope) String() string { return proto.CompactTextString(m) }
func (*IPCEnvelope) ProtoMessage()    {}
func (*IPCEnvelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_6eb2dac662210585, []int{0}
}

func (m *IPCEnvelope) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IPCEnvelope.Unmarshal(m, b)
}
func (m *IPCEnvelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IPCEnvelope.Marshal(b, m, deterministic)
}
func (m *IPCEnvelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPCEnvelope.Merge(m, src)
}
func (m *IPCEnvelope) XXX_Size() int {
	return xxx_messageInfo_IPCEnvelope.Size(m)
}
func (m *IPCEnvelope) XXX_DiscardUnknown() {
	xxx_messageInfo_IPCEnvelope.DiscardUnknown(m)
}

var xxx_messageInfo_IPCEnvelope proto.InternalMessageInfo

func (m *IPCEnvelope) GetType() IPCEnvelope_MsgType {
	if m != nil {
		return m.Type
	}
	return IPCEnvelope_GET_BLOCK_NUMBER
}

func (m *IPCEnvelope) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Append struct {
	Key                  []byte   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Append) Reset()         { *m = Append{} }
func (m *Append) String() string { return proto.CompactTextString(m) }
func (*Append) ProtoMessage()    {}
func (*Append) Descriptor() ([]byte, []int) {
	return fileDescriptor_6eb2dac662210585, []int{1}
}

func (m *Append) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Append.Unmarshal(m, b)
}
func (m *Append) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Append.Marshal(b, m, deterministic)
}
func (m *Append) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Append.Merge(m, src)
}
func (m *Append) XXX_Size() int {
	return xxx_messageInfo_Append.Size(m)
}
func (m *Append) XXX_DiscardUnknown() {
	xxx_messageInfo_Append.DiscardUnknown(m)
}

var xxx_messageInfo_Append proto.InternalMessageInfo

func (m *Append) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Append) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*IPCEnvelope)(nil), "IPCEnvelope")
	proto.RegisterType((*Append)(nil), "Append")
	proto.RegisterEnum("IPCEnvelope_MsgType", IPCEnvelope_MsgType_name, IPCEnvelope_MsgType_value)
}

func init() { proto.RegisterFile("ipc.proto", fileDescriptor_6eb2dac662210585) }

var fileDescriptor_6eb2dac662210585 = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0xdd, 0x26, 0x4d, 0x70, 0x5a, 0xcb, 0x38, 0xe4, 0xd0, 0x63, 0xc9, 0x29, 0xa7, 0x20,
	0xfa, 0x0b, 0xd2, 0x64, 0x68, 0x83, 0x36, 0x86, 0xcd, 0xd6, 0x6b, 0x89, 0x76, 0x11, 0xb1, 0x34,
	0x8b, 0xc6, 0x42, 0xc0, 0x3f, 0xe8, 0xbf, 0x92, 0x8d, 0x41, 0x7a, 0xfb, 0xe6, 0xcd, 0x7b, 0x30,
	0xf3, 0xe0, 0xf2, 0xcd, 0xbc, 0xc4, 0xe6, 0xa3, 0x69, 0x9b, 0xf0, 0x47, 0xc0, 0x24, 0x2f, 0x53,
	0x3e, 0x9e, 0xf4, 0xa1, 0x31, 0x9a, 0x22, 0x70, 0xdb, 0xce, 0xe8, 0xb9, 0x58, 0x88, 0x68, 0x76,
	0x1b, 0xc4, 0x67, 0xbb, 0x78, 0xf3, 0xf9, 0xaa, 0x3a, 0xa3, 0x65, 0xef, 0x20, 0x02, 0x77, 0x5f,
	0xb7, 0xf5, 0x7c, 0xb4, 0x10, 0xd1, 0x54, 0xf6, 0x1c, 0x7e, 0x83, 0x3f, 0x98, 0x28, 0x00, 0x5c,
	0xb1, 0xda, 0x2d, 0x1f, 0x1e, 0xd3, 0xfb, 0x5d, 0xb1, 0xdd, 0x2c, 0x59, 0xe2, 0x05, 0x01, 0x78,
	0x56, 0x55, 0x15, 0x0a, 0x9a, 0x01, 0x58, 0xae, 0xb8, 0xc8, 0x58, 0xe2, 0x88, 0xae, 0xe1, 0xea,
	0x3f, 0xb1, 0x4e, 0xaa, 0x35, 0x3a, 0x34, 0x01, 0xbf, 0xda, 0xe6, 0x69, 0x9e, 0x31, 0xba, 0x36,
	0x2b, 0xf9, 0x89, 0xa5, 0xc2, 0x31, 0xf9, 0xe0, 0xac, 0x58, 0xa1, 0x67, 0xc5, 0xa4, 0x2c, 0xb9,
	0xc8, 0xd0, 0x0f, 0x6f, 0xc0, 0x4b, 0x8c, 0xd1, 0xc7, 0x3d, 0x21, 0x38, 0xef, 0xba, 0xeb, 0x9f,
	0x98, 0x4a, 0x8b, 0x14, 0xc0, 0xf8, 0x54, 0x1f, 0xbe, 0xf4, 0x70, 0xee, 0xdf, 0xf0, 0xec, 0xf5,
	0x25, 0xdc, 0xfd, 0x06, 0x00, 0x00, 0xff, 0xff, 0x42, 0x5d, 0x7d, 0xf9, 0x11, 0x01, 0x00, 0x00,
}
