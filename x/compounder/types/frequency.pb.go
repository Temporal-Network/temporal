// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: temporal/compounder/frequency.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type Frequency struct {
	OnceEvery       github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=onceEvery,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"onceEvery"`
	SecondsOrBlocks string                                 `protobuf:"bytes,2,opt,name=secondsOrBlocks,proto3" json:"secondsOrBlocks,omitempty"`
}

func (m *Frequency) Reset()         { *m = Frequency{} }
func (m *Frequency) String() string { return proto.CompactTextString(m) }
func (*Frequency) ProtoMessage()    {}
func (*Frequency) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1c44afb22d08196, []int{0}
}
func (m *Frequency) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Frequency) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Frequency.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Frequency) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Frequency.Merge(m, src)
}
func (m *Frequency) XXX_Size() int {
	return m.Size()
}
func (m *Frequency) XXX_DiscardUnknown() {
	xxx_messageInfo_Frequency.DiscardUnknown(m)
}

var xxx_messageInfo_Frequency proto.InternalMessageInfo

func (m *Frequency) GetSecondsOrBlocks() string {
	if m != nil {
		return m.SecondsOrBlocks
	}
	return ""
}

func init() {
	proto.RegisterType((*Frequency)(nil), "temporal.compounder.Frequency")
}

func init() {
	proto.RegisterFile("temporal/compounder/frequency.proto", fileDescriptor_a1c44afb22d08196)
}

var fileDescriptor_a1c44afb22d08196 = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2e, 0x49, 0xcd, 0x2d,
	0xc8, 0x2f, 0x4a, 0xcc, 0xd1, 0x4f, 0xce, 0xcf, 0x2d, 0xc8, 0x2f, 0xcd, 0x4b, 0x49, 0x2d, 0xd2,
	0x4f, 0x2b, 0x4a, 0x2d, 0x2c, 0x4d, 0xcd, 0x4b, 0xae, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x12, 0x86, 0x29, 0xd2, 0x43, 0x28, 0x92, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0xcb, 0xeb, 0x83,
	0x58, 0x10, 0xa5, 0x4a, 0xcd, 0x8c, 0x5c, 0x9c, 0x6e, 0x30, 0xed, 0x42, 0x3e, 0x5c, 0x9c, 0xf9,
	0x79, 0xc9, 0xa9, 0xae, 0x65, 0xa9, 0x45, 0x95, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x4e, 0x7a,
	0x27, 0xee, 0xc9, 0x33, 0xdc, 0xba, 0x27, 0xaf, 0x96, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0x04, 0x32,
	0x51, 0x3f, 0x39, 0xbf, 0x38, 0x37, 0xbf, 0x18, 0x4a, 0xe9, 0x16, 0xa7, 0x64, 0xeb, 0x97, 0x54,
	0x16, 0xa4, 0x16, 0xeb, 0x79, 0xe6, 0x95, 0x04, 0x21, 0x0c, 0x10, 0xd2, 0xe0, 0xe2, 0x2f, 0x4e,
	0x4d, 0xce, 0xcf, 0x4b, 0x29, 0xf6, 0x2f, 0x72, 0xca, 0xc9, 0x4f, 0xce, 0x2e, 0x96, 0x60, 0x02,
	0x99, 0x19, 0x84, 0x2e, 0xec, 0x64, 0x7a, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f,
	0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c,
	0x51, 0xd2, 0x70, 0xff, 0x56, 0x20, 0xfb, 0x18, 0x6c, 0x5f, 0x12, 0x1b, 0xd8, 0x0f, 0xc6, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x96, 0x94, 0x20, 0xfa, 0x15, 0x01, 0x00, 0x00,
}

func (m *Frequency) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Frequency) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Frequency) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SecondsOrBlocks) > 0 {
		i -= len(m.SecondsOrBlocks)
		copy(dAtA[i:], m.SecondsOrBlocks)
		i = encodeVarintFrequency(dAtA, i, uint64(len(m.SecondsOrBlocks)))
		i--
		dAtA[i] = 0x12
	}
	{
		size := m.OnceEvery.Size()
		i -= size
		if _, err := m.OnceEvery.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintFrequency(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintFrequency(dAtA []byte, offset int, v uint64) int {
	offset -= sovFrequency(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Frequency) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.OnceEvery.Size()
	n += 1 + l + sovFrequency(uint64(l))
	l = len(m.SecondsOrBlocks)
	if l > 0 {
		n += 1 + l + sovFrequency(uint64(l))
	}
	return n
}

func sovFrequency(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFrequency(x uint64) (n int) {
	return sovFrequency(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Frequency) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFrequency
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
			return fmt.Errorf("proto: Frequency: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Frequency: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OnceEvery", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFrequency
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
				return ErrInvalidLengthFrequency
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFrequency
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.OnceEvery.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SecondsOrBlocks", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFrequency
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
				return ErrInvalidLengthFrequency
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFrequency
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SecondsOrBlocks = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFrequency(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFrequency
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
func skipFrequency(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFrequency
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
					return 0, ErrIntOverflowFrequency
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
					return 0, ErrIntOverflowFrequency
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
				return 0, ErrInvalidLengthFrequency
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFrequency
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFrequency
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFrequency        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFrequency          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFrequency = fmt.Errorf("proto: unexpected end of group")
)
