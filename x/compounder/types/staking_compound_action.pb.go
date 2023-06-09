// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: temporal/compounder/staking_compound_action.proto

package types

import (
	fmt "fmt"
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

type StakingCompoundAction struct {
	Delegator        string     `protobuf:"bytes,1,opt,name=delegator,proto3" json:"delegator,omitempty"`
	ValidatorAddress string     `protobuf:"bytes,2,opt,name=validatorAddress,proto3" json:"validatorAddress,omitempty"`
	Balance          types.Coin `protobuf:"bytes,3,opt,name=balance,proto3" json:"balance"`
}

func (m *StakingCompoundAction) Reset()         { *m = StakingCompoundAction{} }
func (m *StakingCompoundAction) String() string { return proto.CompactTextString(m) }
func (*StakingCompoundAction) ProtoMessage()    {}
func (*StakingCompoundAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_f671e657b7468fa4, []int{0}
}
func (m *StakingCompoundAction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StakingCompoundAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StakingCompoundAction.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StakingCompoundAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StakingCompoundAction.Merge(m, src)
}
func (m *StakingCompoundAction) XXX_Size() int {
	return m.Size()
}
func (m *StakingCompoundAction) XXX_DiscardUnknown() {
	xxx_messageInfo_StakingCompoundAction.DiscardUnknown(m)
}

var xxx_messageInfo_StakingCompoundAction proto.InternalMessageInfo

func (m *StakingCompoundAction) GetDelegator() string {
	if m != nil {
		return m.Delegator
	}
	return ""
}

func (m *StakingCompoundAction) GetValidatorAddress() string {
	if m != nil {
		return m.ValidatorAddress
	}
	return ""
}

func (m *StakingCompoundAction) GetBalance() types.Coin {
	if m != nil {
		return m.Balance
	}
	return types.Coin{}
}

func init() {
	proto.RegisterType((*StakingCompoundAction)(nil), "temporal.compounder.StakingCompoundAction")
}

func init() {
	proto.RegisterFile("temporal/compounder/staking_compound_action.proto", fileDescriptor_f671e657b7468fa4)
}

var fileDescriptor_f671e657b7468fa4 = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x2c, 0x49, 0xcd, 0x2d,
	0xc8, 0x2f, 0x4a, 0xcc, 0xd1, 0x4f, 0xce, 0xcf, 0x2d, 0xc8, 0x2f, 0xcd, 0x4b, 0x49, 0x2d, 0xd2,
	0x2f, 0x2e, 0x49, 0xcc, 0xce, 0xcc, 0x4b, 0x8f, 0x87, 0x09, 0xc5, 0x27, 0x26, 0x97, 0x64, 0xe6,
	0xe7, 0xe9, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x09, 0xc3, 0xb4, 0xe8, 0x21, 0xb4, 0x48, 0x89,
	0xa4, 0xe7, 0xa7, 0xe7, 0x83, 0xe5, 0xf5, 0x41, 0x2c, 0x88, 0x52, 0x29, 0xb9, 0xe4, 0xfc, 0xe2,
	0xdc, 0xfc, 0x62, 0xfd, 0xa4, 0xc4, 0xe2, 0x54, 0xfd, 0x32, 0xc3, 0xa4, 0xd4, 0x92, 0x44, 0x43,
	0xfd, 0xe4, 0xfc, 0x4c, 0xa8, 0x51, 0x4a, 0x73, 0x18, 0xb9, 0x44, 0x83, 0x21, 0x96, 0x39, 0x43,
	0xcd, 0x72, 0x04, 0x5b, 0x25, 0x24, 0xc3, 0xc5, 0x99, 0x92, 0x9a, 0x93, 0x9a, 0x9e, 0x58, 0x92,
	0x5f, 0x24, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x84, 0x10, 0x10, 0xd2, 0xe2, 0x12, 0x28, 0x4b,
	0xcc, 0xc9, 0x4c, 0x01, 0x71, 0x1c, 0x53, 0x52, 0x8a, 0x52, 0x8b, 0x8b, 0x25, 0x98, 0xc0, 0x8a,
	0x30, 0xc4, 0x85, 0x2c, 0xb9, 0xd8, 0x93, 0x12, 0x73, 0x12, 0xf3, 0x92, 0x53, 0x25, 0x98, 0x15,
	0x18, 0x35, 0xb8, 0x8d, 0x24, 0xf5, 0x20, 0xae, 0xd2, 0x03, 0xb9, 0x4a, 0x0f, 0xea, 0x2a, 0x3d,
	0xe7, 0xfc, 0xcc, 0x3c, 0x27, 0x96, 0x13, 0xf7, 0xe4, 0x19, 0x82, 0x60, 0xea, 0x9d, 0xfc, 0x4e,
	0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18,
	0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0xca, 0x24, 0x3d, 0xb3, 0x24, 0xa3, 0x34,
	0x09, 0x14, 0x02, 0xfa, 0xb0, 0xe0, 0xd0, 0xad, 0xca, 0xcf, 0x4b, 0x85, 0xf3, 0xf4, 0x2b, 0x90,
	0x43, 0xb4, 0xa4, 0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d, 0xec, 0x6b, 0x63, 0x40, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x08, 0xed, 0x07, 0xad, 0x75, 0x01, 0x00, 0x00,
}

func (m *StakingCompoundAction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StakingCompoundAction) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StakingCompoundAction) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Balance.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintStakingCompoundAction(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintStakingCompoundAction(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Delegator) > 0 {
		i -= len(m.Delegator)
		copy(dAtA[i:], m.Delegator)
		i = encodeVarintStakingCompoundAction(dAtA, i, uint64(len(m.Delegator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintStakingCompoundAction(dAtA []byte, offset int, v uint64) int {
	offset -= sovStakingCompoundAction(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StakingCompoundAction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Delegator)
	if l > 0 {
		n += 1 + l + sovStakingCompoundAction(uint64(l))
	}
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovStakingCompoundAction(uint64(l))
	}
	l = m.Balance.Size()
	n += 1 + l + sovStakingCompoundAction(uint64(l))
	return n
}

func sovStakingCompoundAction(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStakingCompoundAction(x uint64) (n int) {
	return sovStakingCompoundAction(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StakingCompoundAction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStakingCompoundAction
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
			return fmt.Errorf("proto: StakingCompoundAction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StakingCompoundAction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Delegator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakingCompoundAction
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
				return ErrInvalidLengthStakingCompoundAction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStakingCompoundAction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Delegator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakingCompoundAction
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
				return ErrInvalidLengthStakingCompoundAction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStakingCompoundAction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balance", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakingCompoundAction
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
				return ErrInvalidLengthStakingCompoundAction
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStakingCompoundAction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Balance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStakingCompoundAction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStakingCompoundAction
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
func skipStakingCompoundAction(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStakingCompoundAction
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
					return 0, ErrIntOverflowStakingCompoundAction
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
					return 0, ErrIntOverflowStakingCompoundAction
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
				return 0, ErrInvalidLengthStakingCompoundAction
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStakingCompoundAction
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStakingCompoundAction
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStakingCompoundAction        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStakingCompoundAction          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStakingCompoundAction = fmt.Errorf("proto: unexpected end of group")
)
