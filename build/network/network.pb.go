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
	Request_HASH_EXIST         Request_MessageTypes = 8
	Request_GET_WALLET         Request_MessageTypes = 9
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
	8: "HASH_EXIST",
	9: "GET_WALLET",
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
	"HASH_EXIST":         8,
	"GET_WALLET":         9,
}

func (x Request_MessageTypes) String() string {
	return proto.EnumName(Request_MessageTypes_name, int32(x))
}

func (Request_MessageTypes) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{0, 0}
}

type Envelope_ContentType int32

const (
	Envelope_BROADCAST           Envelope_ContentType = 0
	Envelope_OTHER               Envelope_ContentType = 1
	Envelope_REQUEST             Envelope_ContentType = 2
	Envelope_INTERESTS           Envelope_ContentType = 3
	Envelope_NEIGHBOUR_INTERESTS Envelope_ContentType = 4
)

var Envelope_ContentType_name = map[int32]string{
	0: "BROADCAST",
	1: "OTHER",
	2: "REQUEST",
	3: "INTERESTS",
	4: "NEIGHBOUR_INTERESTS",
}

var Envelope_ContentType_value = map[string]int32{
	"BROADCAST":           0,
	"OTHER":               1,
	"REQUEST":             2,
	"INTERESTS":           3,
	"NEIGHBOUR_INTERESTS": 4,
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
	Broadcast_MONEY_WITHDRAW      Broadcast_BroadcastType = 9
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
	9: "MONEY_WITHDRAW",
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
	"MONEY_WITHDRAW":      9,
}

func (x Broadcast_BroadcastType) String() string {
	return proto.EnumName(Broadcast_BroadcastType_name, int32(x))
}

func (Broadcast_BroadcastType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{2, 0}
}

type Request struct {
	Type                 Request_MessageTypes `protobuf:"varint,1,opt,name=type,proto3,enum=Request_MessageTypes" json:"type,omitempty"`
	Data                 *Envelope            `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
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

func (m *Request) GetData() *Envelope {
	if m != nil {
		return m.Data
	}
	return nil
}

type Envelope struct {
	Type                 Envelope_ContentType `protobuf:"varint,1,opt,name=type,proto3,enum=Envelope_ContentType" json:"type,omitempty"`
	Data                 []byte               `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Shard                uint32               `protobuf:"varint,3,opt,name=shard,proto3" json:"shard,omitempty"`
	Identity             *Signature           `protobuf:"bytes,4,opt,name=identity,proto3" json:"identity,omitempty"`
	TTL                  uint32               `protobuf:"varint,5,opt,name=TTL,proto3" json:"TTL,omitempty"`
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

func (m *Envelope) GetIdentity() *Signature {
	if m != nil {
		return m.Identity
	}
	return nil
}

func (m *Envelope) GetTTL() uint32 {
	if m != nil {
		return m.TTL
	}
	return 0
}

type Broadcast struct {
	Type                 Broadcast_BroadcastType `protobuf:"varint,1,opt,name=type,proto3,enum=Broadcast_BroadcastType" json:"type,omitempty"`
	Data                 []byte                  `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Address              []byte                  `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	ShardAddress         uint32                  `protobuf:"varint,5,opt,name=shardAddress,proto3" json:"shardAddress,omitempty"`
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

func (m *Broadcast) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Broadcast) GetShardAddress() uint32 {
	if m != nil {
		return m.ShardAddress
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

type PeersAndInterests struct {
	Keys                 []string `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
	Ips                  []string `protobuf:"bytes,2,rep,name=ips,proto3" json:"ips,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PeersAndInterests) Reset()         { *m = PeersAndInterests{} }
func (m *PeersAndInterests) String() string { return proto.CompactTextString(m) }
func (*PeersAndInterests) ProtoMessage()    {}
func (*PeersAndInterests) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{6}
}

func (m *PeersAndInterests) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PeersAndInterests.Unmarshal(m, b)
}
func (m *PeersAndInterests) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PeersAndInterests.Marshal(b, m, deterministic)
}
func (m *PeersAndInterests) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeersAndInterests.Merge(m, src)
}
func (m *PeersAndInterests) XXX_Size() int {
	return xxx_messageInfo_PeersAndInterests.Size(m)
}
func (m *PeersAndInterests) XXX_DiscardUnknown() {
	xxx_messageInfo_PeersAndInterests.DiscardUnknown(m)
}

var xxx_messageInfo_PeersAndInterests proto.InternalMessageInfo

func (m *PeersAndInterests) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *PeersAndInterests) GetIps() []string {
	if m != nil {
		return m.Ips
	}
	return nil
}

type RandomMessage struct {
	Pubkey               []byte   `protobuf:"bytes,1,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RandomMessage) Reset()         { *m = RandomMessage{} }
func (m *RandomMessage) String() string { return proto.CompactTextString(m) }
func (*RandomMessage) ProtoMessage()    {}
func (*RandomMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{7}
}

func (m *RandomMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RandomMessage.Unmarshal(m, b)
}
func (m *RandomMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RandomMessage.Marshal(b, m, deterministic)
}
func (m *RandomMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RandomMessage.Merge(m, src)
}
func (m *RandomMessage) XXX_Size() int {
	return xxx_messageInfo_RandomMessage.Size(m)
}
func (m *RandomMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_RandomMessage.DiscardUnknown(m)
}

var xxx_messageInfo_RandomMessage proto.InternalMessageInfo

func (m *RandomMessage) GetPubkey() []byte {
	if m != nil {
		return m.Pubkey
	}
	return nil
}

func (m *RandomMessage) GetData() []byte {
	if m != nil {
		return m.Data
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
	proto.RegisterType((*PeersAndInterests)(nil), "PeersAndInterests")
	proto.RegisterType((*RandomMessage)(nil), "RandomMessage")
	proto.RegisterEnum("Request_MessageTypes", Request_MessageTypes_name, Request_MessageTypes_value)
	proto.RegisterEnum("Envelope_ContentType", Envelope_ContentType_name, Envelope_ContentType_value)
	proto.RegisterEnum("Broadcast_BroadcastType", Broadcast_BroadcastType_name, Broadcast_BroadcastType_value)
}

func init() { proto.RegisterFile("network.proto", fileDescriptor_8571034d60397816) }

var fileDescriptor_8571034d60397816 = []byte{
	// 676 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0x4d, 0x6f, 0xd3, 0x4c,
	0x10, 0xc7, 0x6b, 0xe7, 0xd5, 0x13, 0x27, 0xdd, 0x6e, 0x9f, 0xb6, 0xbe, 0x3c, 0x7a, 0x22, 0x1f,
	0x1e, 0x05, 0x09, 0xe5, 0x50, 0x4e, 0x88, 0x93, 0xe3, 0x2c, 0xb5, 0xd5, 0xd4, 0x1b, 0x76, 0xb7,
	0x0d, 0x48, 0x48, 0x96, 0x8b, 0x57, 0x25, 0x2a, 0x38, 0xc6, 0x76, 0x41, 0xf9, 0x3e, 0x7c, 0x17,
	0x6e, 0xdc, 0xf8, 0x2c, 0x5c, 0xd1, 0xba, 0x8e, 0x93, 0x8a, 0xb7, 0xdb, 0xcc, 0x7f, 0x66, 0x67,
	0x32, 0xbf, 0xfc, 0x13, 0xe8, 0x27, 0xb2, 0xf8, 0xb4, 0xca, 0x6e, 0xc7, 0x69, 0xb6, 0x2a, 0x56,
	0xf6, 0x67, 0x1d, 0x3a, 0x4c, 0x7e, 0xb8, 0x93, 0x79, 0x81, 0x1f, 0x41, 0xb3, 0x58, 0xa7, 0xd2,
	0xd2, 0x86, 0xda, 0x68, 0x70, 0x7a, 0x34, 0xae, 0xf4, 0xf1, 0x85, 0xcc, 0xf3, 0xe8, 0x46, 0x8a,
	0x75, 0x2a, 0x73, 0x56, 0xb6, 0xe0, 0x7f, 0xa1, 0x19, 0x47, 0x45, 0x64, 0xe9, 0x43, 0x6d, 0xd4,
	0x3b, 0x35, 0xc6, 0x24, 0xf9, 0x28, 0xdf, 0xad, 0x52, 0xc9, 0x4a, 0xd9, 0xfe, 0xaa, 0x81, 0xb9,
	0xfb, 0x0a, 0x1f, 0x03, 0x3e, 0x23, 0x22, 0x9c, 0xcc, 0xa8, 0x7b, 0xee, 0x7a, 0x8e, 0x1f, 0x84,
	0x33, 0x12, 0xa0, 0x3d, 0xdc, 0x07, 0xa3, 0xd6, 0x91, 0xb6, 0x49, 0xe7, 0x84, 0x30, 0x8e, 0x74,
	0x7c, 0x04, 0x07, 0x2a, 0x5d, 0x38, 0xb3, 0x19, 0x11, 0x21, 0x17, 0x8e, 0xb8, 0xe4, 0xa8, 0x81,
	0xf7, 0xa1, 0xa7, 0xe4, 0x2b, 0xc2, 0xb8, 0x4f, 0x03, 0xd4, 0xdc, 0xf4, 0xb9, 0x34, 0x10, 0xcc,
	0x71, 0x55, 0x30, 0x25, 0xa8, 0xb5, 0x59, 0x5a, 0xcb, 0x6a, 0x00, 0x41, 0x6d, 0x7c, 0x00, 0x7d,
	0xa5, 0xfb, 0x81, 0x20, 0x8c, 0x70, 0xc1, 0x51, 0x07, 0x0f, 0x00, 0x3c, 0x87, 0x7b, 0x21, 0x79,
	0xe9, 0x73, 0x81, 0xba, 0x2a, 0xdf, 0x6e, 0x46, 0x86, 0xfd, 0x5d, 0x83, 0xee, 0xe6, 0xc6, 0x9f,
	0x38, 0x6d, 0x0a, 0x63, 0x77, 0x95, 0x14, 0x32, 0x29, 0xd4, 0xc9, 0x15, 0x27, 0xbc, 0xc3, 0xc9,
	0xbc, 0x87, 0x83, 0xff, 0x81, 0x56, 0xfe, 0x36, 0xca, 0x62, 0xab, 0x31, 0xd4, 0x46, 0x7d, 0x76,
	0x9f, 0xe0, 0xff, 0xa1, 0xbb, 0x8c, 0x65, 0x52, 0x2c, 0x8b, 0xb5, 0xd5, 0x2c, 0xa9, 0xc2, 0x98,
	0x2f, 0x6f, 0x92, 0xa8, 0xb8, 0xcb, 0x24, 0xab, 0x6b, 0x18, 0x41, 0x43, 0x88, 0x99, 0xd5, 0x2a,
	0xdf, 0xaa, 0xd0, 0x7e, 0x0d, 0xbd, 0x9d, 0xc5, 0x8a, 0xe1, 0x84, 0x51, 0x67, 0xea, 0x3a, 0x5c,
	0xa0, 0x3d, 0x6c, 0x40, 0x8b, 0x0a, 0x8f, 0x30, 0xa4, 0xe1, 0x1e, 0x74, 0x18, 0x79, 0x71, 0x49,
	0xb8, 0x40, 0xba, 0x6a, 0xdb, 0x02, 0x68, 0xe0, 0x13, 0x38, 0x0c, 0x88, 0x7f, 0xe6, 0x4d, 0xe8,
	0x25, 0xdb, 0x21, 0xd3, 0xb4, 0xbf, 0xe9, 0x60, 0x4c, 0xb2, 0x55, 0x14, 0xbf, 0x89, 0xf2, 0x02,
	0x3f, 0x7e, 0x70, 0xba, 0x35, 0xae, 0x2b, 0xdb, 0xe8, 0x2f, 0xd7, 0x5b, 0xd0, 0x89, 0xe2, 0x38,
	0x93, 0x79, 0x5e, 0x9e, 0x69, 0xb2, 0x4d, 0x8a, 0x6d, 0x30, 0x4b, 0x14, 0x4e, 0x55, 0xbe, 0x3f,
	0xf1, 0x81, 0x66, 0x7f, 0xd1, 0xa0, 0xff, 0x60, 0x93, 0x32, 0x83, 0x60, 0x4e, 0xc0, 0x1d, 0x57,
	0x28, 0x33, 0xec, 0x61, 0x0c, 0x83, 0xd2, 0x4e, 0xe1, 0x9c, 0xd1, 0x39, 0xe5, 0xce, 0x0c, 0x69,
	0xf8, 0x10, 0xf6, 0x5d, 0x8f, 0xb8, 0xe7, 0x73, 0xea, 0x07, 0x22, 0xbc, 0xa2, 0x82, 0x20, 0x1d,
	0x23, 0x30, 0x03, 0xb2, 0xa8, 0xed, 0x81, 0x1a, 0xd8, 0x84, 0xee, 0xc2, 0x17, 0xde, 0x94, 0x39,
	0x0b, 0xd4, 0x54, 0x48, 0x2e, 0x08, 0x3b, 0x9f, 0x91, 0x90, 0x51, 0x2a, 0x78, 0xc8, 0xfd, 0xb3,
	0x80, 0x4c, 0x51, 0x4b, 0x71, 0xe4, 0xae, 0x17, 0x50, 0xc6, 0x50, 0x5b, 0x4d, 0x51, 0x85, 0x70,
	0xa3, 0x74, 0x94, 0x52, 0xbd, 0x9b, 0x33, 0x4a, 0x9f, 0xa3, 0xae, 0xfa, 0x48, 0x17, 0x34, 0x20,
	0xaf, 0xc2, 0x7a, 0xba, 0x61, 0x73, 0x30, 0xea, 0xaf, 0x17, 0x1f, 0x43, 0x3b, 0xbd, 0xbb, 0xbe,
	0x95, 0xeb, 0x12, 0xac, 0xc9, 0xaa, 0x0c, 0x9b, 0xa0, 0x65, 0x15, 0x3d, 0x2d, 0x53, 0x59, 0x5e,
	0x9a, 0xc6, 0x64, 0x5a, 0x5e, 0xc3, 0x6d, 0x6e, 0xe1, 0xda, 0xff, 0x81, 0xe1, 0x27, 0x85, 0xcc,
	0x64, 0x5e, 0x94, 0x0d, 0xb7, 0x72, 0x9d, 0x5b, 0xda, 0xb0, 0x31, 0x32, 0x58, 0x19, 0xdb, 0x27,
	0xd0, 0x9a, 0x4b, 0x99, 0xe5, 0x78, 0x00, 0xfa, 0x32, 0xad, 0x4a, 0xfa, 0x32, 0xb5, 0x9f, 0xc2,
	0x41, 0x59, 0x70, 0x92, 0xf8, 0x8f, 0x13, 0x94, 0xff, 0x96, 0x69, 0x6e, 0xe9, 0xa5, 0xa4, 0x42,
	0xfb, 0x19, 0xf4, 0x59, 0x94, 0xc4, 0xab, 0xf7, 0xd5, 0x2f, 0xfe, 0xb7, 0xd7, 0xfc, 0xc2, 0x0e,
	0xd7, 0xed, 0xf2, 0x6f, 0xe8, 0xc9, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa5, 0x3e, 0xea, 0x85,
	0x97, 0x04, 0x00, 0x00,
}
