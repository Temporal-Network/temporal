// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: temporal/icayieldmos/contract_remote_zone.proto

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

type ContractRemoteZone struct {
	Id            uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ConnectionId  string `protobuf:"bytes,2,opt,name=connectionId,proto3" json:"connectionId,omitempty"`
	RemoteChainId string `protobuf:"bytes,3,opt,name=remoteChainId,proto3" json:"remoteChainId,omitempty"`
	Active        bool   `protobuf:"varint,4,opt,name=active,proto3" json:"active,omitempty"`
	Creator       string `protobuf:"bytes,5,opt,name=creator,proto3" json:"creator,omitempty"`
	Bech32Prefix  string `protobuf:"bytes,6,opt,name=bech32Prefix,proto3" json:"bech32Prefix,omitempty"`
}

func (m *ContractRemoteZone) Reset()         { *m = ContractRemoteZone{} }
func (m *ContractRemoteZone) String() string { return proto.CompactTextString(m) }
func (*ContractRemoteZone) ProtoMessage()    {}
func (*ContractRemoteZone) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb7167c69c723915, []int{0}
}
func (m *ContractRemoteZone) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContractRemoteZone) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContractRemoteZone.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContractRemoteZone) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractRemoteZone.Merge(m, src)
}
func (m *ContractRemoteZone) XXX_Size() int {
	return m.Size()
}
func (m *ContractRemoteZone) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractRemoteZone.DiscardUnknown(m)
}

var xxx_messageInfo_ContractRemoteZone proto.InternalMessageInfo

func (m *ContractRemoteZone) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ContractRemoteZone) GetConnectionId() string {
	if m != nil {
		return m.ConnectionId
	}
	return ""
}

func (m *ContractRemoteZone) GetRemoteChainId() string {
	if m != nil {
		return m.RemoteChainId
	}
	return ""
}

func (m *ContractRemoteZone) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *ContractRemoteZone) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *ContractRemoteZone) GetBech32Prefix() string {
	if m != nil {
		return m.Bech32Prefix
	}
	return ""
}

func init() {
	proto.RegisterType((*ContractRemoteZone)(nil), "temporal.icayieldmos.ContractRemoteZone")
}

func init() {
	proto.RegisterFile("temporal/icayieldmos/contract_remote_zone.proto", fileDescriptor_eb7167c69c723915)
}

var fileDescriptor_eb7167c69c723915 = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x31, 0x4e, 0xc3, 0x30,
	0x18, 0x85, 0xe3, 0x50, 0x02, 0x58, 0xc0, 0x60, 0x21, 0xe4, 0xc9, 0x8a, 0x2a, 0x86, 0x2c, 0x24,
	0x12, 0x15, 0x17, 0xa0, 0x13, 0x13, 0x28, 0x63, 0x97, 0xca, 0x71, 0x7e, 0x88, 0xa5, 0xc6, 0x7f,
	0xe4, 0x1a, 0xd4, 0x72, 0x0a, 0x0e, 0xc5, 0xc0, 0xd8, 0x91, 0x11, 0x25, 0x17, 0x41, 0x49, 0x5b,
	0x9a, 0x8c, 0xef, 0xe9, 0x7d, 0x7a, 0xd2, 0x47, 0x13, 0x07, 0x65, 0x85, 0x56, 0x2e, 0x12, 0xad,
	0xe4, 0x5a, 0xc3, 0x22, 0x2f, 0x71, 0x99, 0x28, 0x34, 0xce, 0x4a, 0xe5, 0xe6, 0x16, 0x4a, 0x74,
	0x30, 0xff, 0x40, 0x03, 0x71, 0x65, 0xd1, 0x21, 0xbb, 0xda, 0x03, 0x71, 0x0f, 0x18, 0x7f, 0x11,
	0xca, 0xa6, 0x3b, 0x28, 0xed, 0x98, 0x19, 0x1a, 0x60, 0x97, 0xd4, 0xd7, 0x39, 0x27, 0x21, 0x89,
	0x46, 0xa9, 0xaf, 0x73, 0x36, 0xa6, 0xe7, 0x0a, 0x8d, 0x01, 0xe5, 0x34, 0x9a, 0xc7, 0x9c, 0xfb,
	0x21, 0x89, 0xce, 0xd2, 0x41, 0xc7, 0x6e, 0xe8, 0xc5, 0xf6, 0x75, 0x5a, 0x48, 0xdd, 0x8e, 0x8e,
	0xba, 0xd1, 0xb0, 0x64, 0xd7, 0x34, 0x90, 0xca, 0xe9, 0x77, 0xe0, 0xa3, 0x90, 0x44, 0xa7, 0xe9,
	0x2e, 0x31, 0x4e, 0x4f, 0x94, 0x05, 0xe9, 0xd0, 0xf2, 0xe3, 0x8e, 0xdb, 0xc7, 0xf6, 0x3b, 0x03,
	0x55, 0x4c, 0xee, 0x9e, 0x2d, 0xbc, 0xe8, 0x15, 0x0f, 0xb6, 0xdf, 0xfd, 0xee, 0xe1, 0xe9, 0xbb,
	0x16, 0x64, 0x53, 0x0b, 0xf2, 0x5b, 0x0b, 0xf2, 0xd9, 0x08, 0x6f, 0xd3, 0x08, 0xef, 0xa7, 0x11,
	0xde, 0xec, 0xfe, 0x55, 0xbb, 0xe2, 0x2d, 0x8b, 0x15, 0x96, 0xff, 0xca, 0x6e, 0x5b, 0x2d, 0x07,
	0x81, 0xab, 0x81, 0x42, 0xb7, 0xae, 0x60, 0x99, 0x05, 0x9d, 0xb4, 0xc9, 0x5f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x27, 0xea, 0x1d, 0x75, 0x67, 0x01, 0x00, 0x00,
}

func (m *ContractRemoteZone) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContractRemoteZone) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ContractRemoteZone) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Bech32Prefix) > 0 {
		i -= len(m.Bech32Prefix)
		copy(dAtA[i:], m.Bech32Prefix)
		i = encodeVarintContractRemoteZone(dAtA, i, uint64(len(m.Bech32Prefix)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintContractRemoteZone(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Active {
		i--
		if m.Active {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.RemoteChainId) > 0 {
		i -= len(m.RemoteChainId)
		copy(dAtA[i:], m.RemoteChainId)
		i = encodeVarintContractRemoteZone(dAtA, i, uint64(len(m.RemoteChainId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ConnectionId) > 0 {
		i -= len(m.ConnectionId)
		copy(dAtA[i:], m.ConnectionId)
		i = encodeVarintContractRemoteZone(dAtA, i, uint64(len(m.ConnectionId)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintContractRemoteZone(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintContractRemoteZone(dAtA []byte, offset int, v uint64) int {
	offset -= sovContractRemoteZone(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ContractRemoteZone) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovContractRemoteZone(uint64(m.Id))
	}
	l = len(m.ConnectionId)
	if l > 0 {
		n += 1 + l + sovContractRemoteZone(uint64(l))
	}
	l = len(m.RemoteChainId)
	if l > 0 {
		n += 1 + l + sovContractRemoteZone(uint64(l))
	}
	if m.Active {
		n += 2
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovContractRemoteZone(uint64(l))
	}
	l = len(m.Bech32Prefix)
	if l > 0 {
		n += 1 + l + sovContractRemoteZone(uint64(l))
	}
	return n
}

func sovContractRemoteZone(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozContractRemoteZone(x uint64) (n int) {
	return sovContractRemoteZone(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ContractRemoteZone) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowContractRemoteZone
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
			return fmt.Errorf("proto: ContractRemoteZone: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContractRemoteZone: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContractRemoteZone
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContractRemoteZone
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
				return ErrInvalidLengthContractRemoteZone
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthContractRemoteZone
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConnectionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemoteChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContractRemoteZone
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
				return ErrInvalidLengthContractRemoteZone
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthContractRemoteZone
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RemoteChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Active", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContractRemoteZone
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
			m.Active = bool(v != 0)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContractRemoteZone
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
				return ErrInvalidLengthContractRemoteZone
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthContractRemoteZone
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bech32Prefix", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContractRemoteZone
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
				return ErrInvalidLengthContractRemoteZone
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthContractRemoteZone
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bech32Prefix = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipContractRemoteZone(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthContractRemoteZone
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
func skipContractRemoteZone(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowContractRemoteZone
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
					return 0, ErrIntOverflowContractRemoteZone
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
					return 0, ErrIntOverflowContractRemoteZone
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
				return 0, ErrInvalidLengthContractRemoteZone
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupContractRemoteZone
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthContractRemoteZone
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthContractRemoteZone        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowContractRemoteZone          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupContractRemoteZone = fmt.Errorf("proto: unexpected end of group")
)
