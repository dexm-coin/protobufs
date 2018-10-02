// Code generated by protoc-gen-go. DO NOT EDIT.
// source: network.proto

package network

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

type Request_MessageTypes int32

const (
	Request_GET_BLOCKCHAIN_LEN Request_MessageTypes = 0
	Request_GET_BLOCK          Request_MessageTypes = 1
	Request_GET_PEERS          Request_MessageTypes = 2
	Request_GET_WALLET_STATUS  Request_MessageTypes = 3
	Request_GET_VERSION        Request_MessageTypes = 4
	Request_GET_CONTRACT_CODE  Request_MessageTypes = 5
	Request_GET_CONTRACT_STATE Request_MessageTypes = 6
	Request_GET_INTERESTS      Request_MessageTypes = 7
)

var Request_MessageTypes_name = map[int32]string{
	0: "GET_BLOCKCHAIN_LEN",
	1: "GET_BLOCK",
	2: "GET_PEERS",
	3: "GET_WALLET_STATUS",
	4: "GET_VERSION",
	5: "GET_CONTRACT_CODE",
	6: "GET_CONTRACT_STATE",
	7: "GET_INTERESTS",
}

var Request_MessageTypes_value = map[string]int32{
	"GET_BLOCKCHAIN_LEN": 0,
	"GET_BLOCK":          1,
	"GET_PEERS":          2,
	"GET_WALLET_STATUS":  3,
	"GET_VERSION":        4,
	"GET_CONTRACT_CODE":  5,
	"GET_CONTRACT_STATE": 6,
	"GET_INTERESTS":      7,
}

func (x Request_MessageTypes) String() string {
	return proto.EnumName(Request_MessageTypes_name, int32(x))
}

func (Request_MessageTypes) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{0, 0}
}

type Envelope_ContentType int32

const (
	Envelope_BROADCAST Envelope_ContentType = 0
	Envelope_OTHER     Envelope_ContentType = 1
	Envelope_REQUEST   Envelope_ContentType = 2
	Envelope_INTERESTS Envelope_ContentType = 3
)

var Envelope_ContentType_name = map[int32]string{
	0: "BROADCAST",
	1: "OTHER",
	2: "REQUEST",
	3: "INTERESTS",
}

var Envelope_ContentType_value = map[string]int32{
	"BROADCAST": 0,
	"OTHER":     1,
	"REQUEST":   2,
	"INTERESTS": 3,
}

func (x Envelope_ContentType) String() string {
	return proto.EnumName(Envelope_ContentType_name, int32(x))
}

func (Envelope_ContentType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{1, 0}
}

type Broadcast_BroadcastType int32

const (
	Broadcast_TRANSACTION         Broadcast_BroadcastType = 0
	Broadcast_BLOCK_PROPOSAL      Broadcast_BroadcastType = 1
	Broadcast_CHECKPOINT_VOTE     Broadcast_BroadcastType = 2
	Broadcast_NEW_CONTRACT        Broadcast_BroadcastType = 3
	Broadcast_WITHDRAW            Broadcast_BroadcastType = 4
	Broadcast_MERKLE_ROOTS_SIGNED Broadcast_BroadcastType = 5
	Broadcast_SCHNORR             Broadcast_BroadcastType = 6
	Broadcast_SIGN_SCHNORR        Broadcast_BroadcastType = 7
	Broadcast_MERKLE_PROOF        Broadcast_BroadcastType = 8
)

var Broadcast_BroadcastType_name = map[int32]string{
	0: "TRANSACTION",
	1: "BLOCK_PROPOSAL",
	2: "CHECKPOINT_VOTE",
	3: "NEW_CONTRACT",
	4: "WITHDRAW",
	5: "MERKLE_ROOTS_SIGNED",
	6: "SCHNORR",
	7: "SIGN_SCHNORR",
	8: "MERKLE_PROOF",
}

var Broadcast_BroadcastType_value = map[string]int32{
	"TRANSACTION":         0,
	"BLOCK_PROPOSAL":      1,
	"CHECKPOINT_VOTE":     2,
	"NEW_CONTRACT":        3,
	"WITHDRAW":            4,
	"MERKLE_ROOTS_SIGNED": 5,
	"SCHNORR":             6,
	"SIGN_SCHNORR":        7,
	"MERKLE_PROOF":        8,
}

func (x Broadcast_BroadcastType) String() string {
	return proto.EnumName(Broadcast_BroadcastType_name, int32(x))
}

func (Broadcast_BroadcastType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{2, 0}
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
	return fileDescriptor_8571034d60397816, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
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
	Shard                uint32               `protobuf:"varint,3,opt,name=shard,proto3" json:"shard,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Envelope) Reset()         { *m = Envelope{} }
func (m *Envelope) String() string { return proto.CompactTextString(m) }
func (*Envelope) ProtoMessage()    {}
func (*Envelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{1}
}

func (m *Envelope) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Envelope.Unmarshal(m, b)
}
func (m *Envelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Envelope.Marshal(b, m, deterministic)
}
func (m *Envelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Envelope.Merge(m, src)
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

func (m *Envelope) GetShard() uint32 {
	if m != nil {
		return m.Shard
	}
	return 0
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
	return fileDescriptor_8571034d60397816, []int{2}
}

func (m *Broadcast) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Broadcast.Unmarshal(m, b)
}
func (m *Broadcast) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Broadcast.Marshal(b, m, deterministic)
}
func (m *Broadcast) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Broadcast.Merge(m, src)
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
	return fileDescriptor_8571034d60397816, []int{3}
}

func (m *Signature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Signature.Unmarshal(m, b)
}
func (m *Signature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Signature.Marshal(b, m, deterministic)
}
func (m *Signature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Signature.Merge(m, src)
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

type Interests struct {
	Keys                 []string `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Interests) Reset()         { *m = Interests{} }
func (m *Interests) String() string { return proto.CompactTextString(m) }
func (*Interests) ProtoMessage()    {}
func (*Interests) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{4}
}

func (m *Interests) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Interests.Unmarshal(m, b)
}
func (m *Interests) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Interests.Marshal(b, m, deterministic)
}
func (m *Interests) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Interests.Merge(m, src)
}
func (m *Interests) XXX_Size() int {
	return xxx_messageInfo_Interests.Size(m)
}
func (m *Interests) XXX_DiscardUnknown() {
	xxx_messageInfo_Interests.DiscardUnknown(m)
}

var xxx_messageInfo_Interests proto.InternalMessageInfo

func (m *Interests) GetKeys() []string {
	if m != nil {
		return m.Keys
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
	return fileDescriptor_8571034d60397816, []int{5}
}

func (m *Peers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Peers.Unmarshal(m, b)
}
func (m *Peers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Peers.Marshal(b, m, deterministic)
}
func (m *Peers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Peers.Merge(m, src)
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
	proto.RegisterType((*Interests)(nil), "Interests")
	proto.RegisterType((*Peers)(nil), "Peers")
	proto.RegisterEnum("Request_MessageTypes", Request_MessageTypes_name, Request_MessageTypes_value)
	proto.RegisterEnum("Envelope_ContentType", Envelope_ContentType_name, Envelope_ContentType_value)
	proto.RegisterEnum("Broadcast_BroadcastType", Broadcast_BroadcastType_name, Broadcast_BroadcastType_value)
}

func init() { proto.RegisterFile("network.proto", fileDescriptor_8571034d60397816) }

var fileDescriptor_8571034d60397816 = []byte{
	// 569 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xcd, 0x52, 0xdb, 0x3c,
	0x14, 0x45, 0x8e, 0xf3, 0x77, 0xe3, 0x80, 0x10, 0x1f, 0x90, 0xdd, 0x97, 0xf1, 0xa2, 0x93, 0xce,
	0x74, 0xb2, 0xa0, 0x4f, 0x60, 0x1c, 0x95, 0x78, 0x30, 0x56, 0x2a, 0x09, 0xb2, 0xf4, 0x98, 0x46,
	0x43, 0x3d, 0x74, 0x6c, 0xd7, 0x16, 0x6d, 0xf3, 0x3e, 0x5d, 0x77, 0xd3, 0x75, 0x9f, 0xac, 0x9b,
	0x8e, 0x8c, 0x63, 0x60, 0xba, 0xe8, 0xee, 0x9e, 0x73, 0x8f, 0xa4, 0x7b, 0xce, 0x1d, 0xc1, 0x38,
	0x53, 0xfa, 0x6b, 0x5e, 0xde, 0xcf, 0x8b, 0x32, 0xd7, 0xb9, 0xfb, 0x1b, 0x41, 0x9f, 0xab, 0xcf,
	0x0f, 0xaa, 0xd2, 0xe4, 0x35, 0xd8, 0x7a, 0x5b, 0xa8, 0x09, 0x9a, 0xa2, 0xd9, 0xfe, 0xd9, 0xf1,
	0xbc, 0xe1, 0xe7, 0x57, 0xaa, 0xaa, 0x92, 0x3b, 0x25, 0xb7, 0x85, 0xaa, 0x78, 0x2d, 0x21, 0xff,
	0x41, 0x37, 0xcd, 0x36, 0xea, 0xdb, 0xc4, 0x9a, 0xa2, 0x99, 0xcd, 0x1f, 0x81, 0xfb, 0x03, 0x81,
	0xf3, 0x5c, 0x4c, 0x4e, 0x80, 0x5c, 0x50, 0x19, 0x9f, 0x87, 0xcc, 0xbf, 0xf4, 0x97, 0x5e, 0x10,
	0xc5, 0x21, 0x8d, 0xf0, 0x1e, 0x19, 0xc3, 0xb0, 0xe5, 0x31, 0xda, 0xc1, 0x15, 0xa5, 0x5c, 0x60,
	0x8b, 0x1c, 0xc3, 0xa1, 0x81, 0x6b, 0x2f, 0x0c, 0xa9, 0x8c, 0x85, 0xf4, 0xe4, 0xb5, 0xc0, 0x1d,
	0x72, 0x00, 0x23, 0x43, 0xdf, 0x50, 0x2e, 0x02, 0x16, 0x61, 0x7b, 0xa7, 0xf3, 0x59, 0x24, 0xb9,
	0xe7, 0x9b, 0x62, 0x41, 0x71, 0x77, 0xf7, 0x68, 0x4b, 0x9b, 0x0b, 0x28, 0xee, 0x91, 0x43, 0x18,
	0x1b, 0x3e, 0x88, 0x24, 0xe5, 0x54, 0x48, 0x81, 0xfb, 0xee, 0x77, 0x04, 0x03, 0x9a, 0x7d, 0x51,
	0x9f, 0xf2, 0x42, 0xfd, 0x65, 0x7f, 0xd7, 0x98, 0xfb, 0x79, 0xa6, 0x55, 0xa6, 0x8d, 0xa5, 0xc6,
	0x3e, 0x01, 0x7b, 0x93, 0xe8, 0xa4, 0x76, 0xef, 0xf0, 0xba, 0x36, 0x91, 0x54, 0x1f, 0x93, 0x72,
	0x33, 0xe9, 0x4c, 0xd1, 0x6c, 0xcc, 0x1f, 0x81, 0xeb, 0xc3, 0xe8, 0xd9, 0x71, 0xe3, 0xf4, 0x9c,
	0x33, 0x6f, 0xe1, 0x7b, 0x42, 0xe2, 0x3d, 0x32, 0x84, 0x2e, 0x93, 0x4b, 0xca, 0x31, 0x22, 0x23,
	0xe8, 0x73, 0xfa, 0xfe, 0x9a, 0x0a, 0x89, 0x2d, 0x23, 0x7b, 0x1a, 0xb3, 0xe3, 0xfe, 0xb2, 0x60,
	0x78, 0x5e, 0xe6, 0xc9, 0xe6, 0x43, 0x52, 0x69, 0xf2, 0xe6, 0xc5, 0x9c, 0x93, 0x79, 0xdb, 0x79,
	0xaa, 0xfe, 0x31, 0xea, 0x2b, 0x18, 0xa4, 0x1b, 0x95, 0xe9, 0x54, 0x6f, 0xeb, 0x69, 0x47, 0x67,
	0x30, 0x17, 0xe9, 0x5d, 0x96, 0xe8, 0x87, 0x52, 0xf1, 0xb6, 0x47, 0x30, 0x74, 0xa4, 0x0c, 0x27,
	0x76, 0x6d, 0xc8, 0x94, 0xee, 0x4f, 0x04, 0xe3, 0x17, 0xaf, 0x98, 0xad, 0x48, 0xee, 0x45, 0xc2,
	0xf3, 0xa5, 0xd9, 0xca, 0x1e, 0x21, 0xb0, 0x5f, 0xef, 0x35, 0x5e, 0x71, 0xb6, 0x62, 0xc2, 0x0b,
	0x31, 0x22, 0x47, 0x70, 0xe0, 0x2f, 0xa9, 0x7f, 0xb9, 0x62, 0x41, 0x24, 0xe3, 0x1b, 0x26, 0x29,
	0xb6, 0x08, 0x06, 0x27, 0xa2, 0xeb, 0x76, 0x4f, 0xb8, 0x43, 0x1c, 0x18, 0xac, 0x03, 0xb9, 0x5c,
	0x70, 0x6f, 0x8d, 0x6d, 0x72, 0x0a, 0x47, 0x57, 0x94, 0x5f, 0x86, 0x34, 0xe6, 0x8c, 0x49, 0x11,
	0x8b, 0xe0, 0x22, 0xa2, 0x0b, 0xdc, 0x35, 0x51, 0x09, 0x7f, 0x19, 0x31, 0xce, 0x71, 0xcf, 0xdc,
	0x62, 0x1a, 0xf1, 0x8e, 0xe9, 0x1b, 0xa6, 0x39, 0xb7, 0xe2, 0x8c, 0xbd, 0xc3, 0x03, 0x57, 0xc0,
	0xb0, 0xb5, 0x47, 0x4e, 0xa0, 0x57, 0x3c, 0xdc, 0xde, 0xab, 0x6d, 0x1d, 0xa0, 0xc3, 0x1b, 0x44,
	0x1c, 0x40, 0x65, 0x93, 0x12, 0x2a, 0x0d, 0xaa, 0xea, 0x6c, 0x1c, 0x8e, 0xaa, 0x36, 0x44, 0xfb,
	0x29, 0x44, 0xf7, 0x7f, 0x18, 0x06, 0x99, 0x56, 0xa5, 0xaa, 0x74, 0x2d, 0xb8, 0x57, 0xdb, 0x6a,
	0x82, 0xa6, 0x9d, 0xd9, 0x90, 0xd7, 0xb5, 0x7b, 0x0a, 0xdd, 0x95, 0x52, 0x65, 0x45, 0xf6, 0xc1,
	0x4a, 0x8b, 0xa6, 0x65, 0xa5, 0xc5, 0x6d, 0xaf, 0xfe, 0x7a, 0x6f, 0xff, 0x04, 0x00, 0x00, 0xff,
	0xff, 0xc1, 0x8c, 0x62, 0x9f, 0x8b, 0x03, 0x00, 0x00,
}
