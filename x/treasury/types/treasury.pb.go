// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: terra/treasury/v1beta1/treasury.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
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

// Params defines the parameters for the oracle module.
type Params struct {
	TaxPolicy               PolicyConstraints                      `protobuf:"bytes,1,opt,name=tax_policy,json=taxPolicy,proto3" json:"tax_policy" yaml:"tax_policy"`
	RewardPolicy            PolicyConstraints                      `protobuf:"bytes,2,opt,name=reward_policy,json=rewardPolicy,proto3" json:"reward_policy" yaml:"reward_policy"`
	SeigniorageBurdenTarget github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=seigniorage_burden_target,json=seigniorageBurdenTarget,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"seigniorage_burden_target" yaml:"seigniorage_burden_target"`
	MiningIncrement         github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=mining_increment,json=miningIncrement,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"mining_increment" yaml:"mining_increment"`
	WindowShort             uint64                                 `protobuf:"varint,5,opt,name=window_short,json=windowShort,proto3" json:"window_short,omitempty" yaml:"window_short"`
	WindowLong              uint64                                 `protobuf:"varint,6,opt,name=window_long,json=windowLong,proto3" json:"window_long,omitempty" yaml:"window_long"`
	WindowProbation         uint64                                 `protobuf:"varint,7,opt,name=window_probation,json=windowProbation,proto3" json:"window_probation,omitempty" yaml:"window_probation"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_353bb3a9c554268e, []int{0}
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

func (m *Params) GetTaxPolicy() PolicyConstraints {
	if m != nil {
		return m.TaxPolicy
	}
	return PolicyConstraints{}
}

func (m *Params) GetRewardPolicy() PolicyConstraints {
	if m != nil {
		return m.RewardPolicy
	}
	return PolicyConstraints{}
}

func (m *Params) GetWindowShort() uint64 {
	if m != nil {
		return m.WindowShort
	}
	return 0
}

func (m *Params) GetWindowLong() uint64 {
	if m != nil {
		return m.WindowLong
	}
	return 0
}

func (m *Params) GetWindowProbation() uint64 {
	if m != nil {
		return m.WindowProbation
	}
	return 0
}

// PolicyConstraints - defines policy constraints can be applied in tax & reward policies
type PolicyConstraints struct {
	RateMin       github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=rate_min,json=rateMin,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"rate_min" yaml:"rate_min"`
	RateMax       github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=rate_max,json=rateMax,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"rate_max" yaml:"rate_max"`
	Cap           types.Coin                             `protobuf:"bytes,3,opt,name=cap,proto3" json:"cap" yaml:"cap"`
	ChangeRateMax github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=change_rate_max,json=changeRateMax,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"change_rate_max" yaml:"change_rate_max"`
}

func (m *PolicyConstraints) Reset()      { *m = PolicyConstraints{} }
func (*PolicyConstraints) ProtoMessage() {}
func (*PolicyConstraints) Descriptor() ([]byte, []int) {
	return fileDescriptor_353bb3a9c554268e, []int{1}
}
func (m *PolicyConstraints) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PolicyConstraints) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PolicyConstraints.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PolicyConstraints) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PolicyConstraints.Merge(m, src)
}
func (m *PolicyConstraints) XXX_Size() int {
	return m.Size()
}
func (m *PolicyConstraints) XXX_DiscardUnknown() {
	xxx_messageInfo_PolicyConstraints.DiscardUnknown(m)
}

var xxx_messageInfo_PolicyConstraints proto.InternalMessageInfo

func (m *PolicyConstraints) GetCap() types.Coin {
	if m != nil {
		return m.Cap
	}
	return types.Coin{}
}

// EpochTaxProceeds represents the tax amount
// collected at the current epoch
type EpochTaxProceeds struct {
	TaxProceeds github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=tax_proceeds,json=taxProceeds,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"tax_proceeds" yaml:"tax_proceeds"`
}

func (m *EpochTaxProceeds) Reset()         { *m = EpochTaxProceeds{} }
func (m *EpochTaxProceeds) String() string { return proto.CompactTextString(m) }
func (*EpochTaxProceeds) ProtoMessage()    {}
func (*EpochTaxProceeds) Descriptor() ([]byte, []int) {
	return fileDescriptor_353bb3a9c554268e, []int{2}
}
func (m *EpochTaxProceeds) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EpochTaxProceeds) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EpochTaxProceeds.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EpochTaxProceeds) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EpochTaxProceeds.Merge(m, src)
}
func (m *EpochTaxProceeds) XXX_Size() int {
	return m.Size()
}
func (m *EpochTaxProceeds) XXX_DiscardUnknown() {
	xxx_messageInfo_EpochTaxProceeds.DiscardUnknown(m)
}

var xxx_messageInfo_EpochTaxProceeds proto.InternalMessageInfo

func (m *EpochTaxProceeds) GetTaxProceeds() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.TaxProceeds
	}
	return nil
}

// EpochInitialIssuance represents initial issuance
// of the currrent epoch
type EpochInitialIssuance struct {
	Issuance github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=issuance,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"issuance" yaml:"issuance"`
}

func (m *EpochInitialIssuance) Reset()         { *m = EpochInitialIssuance{} }
func (m *EpochInitialIssuance) String() string { return proto.CompactTextString(m) }
func (*EpochInitialIssuance) ProtoMessage()    {}
func (*EpochInitialIssuance) Descriptor() ([]byte, []int) {
	return fileDescriptor_353bb3a9c554268e, []int{3}
}
func (m *EpochInitialIssuance) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EpochInitialIssuance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EpochInitialIssuance.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EpochInitialIssuance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EpochInitialIssuance.Merge(m, src)
}
func (m *EpochInitialIssuance) XXX_Size() int {
	return m.Size()
}
func (m *EpochInitialIssuance) XXX_DiscardUnknown() {
	xxx_messageInfo_EpochInitialIssuance.DiscardUnknown(m)
}

var xxx_messageInfo_EpochInitialIssuance proto.InternalMessageInfo

func (m *EpochInitialIssuance) GetIssuance() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Issuance
	}
	return nil
}

func init() {
	proto.RegisterType((*Params)(nil), "terra.treasury.v1beta1.Params")
	proto.RegisterType((*PolicyConstraints)(nil), "terra.treasury.v1beta1.PolicyConstraints")
	proto.RegisterType((*EpochTaxProceeds)(nil), "terra.treasury.v1beta1.EpochTaxProceeds")
	proto.RegisterType((*EpochInitialIssuance)(nil), "terra.treasury.v1beta1.EpochInitialIssuance")
}

func init() {
	proto.RegisterFile("terra/treasury/v1beta1/treasury.proto", fileDescriptor_353bb3a9c554268e)
}

var fileDescriptor_353bb3a9c554268e = []byte{
	// 700 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0x4b, 0x6f, 0x13, 0x3b,
	0x14, 0xc7, 0x33, 0x37, 0x7d, 0x3a, 0xad, 0xd2, 0x4e, 0xab, 0x36, 0xed, 0xbd, 0xca, 0x44, 0x96,
	0xee, 0x55, 0xee, 0xa2, 0x33, 0x6a, 0x59, 0x20, 0x75, 0x83, 0x48, 0x79, 0x34, 0x12, 0x48, 0xd1,
	0xd0, 0x15, 0x42, 0x1a, 0x39, 0x8e, 0x35, 0x31, 0x24, 0xf6, 0xc8, 0x76, 0x68, 0xc2, 0x9e, 0x1d,
	0x42, 0x88, 0x15, 0x62, 0xd5, 0x2d, 0x7c, 0x92, 0x2e, 0xbb, 0x44, 0x2c, 0x02, 0x6a, 0x37, 0xac,
	0xf3, 0x09, 0xd0, 0xd8, 0xce, 0xa3, 0xe5, 0x19, 0xb1, 0x8a, 0xcf, 0xeb, 0x77, 0xfe, 0x1e, 0x1f,
	0x3b, 0xe0, 0x5f, 0x45, 0x84, 0x40, 0x81, 0x12, 0x04, 0xc9, 0x8e, 0xe8, 0x05, 0x4f, 0x77, 0xeb,
	0x44, 0xa1, 0xdd, 0x91, 0xc3, 0x4f, 0x04, 0x57, 0xdc, 0xdd, 0xd0, 0x69, 0xfe, 0xc8, 0x6b, 0xd3,
	0xb6, 0xd7, 0x63, 0x1e, 0x73, 0x9d, 0x12, 0xa4, 0x2b, 0x93, 0xbd, 0x5d, 0xc4, 0x5c, 0xb6, 0xb9,
	0x0c, 0xea, 0x48, 0x92, 0x11, 0x11, 0x73, 0xca, 0x4c, 0x1c, 0xbe, 0x9b, 0x05, 0x73, 0x35, 0x24,
	0x50, 0x5b, 0xba, 0x18, 0x00, 0x85, 0xba, 0x51, 0xc2, 0x5b, 0x14, 0xf7, 0x0a, 0x4e, 0xc9, 0x29,
	0xe7, 0xf6, 0xfe, 0xf7, 0xbf, 0xdf, 0xcd, 0xaf, 0xe9, 0xac, 0x03, 0xce, 0xa4, 0x12, 0x88, 0x32,
	0x25, 0x2b, 0x5b, 0xa7, 0x7d, 0x2f, 0x33, 0xe8, 0x7b, 0xab, 0x3d, 0xd4, 0x6e, 0xed, 0xc3, 0x31,
	0x0a, 0x86, 0x8b, 0x0a, 0x75, 0x4d, 0x81, 0xdb, 0x02, 0xcb, 0x82, 0x1c, 0x23, 0xd1, 0x18, 0xf6,
	0xf9, 0x6b, 0xda, 0x3e, 0xff, 0xd8, 0x3e, 0xeb, 0xa6, 0xcf, 0x25, 0x1a, 0x0c, 0x97, 0x8c, 0x6d,
	0xbb, 0xbd, 0x74, 0xc0, 0x96, 0x24, 0x34, 0x66, 0x94, 0x0b, 0x14, 0x93, 0xa8, 0xde, 0x11, 0x0d,
	0xc2, 0x22, 0x85, 0x44, 0x4c, 0x54, 0x21, 0x5b, 0x72, 0xca, 0x8b, 0x95, 0x30, 0xe5, 0x7d, 0xec,
	0x7b, 0xff, 0xc5, 0x54, 0x35, 0x3b, 0x75, 0x1f, 0xf3, 0x76, 0x60, 0x3f, 0x9a, 0xf9, 0xd9, 0x91,
	0x8d, 0x27, 0x81, 0xea, 0x25, 0x44, 0xfa, 0xb7, 0x08, 0x1e, 0xf4, 0xbd, 0x92, 0xe9, 0xfc, 0x43,
	0x30, 0x0c, 0x37, 0x27, 0x62, 0x15, 0x1d, 0x3a, 0xd2, 0x11, 0x57, 0x81, 0x95, 0x36, 0x65, 0x94,
	0xc5, 0x11, 0x65, 0x58, 0x90, 0x36, 0x61, 0xaa, 0x30, 0xa3, 0x65, 0x54, 0xa7, 0x96, 0xb1, 0x69,
	0x64, 0x5c, 0xe5, 0xc1, 0x30, 0x6f, 0x5c, 0xd5, 0xa1, 0xc7, 0xdd, 0x07, 0x4b, 0xc7, 0x94, 0x35,
	0xf8, 0x71, 0x24, 0x9b, 0x5c, 0xa8, 0xc2, 0x6c, 0xc9, 0x29, 0xcf, 0x54, 0x36, 0x07, 0x7d, 0x6f,
	0xcd, 0x30, 0x26, 0xa3, 0x30, 0xcc, 0x19, 0xf3, 0x41, 0x6a, 0xb9, 0xd7, 0x81, 0x35, 0xa3, 0x16,
	0x67, 0x71, 0x61, 0x4e, 0x97, 0x6e, 0x0c, 0xfa, 0x9e, 0x7b, 0xa9, 0x34, 0x0d, 0xc2, 0x10, 0x18,
	0xeb, 0x1e, 0x67, 0xb1, 0x7b, 0x07, 0xac, 0xd8, 0x58, 0x22, 0x78, 0x1d, 0x29, 0xca, 0x59, 0x61,
	0x5e, 0x57, 0xff, 0x3d, 0x16, 0x7f, 0x35, 0x03, 0x86, 0x79, 0xe3, 0xaa, 0x0d, 0x3d, 0xfb, 0x0b,
	0x6f, 0x4e, 0xbc, 0xcc, 0x97, 0x13, 0xcf, 0x81, 0x2f, 0xb2, 0x60, 0xf5, 0x9b, 0x79, 0x70, 0x1f,
	0x81, 0x05, 0x81, 0x14, 0x89, 0xda, 0x94, 0xe9, 0xa1, 0x5d, 0xac, 0xdc, 0x9c, 0xfa, 0x53, 0xe6,
	0xed, 0x2c, 0x59, 0x0e, 0x0c, 0xe7, 0xd3, 0xe5, 0x7d, 0xca, 0xc6, 0x74, 0xd4, 0xd5, 0xa3, 0xfa,
	0xc7, 0x74, 0xd4, 0x1d, 0xd2, 0x51, 0xd7, 0xbd, 0x01, 0xb2, 0x18, 0x25, 0x7a, 0x10, 0x73, 0x7b,
	0x5b, 0xbe, 0xa9, 0xf7, 0xd3, 0xbb, 0x3a, 0xba, 0x00, 0x07, 0x9c, 0xb2, 0x8a, 0x6b, 0x67, 0x1e,
	0x18, 0x12, 0x46, 0x09, 0x0c, 0xd3, 0x4a, 0x37, 0x01, 0x79, 0xdc, 0x44, 0x2c, 0x26, 0xd1, 0x48,
	0xa5, 0x19, 0xa7, 0xc3, 0xa9, 0x55, 0x6e, 0x58, 0xf6, 0x65, 0x1c, 0x0c, 0x97, 0x8d, 0x27, 0x34,
	0x92, 0x27, 0x8e, 0xe3, 0xad, 0x03, 0x56, 0x6e, 0x27, 0x1c, 0x37, 0x8f, 0x50, 0xb7, 0x26, 0x38,
	0x26, 0xa4, 0x21, 0xdd, 0xe7, 0x0e, 0x58, 0xd2, 0x57, 0xdf, 0x3a, 0x0a, 0x4e, 0x29, 0xfb, 0xf3,
	0xbd, 0xdd, 0xb5, 0x7b, 0x5b, 0x9b, 0x78, 0x37, 0x6c, 0x31, 0x7c, 0xff, 0xc9, 0x2b, 0xff, 0xc6,
	0x06, 0x52, 0x8e, 0x0c, 0x73, 0x6a, 0xac, 0x03, 0xbe, 0x76, 0xc0, 0xba, 0x16, 0x57, 0x65, 0x54,
	0x51, 0xd4, 0xaa, 0x4a, 0xd9, 0x41, 0x0c, 0x13, 0xf7, 0x19, 0x58, 0xa0, 0x76, 0xfd, 0x6b, 0x6d,
	0x07, 0x56, 0x9b, 0x3d, 0xc1, 0x61, 0xe1, 0x74, 0xba, 0x46, 0xfd, 0x2a, 0x87, 0xa7, 0xe7, 0x45,
	0xe7, 0xec, 0xbc, 0xe8, 0x7c, 0x3e, 0x2f, 0x3a, 0xaf, 0x2e, 0x8a, 0x99, 0xb3, 0x8b, 0x62, 0xe6,
	0xc3, 0x45, 0x31, 0xf3, 0xd0, 0x9f, 0xa0, 0xe9, 0x97, 0x70, 0x27, 0x11, 0xfc, 0x31, 0xc1, 0x2a,
	0xc0, 0x5c, 0x90, 0xa0, 0x3b, 0xfe, 0x57, 0xd0, 0xe4, 0xfa, 0x9c, 0x7e, 0xbd, 0xaf, 0x7d, 0x0d,
	0x00, 0x00, 0xff, 0xff, 0x24, 0x78, 0x10, 0xa7, 0x34, 0x06, 0x00, 0x00,
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
	if !this.TaxPolicy.Equal(&that1.TaxPolicy) {
		return false
	}
	if !this.RewardPolicy.Equal(&that1.RewardPolicy) {
		return false
	}
	if !this.SeigniorageBurdenTarget.Equal(that1.SeigniorageBurdenTarget) {
		return false
	}
	if !this.MiningIncrement.Equal(that1.MiningIncrement) {
		return false
	}
	if this.WindowShort != that1.WindowShort {
		return false
	}
	if this.WindowLong != that1.WindowLong {
		return false
	}
	if this.WindowProbation != that1.WindowProbation {
		return false
	}
	return true
}
func (this *PolicyConstraints) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PolicyConstraints)
	if !ok {
		that2, ok := that.(PolicyConstraints)
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
	if !this.RateMin.Equal(that1.RateMin) {
		return false
	}
	if !this.RateMax.Equal(that1.RateMax) {
		return false
	}
	if !this.Cap.Equal(&that1.Cap) {
		return false
	}
	if !this.ChangeRateMax.Equal(that1.ChangeRateMax) {
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
	if m.WindowProbation != 0 {
		i = encodeVarintTreasury(dAtA, i, uint64(m.WindowProbation))
		i--
		dAtA[i] = 0x38
	}
	if m.WindowLong != 0 {
		i = encodeVarintTreasury(dAtA, i, uint64(m.WindowLong))
		i--
		dAtA[i] = 0x30
	}
	if m.WindowShort != 0 {
		i = encodeVarintTreasury(dAtA, i, uint64(m.WindowShort))
		i--
		dAtA[i] = 0x28
	}
	{
		size := m.MiningIncrement.Size()
		i -= size
		if _, err := m.MiningIncrement.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTreasury(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.SeigniorageBurdenTarget.Size()
		i -= size
		if _, err := m.SeigniorageBurdenTarget.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTreasury(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.RewardPolicy.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTreasury(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.TaxPolicy.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTreasury(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *PolicyConstraints) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PolicyConstraints) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PolicyConstraints) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.ChangeRateMax.Size()
		i -= size
		if _, err := m.ChangeRateMax.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTreasury(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size, err := m.Cap.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTreasury(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.RateMax.Size()
		i -= size
		if _, err := m.RateMax.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTreasury(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.RateMin.Size()
		i -= size
		if _, err := m.RateMin.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTreasury(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *EpochTaxProceeds) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EpochTaxProceeds) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EpochTaxProceeds) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TaxProceeds) > 0 {
		for iNdEx := len(m.TaxProceeds) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TaxProceeds[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTreasury(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *EpochInitialIssuance) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EpochInitialIssuance) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EpochInitialIssuance) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Issuance) > 0 {
		for iNdEx := len(m.Issuance) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Issuance[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTreasury(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintTreasury(dAtA []byte, offset int, v uint64) int {
	offset -= sovTreasury(v)
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
	l = m.TaxPolicy.Size()
	n += 1 + l + sovTreasury(uint64(l))
	l = m.RewardPolicy.Size()
	n += 1 + l + sovTreasury(uint64(l))
	l = m.SeigniorageBurdenTarget.Size()
	n += 1 + l + sovTreasury(uint64(l))
	l = m.MiningIncrement.Size()
	n += 1 + l + sovTreasury(uint64(l))
	if m.WindowShort != 0 {
		n += 1 + sovTreasury(uint64(m.WindowShort))
	}
	if m.WindowLong != 0 {
		n += 1 + sovTreasury(uint64(m.WindowLong))
	}
	if m.WindowProbation != 0 {
		n += 1 + sovTreasury(uint64(m.WindowProbation))
	}
	return n
}

func (m *PolicyConstraints) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.RateMin.Size()
	n += 1 + l + sovTreasury(uint64(l))
	l = m.RateMax.Size()
	n += 1 + l + sovTreasury(uint64(l))
	l = m.Cap.Size()
	n += 1 + l + sovTreasury(uint64(l))
	l = m.ChangeRateMax.Size()
	n += 1 + l + sovTreasury(uint64(l))
	return n
}

func (m *EpochTaxProceeds) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.TaxProceeds) > 0 {
		for _, e := range m.TaxProceeds {
			l = e.Size()
			n += 1 + l + sovTreasury(uint64(l))
		}
	}
	return n
}

func (m *EpochInitialIssuance) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Issuance) > 0 {
		for _, e := range m.Issuance {
			l = e.Size()
			n += 1 + l + sovTreasury(uint64(l))
		}
	}
	return n
}

func sovTreasury(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTreasury(x uint64) (n int) {
	return sovTreasury(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTreasury
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TaxPolicy", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreasury
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
				return ErrInvalidLengthTreasury
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTreasury
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TaxPolicy.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardPolicy", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreasury
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
				return ErrInvalidLengthTreasury
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTreasury
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RewardPolicy.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SeigniorageBurdenTarget", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreasury
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
				return ErrInvalidLengthTreasury
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTreasury
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SeigniorageBurdenTarget.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MiningIncrement", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreasury
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
				return ErrInvalidLengthTreasury
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTreasury
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MiningIncrement.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WindowShort", wireType)
			}
			m.WindowShort = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreasury
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WindowShort |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WindowLong", wireType)
			}
			m.WindowLong = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreasury
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WindowLong |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WindowProbation", wireType)
			}
			m.WindowProbation = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreasury
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WindowProbation |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTreasury(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTreasury
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
func (m *PolicyConstraints) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTreasury
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
			return fmt.Errorf("proto: PolicyConstraints: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PolicyConstraints: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RateMin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreasury
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
				return ErrInvalidLengthTreasury
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTreasury
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RateMin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RateMax", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreasury
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
				return ErrInvalidLengthTreasury
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTreasury
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RateMax.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cap", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreasury
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
				return ErrInvalidLengthTreasury
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTreasury
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Cap.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChangeRateMax", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreasury
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
				return ErrInvalidLengthTreasury
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTreasury
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ChangeRateMax.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTreasury(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTreasury
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
func (m *EpochTaxProceeds) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTreasury
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
			return fmt.Errorf("proto: EpochTaxProceeds: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EpochTaxProceeds: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TaxProceeds", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreasury
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
				return ErrInvalidLengthTreasury
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTreasury
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TaxProceeds = append(m.TaxProceeds, types.Coin{})
			if err := m.TaxProceeds[len(m.TaxProceeds)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTreasury(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTreasury
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
func (m *EpochInitialIssuance) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTreasury
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
			return fmt.Errorf("proto: EpochInitialIssuance: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EpochInitialIssuance: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Issuance", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreasury
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
				return ErrInvalidLengthTreasury
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTreasury
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Issuance = append(m.Issuance, types.Coin{})
			if err := m.Issuance[len(m.Issuance)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTreasury(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTreasury
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
func skipTreasury(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTreasury
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
					return 0, ErrIntOverflowTreasury
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
					return 0, ErrIntOverflowTreasury
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
				return 0, ErrInvalidLengthTreasury
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTreasury
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTreasury
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTreasury        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTreasury          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTreasury = fmt.Errorf("proto: unexpected end of group")
)