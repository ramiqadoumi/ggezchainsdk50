// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ggezchain/trade/stored_temp_trade.proto

package types

import (
	fmt "fmt"
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

type StoredTempTrade struct {
	TradeIndex     uint64 `protobuf:"varint,1,opt,name=tradeIndex,proto3" json:"tradeIndex,omitempty"`
	TempTradeIndex uint64 `protobuf:"varint,2,opt,name=tempTradeIndex,proto3" json:"tempTradeIndex,omitempty"`
	CreateDate     string `protobuf:"bytes,3,opt,name=createDate,proto3" json:"createDate,omitempty"`
}

func (m *StoredTempTrade) Reset()         { *m = StoredTempTrade{} }
func (m *StoredTempTrade) String() string { return proto.CompactTextString(m) }
func (*StoredTempTrade) ProtoMessage()    {}
func (*StoredTempTrade) Descriptor() ([]byte, []int) {
	return fileDescriptor_42ef1e68a70d7b5c, []int{0}
}
func (m *StoredTempTrade) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StoredTempTrade) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StoredTempTrade.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StoredTempTrade) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoredTempTrade.Merge(m, src)
}
func (m *StoredTempTrade) XXX_Size() int {
	return m.Size()
}
func (m *StoredTempTrade) XXX_DiscardUnknown() {
	xxx_messageInfo_StoredTempTrade.DiscardUnknown(m)
}

var xxx_messageInfo_StoredTempTrade proto.InternalMessageInfo

func (m *StoredTempTrade) GetTradeIndex() uint64 {
	if m != nil {
		return m.TradeIndex
	}
	return 0
}

func (m *StoredTempTrade) GetTempTradeIndex() uint64 {
	if m != nil {
		return m.TempTradeIndex
	}
	return 0
}

func (m *StoredTempTrade) GetCreateDate() string {
	if m != nil {
		return m.CreateDate
	}
	return ""
}

func init() {
	proto.RegisterType((*StoredTempTrade)(nil), "ggezchain.trade.StoredTempTrade")
}

func init() {
	proto.RegisterFile("ggezchain/trade/stored_temp_trade.proto", fileDescriptor_42ef1e68a70d7b5c)
}

var fileDescriptor_42ef1e68a70d7b5c = []byte{
	// 203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4f, 0x4f, 0x4f, 0xad,
	0x4a, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x2f, 0x29, 0x4a, 0x4c, 0x49, 0xd5, 0x2f, 0x2e, 0xc9, 0x2f,
	0x4a, 0x4d, 0x89, 0x2f, 0x49, 0xcd, 0x2d, 0x88, 0x07, 0x8b, 0xe8, 0x15, 0x14, 0xe5, 0x97, 0xe4,
	0x0b, 0xf1, 0xc3, 0x15, 0xea, 0x81, 0x85, 0x95, 0x2a, 0xb9, 0xf8, 0x83, 0xc1, 0x6a, 0x43, 0x52,
	0x73, 0x0b, 0x42, 0x40, 0x42, 0x42, 0x72, 0x5c, 0x5c, 0x60, 0x39, 0xcf, 0xbc, 0x94, 0xd4, 0x0a,
	0x09, 0x46, 0x05, 0x46, 0x0d, 0x96, 0x20, 0x24, 0x11, 0x21, 0x35, 0x2e, 0xbe, 0x12, 0x98, 0x62,
	0x88, 0x1a, 0x26, 0xb0, 0x1a, 0x34, 0x51, 0x90, 0x39, 0xc9, 0x45, 0xa9, 0x89, 0x25, 0xa9, 0x2e,
	0x89, 0x25, 0xa9, 0x12, 0xcc, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x48, 0x22, 0x4e, 0xae, 0x27, 0x1e,
	0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17,
	0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0xa5, 0x9d, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4,
	0x97, 0x9c, 0x9f, 0xab, 0xef, 0xee, 0xee, 0x1a, 0xe5, 0x93, 0x98, 0x54, 0xac, 0x8f, 0xf0, 0x62,
	0x05, 0xd4, 0x93, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0x60, 0x9f, 0x19, 0x03, 0x02, 0x00,
	0x00, 0xff, 0xff, 0xd5, 0x67, 0x3b, 0x58, 0x04, 0x01, 0x00, 0x00,
}

func (m *StoredTempTrade) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StoredTempTrade) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StoredTempTrade) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CreateDate) > 0 {
		i -= len(m.CreateDate)
		copy(dAtA[i:], m.CreateDate)
		i = encodeVarintStoredTempTrade(dAtA, i, uint64(len(m.CreateDate)))
		i--
		dAtA[i] = 0x1a
	}
	if m.TempTradeIndex != 0 {
		i = encodeVarintStoredTempTrade(dAtA, i, uint64(m.TempTradeIndex))
		i--
		dAtA[i] = 0x10
	}
	if m.TradeIndex != 0 {
		i = encodeVarintStoredTempTrade(dAtA, i, uint64(m.TradeIndex))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintStoredTempTrade(dAtA []byte, offset int, v uint64) int {
	offset -= sovStoredTempTrade(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StoredTempTrade) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TradeIndex != 0 {
		n += 1 + sovStoredTempTrade(uint64(m.TradeIndex))
	}
	if m.TempTradeIndex != 0 {
		n += 1 + sovStoredTempTrade(uint64(m.TempTradeIndex))
	}
	l = len(m.CreateDate)
	if l > 0 {
		n += 1 + l + sovStoredTempTrade(uint64(l))
	}
	return n
}

func sovStoredTempTrade(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStoredTempTrade(x uint64) (n int) {
	return sovStoredTempTrade(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StoredTempTrade) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStoredTempTrade
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
			return fmt.Errorf("proto: StoredTempTrade: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StoredTempTrade: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TradeIndex", wireType)
			}
			m.TradeIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredTempTrade
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TradeIndex |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TempTradeIndex", wireType)
			}
			m.TempTradeIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredTempTrade
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TempTradeIndex |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreateDate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredTempTrade
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
				return ErrInvalidLengthStoredTempTrade
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredTempTrade
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CreateDate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStoredTempTrade(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStoredTempTrade
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
func skipStoredTempTrade(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStoredTempTrade
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
					return 0, ErrIntOverflowStoredTempTrade
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
					return 0, ErrIntOverflowStoredTempTrade
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
				return 0, ErrInvalidLengthStoredTempTrade
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStoredTempTrade
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStoredTempTrade
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStoredTempTrade        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStoredTempTrade          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStoredTempTrade = fmt.Errorf("proto: unexpected end of group")
)
