// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: temporal/yieldmos/delegation_history.proto

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

type DelegationHistory struct {
	Address string                 `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	History []*DelegationTimestamp `protobuf:"bytes,2,rep,name=history,proto3" json:"history,omitempty"`
}

func (m *DelegationHistory) Reset()         { *m = DelegationHistory{} }
func (m *DelegationHistory) String() string { return proto.CompactTextString(m) }
func (*DelegationHistory) ProtoMessage()    {}
func (*DelegationHistory) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4af6d208c016d9e, []int{0}
}
func (m *DelegationHistory) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DelegationHistory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DelegationHistory.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DelegationHistory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelegationHistory.Merge(m, src)
}
func (m *DelegationHistory) XXX_Size() int {
	return m.Size()
}
func (m *DelegationHistory) XXX_DiscardUnknown() {
	xxx_messageInfo_DelegationHistory.DiscardUnknown(m)
}

var xxx_messageInfo_DelegationHistory proto.InternalMessageInfo

func (m *DelegationHistory) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *DelegationHistory) GetHistory() []*DelegationTimestamp {
	if m != nil {
		return m.History
	}
	return nil
}

func init() {
	proto.RegisterType((*DelegationHistory)(nil), "temporal.yieldmos.DelegationHistory")
}

func init() {
	proto.RegisterFile("temporal/yieldmos/delegation_history.proto", fileDescriptor_b4af6d208c016d9e)
}

var fileDescriptor_b4af6d208c016d9e = []byte{
	// 211 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2a, 0x49, 0xcd, 0x2d,
	0xc8, 0x2f, 0x4a, 0xcc, 0xd1, 0xaf, 0xcc, 0x4c, 0xcd, 0x49, 0xc9, 0xcd, 0x2f, 0xd6, 0x4f, 0x49,
	0xcd, 0x49, 0x4d, 0x4f, 0x2c, 0xc9, 0xcc, 0xcf, 0x8b, 0xcf, 0xc8, 0x2c, 0x2e, 0xc9, 0x2f, 0xaa,
	0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x84, 0xa9, 0xd5, 0x83, 0xa9, 0x95, 0xd2, 0xc1,
	0xab, 0xbd, 0x24, 0x33, 0x37, 0xb5, 0xb8, 0x24, 0x31, 0xb7, 0x00, 0x62, 0x80, 0x52, 0x3e, 0x97,
	0xa0, 0x0b, 0x5c, 0xd6, 0x03, 0x62, 0xb6, 0x90, 0x04, 0x17, 0x7b, 0x62, 0x4a, 0x4a, 0x51, 0x6a,
	0x71, 0xb1, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x8c, 0x2b, 0xe4, 0xc0, 0xc5, 0x0e, 0x75,
	0x80, 0x04, 0x93, 0x02, 0xb3, 0x06, 0xb7, 0x91, 0x9a, 0x1e, 0x86, 0x0b, 0xf4, 0x10, 0x06, 0x86,
	0xc0, 0x6c, 0x0b, 0x82, 0x69, 0x73, 0xf2, 0x39, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6,
	0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39,
	0x86, 0x28, 0xa3, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0x98, 0xa1,
	0xba, 0x55, 0xf9, 0x79, 0xa9, 0x70, 0x9e, 0x7e, 0x05, 0xc2, 0x4f, 0x25, 0x95, 0x05, 0xa9, 0xc5,
	0x49, 0x6c, 0x60, 0x5f, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x3e, 0x30, 0x7c, 0xfd, 0x34,
	0x01, 0x00, 0x00,
}

func (m *DelegationHistory) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DelegationHistory) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DelegationHistory) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.History) > 0 {
		for iNdEx := len(m.History) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.History[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDelegationHistory(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintDelegationHistory(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDelegationHistory(dAtA []byte, offset int, v uint64) int {
	offset -= sovDelegationHistory(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DelegationHistory) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovDelegationHistory(uint64(l))
	}
	if len(m.History) > 0 {
		for _, e := range m.History {
			l = e.Size()
			n += 1 + l + sovDelegationHistory(uint64(l))
		}
	}
	return n
}

func sovDelegationHistory(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDelegationHistory(x uint64) (n int) {
	return sovDelegationHistory(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DelegationHistory) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDelegationHistory
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
			return fmt.Errorf("proto: DelegationHistory: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DelegationHistory: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelegationHistory
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
				return ErrInvalidLengthDelegationHistory
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDelegationHistory
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field History", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelegationHistory
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
				return ErrInvalidLengthDelegationHistory
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDelegationHistory
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.History = append(m.History, &DelegationTimestamp{})
			if err := m.History[len(m.History)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDelegationHistory(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDelegationHistory
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
func skipDelegationHistory(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDelegationHistory
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
					return 0, ErrIntOverflowDelegationHistory
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
					return 0, ErrIntOverflowDelegationHistory
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
				return 0, ErrInvalidLengthDelegationHistory
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDelegationHistory
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDelegationHistory
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDelegationHistory        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDelegationHistory          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDelegationHistory = fmt.Errorf("proto: unexpected end of group")
)
