// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: temporal/icayieldmos/msg_execute_contract.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
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

// MsgExecuteContract submits the given message data to a smart contract
type MsgExecuteContract struct {
	// Sender is the that actor that signed the messages
	Sender string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	// Contract is the address of the smart contract
	Contract string `protobuf:"bytes,2,opt,name=contract,proto3" json:"contract,omitempty"`
	// Msg json encoded message to be passed to the contract
	Msg RawContractMessage `protobuf:"bytes,3,opt,name=msg,proto3,casttype=RawContractMessage" json:"msg,omitempty"`
	// Funds coins that are transferred to the contract on execution
	Funds github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,5,rep,name=funds,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"funds"`
}

func (m *MsgExecuteContract) Reset()         { *m = MsgExecuteContract{} }
func (m *MsgExecuteContract) String() string { return proto.CompactTextString(m) }
func (*MsgExecuteContract) ProtoMessage()    {}
func (*MsgExecuteContract) Descriptor() ([]byte, []int) {
	return fileDescriptor_d613da6aeeee03b1, []int{0}
}
func (m *MsgExecuteContract) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgExecuteContract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgExecuteContract.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgExecuteContract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgExecuteContract.Merge(m, src)
}
func (m *MsgExecuteContract) XXX_Size() int {
	return m.Size()
}
func (m *MsgExecuteContract) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgExecuteContract.DiscardUnknown(m)
}

var xxx_messageInfo_MsgExecuteContract proto.InternalMessageInfo

func (m *MsgExecuteContract) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *MsgExecuteContract) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func (m *MsgExecuteContract) GetMsg() RawContractMessage {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (m *MsgExecuteContract) GetFunds() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Funds
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgExecuteContract)(nil), "cosmwasm.wasm.v1.MsgExecuteContract")
}

func init() {
	proto.RegisterFile("temporal/icayieldmos/msg_execute_contract.proto", fileDescriptor_d613da6aeeee03b1)
}

var fileDescriptor_d613da6aeeee03b1 = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x51, 0x3d, 0x6f, 0xea, 0x30,
	0x14, 0x4d, 0x1e, 0x02, 0xbd, 0x17, 0xde, 0x64, 0x21, 0x5e, 0x60, 0x30, 0xe8, 0x4d, 0x51, 0x25,
	0xe2, 0xd2, 0x2e, 0x55, 0x47, 0x50, 0x47, 0x3a, 0x44, 0x9d, 0xba, 0x20, 0xc7, 0x71, 0xdd, 0x08,
	0x12, 0xa3, 0x5c, 0xf3, 0xb5, 0xf6, 0x17, 0xf4, 0x77, 0x74, 0xea, 0xcf, 0x60, 0x64, 0xec, 0x44,
	0xab, 0x30, 0xf4, 0x3f, 0x74, 0xaa, 0x12, 0x9b, 0x96, 0x4e, 0xf7, 0xde, 0x9c, 0x7b, 0x6e, 0xce,
	0xf1, 0x71, 0x88, 0xe2, 0xc9, 0x4c, 0x66, 0x74, 0x4a, 0x62, 0x46, 0xd7, 0x31, 0x9f, 0x46, 0x89,
	0x04, 0x92, 0x80, 0x18, 0xf3, 0x15, 0x67, 0x73, 0xc5, 0xc7, 0x4c, 0xa6, 0x2a, 0xa3, 0x4c, 0xf9,
	0xb3, 0x4c, 0x2a, 0x89, 0x1a, 0x07, 0x82, 0x7f, 0x44, 0x68, 0x63, 0x26, 0xa1, 0x20, 0x86, 0x14,
	0x38, 0x59, 0xf4, 0x43, 0xae, 0x68, 0x9f, 0x30, 0x19, 0xa7, 0x9a, 0xd5, 0xfe, 0x67, 0xf0, 0x04,
	0x04, 0x59, 0xf4, 0x8b, 0x62, 0x80, 0x86, 0x90, 0x42, 0x96, 0x2d, 0x29, 0x3a, 0xf3, 0xb5, 0xa5,
	0xd7, 0xc7, 0x1a, 0xd0, 0x83, 0x86, 0xfe, 0xe7, 0xb6, 0x83, 0x46, 0x20, 0xae, 0xb4, 0xba, 0xa1,
	0x11, 0x87, 0x9a, 0x4e, 0x0d, 0x78, 0x1a, 0xf1, 0xcc, 0xb5, 0xbb, 0xb6, 0xf7, 0x27, 0x30, 0x13,
	0x6a, 0x3b, 0xbf, 0x0f, 0x06, 0xdc, 0x5f, 0x25, 0xf2, 0x35, 0x23, 0xcf, 0xa9, 0x24, 0x20, 0xdc,
	0x4a, 0xd7, 0xf6, 0xfe, 0x0e, 0x9a, 0x1f, 0xbb, 0x0e, 0x0a, 0xe8, 0xf2, 0x70, 0x71, 0xc4, 0x01,
	0xa8, 0xe0, 0x41, 0xb1, 0x82, 0xa8, 0x53, 0xbd, 0x9b, 0xa7, 0x11, 0xb8, 0xd5, 0x6e, 0xc5, 0xab,
	0x9f, 0xb5, 0x7c, 0x23, 0xa9, 0xb0, 0xeb, 0x1b, 0xbb, 0xfe, 0x50, 0xc6, 0xe9, 0xe0, 0x74, 0xb3,
	0xeb, 0x58, 0x4f, 0xaf, 0x1d, 0x4f, 0xc4, 0xea, 0x7e, 0x1e, 0xfa, 0x4c, 0x26, 0x46, 0xbf, 0x29,
	0x3d, 0x88, 0x26, 0x44, 0xad, 0x67, 0x1c, 0x4a, 0x02, 0x04, 0xfa, 0xf2, 0x65, 0xfd, 0xe1, 0xfd,
	0xf9, 0xc4, 0xa8, 0x1e, 0x04, 0x9b, 0x1c, 0xdb, 0xdb, 0x1c, 0xdb, 0x6f, 0x39, 0xb6, 0x1f, 0xf7,
	0xd8, 0xda, 0xee, 0xb1, 0xf5, 0xb2, 0xc7, 0xd6, 0xed, 0xc5, 0xd1, 0xdd, 0x1b, 0x93, 0x44, 0xef,
	0x9a, 0xab, 0xa5, 0xcc, 0x26, 0xdf, 0x59, 0xae, 0x7e, 0xa4, 0x59, 0xfe, 0x2d, 0xac, 0x95, 0xef,
	0x77, 0xfe, 0x19, 0x00, 0x00, 0xff, 0xff, 0x41, 0xe6, 0xb1, 0xdd, 0xf2, 0x01, 0x00, 0x00,
}

func (m *MsgExecuteContract) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgExecuteContract) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgExecuteContract) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Funds) > 0 {
		for iNdEx := len(m.Funds) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Funds[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMsgExecuteContract(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Msg) > 0 {
		i -= len(m.Msg)
		copy(dAtA[i:], m.Msg)
		i = encodeVarintMsgExecuteContract(dAtA, i, uint64(len(m.Msg)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Contract) > 0 {
		i -= len(m.Contract)
		copy(dAtA[i:], m.Contract)
		i = encodeVarintMsgExecuteContract(dAtA, i, uint64(len(m.Contract)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintMsgExecuteContract(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMsgExecuteContract(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsgExecuteContract(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgExecuteContract) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovMsgExecuteContract(uint64(l))
	}
	l = len(m.Contract)
	if l > 0 {
		n += 1 + l + sovMsgExecuteContract(uint64(l))
	}
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovMsgExecuteContract(uint64(l))
	}
	if len(m.Funds) > 0 {
		for _, e := range m.Funds {
			l = e.Size()
			n += 1 + l + sovMsgExecuteContract(uint64(l))
		}
	}
	return n
}

func sovMsgExecuteContract(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsgExecuteContract(x uint64) (n int) {
	return sovMsgExecuteContract(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgExecuteContract) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgExecuteContract
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
			return fmt.Errorf("proto: MsgExecuteContract: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgExecuteContract: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgExecuteContract
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
				return ErrInvalidLengthMsgExecuteContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgExecuteContract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contract", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgExecuteContract
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
				return ErrInvalidLengthMsgExecuteContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgExecuteContract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contract = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgExecuteContract
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMsgExecuteContract
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgExecuteContract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msg = append(m.Msg[:0], dAtA[iNdEx:postIndex]...)
			if m.Msg == nil {
				m.Msg = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Funds", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgExecuteContract
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
				return ErrInvalidLengthMsgExecuteContract
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsgExecuteContract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Funds = append(m.Funds, types.Coin{})
			if err := m.Funds[len(m.Funds)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsgExecuteContract(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgExecuteContract
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
func skipMsgExecuteContract(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsgExecuteContract
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
					return 0, ErrIntOverflowMsgExecuteContract
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
					return 0, ErrIntOverflowMsgExecuteContract
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
				return 0, ErrInvalidLengthMsgExecuteContract
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsgExecuteContract
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsgExecuteContract
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsgExecuteContract        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsgExecuteContract          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsgExecuteContract = fmt.Errorf("proto: unexpected end of group")
)