// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: zetacore/genesis.proto

package types

import (
	fmt "fmt"
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

// GenesisState defines the metacore module's genesis state.
type GenesisState struct {
	// this line is used by starport scaffolding # genesis/proto/state
	Keygen              *Keygen            `protobuf:"bytes,12,opt,name=keygen,proto3" json:"keygen,omitempty"`
	TSSVoterList        []*TSSVoter        `protobuf:"bytes,11,rep,name=tSSVoterList,proto3" json:"tSSVoterList,omitempty"`
	TSSList             []*TSS             `protobuf:"bytes,10,rep,name=tSSList,proto3" json:"tSSList,omitempty"`
	InTxList            []*InTx            `protobuf:"bytes,9,rep,name=inTxList,proto3" json:"inTxList,omitempty"`
	TxList              *TxList            `protobuf:"bytes,8,opt,name=txList,proto3" json:"txList,omitempty"`
	GasBalanceList      []*GasBalance      `protobuf:"bytes,7,rep,name=gasBalanceList,proto3" json:"gasBalanceList,omitempty"`
	GasPriceList        []*GasPrice        `protobuf:"bytes,6,rep,name=gasPriceList,proto3" json:"gasPriceList,omitempty"`
	ChainNoncesList     []*ChainNonces     `protobuf:"bytes,5,rep,name=chainNoncesList,proto3" json:"chainNoncesList,omitempty"`
	LastBlockHeightList []*LastBlockHeight `protobuf:"bytes,4,rep,name=lastBlockHeightList,proto3" json:"lastBlockHeightList,omitempty"`
	SendList            []*Send            `protobuf:"bytes,1,rep,name=sendList,proto3" json:"sendList,omitempty"`
	ReceiveList         []*Receive         `protobuf:"bytes,2,rep,name=receiveList,proto3" json:"receiveList,omitempty"`
	NodeAccountList     []*NodeAccount     `protobuf:"bytes,3,rep,name=nodeAccountList,proto3" json:"nodeAccountList,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_14d0e1067b34301c, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetKeygen() *Keygen {
	if m != nil {
		return m.Keygen
	}
	return nil
}

func (m *GenesisState) GetTSSVoterList() []*TSSVoter {
	if m != nil {
		return m.TSSVoterList
	}
	return nil
}

func (m *GenesisState) GetTSSList() []*TSS {
	if m != nil {
		return m.TSSList
	}
	return nil
}

func (m *GenesisState) GetInTxList() []*InTx {
	if m != nil {
		return m.InTxList
	}
	return nil
}

func (m *GenesisState) GetTxList() *TxList {
	if m != nil {
		return m.TxList
	}
	return nil
}

func (m *GenesisState) GetGasBalanceList() []*GasBalance {
	if m != nil {
		return m.GasBalanceList
	}
	return nil
}

func (m *GenesisState) GetGasPriceList() []*GasPrice {
	if m != nil {
		return m.GasPriceList
	}
	return nil
}

func (m *GenesisState) GetChainNoncesList() []*ChainNonces {
	if m != nil {
		return m.ChainNoncesList
	}
	return nil
}

func (m *GenesisState) GetLastBlockHeightList() []*LastBlockHeight {
	if m != nil {
		return m.LastBlockHeightList
	}
	return nil
}

func (m *GenesisState) GetSendList() []*Send {
	if m != nil {
		return m.SendList
	}
	return nil
}

func (m *GenesisState) GetReceiveList() []*Receive {
	if m != nil {
		return m.ReceiveList
	}
	return nil
}

func (m *GenesisState) GetNodeAccountList() []*NodeAccount {
	if m != nil {
		return m.NodeAccountList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "zetachain.zetacore.zetacore.GenesisState")
}

func init() { proto.RegisterFile("zetacore/genesis.proto", fileDescriptor_14d0e1067b34301c) }

var fileDescriptor_14d0e1067b34301c = []byte{
	// 515 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x1b, 0x06, 0xdd, 0x70, 0x2b, 0x90, 0x3c, 0x40, 0x55, 0x27, 0x45, 0xe5, 0x97, 0xe8,
	0x01, 0x52, 0x69, 0xdc, 0x40, 0x1c, 0x28, 0x12, 0xa5, 0x62, 0x1a, 0xc8, 0xa9, 0x38, 0x70, 0x20,
	0x72, 0x53, 0x2b, 0xb5, 0x56, 0xec, 0x2a, 0xf6, 0xa6, 0x8c, 0xbf, 0x82, 0x3f, 0x8b, 0xe3, 0x8e,
	0x1c, 0x51, 0xfb, 0x57, 0x70, 0x43, 0x7e, 0x71, 0x9d, 0xfe, 0x40, 0x59, 0x6e, 0xd1, 0x7b, 0xfe,
	0x7c, 0xe2, 0x7c, 0xf3, 0x1e, 0x7a, 0xf0, 0x83, 0x69, 0x1a, 0xcb, 0x94, 0xf5, 0x12, 0x26, 0x98,
	0xe2, 0x2a, 0x98, 0xa7, 0x52, 0x4b, 0x7c, 0x04, 0xf5, 0x29, 0xe5, 0x22, 0x58, 0x9d, 0x70, 0x0f,
	0xed, 0xfb, 0x0e, 0x3a, 0x63, 0x97, 0x09, 0x13, 0x39, 0xd3, 0x6e, 0xb9, 0xb2, 0x56, 0x2a, 0xba,
	0x90, 0x9a, 0xa5, 0xb6, 0x83, 0xd7, 0x3b, 0xb6, 0x76, 0xcf, 0xd5, 0xb8, 0x88, 0x74, 0x66, 0xab,
	0xc5, 0x7d, 0x74, 0x16, 0xcd, 0xb8, 0xd2, 0xb6, 0xde, 0x2e, 0xee, 0x49, 0x55, 0x34, 0xa6, 0x33,
	0x2a, 0x62, 0xb6, 0xf3, 0x5e, 0xd3, 0x9b, 0xa7, 0xdc, 0x75, 0x8e, 0x5c, 0x07, 0x3e, 0x25, 0x12,
	0x52, 0xc4, 0x6c, 0x75, 0x81, 0x8e, 0x6b, 0xce, 0xa8, 0xd2, 0xd1, 0x78, 0x26, 0xe3, 0xb3, 0x68,
	0xca, 0x78, 0x32, 0xd5, 0x3b, 0x97, 0x49, 0x59, 0xcc, 0xf8, 0xc5, 0x4a, 0x7b, 0xe8, 0xea, 0x8a,
	0x89, 0xc9, 0xce, 0xbb, 0x84, 0x9c, 0xb0, 0x88, 0xc6, 0xb1, 0x3c, 0x17, 0xd6, 0xf4, 0xe8, 0x6f,
	0x1d, 0x35, 0x07, 0x79, 0xc0, 0xa1, 0xa6, 0x9a, 0xe1, 0xd7, 0xa8, 0x9e, 0x67, 0xd7, 0x6a, 0x76,
	0xbc, 0x6e, 0xe3, 0xf8, 0x71, 0x50, 0x12, 0x78, 0xf0, 0x11, 0x8e, 0x12, 0x8b, 0xe0, 0x21, 0x6a,
	0xea, 0x30, 0xfc, 0x62, 0x02, 0x3e, 0xe1, 0x4a, 0xb7, 0x1a, 0x9d, 0xbd, 0x6e, 0xe3, 0xf8, 0x69,
	0xa9, 0x62, 0x64, 0x01, 0xb2, 0x81, 0xe2, 0x57, 0x68, 0x5f, 0x87, 0x21, 0x58, 0x10, 0x58, 0x3a,
	0xd7, 0x59, 0xc8, 0x0a, 0xc0, 0x6f, 0xd0, 0x01, 0x17, 0xa3, 0x0c, 0xe0, 0xdb, 0x00, 0x3f, 0x2c,
	0x85, 0x87, 0x62, 0x94, 0x11, 0x87, 0x98, 0x08, 0x74, 0x0e, 0x1f, 0x54, 0x88, 0x20, 0x87, 0x88,
	0x45, 0xf0, 0x27, 0x74, 0x27, 0xa1, 0xaa, 0x9f, 0xcf, 0x01, 0x48, 0xf6, 0xe1, 0x06, 0xcf, 0x4a,
	0x25, 0x03, 0x87, 0x90, 0x2d, 0xdc, 0x64, 0x9a, 0x50, 0xf5, 0xd9, 0x0c, 0x0f, 0xe8, 0xea, 0x15,
	0x32, 0x1d, 0x58, 0x80, 0x6c, 0xa0, 0x98, 0xa0, 0xbb, 0x40, 0x9c, 0xc2, 0xb4, 0x81, 0xed, 0x16,
	0xd8, 0xba, 0xa5, 0xb6, 0x77, 0x05, 0x43, 0xb6, 0x05, 0xf8, 0x1b, 0x3a, 0x34, 0x53, 0xda, 0x37,
	0x43, 0xfa, 0x01, 0x66, 0x14, 0xbc, 0x37, 0xc1, 0xfb, 0xbc, 0xd4, 0x7b, 0xb2, 0xc9, 0x91, 0xff,
	0x89, 0xcc, 0xbf, 0x34, 0xb3, 0x0c, 0x52, 0xaf, 0xc2, 0xbf, 0x0c, 0x99, 0x98, 0x10, 0x87, 0xe0,
	0xf7, 0xa8, 0x61, 0x57, 0x04, 0x0c, 0x37, 0xc0, 0xf0, 0xa4, 0xd4, 0x40, 0xf2, 0xf3, 0x64, 0x1d,
	0x34, 0xd1, 0x99, 0xed, 0x79, 0x9b, 0x2f, 0x0f, 0xb8, 0xf6, 0x2a, 0x44, 0x77, 0x5a, 0x30, 0x64,
	0x5b, 0xd0, 0x1f, 0xfe, 0x5a, 0xf8, 0xde, 0xd5, 0xc2, 0xf7, 0xfe, 0x2c, 0x7c, 0xef, 0xe7, 0xd2,
	0xaf, 0x5d, 0x2d, 0xfd, 0xda, 0xef, 0xa5, 0x5f, 0xfb, 0xda, 0x4b, 0xb8, 0x9e, 0x9e, 0x8f, 0x83,
	0x58, 0x7e, 0xef, 0x19, 0xd7, 0x0b, 0xf0, 0xf7, 0xdc, 0x22, 0x67, 0xc5, 0xa3, 0xbe, 0x9c, 0x33,
	0x35, 0xae, 0xc3, 0x36, 0xbf, 0xfc, 0x17, 0x00, 0x00, 0xff, 0xff, 0x4c, 0xc7, 0xba, 0xd9, 0x36,
	0x05, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Keygen != nil {
		{
			size, err := m.Keygen.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x62
	}
	if len(m.TSSVoterList) > 0 {
		for iNdEx := len(m.TSSVoterList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TSSVoterList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x5a
		}
	}
	if len(m.TSSList) > 0 {
		for iNdEx := len(m.TSSList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TSSList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	if len(m.InTxList) > 0 {
		for iNdEx := len(m.InTxList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.InTxList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if m.TxList != nil {
		{
			size, err := m.TxList.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if len(m.GasBalanceList) > 0 {
		for iNdEx := len(m.GasBalanceList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GasBalanceList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.GasPriceList) > 0 {
		for iNdEx := len(m.GasPriceList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GasPriceList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.ChainNoncesList) > 0 {
		for iNdEx := len(m.ChainNoncesList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ChainNoncesList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.LastBlockHeightList) > 0 {
		for iNdEx := len(m.LastBlockHeightList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LastBlockHeightList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.NodeAccountList) > 0 {
		for iNdEx := len(m.NodeAccountList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.NodeAccountList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.ReceiveList) > 0 {
		for iNdEx := len(m.ReceiveList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ReceiveList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.SendList) > 0 {
		for iNdEx := len(m.SendList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SendList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.SendList) > 0 {
		for _, e := range m.SendList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ReceiveList) > 0 {
		for _, e := range m.ReceiveList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.NodeAccountList) > 0 {
		for _, e := range m.NodeAccountList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.LastBlockHeightList) > 0 {
		for _, e := range m.LastBlockHeightList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ChainNoncesList) > 0 {
		for _, e := range m.ChainNoncesList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.GasPriceList) > 0 {
		for _, e := range m.GasPriceList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.GasBalanceList) > 0 {
		for _, e := range m.GasBalanceList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.TxList != nil {
		l = m.TxList.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.InTxList) > 0 {
		for _, e := range m.InTxList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.TSSList) > 0 {
		for _, e := range m.TSSList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.TSSVoterList) > 0 {
		for _, e := range m.TSSVoterList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.Keygen != nil {
		l = m.Keygen.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SendList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SendList = append(m.SendList, &Send{})
			if err := m.SendList[len(m.SendList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReceiveList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReceiveList = append(m.ReceiveList, &Receive{})
			if err := m.ReceiveList[len(m.ReceiveList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeAccountList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeAccountList = append(m.NodeAccountList, &NodeAccount{})
			if err := m.NodeAccountList[len(m.NodeAccountList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastBlockHeightList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LastBlockHeightList = append(m.LastBlockHeightList, &LastBlockHeight{})
			if err := m.LastBlockHeightList[len(m.LastBlockHeightList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainNoncesList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainNoncesList = append(m.ChainNoncesList, &ChainNonces{})
			if err := m.ChainNoncesList[len(m.ChainNoncesList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasPriceList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GasPriceList = append(m.GasPriceList, &GasPrice{})
			if err := m.GasPriceList[len(m.GasPriceList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasBalanceList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GasBalanceList = append(m.GasBalanceList, &GasBalance{})
			if err := m.GasBalanceList[len(m.GasBalanceList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TxList == nil {
				m.TxList = &TxList{}
			}
			if err := m.TxList.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InTxList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InTxList = append(m.InTxList, &InTx{})
			if err := m.InTxList[len(m.InTxList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TSSList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TSSList = append(m.TSSList, &TSS{})
			if err := m.TSSList[len(m.TSSList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TSSVoterList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TSSVoterList = append(m.TSSVoterList, &TSSVoter{})
			if err := m.TSSVoterList[len(m.TSSVoterList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Keygen", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Keygen == nil {
				m.Keygen = &Keygen{}
			}
			if err := m.Keygen.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
