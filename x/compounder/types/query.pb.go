// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: temporal/compounder/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_528d2acfcf7783e7, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_528d2acfcf7783e7, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

type QueryGetCompoundSettingsRequest struct {
	Delegator string `protobuf:"bytes,1,opt,name=delegator,proto3" json:"delegator,omitempty"`
}

func (m *QueryGetCompoundSettingsRequest) Reset()         { *m = QueryGetCompoundSettingsRequest{} }
func (m *QueryGetCompoundSettingsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetCompoundSettingsRequest) ProtoMessage()    {}
func (*QueryGetCompoundSettingsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_528d2acfcf7783e7, []int{2}
}
func (m *QueryGetCompoundSettingsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetCompoundSettingsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetCompoundSettingsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetCompoundSettingsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetCompoundSettingsRequest.Merge(m, src)
}
func (m *QueryGetCompoundSettingsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetCompoundSettingsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetCompoundSettingsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetCompoundSettingsRequest proto.InternalMessageInfo

func (m *QueryGetCompoundSettingsRequest) GetDelegator() string {
	if m != nil {
		return m.Delegator
	}
	return ""
}

type QueryGetCompoundSettingsResponse struct {
	CompoundSettings CompoundSettings `protobuf:"bytes,1,opt,name=compoundSettings,proto3" json:"compoundSettings"`
}

func (m *QueryGetCompoundSettingsResponse) Reset()         { *m = QueryGetCompoundSettingsResponse{} }
func (m *QueryGetCompoundSettingsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetCompoundSettingsResponse) ProtoMessage()    {}
func (*QueryGetCompoundSettingsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_528d2acfcf7783e7, []int{3}
}
func (m *QueryGetCompoundSettingsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetCompoundSettingsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetCompoundSettingsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetCompoundSettingsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetCompoundSettingsResponse.Merge(m, src)
}
func (m *QueryGetCompoundSettingsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetCompoundSettingsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetCompoundSettingsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetCompoundSettingsResponse proto.InternalMessageInfo

func (m *QueryGetCompoundSettingsResponse) GetCompoundSettings() CompoundSettings {
	if m != nil {
		return m.CompoundSettings
	}
	return CompoundSettings{}
}

type QueryAllCompoundSettingsRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllCompoundSettingsRequest) Reset()         { *m = QueryAllCompoundSettingsRequest{} }
func (m *QueryAllCompoundSettingsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllCompoundSettingsRequest) ProtoMessage()    {}
func (*QueryAllCompoundSettingsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_528d2acfcf7783e7, []int{4}
}
func (m *QueryAllCompoundSettingsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllCompoundSettingsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllCompoundSettingsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllCompoundSettingsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllCompoundSettingsRequest.Merge(m, src)
}
func (m *QueryAllCompoundSettingsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllCompoundSettingsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllCompoundSettingsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllCompoundSettingsRequest proto.InternalMessageInfo

func (m *QueryAllCompoundSettingsRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllCompoundSettingsResponse struct {
	CompoundSettings []CompoundSettings  `protobuf:"bytes,1,rep,name=compoundSettings,proto3" json:"compoundSettings"`
	Pagination       *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllCompoundSettingsResponse) Reset()         { *m = QueryAllCompoundSettingsResponse{} }
func (m *QueryAllCompoundSettingsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllCompoundSettingsResponse) ProtoMessage()    {}
func (*QueryAllCompoundSettingsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_528d2acfcf7783e7, []int{5}
}
func (m *QueryAllCompoundSettingsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllCompoundSettingsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllCompoundSettingsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllCompoundSettingsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllCompoundSettingsResponse.Merge(m, src)
}
func (m *QueryAllCompoundSettingsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllCompoundSettingsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllCompoundSettingsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllCompoundSettingsResponse proto.InternalMessageInfo

func (m *QueryAllCompoundSettingsResponse) GetCompoundSettings() []CompoundSettings {
	if m != nil {
		return m.CompoundSettings
	}
	return nil
}

func (m *QueryAllCompoundSettingsResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "temporal.compounder.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "temporal.compounder.QueryParamsResponse")
	proto.RegisterType((*QueryGetCompoundSettingsRequest)(nil), "temporal.compounder.QueryGetCompoundSettingsRequest")
	proto.RegisterType((*QueryGetCompoundSettingsResponse)(nil), "temporal.compounder.QueryGetCompoundSettingsResponse")
	proto.RegisterType((*QueryAllCompoundSettingsRequest)(nil), "temporal.compounder.QueryAllCompoundSettingsRequest")
	proto.RegisterType((*QueryAllCompoundSettingsResponse)(nil), "temporal.compounder.QueryAllCompoundSettingsResponse")
}

func init() { proto.RegisterFile("temporal/compounder/query.proto", fileDescriptor_528d2acfcf7783e7) }

var fileDescriptor_528d2acfcf7783e7 = []byte{
	// 531 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0x41, 0x6b, 0x13, 0x41,
	0x1c, 0xc5, 0x33, 0xad, 0x06, 0x3a, 0x5e, 0xca, 0xa4, 0x07, 0x49, 0xea, 0x26, 0x6c, 0xb1, 0x0d,
	0x0a, 0x33, 0x34, 0xb6, 0x07, 0xbd, 0x48, 0x2b, 0xd8, 0x6b, 0x8c, 0x07, 0xc1, 0x8b, 0x4c, 0x36,
	0xe3, 0xb2, 0xb0, 0xd9, 0xd9, 0xee, 0x4c, 0x8a, 0xa1, 0x08, 0xe2, 0x27, 0x10, 0xfc, 0x26, 0x5e,
	0xbd, 0x7a, 0xe8, 0xb1, 0xe0, 0xc5, 0x93, 0x48, 0xe2, 0xd7, 0x10, 0x64, 0x67, 0xfe, 0xdb, 0xa4,
	0xdb, 0xdd, 0x6d, 0xc5, 0xdb, 0xb2, 0xf3, 0xfe, 0xef, 0xff, 0x7e, 0x93, 0x97, 0xc5, 0x6d, 0x2d,
	0xc6, 0xb1, 0x4c, 0x78, 0xc8, 0x3c, 0x39, 0x8e, 0xe5, 0x24, 0x1a, 0x89, 0x84, 0x1d, 0x4f, 0x44,
	0x32, 0xa5, 0x71, 0x22, 0xb5, 0x24, 0x8d, 0x4c, 0x40, 0x17, 0x82, 0xe6, 0x86, 0x2f, 0x7d, 0x69,
	0xce, 0x59, 0xfa, 0x64, 0xa5, 0xcd, 0x4d, 0x5f, 0x4a, 0x3f, 0x14, 0x8c, 0xc7, 0x01, 0xe3, 0x51,
	0x24, 0x35, 0xd7, 0x81, 0x8c, 0x14, 0x9c, 0x3e, 0xf0, 0xa4, 0x1a, 0x4b, 0xc5, 0x86, 0x5c, 0x09,
	0xbb, 0x81, 0x9d, 0xec, 0x0e, 0x85, 0xe6, 0xbb, 0x2c, 0xe6, 0x7e, 0x10, 0x19, 0x31, 0x68, 0x3b,
	0x45, 0xa9, 0x62, 0x9e, 0xf0, 0x71, 0xe6, 0xf6, 0xb0, 0x48, 0x91, 0x3d, 0xbe, 0x51, 0x42, 0xeb,
	0x20, 0xf2, 0x33, 0xf1, 0x56, 0x91, 0xf8, 0x6d, 0x22, 0x8e, 0x27, 0x22, 0xf2, 0xa6, 0x55, 0x8e,
	0x27, 0x3c, 0x0c, 0x46, 0x5c, 0xcb, 0x24, 0xb3, 0x04, 0xb1, 0xb3, 0x0c, 0x93, 0x61, 0x78, 0x32,
	0x00, 0x00, 0x77, 0x03, 0x93, 0x17, 0x29, 0x62, 0xdf, 0x64, 0x1e, 0xa4, 0x9b, 0x94, 0x76, 0xfb,
	0xb8, 0x71, 0xe9, 0xad, 0x8a, 0x65, 0xa4, 0x04, 0x79, 0x8c, 0xeb, 0x96, 0xed, 0x2e, 0xea, 0xa0,
	0xee, 0x9d, 0x5e, 0x8b, 0x16, 0xdc, 0x39, 0xb5, 0x43, 0x87, 0xb7, 0xce, 0x7e, 0xb6, 0x6b, 0x03,
	0x18, 0x70, 0x9f, 0xe2, 0xb6, 0x71, 0x3c, 0x12, 0xfa, 0x19, 0x48, 0x5f, 0x02, 0x3b, 0x2c, 0x25,
	0x9b, 0x78, 0x6d, 0x24, 0x42, 0xe1, 0xa7, 0x14, 0x66, 0xc1, 0xda, 0x60, 0xf1, 0xc2, 0x3d, 0xc5,
	0x9d, 0x72, 0x03, 0xc8, 0xf7, 0x0a, 0xaf, 0x7b, 0xb9, 0x33, 0x48, 0x7a, 0xbf, 0x30, 0x69, 0xde,
	0x08, 0x32, 0x5f, 0x31, 0x71, 0x03, 0x48, 0x7f, 0x10, 0x86, 0x65, 0xe9, 0x9f, 0x63, 0xbc, 0x68,
	0x07, 0x6c, 0xdd, 0xa6, 0xf6, 0xf6, 0x69, 0x7a, 0xfb, 0xd4, 0x96, 0x15, 0x7e, 0x03, 0xda, 0xe7,
	0xbe, 0x80, 0xd9, 0xc1, 0xd2, 0xa4, 0xfb, 0x0d, 0x01, 0x68, 0xe1, 0xae, 0x4a, 0xd0, 0xd5, 0xff,
	0x06, 0x25, 0x47, 0x97, 0x28, 0x56, 0x0c, 0xc5, 0xce, 0xb5, 0x14, 0x36, 0xd5, 0x32, 0x46, 0xef,
	0xcf, 0x2a, 0xbe, 0x6d, 0x30, 0xc8, 0x07, 0x84, 0xeb, 0xb6, 0x12, 0x64, 0xa7, 0x30, 0xdc, 0xd5,
	0xfe, 0x35, 0xbb, 0xd7, 0x0b, 0xed, 0x4e, 0x77, 0xeb, 0xe3, 0xf7, 0xdf, 0x9f, 0x57, 0xee, 0x91,
	0x16, 0x2b, 0xff, 0x27, 0x92, 0xaf, 0x08, 0xaf, 0xe7, 0xaf, 0x80, 0xec, 0x95, 0xef, 0x28, 0x2f,
	0x69, 0x73, 0xff, 0x1f, 0xa7, 0x20, 0xe6, 0x13, 0x13, 0x73, 0x8f, 0xf4, 0xd8, 0x8d, 0x3e, 0x07,
	0xec, 0xf4, 0xa2, 0xf8, 0xef, 0xc9, 0x17, 0x84, 0x1b, 0x79, 0xe3, 0x83, 0x30, 0xac, 0x02, 0x28,
	0xef, 0x69, 0x15, 0x40, 0x45, 0xe3, 0x5c, 0x6a, 0x00, 0xba, 0x64, 0xfb, 0x66, 0x00, 0x87, 0xfb,
	0x67, 0x33, 0x07, 0x9d, 0xcf, 0x1c, 0xf4, 0x6b, 0xe6, 0xa0, 0x4f, 0x73, 0xa7, 0x76, 0x3e, 0x77,
	0x6a, 0x3f, 0xe6, 0x4e, 0xed, 0x75, 0xeb, 0xc2, 0xe0, 0xdd, 0xb2, 0x85, 0x9e, 0xc6, 0x42, 0x0d,
	0xeb, 0xe6, 0xab, 0xf4, 0xe8, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2b, 0xa3, 0xa0, 0x1a, 0xee,
	0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of CompoundSettings items.
	CompoundSettings(ctx context.Context, in *QueryGetCompoundSettingsRequest, opts ...grpc.CallOption) (*QueryGetCompoundSettingsResponse, error)
	CompoundSettingsAll(ctx context.Context, in *QueryAllCompoundSettingsRequest, opts ...grpc.CallOption) (*QueryAllCompoundSettingsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/temporal.compounder.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) CompoundSettings(ctx context.Context, in *QueryGetCompoundSettingsRequest, opts ...grpc.CallOption) (*QueryGetCompoundSettingsResponse, error) {
	out := new(QueryGetCompoundSettingsResponse)
	err := c.cc.Invoke(ctx, "/temporal.compounder.Query/CompoundSettings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) CompoundSettingsAll(ctx context.Context, in *QueryAllCompoundSettingsRequest, opts ...grpc.CallOption) (*QueryAllCompoundSettingsResponse, error) {
	out := new(QueryAllCompoundSettingsResponse)
	err := c.cc.Invoke(ctx, "/temporal.compounder.Query/CompoundSettingsAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a list of CompoundSettings items.
	CompoundSettings(context.Context, *QueryGetCompoundSettingsRequest) (*QueryGetCompoundSettingsResponse, error)
	CompoundSettingsAll(context.Context, *QueryAllCompoundSettingsRequest) (*QueryAllCompoundSettingsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) CompoundSettings(ctx context.Context, req *QueryGetCompoundSettingsRequest) (*QueryGetCompoundSettingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompoundSettings not implemented")
}
func (*UnimplementedQueryServer) CompoundSettingsAll(ctx context.Context, req *QueryAllCompoundSettingsRequest) (*QueryAllCompoundSettingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompoundSettingsAll not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/temporal.compounder.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_CompoundSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetCompoundSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CompoundSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/temporal.compounder.Query/CompoundSettings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CompoundSettings(ctx, req.(*QueryGetCompoundSettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_CompoundSettingsAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllCompoundSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CompoundSettingsAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/temporal.compounder.Query/CompoundSettingsAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CompoundSettingsAll(ctx, req.(*QueryAllCompoundSettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "temporal.compounder.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "CompoundSettings",
			Handler:    _Query_CompoundSettings_Handler,
		},
		{
			MethodName: "CompoundSettingsAll",
			Handler:    _Query_CompoundSettingsAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "temporal/compounder/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryGetCompoundSettingsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetCompoundSettingsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetCompoundSettingsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Delegator) > 0 {
		i -= len(m.Delegator)
		copy(dAtA[i:], m.Delegator)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Delegator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetCompoundSettingsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetCompoundSettingsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetCompoundSettingsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.CompoundSettings.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryAllCompoundSettingsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllCompoundSettingsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllCompoundSettingsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllCompoundSettingsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllCompoundSettingsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllCompoundSettingsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.CompoundSettings) > 0 {
		for iNdEx := len(m.CompoundSettings) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CompoundSettings[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryGetCompoundSettingsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Delegator)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetCompoundSettingsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.CompoundSettings.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryAllCompoundSettingsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllCompoundSettingsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.CompoundSettings) > 0 {
		for _, e := range m.CompoundSettings {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryGetCompoundSettingsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetCompoundSettingsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetCompoundSettingsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Delegator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Delegator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryGetCompoundSettingsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetCompoundSettingsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetCompoundSettingsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CompoundSettings", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CompoundSettings.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAllCompoundSettingsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllCompoundSettingsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllCompoundSettingsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAllCompoundSettingsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllCompoundSettingsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllCompoundSettingsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CompoundSettings", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CompoundSettings = append(m.CompoundSettings, CompoundSettings{})
			if err := m.CompoundSettings[len(m.CompoundSettings)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
