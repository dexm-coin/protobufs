// Code generated by protoc-gen-go. DO NOT EDIT.
// source: network.proto

package network

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

type Request_MessageTypes int32

const (
	Request_GET_BLOCKCHAIN_LEN Request_MessageTypes = 0
	Request_GET_BLOCK          Request_MessageTypes = 1
	Request_GET_PEERS          Request_MessageTypes = 2
	Request_GET_WALLET_STATUS  Request_MessageTypes = 3
	Request_GET_VERSION        Request_MessageTypes = 4
)

var Request_MessageTypes_name = map[int32]string{
	0: "GET_BLOCKCHAIN_LEN",
	1: "GET_BLOCK",
	2: "GET_PEERS",
	3: "GET_WALLET_STATUS",
	4: "GET_VERSION",
}
var Request_MessageTypes_value = map[string]int32{
	"GET_BLOCKCHAIN_LEN": 0,
	"GET_BLOCK":          1,
	"GET_PEERS":          2,
	"GET_WALLET_STATUS":  3,
	"GET_VERSION":        4,
}

func (x Request_MessageTypes) String() string {
	return proto.EnumName(Request_MessageTypes_name, int32(x))
}
func (Request_MessageTypes) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_network_80b897d2f8d0b391, []int{0, 0}
}

type Envelope_ContentType int32

const (
	Envelope_BROADCAST Envelope_ContentType = 0
	Envelope_OTHER     Envelope_ContentType = 1
	Envelope_REQUEST   Envelope_ContentType = 2
)

var Envelope_ContentType_name = map[int32]string{
	0: "BROADCAST",
	1: "OTHER",
	2: "REQUEST",
}
var Envelope_ContentType_value = map[string]int32{
	"BROADCAST": 0,
	"OTHER":     1,
	"REQUEST":   2,
}

func (x Envelope_ContentType) String() string {
	return proto.EnumName(Envelope_ContentType_name, int32(x))
}
func (Envelope_ContentType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_network_80b897d2f8d0b391, []int{1, 0}
}

type Broadcast_BroadcastType int32

const (
	Broadcast_TRANSACTION     Broadcast_BroadcastType = 0
	Broadcast_BLOCK_PROPOSAL  Broadcast_BroadcastType = 1
	Broadcast_CHECKPOINT_VOTE Broadcast_BroadcastType = 2
	Broadcast_NEW_CONTRACT    Broadcast_BroadcastType = 3
	Broadcast_WITHDRAW        Broadcast_BroadcastType = 4
)

var Broadcast_BroadcastType_name = map[int32]string{
	0: "TRANSACTION",
	1: "BLOCK_PROPOSAL",
	2: "CHECKPOINT_VOTE",
	3: "NEW_CONTRACT",
	4: "WITHDRAW",
}
var Broadcast_BroadcastType_value = map[string]int32{
	"TRANSACTION":     0,
	"BLOCK_PROPOSAL":  1,
	"CHECKPOINT_VOTE": 2,
	"NEW_CONTRACT":    3,
	"WITHDRAW":        4,
}

func (x Broadcast_BroadcastType) String() string {
	return proto.EnumName(Broadcast_BroadcastType_name, int32(x))
}
func (Broadcast_BroadcastType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_network_80b897d2f8d0b391, []int{2, 0}
}

type Request struct {
	Type                 Request_MessageTypes `protobuf:"varint,1,opt,name=type,proto3,enum=Request_MessageTypes" json:"type,omitempty"`
	Index                uint64               `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_network_80b897d2f8d0b391, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetType() Request_MessageTypes {
	if m != nil {
		return m.Type
	}
	return Request_GET_BLOCKCHAIN_LEN
}

func (m *Request) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

type Envelope struct {
	Type                 Envelope_ContentType `protobuf:"varint,1,opt,name=type,proto3,enum=Envelope_ContentType" json:"type,omitempty"`
	Data                 []byte               `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Envelope) Reset()         { *m = Envelope{} }
func (m *Envelope) String() string { return proto.CompactTextString(m) }
func (*Envelope) ProtoMessage()    {}
func (*Envelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_network_80b897d2f8d0b391, []int{1}
}
func (m *Envelope) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Envelope.Unmarshal(m, b)
}
func (m *Envelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Envelope.Marshal(b, m, deterministic)
}
func (dst *Envelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Envelope.Merge(dst, src)
}
func (m *Envelope) XXX_Size() int {
	return xxx_messageInfo_Envelope.Size(m)
}
func (m *Envelope) XXX_DiscardUnknown() {
	xxx_messageInfo_Envelope.DiscardUnknown(m)
}

var xxx_messageInfo_Envelope proto.InternalMessageInfo

func (m *Envelope) GetType() Envelope_ContentType {
	if m != nil {
		return m.Type
	}
	return Envelope_BROADCAST
}

func (m *Envelope) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Broadcast struct {
	Type                 Broadcast_BroadcastType `protobuf:"varint,1,opt,name=type,proto3,enum=Broadcast_BroadcastType" json:"type,omitempty"`
	Data                 []byte                  `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Identity             *Signature              `protobuf:"bytes,3,opt,name=identity,proto3" json:"identity,omitempty"`
	TTL                  uint32                  `protobuf:"varint,4,opt,name=TTL,proto3" json:"TTL,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *Broadcast) Reset()         { *m = Broadcast{} }
func (m *Broadcast) String() string { return proto.CompactTextString(m) }
func (*Broadcast) ProtoMessage()    {}
func (*Broadcast) Descriptor() ([]byte, []int) {
	return fileDescriptor_network_80b897d2f8d0b391, []int{2}
}
func (m *Broadcast) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Broadcast.Unmarshal(m, b)
}
func (m *Broadcast) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Broadcast.Marshal(b, m, deterministic)
}
func (dst *Broadcast) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Broadcast.Merge(dst, src)
}
func (m *Broadcast) XXX_Size() int {
	return xxx_messageInfo_Broadcast.Size(m)
}
func (m *Broadcast) XXX_DiscardUnknown() {
	xxx_messageInfo_Broadcast.DiscardUnknown(m)
}

var xxx_messageInfo_Broadcast proto.InternalMessageInfo

func (m *Broadcast) GetType() Broadcast_BroadcastType {
	if m != nil {
		return m.Type
	}
	return Broadcast_TRANSACTION
}

func (m *Broadcast) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Broadcast) GetIdentity() *Signature {
	if m != nil {
		return m.Identity
	}
	return nil
}

func (m *Broadcast) GetTTL() uint32 {
	if m != nil {
		return m.TTL
	}
	return 0
}

type Signature struct {
	Pubkey               []byte   `protobuf:"bytes,1,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	R                    []byte   `protobuf:"bytes,2,opt,name=r,proto3" json:"r,omitempty"`
	S                    []byte   `protobuf:"bytes,3,opt,name=s,proto3" json:"s,omitempty"`
	Data                 []byte   `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Signature) Reset()         { *m = Signature{} }
func (m *Signature) String() string { return proto.CompactTextString(m) }
func (*Signature) ProtoMessage()    {}
func (*Signature) Descriptor() ([]byte, []int) {
	return fileDescriptor_network_80b897d2f8d0b391, []int{3}
}
func (m *Signature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Signature.Unmarshal(m, b)
}
func (m *Signature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Signature.Marshal(b, m, deterministic)
}
func (dst *Signature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Signature.Merge(dst, src)
}
func (m *Signature) XXX_Size() int {
	return xxx_messageInfo_Signature.Size(m)
}
func (m *Signature) XXX_DiscardUnknown() {
	xxx_messageInfo_Signature.DiscardUnknown(m)
}

var xxx_messageInfo_Signature proto.InternalMessageInfo

func (m *Signature) GetPubkey() []byte {
	if m != nil {
		return m.Pubkey
	}
	return nil
}

func (m *Signature) GetR() []byte {
	if m != nil {
		return m.R
	}
	return nil
}

func (m *Signature) GetS() []byte {
	if m != nil {
		return m.S
	}
	return nil
}

func (m *Signature) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Peers struct {
	Ip                   []string `protobuf:"bytes,1,rep,name=ip,proto3" json:"ip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Peers) Reset()         { *m = Peers{} }
func (m *Peers) String() string { return proto.CompactTextString(m) }
func (*Peers) ProtoMessage()    {}
func (*Peers) Descriptor() ([]byte, []int) {
	return fileDescriptor_network_80b897d2f8d0b391, []int{4}
}
func (m *Peers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Peers.Unmarshal(m, b)
}
func (m *Peers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Peers.Marshal(b, m, deterministic)
}
func (dst *Peers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Peers.Merge(dst, src)
}
func (m *Peers) XXX_Size() int {
	return xxx_messageInfo_Peers.Size(m)
}
func (m *Peers) XXX_DiscardUnknown() {
	xxx_messageInfo_Peers.DiscardUnknown(m)
}

var xxx_messageInfo_Peers proto.InternalMessageInfo

func (m *Peers) GetIp() []string {
	if m != nil {
		return m.Ip
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "Request")
	proto.RegisterType((*Envelope)(nil), "Envelope")
	proto.RegisterType((*Broadcast)(nil), "Broadcast")
	proto.RegisterType((*Signature)(nil), "Signature")
	proto.RegisterType((*Peers)(nil), "Peers")
	proto.RegisterEnum("Request_MessageTypes", Request_MessageTypes_name, Request_MessageTypes_value)
	proto.RegisterEnum("Envelope_ContentType", Envelope_ContentType_name, Envelope_ContentType_value)
	proto.RegisterEnum("Broadcast_BroadcastType", Broadcast_BroadcastType_name, Broadcast_BroadcastType_value)
}

func init() { proto.RegisterFile("network.proto", fileDescriptor_network_80b897d2f8d0b391) }

var fileDescriptor_network_80b897d2f8d0b391 = []byte{
	// 453 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xc1, 0x6e, 0xd3, 0x4c,
	0x14, 0x85, 0x33, 0x8e, 0xd3, 0x36, 0x37, 0x4e, 0x3a, 0xff, 0xfd, 0x69, 0xf1, 0xd2, 0xf2, 0x02,
	0x19, 0x09, 0x79, 0x51, 0x78, 0x01, 0xc7, 0x8c, 0x48, 0x54, 0x63, 0x9b, 0x99, 0x69, 0xbb, 0x60,
	0x61, 0xb9, 0x64, 0x54, 0x59, 0xad, 0x6c, 0x63, 0x4f, 0x80, 0xac, 0x78, 0x2d, 0x1e, 0x88, 0x07,
	0x41, 0x36, 0xa9, 0x49, 0x85, 0xc4, 0xee, 0x9e, 0x7b, 0x8f, 0x34, 0xdf, 0x39, 0x1a, 0x98, 0x97,
	0x4a, 0x7f, 0xad, 0x9a, 0x7b, 0xbf, 0x6e, 0x2a, 0x5d, 0xb9, 0x3f, 0x08, 0x1c, 0x73, 0xf5, 0x79,
	0xab, 0x5a, 0x8d, 0x2f, 0xc1, 0xd4, 0xbb, 0x5a, 0xd9, 0xc4, 0x21, 0xde, 0xe2, 0xe2, 0xcc, 0xdf,
	0xef, 0xfd, 0xf7, 0xaa, 0x6d, 0xf3, 0x3b, 0x25, 0x77, 0xb5, 0x6a, 0x79, 0x6f, 0xc1, 0x67, 0x30,
	0x29, 0xca, 0x8d, 0xfa, 0x66, 0x1b, 0x0e, 0xf1, 0x4c, 0xfe, 0x5b, 0xb8, 0x0f, 0x60, 0x1d, 0x7a,
	0xf1, 0x1c, 0xf0, 0x1d, 0x93, 0xd9, 0x32, 0x4a, 0xc2, 0xcb, 0x70, 0x15, 0xac, 0xe3, 0x2c, 0x62,
	0x31, 0x1d, 0xe1, 0x1c, 0xa6, 0xc3, 0x9e, 0x92, 0x47, 0x99, 0x32, 0xc6, 0x05, 0x35, 0xf0, 0x0c,
	0xfe, 0xeb, 0xe4, 0x4d, 0x10, 0x45, 0x4c, 0x66, 0x42, 0x06, 0xf2, 0x4a, 0xd0, 0x31, 0x9e, 0xc2,
	0xac, 0x5b, 0x5f, 0x33, 0x2e, 0xd6, 0x49, 0x4c, 0x4d, 0xf7, 0x3b, 0x9c, 0xb0, 0xf2, 0x8b, 0x7a,
	0xa8, 0x6a, 0xf5, 0x17, 0xfa, 0xe3, 0xc1, 0x0f, 0xab, 0x52, 0xab, 0x52, 0x77, 0x3c, 0x7b, 0x74,
	0x04, 0x73, 0x93, 0xeb, 0xbc, 0x27, 0xb7, 0x78, 0x3f, 0xbb, 0x6f, 0x60, 0x76, 0x60, 0xec, 0x80,
	0x96, 0x3c, 0x09, 0xde, 0x86, 0x81, 0x90, 0x74, 0x84, 0x53, 0x98, 0x24, 0x72, 0xc5, 0x38, 0x25,
	0x38, 0x83, 0x63, 0xce, 0x3e, 0x5c, 0x31, 0x21, 0xa9, 0xe1, 0xfe, 0x24, 0x30, 0x5d, 0x36, 0x55,
	0xbe, 0xf9, 0x94, 0xb7, 0x1a, 0x5f, 0x3d, 0x41, 0xb0, 0xfd, 0xe1, 0xf2, 0x67, 0xfa, 0x37, 0x05,
	0xbe, 0x80, 0x93, 0x62, 0xa3, 0x4a, 0x5d, 0xe8, 0x9d, 0x3d, 0x76, 0x88, 0x37, 0xbb, 0x00, 0x5f,
	0x14, 0x77, 0x65, 0xae, 0xb7, 0x8d, 0xe2, 0xc3, 0x0d, 0x29, 0x8c, 0xa5, 0x8c, 0x6c, 0xd3, 0x21,
	0xde, 0x9c, 0x77, 0xa3, 0xfb, 0x11, 0xe6, 0x4f, 0x1e, 0xe9, 0xca, 0x92, 0x3c, 0x88, 0x45, 0x10,
	0xca, 0xae, 0xac, 0x11, 0x22, 0x2c, 0xfa, 0xba, 0xb3, 0x94, 0x27, 0x69, 0x22, 0x82, 0x88, 0x12,
	0xfc, 0x1f, 0x4e, 0xc3, 0x15, 0x0b, 0x2f, 0xd3, 0x64, 0x1d, 0xcb, 0xec, 0x3a, 0x91, 0x8c, 0x1a,
	0x48, 0xc1, 0x8a, 0xd9, 0x4d, 0x16, 0x26, 0xb1, 0xe4, 0x41, 0x28, 0xe9, 0xd8, 0x15, 0x30, 0x1d,
	0x28, 0xf0, 0x1c, 0x8e, 0xea, 0xed, 0xed, 0xbd, 0xda, 0xf5, 0x39, 0x2d, 0xbe, 0x57, 0x68, 0x01,
	0x69, 0xf6, 0x61, 0x48, 0xd3, 0xa9, 0xb6, 0x8f, 0x60, 0x71, 0xd2, 0x0e, 0x59, 0xcd, 0x83, 0xc6,
	0x9f, 0xc3, 0x24, 0x55, 0xaa, 0x69, 0x71, 0x01, 0x46, 0x51, 0xdb, 0xc4, 0x19, 0x7b, 0x53, 0x6e,
	0x14, 0xf5, 0xed, 0x51, 0xff, 0x2f, 0x5f, 0xff, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x7e, 0x6a, 0x05,
	0xe8, 0xa8, 0x02, 0x00, 0x00,
}
