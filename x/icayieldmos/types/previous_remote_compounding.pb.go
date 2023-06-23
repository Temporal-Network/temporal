// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: temporal/icayieldmos/previous_remote_compounding.proto

package types

import (
	fmt "fmt"
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

type PreviousRemoteCompounding struct {
	RemoteContractCompoundSetting uint64    `protobuf:"varint,1,opt,name=remoteContractCompoundSetting,proto3" json:"remoteContractCompoundSetting,omitempty"`
	Timestamp                     time.Time `protobuf:"bytes,2,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
}

func (m *PreviousRemoteCompounding) Reset()         { *m = PreviousRemoteCompounding{} }
func (m *PreviousRemoteCompounding) String() string { return proto.CompactTextString(m) }
func (*PreviousRemoteCompounding) ProtoMessage()    {}
func (*PreviousRemoteCompounding) Descriptor() ([]byte, []int) {
	return fileDescriptor_18e34a51b41d6112, []int{0}
}
func (m *PreviousRemoteCompounding) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PreviousRemoteCompounding) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PreviousRemoteCompounding.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PreviousRemoteCompounding) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PreviousRemoteCompounding.Merge(m, src)
}
func (m *PreviousRemoteCompounding) XXX_Size() int {
	return m.Size()
}
func (m *PreviousRemoteCompounding) XXX_DiscardUnknown() {
	xxx_messageInfo_PreviousRemoteCompounding.DiscardUnknown(m)
}

var xxx_messageInfo_PreviousRemoteCompounding proto.InternalMessageInfo

func (m *PreviousRemoteCompounding) GetRemoteContractCompoundSetting() uint64 {
	if m != nil {
		return m.RemoteContractCompoundSetting
	}
	return 0
}

func (m *PreviousRemoteCompounding) GetTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Time{}
}

func init() {
	proto.RegisterType((*PreviousRemoteCompounding)(nil), "temporal.icayieldmos.PreviousRemoteCompounding")
}

func init() {
	proto.RegisterFile("temporal/icayieldmos/previous_remote_compounding.proto", fileDescriptor_18e34a51b41d6112)
}

var fileDescriptor_18e34a51b41d6112 = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x63, 0x84, 0x10, 0x84, 0xad, 0xea, 0x50, 0x22, 0xe1, 0x54, 0x4c, 0x5d, 0xb0, 0x25,
	0x10, 0x3c, 0x40, 0x60, 0x07, 0x05, 0x26, 0x96, 0x2a, 0x49, 0x0f, 0x63, 0x29, 0xce, 0x59, 0xce,
	0x05, 0x51, 0x9e, 0xa2, 0x2f, 0xc1, 0xbb, 0x74, 0xec, 0xc8, 0x04, 0x28, 0x79, 0x11, 0x44, 0xd2,
	0xb4, 0xb0, 0xb0, 0xf9, 0x7c, 0xdf, 0xaf, 0xef, 0xf4, 0xfb, 0x97, 0x04, 0xc6, 0xa2, 0x4b, 0x72,
	0xa9, 0xb3, 0x64, 0xae, 0x21, 0x9f, 0x19, 0x2c, 0xa5, 0x75, 0xf0, 0xac, 0xb1, 0x2a, 0xa7, 0x0e,
	0x0c, 0x12, 0x4c, 0x33, 0x34, 0x16, 0xab, 0x62, 0xa6, 0x0b, 0x25, 0xac, 0x43, 0xc2, 0xc1, 0xb0,
	0xcf, 0x89, 0x5f, 0xb9, 0x60, 0xa8, 0x50, 0x61, 0x0b, 0xc8, 0x9f, 0x57, 0xc7, 0x06, 0xa1, 0x42,
	0x54, 0x39, 0xc8, 0x76, 0x4a, 0xab, 0x47, 0x49, 0xda, 0x40, 0x49, 0x89, 0xb1, 0x1d, 0x70, 0xf2,
	0xc6, 0xfc, 0xa3, 0xdb, 0xb5, 0x32, 0x6e, 0x8d, 0x57, 0x5b, 0xe1, 0xe0, 0xda, 0x3f, 0x76, 0xeb,
	0xcf, 0x82, 0x5c, 0x92, 0x51, 0xbf, 0xbc, 0x03, 0x22, 0x5d, 0xa8, 0x11, 0x1b, 0xb3, 0xc9, 0x6e,
	0xfc, 0x3f, 0x34, 0x88, 0xfc, 0x83, 0x8d, 0x76, 0xb4, 0x33, 0x66, 0x93, 0xc3, 0xb3, 0x40, 0x74,
	0x87, 0x89, 0xfe, 0x30, 0x71, 0xdf, 0x13, 0xd1, 0xfe, 0xf2, 0x23, 0xf4, 0x16, 0x9f, 0x21, 0x8b,
	0xb7, 0xb1, 0xe8, 0x66, 0x59, 0x73, 0xb6, 0xaa, 0x39, 0xfb, 0xaa, 0x39, 0x5b, 0x34, 0xdc, 0x5b,
	0x35, 0xdc, 0x7b, 0x6f, 0xb8, 0xf7, 0x70, 0xa1, 0x34, 0x3d, 0x55, 0xa9, 0xc8, 0xd0, 0xc8, 0xbe,
	0x99, 0xd3, 0x57, 0x2c, 0x60, 0x33, 0xc9, 0x97, 0x3f, 0x0d, 0xd3, 0xdc, 0x42, 0x99, 0xee, 0xb5,
	0xe6, 0xf3, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa4, 0x13, 0x67, 0x81, 0x86, 0x01, 0x00, 0x00,
}

func (m *PreviousRemoteCompounding) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PreviousRemoteCompounding) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PreviousRemoteCompounding) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Timestamp, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Timestamp):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintPreviousRemoteCompounding(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x12
	if m.RemoteContractCompoundSetting != 0 {
		i = encodeVarintPreviousRemoteCompounding(dAtA, i, uint64(m.RemoteContractCompoundSetting))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPreviousRemoteCompounding(dAtA []byte, offset int, v uint64) int {
	offset -= sovPreviousRemoteCompounding(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PreviousRemoteCompounding) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RemoteContractCompoundSetting != 0 {
		n += 1 + sovPreviousRemoteCompounding(uint64(m.RemoteContractCompoundSetting))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovPreviousRemoteCompounding(uint64(l))
	return n
}

func sovPreviousRemoteCompounding(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPreviousRemoteCompounding(x uint64) (n int) {
	return sovPreviousRemoteCompounding(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PreviousRemoteCompounding) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPreviousRemoteCompounding
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
			return fmt.Errorf("proto: PreviousRemoteCompounding: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PreviousRemoteCompounding: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemoteContractCompoundSetting", wireType)
			}
			m.RemoteContractCompoundSetting = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPreviousRemoteCompounding
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RemoteContractCompoundSetting |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPreviousRemoteCompounding
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
				return ErrInvalidLengthPreviousRemoteCompounding
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPreviousRemoteCompounding
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Timestamp, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPreviousRemoteCompounding(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPreviousRemoteCompounding
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
func skipPreviousRemoteCompounding(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPreviousRemoteCompounding
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
					return 0, ErrIntOverflowPreviousRemoteCompounding
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
					return 0, ErrIntOverflowPreviousRemoteCompounding
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
				return 0, ErrInvalidLengthPreviousRemoteCompounding
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPreviousRemoteCompounding
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPreviousRemoteCompounding
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPreviousRemoteCompounding        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPreviousRemoteCompounding          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPreviousRemoteCompounding = fmt.Errorf("proto: unexpected end of group")
)
