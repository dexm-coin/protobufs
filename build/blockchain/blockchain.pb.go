// Code generated by protoc-gen-go. DO NOT EDIT.
// source: blockchain.proto

/*
Package blockchain is a generated protocol buffer package.

It is generated from these files:
	blockchain.proto

It has these top-level messages:
	Transaction
	Block
	AccountState
*/
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
	Sender    []byte `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Recipient string `protobuf:"bytes,2,opt,name=recipient" json:"recipient,omitempty"`
	Nonce     uint32 `protobuf:"varint,3,opt,name=nonce" json:"nonce,omitempty"`
	Amount    uint64 `protobuf:"varint,4,opt,name=amount" json:"amount,omitempty"`
	Gas       uint32 `protobuf:"varint,5,opt,name=gas" json:"gas,omitempty"`
	Timestamp uint64 `protobuf:"varint,6,opt,name=timestamp" json:"timestamp,omitempty"`
	R         []byte `protobuf:"bytes,7,opt,name=r,proto3" json:"r,omitempty"`
	S         []byte `protobuf:"bytes,8,opt,name=s,proto3" json:"s,omitempty"`
}

func (m *Transaction) Reset()                    { *m = Transaction{} }
func (m *Transaction) String() string            { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()               {}
func (*Transaction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

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
	Index        uint64         `protobuf:"varint,1,opt,name=index" json:"index,omitempty"`
	Timestamp    uint64         `protobuf:"varint,2,opt,name=timestamp" json:"timestamp,omitempty"`
	PrevHash     []byte         `protobuf:"bytes,3,opt,name=prevHash,proto3" json:"prevHash,omitempty"`
	Miner        string         `protobuf:"bytes,4,opt,name=miner" json:"miner,omitempty"`
	Transactions []*Transaction `protobuf:"bytes,5,rep,name=transactions" json:"transactions,omitempty"`
}

func (m *Block) Reset()                    { *m = Block{} }
func (m *Block) String() string            { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()               {}
func (*Block) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

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

type AccountState struct {
	Balance uint64 `protobuf:"varint,1,opt,name=balance" json:"balance,omitempty"`
	Nonce   uint32 `protobuf:"varint,2,opt,name=nonce" json:"nonce,omitempty"`
}

func (m *AccountState) Reset()                    { *m = AccountState{} }
func (m *AccountState) String() string            { return proto.CompactTextString(m) }
func (*AccountState) ProtoMessage()               {}
func (*AccountState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

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

func init() {
	proto.RegisterType((*Transaction)(nil), "Transaction")
	proto.RegisterType((*Block)(nil), "Block")
	proto.RegisterType((*AccountState)(nil), "AccountState")
}

func init() { proto.RegisterFile("blockchain.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0xcf, 0x4a, 0x03, 0x31,
	0x10, 0xc6, 0xc9, 0xfe, 0x6b, 0x77, 0x1a, 0xa1, 0x04, 0x91, 0x20, 0x1e, 0x96, 0x3d, 0xed, 0xa9,
	0x88, 0xde, 0x05, 0x3d, 0x79, 0x8e, 0xbe, 0x40, 0x36, 0x0d, 0x36, 0xd8, 0x4d, 0x96, 0x24, 0x8a,
	0x6f, 0xe3, 0x93, 0xf8, 0x6e, 0x32, 0xd9, 0xea, 0xb6, 0xbd, 0xe5, 0x37, 0x0c, 0x1f, 0xbf, 0xf9,
	0x02, 0xeb, 0x7e, 0xef, 0xd4, 0xbb, 0xda, 0x49, 0x63, 0x37, 0xa3, 0x77, 0xd1, 0xb5, 0x3f, 0x04,
	0x56, 0xaf, 0x5e, 0xda, 0x20, 0x55, 0x34, 0xce, 0xb2, 0x2b, 0xa8, 0x82, 0xb6, 0x5b, 0xed, 0x39,
	0x69, 0x48, 0x47, 0xc5, 0x81, 0xd8, 0x0d, 0xd4, 0x5e, 0x2b, 0x33, 0x1a, 0x6d, 0x23, 0xcf, 0x1a,
	0xd2, 0xd5, 0x62, 0x1e, 0xb0, 0x4b, 0x28, 0xad, 0xb3, 0x4a, 0xf3, 0xbc, 0x21, 0xdd, 0x85, 0x98,
	0x00, 0xb3, 0xe4, 0xe0, 0x3e, 0x6c, 0xe4, 0x45, 0x43, 0xba, 0x42, 0x1c, 0x88, 0xad, 0x21, 0x7f,
	0x93, 0x81, 0x97, 0x69, 0x17, 0x9f, 0x98, 0x1e, 0xcd, 0xa0, 0x43, 0x94, 0xc3, 0xc8, 0xab, 0xb4,
	0x3c, 0x0f, 0x18, 0x05, 0xe2, 0xf9, 0x22, 0xe9, 0x10, 0x8f, 0x14, 0xf8, 0x72, 0xa2, 0xd0, 0x7e,
	0x13, 0x28, 0x9f, 0xf0, 0x28, 0x74, 0x30, 0x76, 0xab, 0xbf, 0x92, 0x78, 0x21, 0x26, 0x38, 0x4d,
	0xce, 0xce, 0x93, 0xaf, 0x61, 0x39, 0x7a, 0xfd, 0xf9, 0x2c, 0xc3, 0x2e, 0xa9, 0x53, 0xf1, 0xcf,
	0x98, 0x37, 0x18, 0xab, 0x7d, 0x92, 0xaf, 0xc5, 0x04, 0xec, 0x16, 0x68, 0x9c, 0xeb, 0xc2, 0x23,
	0xf2, 0x6e, 0x75, 0x47, 0x37, 0x47, 0x1d, 0x8a, 0x93, 0x8d, 0xf6, 0x01, 0xe8, 0xa3, 0x52, 0x78,
	0xf8, 0x4b, 0x94, 0x51, 0x33, 0x0e, 0x8b, 0x5e, 0xee, 0x25, 0xb6, 0x35, 0x99, 0xfe, 0xe1, 0xdc,
	0x62, 0x76, 0xd4, 0x62, 0x5f, 0xa5, 0x8f, 0xba, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x56, 0xa7,
	0x6e, 0x85, 0xbc, 0x01, 0x00, 0x00,
}
