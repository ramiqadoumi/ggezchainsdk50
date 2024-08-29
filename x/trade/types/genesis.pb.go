// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ggezchain/trade/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// GenesisState defines the trade module's genesis state.
type GenesisState struct {
	// params defines all the parameters of the module.
	Params              Params            `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	TradeIndex          TradeIndex        `protobuf:"bytes,2,opt,name=tradeIndex,proto3" json:"tradeIndex"`
	StoredTradeList     []StoredTrade     `protobuf:"bytes,3,rep,name=storedTradeList,proto3" json:"storedTradeList"`
	StoredTempTradeList []StoredTempTrade `protobuf:"bytes,4,rep,name=storedTempTradeList,proto3" json:"storedTempTradeList"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_23eb2ced9fe723f5, []int{0}
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

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetTradeIndex() TradeIndex {
	if m != nil {
		return m.TradeIndex
	}
	return TradeIndex{}
}

func (m *GenesisState) GetStoredTradeList() []StoredTrade {
	if m != nil {
		return m.StoredTradeList
	}
	return nil
}

func (m *GenesisState) GetStoredTempTradeList() []StoredTempTrade {
	if m != nil {
		return m.StoredTempTradeList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "ggezchain.trade.GenesisState")
}

func init() { proto.RegisterFile("ggezchain/trade/genesis.proto", fileDescriptor_23eb2ced9fe723f5) }

var fileDescriptor_23eb2ced9fe723f5 = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4d, 0x4f, 0x4f, 0xad,
	0x4a, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x2f, 0x29, 0x4a, 0x4c, 0x49, 0xd5, 0x4f, 0x4f, 0xcd, 0x4b,
	0x2d, 0xce, 0x2c, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x87, 0x4b, 0xeb, 0x81, 0xa5,
	0xa5, 0x04, 0x13, 0x73, 0x33, 0xf3, 0xf2, 0xf5, 0xc1, 0x24, 0x44, 0x8d, 0x94, 0x48, 0x7a, 0x7e,
	0x7a, 0x3e, 0x98, 0xa9, 0x0f, 0x62, 0x41, 0x45, 0x65, 0xd0, 0x0d, 0x2e, 0x48, 0x2c, 0x4a, 0xcc,
	0x85, 0x9a, 0x2b, 0xa5, 0x88, 0x2e, 0x0b, 0x26, 0xe3, 0x33, 0xf3, 0x52, 0x52, 0x2b, 0xa0, 0x4a,
	0x94, 0xd0, 0x95, 0x14, 0x97, 0xe4, 0x17, 0xa5, 0xa6, 0xc4, 0x83, 0x39, 0x50, 0x35, 0xea, 0xb8,
	0xd4, 0xa4, 0xe6, 0x16, 0x20, 0x2b, 0x54, 0xda, 0xca, 0xc4, 0xc5, 0xe3, 0x0e, 0xf1, 0x59, 0x70,
	0x49, 0x62, 0x49, 0xaa, 0x90, 0x15, 0x17, 0x1b, 0xc4, 0x41, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc,
	0x46, 0xe2, 0x7a, 0x68, 0x3e, 0xd5, 0x0b, 0x00, 0x4b, 0x3b, 0x71, 0x9e, 0xb8, 0x27, 0xcf, 0xb0,
	0xe2, 0xf9, 0x06, 0x2d, 0xc6, 0x20, 0xa8, 0x0e, 0x21, 0x47, 0x2e, 0x2e, 0xb0, 0x12, 0x4f, 0x90,
	0x6b, 0x25, 0x98, 0xc0, 0xfa, 0xa5, 0x31, 0xf4, 0x87, 0xc0, 0x95, 0x38, 0xb1, 0x80, 0xcc, 0x08,
	0x42, 0xd2, 0x24, 0xe4, 0xc3, 0xc5, 0x0f, 0x71, 0x2a, 0x58, 0x95, 0x4f, 0x66, 0x71, 0x89, 0x04,
	0xb3, 0x02, 0xb3, 0x06, 0xb7, 0x91, 0x0c, 0x86, 0x39, 0xc1, 0x08, 0x75, 0x50, 0x83, 0xd0, 0xb5,
	0x0a, 0x45, 0x70, 0x09, 0x43, 0x85, 0x52, 0x73, 0x0b, 0x10, 0x26, 0xb2, 0x80, 0x4d, 0x54, 0xc0,
	0x65, 0x22, 0x4c, 0x2d, 0xd4, 0x54, 0x6c, 0x46, 0x38, 0x79, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1,
	0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70,
	0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x5e, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae,
	0x7e, 0x51, 0x62, 0x6e, 0x66, 0x61, 0x62, 0x4a, 0x7e, 0x69, 0x6e, 0xa6, 0x3e, 0x22, 0x46, 0x2a,
	0x60, 0x51, 0x5b, 0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06, 0x8e, 0x08, 0x63, 0x40, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x8d, 0x29, 0x4d, 0x8a, 0x71, 0x02, 0x00, 0x00,
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
	if len(m.StoredTempTradeList) > 0 {
		for iNdEx := len(m.StoredTempTradeList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.StoredTempTradeList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.StoredTradeList) > 0 {
		for iNdEx := len(m.StoredTradeList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.StoredTradeList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	{
		size, err := m.TradeIndex.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
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
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.TradeIndex.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.StoredTradeList) > 0 {
		for _, e := range m.StoredTradeList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.StoredTempTradeList) > 0 {
		for _, e := range m.StoredTempTradeList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
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
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
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
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TradeIndex", wireType)
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
			if err := m.TradeIndex.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StoredTradeList", wireType)
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
			m.StoredTradeList = append(m.StoredTradeList, StoredTrade{})
			if err := m.StoredTradeList[len(m.StoredTradeList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StoredTempTradeList", wireType)
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
			m.StoredTempTradeList = append(m.StoredTempTradeList, StoredTempTrade{})
			if err := m.StoredTempTradeList[len(m.StoredTempTradeList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
