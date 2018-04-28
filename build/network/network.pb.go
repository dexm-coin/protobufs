// Code generated by protoc-gen-go. DO NOT EDIT.
// source: network.proto

/*
Package network is a generated protocol buffer package.

It is generated from these files:
	network.proto

It has these top-level messages:
	Request
	Envelope
	Broadcast
	Signature
	Peers
*/
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
	Request_GET_CHAIN_SEGMENT  Request_MessageTypes = 4
)

var Request_MessageTypes_name = map[int32]string{
	0: "GET_BLOCKCHAIN_LEN",
	1: "GET_BLOCK",
	2: "GET_PEERS",
	4: "GET_CHAIN_SEGMENT",
}
var Request_MessageTypes_value = map[string]int32{
	"GET_BLOCKCHAIN_LEN": 0,
	"GET_BLOCK":          1,
	"GET_PEERS":          2,
	"GET_CHAIN_SEGMENT":  4,
}

func (x Request_MessageTypes) String() string {
	return proto.EnumName(Request_MessageTypes_name, int32(x))
}
func (Request_MessageTypes) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type Envelope_ContentType int32

const (
	Envelope_REQUEST   Envelope_ContentType = 0
	Envelope_OTHER     Envelope_ContentType = 1
	Envelope_BROADCAST Envelope_ContentType = 2
)

var Envelope_ContentType_name = map[int32]string{
	0: "REQUEST",
	1: "OTHER",
	2: "BROADCAST",
}
var Envelope_ContentType_value = map[string]int32{
	"REQUEST":   0,
	"OTHER":     1,
	"BROADCAST": 2,
}

func (x Envelope_ContentType) String() string {
	return proto.EnumName(Envelope_ContentType_name, int32(x))
}
func (Envelope_ContentType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type Broadcast_BroadcastType int32

const (
	Broadcast_TRANSACTION     Broadcast_BroadcastType = 0
	Broadcast_BLOCK_PROPOSAL  Broadcast_BroadcastType = 1
	Broadcast_CHECKPOINT_VOTE Broadcast_BroadcastType = 2
)

var Broadcast_BroadcastType_name = map[int32]string{
	0: "TRANSACTION",
	1: "BLOCK_PROPOSAL",
	2: "CHECKPOINT_VOTE",
}
var Broadcast_BroadcastType_value = map[string]int32{
	"TRANSACTION":     0,
	"BLOCK_PROPOSAL":  1,
	"CHECKPOINT_VOTE": 2,
}

func (x Broadcast_BroadcastType) String() string {
	return proto.EnumName(Broadcast_BroadcastType_name, int32(x))
}
func (Broadcast_BroadcastType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type Request struct {
	Type Request_MessageTypes `protobuf:"varint,1,opt,name=type,enum=Request_MessageTypes" json:"type,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetType() Request_MessageTypes {
	if m != nil {
		return m.Type
	}
	return Request_GET_BLOCKCHAIN_LEN
}

type Envelope struct {
	Type Envelope_ContentType `protobuf:"varint,1,opt,name=type,enum=Envelope_ContentType" json:"type,omitempty"`
	Data []byte               `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Envelope) Reset()                    { *m = Envelope{} }
func (m *Envelope) String() string            { return proto.CompactTextString(m) }
func (*Envelope) ProtoMessage()               {}
func (*Envelope) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Envelope) GetType() Envelope_ContentType {
	if m != nil {
		return m.Type
	}
	return Envelope_REQUEST
}

func (m *Envelope) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Broadcast struct {
	Type     Broadcast_BroadcastType `protobuf:"varint,1,opt,name=type,enum=Broadcast_BroadcastType" json:"type,omitempty"`
	Data     []byte                  `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Identity *Signature              `protobuf:"bytes,3,opt,name=identity" json:"identity,omitempty"`
}

func (m *Broadcast) Reset()                    { *m = Broadcast{} }
func (m *Broadcast) String() string            { return proto.CompactTextString(m) }
func (*Broadcast) ProtoMessage()               {}
func (*Broadcast) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

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

type Signature struct {
	Pubkey []byte `protobuf:"bytes,1,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	R      []byte `protobuf:"bytes,2,opt,name=r,proto3" json:"r,omitempty"`
	S      []byte `protobuf:"bytes,3,opt,name=s,proto3" json:"s,omitempty"`
	Data   []byte `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Signature) Reset()                    { *m = Signature{} }
func (m *Signature) String() string            { return proto.CompactTextString(m) }
func (*Signature) ProtoMessage()               {}
func (*Signature) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

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
	Ip []string `protobuf:"bytes,1,rep,name=ip" json:"ip,omitempty"`
}

func (m *Peers) Reset()                    { *m = Peers{} }
func (m *Peers) String() string            { return proto.CompactTextString(m) }
func (*Peers) ProtoMessage()               {}
func (*Peers) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

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

func init() { proto.RegisterFile("network.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 403 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xbb, 0x6e, 0xfa, 0xc7, 0x13, 0x27, 0x5d, 0x06, 0xb5, 0xf8, 0x18, 0xf9, 0x80, 0x82,
	0x84, 0x7c, 0x28, 0xbc, 0x80, 0x63, 0x56, 0x4d, 0xd4, 0xd4, 0x6b, 0x76, 0x17, 0x2e, 0x1c, 0x22,
	0x97, 0xac, 0x2a, 0xab, 0xc8, 0x36, 0xde, 0x0d, 0xc8, 0x27, 0x9e, 0x81, 0x37, 0xe2, 0xd1, 0x90,
	0x9d, 0xd4, 0x24, 0x42, 0xea, 0x6d, 0xbe, 0xf9, 0xc6, 0x9e, 0xdf, 0x7c, 0x5a, 0x18, 0x15, 0xda,
	0xfe, 0x2c, 0xeb, 0xc7, 0xb0, 0xaa, 0x4b, 0x5b, 0x06, 0xbf, 0x09, 0x9c, 0x09, 0xfd, 0x7d, 0xa3,
	0x8d, 0xc5, 0x37, 0x30, 0xb0, 0x4d, 0xa5, 0x7d, 0x32, 0x21, 0xd3, 0xf1, 0xf5, 0x65, 0xb8, 0xeb,
	0x87, 0x77, 0xda, 0x98, 0xec, 0x41, 0xab, 0xa6, 0xd2, 0x46, 0x74, 0x23, 0xc1, 0x17, 0xf0, 0xf6,
	0xbb, 0x78, 0x05, 0x78, 0xc3, 0xd4, 0x6a, 0xb6, 0xe4, 0xf1, 0x6d, 0x3c, 0x8f, 0x16, 0xc9, 0x6a,
	0xc9, 0x12, 0x7a, 0x84, 0x23, 0x70, 0xfb, 0x3e, 0x25, 0x4f, 0x32, 0x65, 0x4c, 0x48, 0xea, 0xe0,
	0x25, 0xbc, 0x68, 0xe5, 0xf6, 0x03, 0xc9, 0x6e, 0xee, 0x58, 0xa2, 0xe8, 0x20, 0xf8, 0x05, 0xe7,
	0xac, 0xf8, 0xa1, 0xbf, 0x95, 0x95, 0xfe, 0x8f, 0xe9, 0xc9, 0x08, 0xe3, 0xb2, 0xb0, 0xba, 0xb0,
	0xed, 0xfa, 0x2d, 0x13, 0x22, 0x0c, 0xd6, 0x99, 0xcd, 0x7c, 0x67, 0x42, 0xa6, 0x9e, 0xe8, 0xea,
	0xe0, 0x3d, 0x0c, 0xf7, 0x06, 0x71, 0x08, 0x67, 0x82, 0x7d, 0xfc, 0xc4, 0xa4, 0xa2, 0x47, 0xe8,
	0xc2, 0x09, 0x57, 0x73, 0x26, 0xb6, 0x5c, 0x33, 0xc1, 0xa3, 0x0f, 0x71, 0x24, 0x15, 0x75, 0x82,
	0x3f, 0x04, 0xdc, 0x59, 0x5d, 0x66, 0xeb, 0xaf, 0x99, 0xb1, 0xf8, 0xf6, 0x00, 0xc1, 0x0f, 0x7b,
	0xe7, 0x5f, 0xf5, 0x3c, 0x05, 0xbe, 0x86, 0xf3, 0x7c, 0xad, 0x0b, 0x9b, 0xdb, 0xc6, 0x3f, 0x9e,
	0x90, 0xe9, 0xf0, 0x1a, 0x42, 0x99, 0x3f, 0x14, 0x99, 0xdd, 0xd4, 0x5a, 0xf4, 0x5e, 0xb0, 0x80,
	0xd1, 0xc1, 0x2f, 0xf1, 0x02, 0x86, 0x4a, 0x44, 0x89, 0x8c, 0x62, 0xb5, 0xe0, 0x6d, 0x9e, 0x08,
	0xe3, 0x2e, 0xcb, 0x55, 0x2a, 0x78, 0xca, 0x65, 0xb4, 0xa4, 0x04, 0x5f, 0xc2, 0x45, 0x3c, 0x67,
	0xf1, 0x6d, 0xca, 0x17, 0x89, 0x5a, 0x7d, 0xe6, 0x8a, 0x51, 0x27, 0x90, 0xe0, 0xf6, 0x1b, 0xf0,
	0x0a, 0x4e, 0xab, 0xcd, 0xfd, 0xa3, 0x6e, 0xba, 0x1b, 0x3c, 0xb1, 0x53, 0xe8, 0x01, 0xa9, 0x77,
	0xa0, 0xa4, 0x6e, 0x95, 0xe9, 0xf0, 0x3c, 0x41, 0x4c, 0x7f, 0xc7, 0x60, 0x2f, 0xcd, 0x57, 0x70,
	0x92, 0x6a, 0x5d, 0x1b, 0x1c, 0x83, 0x93, 0x57, 0x3e, 0x99, 0x1c, 0x4f, 0x5d, 0xe1, 0xe4, 0xd5,
	0xfd, 0x69, 0xf7, 0x98, 0xde, 0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x88, 0x67, 0x08, 0x5c, 0x5d,
	0x02, 0x00, 0x00,
}
