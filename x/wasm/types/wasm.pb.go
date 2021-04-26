// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: terra/wasm/v1beta1/wasm.proto

package types

import (
	bytes "bytes"
	encoding_json "encoding/json"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

// Params defines the parameters for the wasm module.
type Params struct {
	MaxContractSize     uint64      `protobuf:"varint,1,opt,name=max_contract_size,json=maxContractSize,proto3" json:"max_contract_size,omitempty" yaml:"max_contract_size"`
	MaxContractGas      uint64      `protobuf:"varint,2,opt,name=max_contract_gas,json=maxContractGas,proto3" json:"max_contract_gas,omitempty" yaml:"max_contract_gas"`
	MaxContractMsgSize  uint64      `protobuf:"varint,3,opt,name=max_contract_msg_size,json=maxContractMsgSize,proto3" json:"max_contract_msg_size,omitempty" yaml:"max_contract_msg_size"`
	MaxContractDataSize uint64      `protobuf:"varint,4,opt,name=max_contract_data_size,json=maxContractDataSize,proto3" json:"max_contract_data_size,omitempty" yaml:"max_contract_data_size"`
	EventParams         EventParams `protobuf:"bytes,5,opt,name=event_params,json=eventParams,proto3" json:"event_params" yaml:"event_params"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bd5d0123068c880, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetMaxContractSize() uint64 {
	if m != nil {
		return m.MaxContractSize
	}
	return 0
}

func (m *Params) GetMaxContractGas() uint64 {
	if m != nil {
		return m.MaxContractGas
	}
	return 0
}

func (m *Params) GetMaxContractMsgSize() uint64 {
	if m != nil {
		return m.MaxContractMsgSize
	}
	return 0
}

func (m *Params) GetMaxContractDataSize() uint64 {
	if m != nil {
		return m.MaxContractDataSize
	}
	return 0
}

func (m *Params) GetEventParams() EventParams {
	if m != nil {
		return m.EventParams
	}
	return EventParams{}
}

// EventParams defines the event related parameteres
type EventParams struct {
	MaxAttributeNum         uint64 `protobuf:"varint,1,opt,name=max_attribute_num,json=maxAttributeNum,proto3" json:"max_attribute_num,omitempty" yaml:"max_contract_event_attribute_num"`
	MaxAttributeKeyLength   uint64 `protobuf:"varint,2,opt,name=max_attribute_key_length,json=maxAttributeKeyLength,proto3" json:"max_attribute_key_length,omitempty" yaml:"max_contract_event_attribute_key_length"`
	MaxAttributeValueLength uint64 `protobuf:"varint,3,opt,name=max_attribute_value_length,json=maxAttributeValueLength,proto3" json:"max_attribute_value_length,omitempty" yaml:"max_contract_event_attribute_value_length"`
}

func (m *EventParams) Reset()         { *m = EventParams{} }
func (m *EventParams) String() string { return proto.CompactTextString(m) }
func (*EventParams) ProtoMessage()    {}
func (*EventParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bd5d0123068c880, []int{1}
}
func (m *EventParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventParams.Merge(m, src)
}
func (m *EventParams) XXX_Size() int {
	return m.Size()
}
func (m *EventParams) XXX_DiscardUnknown() {
	xxx_messageInfo_EventParams.DiscardUnknown(m)
}

var xxx_messageInfo_EventParams proto.InternalMessageInfo

func (m *EventParams) GetMaxAttributeNum() uint64 {
	if m != nil {
		return m.MaxAttributeNum
	}
	return 0
}

func (m *EventParams) GetMaxAttributeKeyLength() uint64 {
	if m != nil {
		return m.MaxAttributeKeyLength
	}
	return 0
}

func (m *EventParams) GetMaxAttributeValueLength() uint64 {
	if m != nil {
		return m.MaxAttributeValueLength
	}
	return 0
}

// CodeInfo is data for the uploaded contract WASM code
type CodeInfo struct {
	// CodeID is the sequentially increasing unique identifier
	CodeID uint64 `protobuf:"varint,1,opt,name=code_id,json=codeId,proto3" json:"code_id,omitempty" yaml:"code_id"`
	// CodeHash is the unique identifier created by wasmvm
	CodeHash []byte `protobuf:"bytes,2,opt,name=code_hash,json=codeHash,proto3" json:"code_hash,omitempty" yaml:"code_hash"`
	// Creator address who initially stored the code
	Creator string `protobuf:"bytes,3,opt,name=creator,proto3" json:"creator,omitempty" yaml:"creator"`
}

func (m *CodeInfo) Reset()         { *m = CodeInfo{} }
func (m *CodeInfo) String() string { return proto.CompactTextString(m) }
func (*CodeInfo) ProtoMessage()    {}
func (*CodeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bd5d0123068c880, []int{2}
}
func (m *CodeInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CodeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CodeInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CodeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CodeInfo.Merge(m, src)
}
func (m *CodeInfo) XXX_Size() int {
	return m.Size()
}
func (m *CodeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CodeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CodeInfo proto.InternalMessageInfo

func (m *CodeInfo) GetCodeID() uint64 {
	if m != nil {
		return m.CodeID
	}
	return 0
}

func (m *CodeInfo) GetCodeHash() []byte {
	if m != nil {
		return m.CodeHash
	}
	return nil
}

func (m *CodeInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

// ContractInfo stores a WASM contract instance
type ContractInfo struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
	// Owner address that can execute migrations
	Owner string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty" yaml:"owner"`
	// CodeID is the reference to the stored Wasm code
	CodeID uint64 `protobuf:"varint,3,opt,name=code_id,json=codeId,proto3" json:"code_id,omitempty" yaml:"code_id"`
	// InitMsg is the raw message used when instantiating a contract
	InitMsg encoding_json.RawMessage `protobuf:"bytes,4,opt,name=init_msg,json=initMsg,proto3,casttype=encoding/json.RawMessage" json:"init_msg,omitempty" yaml:"init_msg"`
	// Migratable is the flag to specify the contract migratability
	Migratable bool `protobuf:"varint,5,opt,name=migratable,proto3" json:"migratable,omitempty" yaml:"migratable"`
}

func (m *ContractInfo) Reset()         { *m = ContractInfo{} }
func (m *ContractInfo) String() string { return proto.CompactTextString(m) }
func (*ContractInfo) ProtoMessage()    {}
func (*ContractInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bd5d0123068c880, []int{3}
}
func (m *ContractInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContractInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContractInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContractInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractInfo.Merge(m, src)
}
func (m *ContractInfo) XXX_Size() int {
	return m.Size()
}
func (m *ContractInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ContractInfo proto.InternalMessageInfo

func (m *ContractInfo) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ContractInfo) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ContractInfo) GetCodeID() uint64 {
	if m != nil {
		return m.CodeID
	}
	return 0
}

func (m *ContractInfo) GetInitMsg() encoding_json.RawMessage {
	if m != nil {
		return m.InitMsg
	}
	return nil
}

func (m *ContractInfo) GetMigratable() bool {
	if m != nil {
		return m.Migratable
	}
	return false
}

func init() {
	proto.RegisterType((*Params)(nil), "terra.wasm.v1beta1.Params")
	proto.RegisterType((*EventParams)(nil), "terra.wasm.v1beta1.EventParams")
	proto.RegisterType((*CodeInfo)(nil), "terra.wasm.v1beta1.CodeInfo")
	proto.RegisterType((*ContractInfo)(nil), "terra.wasm.v1beta1.ContractInfo")
}

func init() { proto.RegisterFile("terra/wasm/v1beta1/wasm.proto", fileDescriptor_2bd5d0123068c880) }

var fileDescriptor_2bd5d0123068c880 = []byte{
	// 721 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x4d, 0x4f, 0x1b, 0x39,
	0x18, 0xc7, 0x33, 0xbc, 0x24, 0xc1, 0x44, 0xbc, 0x18, 0x58, 0x22, 0x5e, 0x62, 0xd6, 0x87, 0x5d,
	0xb4, 0xcb, 0x66, 0x16, 0x76, 0xb9, 0x70, 0xdb, 0x01, 0xb4, 0xb0, 0x5b, 0xaa, 0xca, 0x48, 0x54,
	0xea, 0x25, 0x72, 0x66, 0xdc, 0xc9, 0x40, 0x66, 0x9c, 0x8e, 0x1d, 0x20, 0x7c, 0x8a, 0x1e, 0x7b,
	0x44, 0x55, 0xbf, 0x40, 0xbf, 0x05, 0x87, 0x1e, 0x38, 0xf6, 0x64, 0x55, 0xe1, 0xd2, 0xf3, 0x1c,
	0x7b, 0xaa, 0x62, 0x67, 0x60, 0x22, 0x72, 0xa0, 0x37, 0xe7, 0xf9, 0xff, 0xfd, 0xfb, 0x3f, 0xf6,
	0x13, 0x0f, 0x58, 0x95, 0x2c, 0x8e, 0xa9, 0x7d, 0x41, 0x45, 0x68, 0x9f, 0x6f, 0xd6, 0x99, 0xa4,
	0x9b, 0xfa, 0x47, 0xb5, 0x15, 0x73, 0xc9, 0x21, 0xd4, 0x72, 0x55, 0x57, 0xfa, 0xf2, 0xd2, 0xbc,
	0xcf, 0x7d, 0xae, 0x65, 0xbb, 0xb7, 0x32, 0xce, 0xa5, 0x8a, 0xcb, 0x45, 0xc8, 0x85, 0x5d, 0xa7,
	0x82, 0xdd, 0x93, 0x5c, 0x1e, 0x44, 0x46, 0xc7, 0x1f, 0x47, 0x41, 0xfe, 0x05, 0x8d, 0x69, 0x28,
	0xe0, 0x01, 0x98, 0x0d, 0xe9, 0x65, 0xcd, 0xe5, 0x91, 0x8c, 0xa9, 0x2b, 0x6b, 0x22, 0xb8, 0x62,
	0x65, 0x6b, 0xcd, 0x5a, 0x1f, 0x73, 0x56, 0x12, 0x85, 0xca, 0x1d, 0x1a, 0x36, 0x77, 0xf0, 0x23,
	0x0b, 0x26, 0xd3, 0x21, 0xbd, 0xdc, 0xed, 0x97, 0x8e, 0x83, 0x2b, 0x06, 0xf7, 0xc1, 0xcc, 0x80,
	0xcd, 0xa7, 0xa2, 0x3c, 0xa2, 0x41, 0xcb, 0x89, 0x42, 0x8b, 0x43, 0x40, 0x3e, 0x15, 0x98, 0x4c,
	0x65, 0x38, 0xff, 0x52, 0x01, 0x8f, 0xc1, 0xc2, 0x80, 0x29, 0x14, 0xbe, 0x69, 0x6a, 0x54, 0xb3,
	0xd6, 0x12, 0x85, 0x56, 0x86, 0xb0, 0x52, 0x1b, 0x26, 0x30, 0x03, 0x3c, 0x12, 0xbe, 0xee, 0xed,
	0x04, 0xfc, 0x34, 0xe0, 0xf6, 0xa8, 0xa4, 0x86, 0x3a, 0xa6, 0xa9, 0x3f, 0x27, 0x0a, 0xad, 0x0e,
	0xa1, 0xde, 0xfb, 0x30, 0x99, 0xcb, 0x60, 0xf7, 0xa8, 0xa4, 0x9a, 0x5b, 0x03, 0x25, 0x76, 0xce,
	0x22, 0x59, 0x6b, 0xe9, 0xdb, 0x2c, 0x8f, 0xaf, 0x59, 0xeb, 0x93, 0x5b, 0xa8, 0xfa, 0x78, 0x52,
	0xd5, 0xfd, 0x9e, 0xcf, 0x5c, 0xba, 0xb3, 0x7c, 0xa3, 0x50, 0x2e, 0x51, 0x68, 0xce, 0x44, 0x66,
	0x11, 0x98, 0x4c, 0xb2, 0x07, 0xe7, 0x4e, 0xf1, 0xdd, 0x35, 0xca, 0x7d, 0xbd, 0x46, 0x16, 0xfe,
	0x34, 0x02, 0x26, 0x33, 0x0c, 0xf8, 0xd2, 0x0c, 0x8e, 0x4a, 0x19, 0x07, 0xf5, 0xb6, 0x64, 0xb5,
	0xa8, 0x1d, 0xf6, 0x07, 0xf7, 0x7b, 0xa2, 0xd0, 0xaf, 0x43, 0x4e, 0x63, 0x72, 0x06, 0x76, 0x98,
	0x39, 0xfe, 0x93, 0x96, 0x9e, 0xb7, 0x43, 0x78, 0x06, 0xca, 0x83, 0xe0, 0x33, 0xd6, 0xa9, 0x35,
	0x59, 0xe4, 0xcb, 0x46, 0x7f, 0x9e, 0x5b, 0x89, 0x42, 0xd5, 0x27, 0xf0, 0x1f, 0x36, 0x62, 0xb2,
	0x90, 0x8d, 0xf9, 0x9f, 0x75, 0x9e, 0xe9, 0x3a, 0x7c, 0x03, 0x96, 0x06, 0xc3, 0xce, 0x69, 0xb3,
	0xcd, 0xd2, 0x38, 0x33, 0xf2, 0xbf, 0x13, 0x85, 0xfe, 0x7c, 0x42, 0x5c, 0x76, 0x2b, 0x26, 0x8b,
	0xd9, 0xc0, 0x93, 0x9e, 0x64, 0x22, 0x77, 0xc6, 0xf4, 0x75, 0xbe, 0xb7, 0x40, 0x71, 0x97, 0x7b,
	0xec, 0x30, 0x7a, 0xcd, 0xe1, 0x36, 0x28, 0xb8, 0xdc, 0x63, 0xb5, 0xc0, 0x4b, 0xff, 0xfa, 0x5d,
	0x85, 0xf2, 0x5a, 0xde, 0x4b, 0x14, 0x9a, 0x32, 0xe1, 0x7d, 0x0b, 0x26, 0xf9, 0xde, 0xea, 0xd0,
	0x83, 0x9b, 0x60, 0x42, 0xd7, 0x1a, 0x54, 0x98, 0xab, 0x29, 0x39, 0xf3, 0x89, 0x42, 0x33, 0x19,
	0x7b, 0x4f, 0xc2, 0xa4, 0xd8, 0x5b, 0x1f, 0x50, 0xd1, 0x80, 0x1b, 0xa0, 0xe0, 0xc6, 0x8c, 0x4a,
	0x1e, 0xeb, 0xc3, 0x4d, 0x38, 0x30, 0xc3, 0x37, 0x02, 0x26, 0xa9, 0x05, 0x7f, 0x18, 0x01, 0xa5,
	0xf4, 0x3f, 0xa7, 0x1b, 0xdd, 0x00, 0x05, 0xea, 0x79, 0x31, 0x13, 0x42, 0x37, 0x3a, 0xb0, 0xbd,
	0x2f, 0x60, 0x92, 0x5a, 0xe0, 0x2f, 0x60, 0x9c, 0x5f, 0x44, 0x2c, 0xd6, 0xbd, 0x4d, 0x38, 0x33,
	0x89, 0x42, 0x25, 0xe3, 0xd5, 0x65, 0x4c, 0x8c, 0x9c, 0x3d, 0xfe, 0xe8, 0x0f, 0x1c, 0xff, 0x3f,
	0x50, 0x0c, 0xa2, 0x40, 0x3f, 0x3d, 0xfd, 0x8c, 0x4a, 0x8e, 0x9d, 0x28, 0x34, 0x6d, 0xdc, 0xa9,
	0x82, 0xbf, 0x29, 0x54, 0x66, 0x91, 0xcb, 0xbd, 0x20, 0xf2, 0xed, 0x53, 0xc1, 0xa3, 0x2a, 0xa1,
	0x17, 0x47, 0x4c, 0x08, 0xea, 0x33, 0x52, 0xe8, 0xd9, 0x8e, 0x84, 0x0f, 0xb7, 0x01, 0x08, 0x03,
	0x3f, 0xa6, 0x92, 0xd6, 0x9b, 0x4c, 0x3f, 0xa3, 0xa2, 0xb3, 0x90, 0x28, 0x34, 0xdb, 0x9f, 0xfb,
	0xbd, 0x86, 0x49, 0xc6, 0x68, 0x66, 0xe9, 0xec, 0xdd, 0x74, 0x2b, 0xd6, 0x6d, 0xb7, 0x62, 0x7d,
	0xe9, 0x56, 0xac, 0xb7, 0x77, 0x95, 0xdc, 0xed, 0x5d, 0x25, 0xf7, 0xf9, 0xae, 0x92, 0x7b, 0xf5,
	0x9b, 0x1f, 0xc8, 0x46, 0xbb, 0x5e, 0x75, 0x79, 0x68, 0xeb, 0x37, 0xf9, 0x47, 0x2b, 0xe6, 0xa7,
	0xcc, 0x95, 0xb6, 0xcb, 0x63, 0x66, 0x5f, 0x9a, 0x6f, 0xad, 0xec, 0xb4, 0x98, 0xa8, 0xe7, 0xf5,
	0xb7, 0xf1, 0xaf, 0xef, 0x01, 0x00, 0x00, 0xff, 0xff, 0x85, 0xee, 0xeb, 0x75, 0x86, 0x05, 0x00,
	0x00,
}

func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.MaxContractSize != that1.MaxContractSize {
		return false
	}
	if this.MaxContractGas != that1.MaxContractGas {
		return false
	}
	if this.MaxContractMsgSize != that1.MaxContractMsgSize {
		return false
	}
	if this.MaxContractDataSize != that1.MaxContractDataSize {
		return false
	}
	if !this.EventParams.Equal(&that1.EventParams) {
		return false
	}
	return true
}
func (this *EventParams) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*EventParams)
	if !ok {
		that2, ok := that.(EventParams)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.MaxAttributeNum != that1.MaxAttributeNum {
		return false
	}
	if this.MaxAttributeKeyLength != that1.MaxAttributeKeyLength {
		return false
	}
	if this.MaxAttributeValueLength != that1.MaxAttributeValueLength {
		return false
	}
	return true
}
func (this *ContractInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ContractInfo)
	if !ok {
		that2, ok := that.(ContractInfo)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Address != that1.Address {
		return false
	}
	if this.Owner != that1.Owner {
		return false
	}
	if this.CodeID != that1.CodeID {
		return false
	}
	if !bytes.Equal(this.InitMsg, that1.InitMsg) {
		return false
	}
	if this.Migratable != that1.Migratable {
		return false
	}
	return true
}
func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.EventParams.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintWasm(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.MaxContractDataSize != 0 {
		i = encodeVarintWasm(dAtA, i, uint64(m.MaxContractDataSize))
		i--
		dAtA[i] = 0x20
	}
	if m.MaxContractMsgSize != 0 {
		i = encodeVarintWasm(dAtA, i, uint64(m.MaxContractMsgSize))
		i--
		dAtA[i] = 0x18
	}
	if m.MaxContractGas != 0 {
		i = encodeVarintWasm(dAtA, i, uint64(m.MaxContractGas))
		i--
		dAtA[i] = 0x10
	}
	if m.MaxContractSize != 0 {
		i = encodeVarintWasm(dAtA, i, uint64(m.MaxContractSize))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EventParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MaxAttributeValueLength != 0 {
		i = encodeVarintWasm(dAtA, i, uint64(m.MaxAttributeValueLength))
		i--
		dAtA[i] = 0x18
	}
	if m.MaxAttributeKeyLength != 0 {
		i = encodeVarintWasm(dAtA, i, uint64(m.MaxAttributeKeyLength))
		i--
		dAtA[i] = 0x10
	}
	if m.MaxAttributeNum != 0 {
		i = encodeVarintWasm(dAtA, i, uint64(m.MaxAttributeNum))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *CodeInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CodeInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CodeInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintWasm(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.CodeHash) > 0 {
		i -= len(m.CodeHash)
		copy(dAtA[i:], m.CodeHash)
		i = encodeVarintWasm(dAtA, i, uint64(len(m.CodeHash)))
		i--
		dAtA[i] = 0x12
	}
	if m.CodeID != 0 {
		i = encodeVarintWasm(dAtA, i, uint64(m.CodeID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ContractInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContractInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ContractInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Migratable {
		i--
		if m.Migratable {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if len(m.InitMsg) > 0 {
		i -= len(m.InitMsg)
		copy(dAtA[i:], m.InitMsg)
		i = encodeVarintWasm(dAtA, i, uint64(len(m.InitMsg)))
		i--
		dAtA[i] = 0x22
	}
	if m.CodeID != 0 {
		i = encodeVarintWasm(dAtA, i, uint64(m.CodeID))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintWasm(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintWasm(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintWasm(dAtA []byte, offset int, v uint64) int {
	offset -= sovWasm(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaxContractSize != 0 {
		n += 1 + sovWasm(uint64(m.MaxContractSize))
	}
	if m.MaxContractGas != 0 {
		n += 1 + sovWasm(uint64(m.MaxContractGas))
	}
	if m.MaxContractMsgSize != 0 {
		n += 1 + sovWasm(uint64(m.MaxContractMsgSize))
	}
	if m.MaxContractDataSize != 0 {
		n += 1 + sovWasm(uint64(m.MaxContractDataSize))
	}
	l = m.EventParams.Size()
	n += 1 + l + sovWasm(uint64(l))
	return n
}

func (m *EventParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaxAttributeNum != 0 {
		n += 1 + sovWasm(uint64(m.MaxAttributeNum))
	}
	if m.MaxAttributeKeyLength != 0 {
		n += 1 + sovWasm(uint64(m.MaxAttributeKeyLength))
	}
	if m.MaxAttributeValueLength != 0 {
		n += 1 + sovWasm(uint64(m.MaxAttributeValueLength))
	}
	return n
}

func (m *CodeInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CodeID != 0 {
		n += 1 + sovWasm(uint64(m.CodeID))
	}
	l = len(m.CodeHash)
	if l > 0 {
		n += 1 + l + sovWasm(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovWasm(uint64(l))
	}
	return n
}

func (m *ContractInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovWasm(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovWasm(uint64(l))
	}
	if m.CodeID != 0 {
		n += 1 + sovWasm(uint64(m.CodeID))
	}
	l = len(m.InitMsg)
	if l > 0 {
		n += 1 + l + sovWasm(uint64(l))
	}
	if m.Migratable {
		n += 2
	}
	return n
}

func sovWasm(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozWasm(x uint64) (n int) {
	return sovWasm(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWasm
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxContractSize", wireType)
			}
			m.MaxContractSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWasm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxContractSize |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxContractGas", wireType)
			}
			m.MaxContractGas = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWasm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxContractGas |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxContractMsgSize", wireType)
			}
			m.MaxContractMsgSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWasm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxContractMsgSize |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxContractDataSize", wireType)
			}
			m.MaxContractDataSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWasm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxContractDataSize |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EventParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWasm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWasm
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWasm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.EventParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWasm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWasm
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *EventParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWasm
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EventParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxAttributeNum", wireType)
			}
			m.MaxAttributeNum = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWasm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxAttributeNum |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxAttributeKeyLength", wireType)
			}
			m.MaxAttributeKeyLength = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWasm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxAttributeKeyLength |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxAttributeValueLength", wireType)
			}
			m.MaxAttributeValueLength = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWasm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxAttributeValueLength |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipWasm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWasm
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CodeInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWasm
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CodeInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CodeInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodeID", wireType)
			}
			m.CodeID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWasm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CodeID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodeHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWasm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthWasm
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthWasm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CodeHash = append(m.CodeHash[:0], dAtA[iNdEx:postIndex]...)
			if m.CodeHash == nil {
				m.CodeHash = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWasm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWasm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWasm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWasm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWasm
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ContractInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWasm
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ContractInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContractInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWasm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWasm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWasm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWasm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWasm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWasm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodeID", wireType)
			}
			m.CodeID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWasm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CodeID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitMsg", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWasm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthWasm
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthWasm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InitMsg = append(m.InitMsg[:0], dAtA[iNdEx:postIndex]...)
			if m.InitMsg == nil {
				m.InitMsg = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Migratable", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWasm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Migratable = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipWasm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWasm
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipWasm(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWasm
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowWasm
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowWasm
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthWasm
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupWasm
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthWasm
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthWasm        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWasm          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupWasm = fmt.Errorf("proto: unexpected end of group")
)
