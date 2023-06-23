// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: temporal/yieldmos/delegation_timestamp.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type DelegationTimestamp struct {
	Timestamp time.Time  `protobuf:"bytes,1,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
	Balance   types.Coin `protobuf:"bytes,2,opt,name=balance,proto3" json:"balance"`
}

func (m *DelegationTimestamp) Reset()         { *m = DelegationTimestamp{} }
func (m *DelegationTimestamp) String() string { return proto.CompactTextString(m) }
func (*DelegationTimestamp) ProtoMessage()    {}
func (*DelegationTimestamp) Descriptor() ([]byte, []int) {
	return fileDescriptor_241987de12034dd1, []int{0}
}
func (m *DelegationTimestamp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DelegationTimestamp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DelegationTimestamp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DelegationTimestamp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelegationTimestamp.Merge(m, src)
}
func (m *DelegationTimestamp) XXX_Size() int {
	return m.Size()
}
func (m *DelegationTimestamp) XXX_DiscardUnknown() {
	xxx_messageInfo_DelegationTimestamp.DiscardUnknown(m)
}

var xxx_messageInfo_DelegationTimestamp proto.InternalMessageInfo

func (m *DelegationTimestamp) GetTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Time{}
}

func (m *DelegationTimestamp) GetBalance() types.Coin {
	if m != nil {
		return m.Balance
	}
	return types.Coin{}
}

func init() {
	proto.RegisterType((*DelegationTimestamp)(nil), "temporal.yieldmos.DelegationTimestamp")
}

func init() {
	proto.RegisterFile("temporal/yieldmos/delegation_timestamp.proto", fileDescriptor_241987de12034dd1)
}

var fileDescriptor_241987de12034dd1 = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0xc6, 0x63, 0x84, 0xf8, 0x13, 0x26, 0x0a, 0x43, 0xc9, 0xe0, 0x20, 0x26, 0x06, 0x38, 0xab,
	0x65, 0x62, 0x0d, 0x8c, 0x4c, 0x15, 0x13, 0x0b, 0xb2, 0xd3, 0xc3, 0x58, 0x4a, 0x72, 0x51, 0xed,
	0x22, 0xca, 0x53, 0x74, 0xe0, 0xa1, 0x3a, 0x76, 0x64, 0x02, 0x94, 0xbc, 0x08, 0x4a, 0x52, 0xa7,
	0x12, 0x9b, 0x4f, 0xdf, 0xcf, 0xdf, 0xfd, 0x74, 0xe1, 0x95, 0xc3, 0xbc, 0xa4, 0x99, 0xcc, 0xc4,
	0xc2, 0x60, 0x36, 0xcd, 0xc9, 0x8a, 0x29, 0x66, 0xa8, 0xa5, 0x33, 0x54, 0x3c, 0x3b, 0x93, 0xa3,
	0x75, 0x32, 0x2f, 0xa1, 0x9c, 0x91, 0xa3, 0xc1, 0xb1, 0xa7, 0xc1, 0xd3, 0xd1, 0xa9, 0x26, 0x4d,
	0x6d, 0x2a, 0x9a, 0x57, 0x07, 0x46, 0x3c, 0x25, 0xdb, 0x74, 0x29, 0x69, 0x51, 0xbc, 0x8d, 0x14,
	0x3a, 0x39, 0x12, 0x29, 0x99, 0x62, 0x93, 0xc7, 0x9a, 0x48, 0x67, 0x28, 0xda, 0x49, 0xcd, 0x5f,
	0xc4, 0xbf, 0x4d, 0x17, 0x9f, 0x2c, 0x3c, 0xb9, 0xef, 0x45, 0x1e, 0x7d, 0x3a, 0x48, 0xc2, 0xc3,
	0x1e, 0x1d, 0xb2, 0x73, 0x76, 0x79, 0x34, 0x8e, 0xa0, 0x2b, 0x03, 0x5f, 0x06, 0x3d, 0x9e, 0x1c,
	0xac, 0xbe, 0xe3, 0x60, 0xf9, 0x13, 0xb3, 0xc9, 0xf6, 0xdb, 0xe0, 0x36, 0xdc, 0x57, 0x32, 0x93,
	0x45, 0x8a, 0xc3, 0x9d, 0xb6, 0xe1, 0x0c, 0x3a, 0x5d, 0x68, 0x74, 0x61, 0xa3, 0x0b, 0x77, 0x64,
	0x8a, 0x64, 0xb7, 0x29, 0x98, 0x78, 0x3e, 0x79, 0x58, 0x55, 0x9c, 0xad, 0x2b, 0xce, 0x7e, 0x2b,
	0xce, 0x96, 0x35, 0x0f, 0xd6, 0x35, 0x0f, 0xbe, 0x6a, 0x1e, 0x3c, 0x8d, 0xb5, 0x71, 0xaf, 0x73,
	0x05, 0x29, 0xe5, 0xc2, 0x5f, 0xe9, 0xfa, 0x83, 0x0a, 0xec, 0x27, 0xf1, 0xbe, 0xbd, 0xb1, 0x5b,
	0x94, 0x68, 0xd5, 0x5e, 0x6b, 0x7c, 0xf3, 0x17, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x25, 0xef, 0xed,
	0x85, 0x01, 0x00, 0x00,
}

func (m *DelegationTimestamp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DelegationTimestamp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DelegationTimestamp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
		i = encodeVarintDelegationTimestamp(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Timestamp, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Timestamp):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintDelegationTimestamp(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintDelegationTimestamp(dAtA []byte, offset int, v uint64) int {
	offset -= sovDelegationTimestamp(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DelegationTimestamp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovDelegationTimestamp(uint64(l))
	l = m.Balance.Size()
	n += 1 + l + sovDelegationTimestamp(uint64(l))
	return n
}

func sovDelegationTimestamp(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDelegationTimestamp(x uint64) (n int) {
	return sovDelegationTimestamp(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DelegationTimestamp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDelegationTimestamp
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
			return fmt.Errorf("proto: DelegationTimestamp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DelegationTimestamp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelegationTimestamp
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
				return ErrInvalidLengthDelegationTimestamp
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDelegationTimestamp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Timestamp, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balance", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelegationTimestamp
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
				return ErrInvalidLengthDelegationTimestamp
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDelegationTimestamp
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
			skippy, err := skipDelegationTimestamp(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDelegationTimestamp
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
func skipDelegationTimestamp(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDelegationTimestamp
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
					return 0, ErrIntOverflowDelegationTimestamp
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
					return 0, ErrIntOverflowDelegationTimestamp
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
				return 0, ErrInvalidLengthDelegationTimestamp
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDelegationTimestamp
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDelegationTimestamp
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDelegationTimestamp        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDelegationTimestamp          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDelegationTimestamp = fmt.Errorf("proto: unexpected end of group")
)
