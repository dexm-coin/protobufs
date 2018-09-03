// Code generated by protoc-gen-go. DO NOT EDIT.
// source: blockchain.proto

package blockchain

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

type Transaction struct {
	Sender               []byte   `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Recipient            string   `protobuf:"bytes,2,opt,name=recipient,proto3" json:"recipient,omitempty"`
	Nonce                uint32   `protobuf:"varint,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Amount               uint64   `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Gas                  uint32   `protobuf:"varint,5,opt,name=gas,proto3" json:"gas,omitempty"`
	Timestamp            uint64   `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	ContractCreation     bool     `protobuf:"varint,7,opt,name=contractCreation,proto3" json:"contractCreation,omitempty"`
	Function             string   `protobuf:"bytes,8,opt,name=function,proto3" json:"function,omitempty"`
	Args                 []uint64 `protobuf:"varint,9,rep,packed,name=args,proto3" json:"args,omitempty"`
	Data                 []byte   `protobuf:"bytes,10,opt,name=data,proto3" json:"data,omitempty"`
	R                    []byte   `protobuf:"bytes,11,opt,name=r,proto3" json:"r,omitempty"`
	S                    []byte   `protobuf:"bytes,12,opt,name=s,proto3" json:"s,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_ac7a21bdb2afe16f, []int{0}
}
func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transaction.Unmarshal(m, b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
}
func (dst *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(dst, src)
}
func (m *Transaction) XXX_Size() int {
	return xxx_messageInfo_Transaction.Size(m)
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetSender() []byte {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *Transaction) GetRecipient() string {
	if m != nil {
		return m.Recipient
	}
	return ""
}

func (m *Transaction) GetNonce() uint32 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *Transaction) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Transaction) GetGas() uint32 {
	if m != nil {
		return m.Gas
	}
	return 0
}

func (m *Transaction) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Transaction) GetContractCreation() bool {
	if m != nil {
		return m.ContractCreation
	}
	return false
}

func (m *Transaction) GetFunction() string {
	if m != nil {
		return m.Function
	}
	return ""
}

func (m *Transaction) GetArgs() []uint64 {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *Transaction) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Transaction) GetR() []byte {
	if m != nil {
		return m.R
	}
	return nil
}

func (m *Transaction) GetS() []byte {
	if m != nil {
		return m.S
	}
	return nil
}

type Block struct {
	Index                uint64         `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Timestamp            uint64         `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	PrevHash             []byte         `protobuf:"bytes,3,opt,name=prevHash,proto3" json:"prevHash,omitempty"`
	Miner                string         `protobuf:"bytes,4,opt,name=miner,proto3" json:"miner,omitempty"`
	Transactions         []*Transaction `protobuf:"bytes,5,rep,name=transactions,proto3" json:"transactions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_ac7a21bdb2afe16f, []int{1}
}
func (m *Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Block.Unmarshal(m, b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Block.Marshal(b, m, deterministic)
}
func (dst *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(dst, src)
}
func (m *Block) XXX_Size() int {
	return xxx_messageInfo_Block.Size(m)
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Block) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Block) GetPrevHash() []byte {
	if m != nil {
		return m.PrevHash
	}
	return nil
}

func (m *Block) GetMiner() string {
	if m != nil {
		return m.Miner
	}
	return ""
}

func (m *Block) GetTransactions() []*Transaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

type Index struct {
	Blocks               []*Block `protobuf:"bytes,1,rep,name=blocks,proto3" json:"blocks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Index) Reset()         { *m = Index{} }
func (m *Index) String() string { return proto.CompactTextString(m) }
func (*Index) ProtoMessage()    {}
func (*Index) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_ac7a21bdb2afe16f, []int{2}
}
func (m *Index) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Index.Unmarshal(m, b)
}
func (m *Index) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Index.Marshal(b, m, deterministic)
}
func (dst *Index) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Index.Merge(dst, src)
}
func (m *Index) XXX_Size() int {
	return xxx_messageInfo_Index.Size(m)
}
func (m *Index) XXX_DiscardUnknown() {
	xxx_messageInfo_Index.DiscardUnknown(m)
}

var xxx_messageInfo_Index proto.InternalMessageInfo

func (m *Index) GetBlocks() []*Block {
	if m != nil {
		return m.Blocks
	}
	return nil
}

type AccountState struct {
	Balance              uint64   `protobuf:"varint,1,opt,name=balance,proto3" json:"balance,omitempty"`
	Nonce                uint32   `protobuf:"varint,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountState) Reset()         { *m = AccountState{} }
func (m *AccountState) String() string { return proto.CompactTextString(m) }
func (*AccountState) ProtoMessage()    {}
func (*AccountState) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_ac7a21bdb2afe16f, []int{3}
}
func (m *AccountState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountState.Unmarshal(m, b)
}
func (m *AccountState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountState.Marshal(b, m, deterministic)
}
func (dst *AccountState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountState.Merge(dst, src)
}
func (m *AccountState) XXX_Size() int {
	return xxx_messageInfo_AccountState.Size(m)
}
func (m *AccountState) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountState.DiscardUnknown(m)
}

var xxx_messageInfo_AccountState proto.InternalMessageInfo

func (m *AccountState) GetBalance() uint64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *AccountState) GetNonce() uint32 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

type ContractState struct {
	Memory               []byte   `protobuf:"bytes,1,opt,name=memory,proto3" json:"memory,omitempty"`
	Globals              []uint64 `protobuf:"varint,2,rep,packed,name=globals,proto3" json:"globals,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContractState) Reset()         { *m = ContractState{} }
func (m *ContractState) String() string { return proto.CompactTextString(m) }
func (*ContractState) ProtoMessage()    {}
func (*ContractState) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_ac7a21bdb2afe16f, []int{4}
}
func (m *ContractState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContractState.Unmarshal(m, b)
}
func (m *ContractState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContractState.Marshal(b, m, deterministic)
}
func (dst *ContractState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractState.Merge(dst, src)
}
func (m *ContractState) XXX_Size() int {
	return xxx_messageInfo_ContractState.Size(m)
}
func (m *ContractState) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractState.DiscardUnknown(m)
}

var xxx_messageInfo_ContractState proto.InternalMessageInfo

func (m *ContractState) GetMemory() []byte {
	if m != nil {
		return m.Memory
	}
	return nil
}

func (m *ContractState) GetGlobals() []uint64 {
	if m != nil {
		return m.Globals
	}
	return nil
}

type CasperVote struct {
	Source               []byte   `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Target               []byte   `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	SourceHeight         uint64   `protobuf:"varint,3,opt,name=sourceHeight,proto3" json:"sourceHeight,omitempty"`
	TargetHeight         uint64   `protobuf:"varint,4,opt,name=targetHeight,proto3" json:"targetHeight,omitempty"`
	R                    []byte   `protobuf:"bytes,5,opt,name=r,proto3" json:"r,omitempty"`
	S                    []byte   `protobuf:"bytes,6,opt,name=s,proto3" json:"s,omitempty"`
	PublicKey            string   `protobuf:"bytes,7,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CasperVote) Reset()         { *m = CasperVote{} }
func (m *CasperVote) String() string { return proto.CompactTextString(m) }
func (*CasperVote) ProtoMessage()    {}
func (*CasperVote) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_ac7a21bdb2afe16f, []int{5}
}
func (m *CasperVote) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CasperVote.Unmarshal(m, b)
}
func (m *CasperVote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CasperVote.Marshal(b, m, deterministic)
}
func (dst *CasperVote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CasperVote.Merge(dst, src)
}
func (m *CasperVote) XXX_Size() int {
	return xxx_messageInfo_CasperVote.Size(m)
}
func (m *CasperVote) XXX_DiscardUnknown() {
	xxx_messageInfo_CasperVote.DiscardUnknown(m)
}

var xxx_messageInfo_CasperVote proto.InternalMessageInfo

func (m *CasperVote) GetSource() []byte {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *CasperVote) GetTarget() []byte {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *CasperVote) GetSourceHeight() uint64 {
	if m != nil {
		return m.SourceHeight
	}
	return 0
}

func (m *CasperVote) GetTargetHeight() uint64 {
	if m != nil {
		return m.TargetHeight
	}
	return 0
}

func (m *CasperVote) GetR() []byte {
	if m != nil {
		return m.R
	}
	return nil
}

func (m *CasperVote) GetS() []byte {
	if m != nil {
		return m.S
	}
	return nil
}

func (m *CasperVote) GetPublicKey() string {
	if m != nil {
		return m.PublicKey
	}
	return ""
}

func init() {
	proto.RegisterType((*Transaction)(nil), "Transaction")
	proto.RegisterType((*Block)(nil), "Block")
	proto.RegisterType((*Index)(nil), "Index")
	proto.RegisterType((*AccountState)(nil), "AccountState")
	proto.RegisterType((*ContractState)(nil), "ContractState")
	proto.RegisterType((*CasperVote)(nil), "CasperVote")
}

func init() { proto.RegisterFile("blockchain.proto", fileDescriptor_blockchain_ac7a21bdb2afe16f) }

var fileDescriptor_blockchain_ac7a21bdb2afe16f = []byte{
	// 466 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x93, 0xc1, 0x8a, 0xdb, 0x30,
	0x10, 0x86, 0x51, 0xe2, 0x78, 0x93, 0x89, 0x16, 0x82, 0x28, 0x8b, 0x58, 0x4a, 0x31, 0xbe, 0xd4,
	0xf4, 0x10, 0x4a, 0x7b, 0x2f, 0x6c, 0x73, 0xd9, 0xd2, 0x9b, 0x5a, 0x7a, 0x97, 0x15, 0xd5, 0x31,
	0xb5, 0x25, 0x23, 0x29, 0xa5, 0xfb, 0x34, 0x7d, 0x82, 0xbe, 0x42, 0x9f, 0xad, 0x68, 0xe4, 0xac,
	0xd7, 0xdd, 0x9b, 0xbe, 0x9f, 0xf1, 0x30, 0xf3, 0xff, 0x63, 0xd8, 0xd5, 0x9d, 0x55, 0x3f, 0xd4,
	0x49, 0xb6, 0x66, 0x3f, 0x38, 0x1b, 0x6c, 0xf9, 0x67, 0x01, 0xdb, 0xaf, 0x4e, 0x1a, 0x2f, 0x55,
	0x68, 0xad, 0x61, 0x37, 0x90, 0x7b, 0x6d, 0x8e, 0xda, 0x71, 0x52, 0x90, 0x8a, 0x8a, 0x91, 0xd8,
	0x4b, 0xd8, 0x38, 0xad, 0xda, 0xa1, 0xd5, 0x26, 0xf0, 0x45, 0x41, 0xaa, 0x8d, 0x98, 0x04, 0xf6,
	0x02, 0x56, 0xc6, 0x1a, 0xa5, 0xf9, 0xb2, 0x20, 0xd5, 0xb5, 0x48, 0x10, 0x7b, 0xc9, 0xde, 0x9e,
	0x4d, 0xe0, 0x59, 0x41, 0xaa, 0x4c, 0x8c, 0xc4, 0x76, 0xb0, 0x6c, 0xa4, 0xe7, 0x2b, 0xac, 0x8d,
	0xcf, 0xd8, 0x3d, 0xb4, 0xbd, 0xf6, 0x41, 0xf6, 0x03, 0xcf, 0xb1, 0x78, 0x12, 0xd8, 0x1b, 0xd8,
	0x29, 0x6b, 0x82, 0x93, 0x2a, 0x1c, 0x9c, 0x96, 0x71, 0x4e, 0x7e, 0x55, 0x90, 0x6a, 0x2d, 0x9e,
	0xe9, 0xec, 0x16, 0xd6, 0xdf, 0xcf, 0x06, 0x77, 0xe1, 0x6b, 0x1c, 0xf3, 0x91, 0x19, 0x83, 0x4c,
	0xba, 0xc6, 0xf3, 0x4d, 0xb1, 0xac, 0x32, 0x81, 0xef, 0xa8, 0x1d, 0x65, 0x90, 0x1c, 0x70, 0x5b,
	0x7c, 0x33, 0x0a, 0xc4, 0xf1, 0x2d, 0x0a, 0xc4, 0x45, 0xf2, 0x9c, 0x26, 0xf2, 0xe5, 0x6f, 0x02,
	0xab, 0x8f, 0xd1, 0xc4, 0xb8, 0x73, 0x6b, 0x8e, 0xfa, 0x17, 0x1a, 0x95, 0x89, 0x04, 0xf3, 0x4d,
	0x16, 0xff, 0x6f, 0x72, 0x0b, 0xeb, 0xc1, 0xe9, 0x9f, 0xf7, 0xd2, 0x9f, 0xd0, 0x2a, 0x2a, 0x1e,
	0x39, 0xf6, 0xeb, 0x5b, 0xa3, 0x1d, 0x9a, 0xb5, 0x11, 0x09, 0xd8, 0x5b, 0xa0, 0x61, 0x8a, 0x27,
	0x9a, 0xb6, 0xac, 0xb6, 0xef, 0xe8, 0xfe, 0x49, 0x66, 0x62, 0x56, 0x51, 0xbe, 0x86, 0xd5, 0x27,
	0x1c, 0xe5, 0x15, 0xe4, 0x18, 0xb7, 0xe7, 0x04, 0x3f, 0xca, 0xf7, 0x38, 0xb8, 0x18, 0xd5, 0xf2,
	0x03, 0xd0, 0x3b, 0xa5, 0x62, 0x22, 0x5f, 0x82, 0x0c, 0x9a, 0x71, 0xb8, 0xaa, 0x65, 0x27, 0x63,
	0x8c, 0x69, 0xa5, 0x0b, 0x4e, 0xf1, 0x2e, 0x9e, 0xc4, 0x5b, 0xde, 0xc1, 0xf5, 0x61, 0xb4, 0x3f,
	0x35, 0xb8, 0x81, 0xbc, 0xd7, 0xbd, 0x75, 0x0f, 0x97, 0xdb, 0x49, 0x14, 0x1b, 0x37, 0x9d, 0xad,
	0x65, 0xe7, 0xf9, 0x02, 0xad, 0xbf, 0x60, 0xf9, 0x97, 0x00, 0x1c, 0xa4, 0x1f, 0xb4, 0xfb, 0x66,
	0x53, 0x03, 0x6f, 0xcf, 0x6e, 0x1c, 0x20, 0x1e, 0x1f, 0x52, 0xd4, 0x83, 0x74, 0x8d, 0x4e, 0x97,
	0x47, 0xc5, 0x48, 0xac, 0x04, 0x9a, 0x2a, 0xee, 0x75, 0xdb, 0x9c, 0x02, 0x5a, 0x9a, 0x89, 0x99,
	0x16, 0x6b, 0x52, 0xf5, 0x58, 0x93, 0x4e, 0x71, 0xa6, 0xa5, 0xc0, 0x57, 0xb3, 0xc0, 0xf3, 0x31,
	0xf0, 0x18, 0xe8, 0x70, 0xae, 0xbb, 0x56, 0x7d, 0xd6, 0x0f, 0x78, 0x75, 0x54, 0x4c, 0x42, 0x9d,
	0xe3, 0x5f, 0xf4, 0xfe, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x19, 0xda, 0x41, 0x65, 0x59, 0x03,
	0x00, 0x00,
}
