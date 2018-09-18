// Code generated by protoc-gen-go. DO NOT EDIT.
// source: blockchain.proto

package blockchain

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
	return fileDescriptor_e9ac6287ce250c9a, []int{0}
}

func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transaction.Unmarshal(m, b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
}
func (m *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(m, src)
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
	return fileDescriptor_e9ac6287ce250c9a, []int{1}
}

func (m *Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Block.Unmarshal(m, b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Block.Marshal(b, m, deterministic)
}
func (m *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(m, src)
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
	return fileDescriptor_e9ac6287ce250c9a, []int{2}
}

func (m *Index) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Index.Unmarshal(m, b)
}
func (m *Index) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Index.Marshal(b, m, deterministic)
}
func (m *Index) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Index.Merge(m, src)
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
	return fileDescriptor_e9ac6287ce250c9a, []int{3}
}

func (m *AccountState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountState.Unmarshal(m, b)
}
func (m *AccountState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountState.Marshal(b, m, deterministic)
}
func (m *AccountState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountState.Merge(m, src)
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
	return fileDescriptor_e9ac6287ce250c9a, []int{4}
}

func (m *ContractState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContractState.Unmarshal(m, b)
}
func (m *ContractState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContractState.Marshal(b, m, deterministic)
}
func (m *ContractState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractState.Merge(m, src)
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
	return fileDescriptor_e9ac6287ce250c9a, []int{5}
}

func (m *CasperVote) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CasperVote.Unmarshal(m, b)
}
func (m *CasperVote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CasperVote.Marshal(b, m, deterministic)
}
func (m *CasperVote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CasperVote.Merge(m, src)
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

type ValidatorWithdraw struct {
	PublicKey            string   `protobuf:"bytes,1,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
	R                    []byte   `protobuf:"bytes,2,opt,name=r,proto3" json:"r,omitempty"`
	S                    []byte   `protobuf:"bytes,3,opt,name=s,proto3" json:"s,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidatorWithdraw) Reset()         { *m = ValidatorWithdraw{} }
func (m *ValidatorWithdraw) String() string { return proto.CompactTextString(m) }
func (*ValidatorWithdraw) ProtoMessage()    {}
func (*ValidatorWithdraw) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9ac6287ce250c9a, []int{6}
}

func (m *ValidatorWithdraw) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidatorWithdraw.Unmarshal(m, b)
}
func (m *ValidatorWithdraw) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidatorWithdraw.Marshal(b, m, deterministic)
}
func (m *ValidatorWithdraw) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorWithdraw.Merge(m, src)
}
func (m *ValidatorWithdraw) XXX_Size() int {
	return xxx_messageInfo_ValidatorWithdraw.Size(m)
}
func (m *ValidatorWithdraw) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorWithdraw.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorWithdraw proto.InternalMessageInfo

func (m *ValidatorWithdraw) GetPublicKey() string {
	if m != nil {
		return m.PublicKey
	}
	return ""
}

func (m *ValidatorWithdraw) GetR() []byte {
	if m != nil {
		return m.R
	}
	return nil
}

func (m *ValidatorWithdraw) GetS() []byte {
	if m != nil {
		return m.S
	}
	return nil
}

type Receipt struct {
	Sender               string   `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Recipient            string   `protobuf:"bytes,2,opt,name=recipient,proto3" json:"recipient,omitempty"`
	Amount               uint64   `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Receipt) Reset()         { *m = Receipt{} }
func (m *Receipt) String() string { return proto.CompactTextString(m) }
func (*Receipt) ProtoMessage()    {}
func (*Receipt) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9ac6287ce250c9a, []int{7}
}

func (m *Receipt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Receipt.Unmarshal(m, b)
}
func (m *Receipt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Receipt.Marshal(b, m, deterministic)
}
func (m *Receipt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Receipt.Merge(m, src)
}
func (m *Receipt) XXX_Size() int {
	return xxx_messageInfo_Receipt.Size(m)
}
func (m *Receipt) XXX_DiscardUnknown() {
	xxx_messageInfo_Receipt.DiscardUnknown(m)
}

var xxx_messageInfo_Receipt proto.InternalMessageInfo

func (m *Receipt) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *Receipt) GetRecipient() string {
	if m != nil {
		return m.Recipient
	}
	return ""
}

func (m *Receipt) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func init() {
	proto.RegisterType((*Transaction)(nil), "Transaction")
	proto.RegisterType((*Block)(nil), "Block")
	proto.RegisterType((*Index)(nil), "Index")
	proto.RegisterType((*AccountState)(nil), "AccountState")
	proto.RegisterType((*ContractState)(nil), "ContractState")
	proto.RegisterType((*CasperVote)(nil), "CasperVote")
	proto.RegisterType((*ValidatorWithdraw)(nil), "ValidatorWithdraw")
	proto.RegisterType((*Receipt)(nil), "Receipt")
}

func init() { proto.RegisterFile("blockchain.proto", fileDescriptor_e9ac6287ce250c9a) }

var fileDescriptor_e9ac6287ce250c9a = []byte{
	// 518 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xc1, 0x8e, 0xd3, 0x3c,
	0x10, 0x96, 0x93, 0x34, 0x6d, 0xa6, 0x59, 0xa9, 0xbf, 0xf5, 0x6b, 0x65, 0xad, 0x10, 0x8a, 0x72,
	0x21, 0xe2, 0x50, 0x21, 0xb8, 0x23, 0x2d, 0xbd, 0x2c, 0x42, 0x5c, 0x0c, 0xda, 0x3d, 0xbb, 0xae,
	0x69, 0x2d, 0x12, 0x3b, 0xb2, 0x5d, 0x60, 0x9f, 0x86, 0x27, 0xe0, 0x15, 0x78, 0x36, 0x64, 0x3b,
	0xdd, 0x34, 0xe5, 0xc2, 0xcd, 0xdf, 0xa7, 0x99, 0xf1, 0xcc, 0xf7, 0xcd, 0xc0, 0x6a, 0xdb, 0x6a,
	0xfe, 0x95, 0x1f, 0x98, 0x54, 0xeb, 0xde, 0x68, 0xa7, 0xeb, 0x5f, 0x09, 0x2c, 0x3f, 0x1b, 0xa6,
	0x2c, 0xe3, 0x4e, 0x6a, 0x85, 0xaf, 0x21, 0xb7, 0x42, 0xed, 0x84, 0x21, 0xa8, 0x42, 0x4d, 0x49,
	0x07, 0x84, 0x9f, 0x41, 0x61, 0x04, 0x97, 0xbd, 0x14, 0xca, 0x91, 0xa4, 0x42, 0x4d, 0x41, 0x47,
	0x02, 0xff, 0x0f, 0x33, 0xa5, 0x15, 0x17, 0x24, 0xad, 0x50, 0x73, 0x45, 0x23, 0xf0, 0xb5, 0x58,
	0xa7, 0x8f, 0xca, 0x91, 0xac, 0x42, 0x4d, 0x46, 0x07, 0x84, 0x57, 0x90, 0xee, 0x99, 0x25, 0xb3,
	0x10, 0xeb, 0x9f, 0xbe, 0xba, 0x93, 0x9d, 0xb0, 0x8e, 0x75, 0x3d, 0xc9, 0x43, 0xf0, 0x48, 0xe0,
	0x97, 0xb0, 0xe2, 0x5a, 0x39, 0xc3, 0xb8, 0xdb, 0x18, 0xc1, 0x7c, 0x9f, 0x64, 0x5e, 0xa1, 0x66,
	0x41, 0xff, 0xe2, 0xf1, 0x0d, 0x2c, 0xbe, 0x1c, 0x55, 0x98, 0x85, 0x2c, 0x42, 0x9b, 0x4f, 0x18,
	0x63, 0xc8, 0x98, 0xd9, 0x5b, 0x52, 0x54, 0x69, 0x93, 0xd1, 0xf0, 0xf6, 0xdc, 0x8e, 0x39, 0x46,
	0x20, 0x4c, 0x1b, 0xde, 0xb8, 0x04, 0x64, 0xc8, 0x32, 0x10, 0xc8, 0x78, 0x64, 0x49, 0x19, 0x91,
	0xad, 0x7f, 0x22, 0x98, 0xbd, 0xf3, 0x22, 0xfa, 0x99, 0xa5, 0xda, 0x89, 0x1f, 0x41, 0xa8, 0x8c,
	0x46, 0x30, 0x9d, 0x24, 0xb9, 0x9c, 0xe4, 0x06, 0x16, 0xbd, 0x11, 0xdf, 0xee, 0x98, 0x3d, 0x04,
	0xa9, 0x4a, 0xfa, 0x84, 0x7d, 0xbd, 0x4e, 0x2a, 0x61, 0x82, 0x58, 0x05, 0x8d, 0x00, 0xbf, 0x82,
	0xd2, 0x8d, 0xf6, 0x78, 0xd1, 0xd2, 0x66, 0xf9, 0xba, 0x5c, 0x9f, 0x79, 0x46, 0x27, 0x11, 0xf5,
	0x0b, 0x98, 0xbd, 0x0f, 0xad, 0x3c, 0x87, 0x3c, 0xd8, 0x6d, 0x09, 0x0a, 0x49, 0xf9, 0x3a, 0x34,
	0x4e, 0x07, 0xb6, 0x7e, 0x0b, 0xe5, 0x2d, 0xe7, 0xde, 0x91, 0x4f, 0x8e, 0x39, 0x81, 0x09, 0xcc,
	0xb7, 0xac, 0x65, 0xde, 0xc6, 0x38, 0xd2, 0x09, 0x8e, 0xf6, 0x26, 0x67, 0xf6, 0xd6, 0xb7, 0x70,
	0xb5, 0x19, 0xe4, 0x8f, 0x05, 0xae, 0x21, 0xef, 0x44, 0xa7, 0xcd, 0xe3, 0x69, 0x77, 0x22, 0xf2,
	0x85, 0xf7, 0xad, 0xde, 0xb2, 0xd6, 0x92, 0x24, 0x48, 0x7f, 0x82, 0xf5, 0x6f, 0x04, 0xb0, 0x61,
	0xb6, 0x17, 0xe6, 0x5e, 0xc7, 0x02, 0x56, 0x1f, 0xcd, 0xd0, 0x80, 0x5f, 0xbe, 0x80, 0x3c, 0xef,
	0x98, 0xd9, 0x8b, 0xb8, 0x79, 0x25, 0x1d, 0x10, 0xae, 0xa1, 0x8c, 0x11, 0x77, 0x42, 0xee, 0x0f,
	0x2e, 0x48, 0x9a, 0xd1, 0x09, 0xe7, 0x63, 0x62, 0xf4, 0x10, 0x13, 0x57, 0x71, 0xc2, 0x45, 0xc3,
	0x67, 0x13, 0xc3, 0xf3, 0xc1, 0x70, 0x6f, 0x68, 0x7f, 0xdc, 0xb6, 0x92, 0x7f, 0x10, 0x8f, 0x61,
	0xeb, 0x0a, 0x3a, 0x12, 0xf5, 0x47, 0xf8, 0xef, 0x9e, 0xb5, 0x72, 0xc7, 0x9c, 0x36, 0x0f, 0xd2,
	0x1d, 0x76, 0x86, 0x7d, 0x9f, 0xa6, 0xa0, 0x8b, 0x94, 0xf8, 0x59, 0x32, 0xf9, 0x2c, 0x3d, 0x6d,
	0xd7, 0x03, 0xcc, 0xa9, 0xe0, 0x42, 0xf6, 0xee, 0xe2, 0x10, 0x8b, 0x7f, 0x3c, 0xc4, 0xf1, 0xe4,
	0xd2, 0xf3, 0x93, 0xdb, 0xe6, 0xe1, 0xda, 0xdf, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0xe4, 0xc9,
	0xa2, 0x00, 0x01, 0x04, 0x00, 0x00,
}
