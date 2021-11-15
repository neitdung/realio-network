// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: asset/query.proto

package types

import (
	context "context"
	fmt "fmt"
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

type QueryGetTokenRequest struct {
	Index string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
}

func (m *QueryGetTokenRequest) Reset()         { *m = QueryGetTokenRequest{} }
func (m *QueryGetTokenRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetTokenRequest) ProtoMessage()    {}
func (*QueryGetTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4785ee6dd031b7c8, []int{0}
}
func (m *QueryGetTokenRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetTokenRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetTokenRequest.Merge(m, src)
}
func (m *QueryGetTokenRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetTokenRequest proto.InternalMessageInfo

func (m *QueryGetTokenRequest) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

type QueryGetTokenResponse struct {
	Token Token `protobuf:"bytes,1,opt,name=token,proto3" json:"token"`
}

func (m *QueryGetTokenResponse) Reset()         { *m = QueryGetTokenResponse{} }
func (m *QueryGetTokenResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetTokenResponse) ProtoMessage()    {}
func (*QueryGetTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4785ee6dd031b7c8, []int{1}
}
func (m *QueryGetTokenResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetTokenResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetTokenResponse.Merge(m, src)
}
func (m *QueryGetTokenResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetTokenResponse proto.InternalMessageInfo

func (m *QueryGetTokenResponse) GetToken() Token {
	if m != nil {
		return m.Token
	}
	return Token{}
}

type QueryAllTokenRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllTokenRequest) Reset()         { *m = QueryAllTokenRequest{} }
func (m *QueryAllTokenRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllTokenRequest) ProtoMessage()    {}
func (*QueryAllTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4785ee6dd031b7c8, []int{2}
}
func (m *QueryAllTokenRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllTokenRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllTokenRequest.Merge(m, src)
}
func (m *QueryAllTokenRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllTokenRequest proto.InternalMessageInfo

func (m *QueryAllTokenRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllTokenResponse struct {
	Token      []Token             `protobuf:"bytes,1,rep,name=token,proto3" json:"token"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllTokenResponse) Reset()         { *m = QueryAllTokenResponse{} }
func (m *QueryAllTokenResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllTokenResponse) ProtoMessage()    {}
func (*QueryAllTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4785ee6dd031b7c8, []int{3}
}
func (m *QueryAllTokenResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllTokenResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllTokenResponse.Merge(m, src)
}
func (m *QueryAllTokenResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllTokenResponse proto.InternalMessageInfo

func (m *QueryAllTokenResponse) GetToken() []Token {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *QueryAllTokenResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryGetTokenRequest)(nil), "realiotech.network.asset.QueryGetTokenRequest")
	proto.RegisterType((*QueryGetTokenResponse)(nil), "realiotech.network.asset.QueryGetTokenResponse")
	proto.RegisterType((*QueryAllTokenRequest)(nil), "realiotech.network.asset.QueryAllTokenRequest")
	proto.RegisterType((*QueryAllTokenResponse)(nil), "realiotech.network.asset.QueryAllTokenResponse")
}

func init() { proto.RegisterFile("asset/query.proto", fileDescriptor_4785ee6dd031b7c8) }

var fileDescriptor_4785ee6dd031b7c8 = []byte{
	// 426 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xb1, 0xae, 0xd3, 0x30,
	0x14, 0x86, 0xe3, 0x42, 0x10, 0x98, 0x09, 0xab, 0x48, 0x55, 0x85, 0x52, 0xc8, 0x40, 0x01, 0x81,
	0xad, 0x96, 0x91, 0xa9, 0x1d, 0xe8, 0xc2, 0x00, 0x51, 0x27, 0x06, 0x24, 0xa7, 0x1c, 0xa5, 0x51,
	0xd3, 0x38, 0x8d, 0x5d, 0x68, 0x85, 0x58, 0x78, 0x02, 0x24, 0x58, 0x18, 0x78, 0x11, 0x9e, 0xa0,
	0x63, 0x25, 0x16, 0x26, 0x84, 0x5a, 0x1e, 0xe4, 0x2a, 0xb6, 0xab, 0xa6, 0xbd, 0xb7, 0xca, 0xbd,
	0xdb, 0x89, 0xf5, 0xff, 0xff, 0xf9, 0x7c, 0x4e, 0x8c, 0xef, 0x70, 0x29, 0x41, 0xb1, 0xd9, 0x1c,
	0xf2, 0x25, 0xcd, 0x72, 0xa1, 0x04, 0x69, 0xe4, 0xc0, 0x93, 0x58, 0x28, 0x18, 0x8d, 0x69, 0x0a,
	0xea, 0xa3, 0xc8, 0x27, 0x54, 0xab, 0x9a, 0xf7, 0x22, 0x21, 0xa2, 0x04, 0x18, 0xcf, 0x62, 0xc6,
	0xd3, 0x54, 0x28, 0xae, 0x62, 0x91, 0x4a, 0xe3, 0x6b, 0x3e, 0x19, 0x09, 0x39, 0x15, 0x92, 0x85,
	0x5c, 0x82, 0x09, 0x64, 0x1f, 0x3a, 0x21, 0x28, 0xde, 0x61, 0x19, 0x8f, 0xe2, 0x54, 0x8b, 0xad,
	0xd6, 0xb6, 0x55, 0x62, 0x02, 0xbb, 0xa3, 0x7a, 0x24, 0x22, 0xa1, 0x4b, 0x56, 0x54, 0xe6, 0xd4,
	0x7f, 0x8a, 0xeb, 0x6f, 0x8a, 0xa8, 0x01, 0xa8, 0x61, 0x21, 0x0e, 0x60, 0x36, 0x07, 0xa9, 0x48,
	0x1d, 0xbb, 0x71, 0xfa, 0x1e, 0x16, 0x0d, 0x74, 0x1f, 0x3d, 0xba, 0x15, 0x98, 0x0f, 0x7f, 0x88,
	0xef, 0x1e, 0xa9, 0x65, 0x26, 0x52, 0x09, 0xe4, 0x05, 0x76, 0x75, 0x2f, 0x2d, 0xbf, 0xdd, 0x6d,
	0xd1, 0x53, 0x77, 0xa4, 0xda, 0xd7, 0xbf, 0xbe, 0xfa, 0xdb, 0x72, 0x02, 0xe3, 0xf1, 0xdf, 0x59,
	0x86, 0x5e, 0x92, 0x1c, 0x30, 0xbc, 0xc4, 0x78, 0x7f, 0x31, 0x9b, 0xfc, 0x90, 0x9a, 0x29, 0xd0,
	0x62, 0x0a, 0xd4, 0x8c, 0xd5, 0x4e, 0x81, 0xbe, 0xe6, 0x11, 0x58, 0x6f, 0x50, 0x72, 0xfa, 0x3f,
	0x91, 0xc5, 0xde, 0x37, 0x38, 0x8f, 0x7d, 0xed, 0xaa, 0xd8, 0x64, 0x70, 0x80, 0x57, 0xd3, 0x78,
	0xed, 0x4a, 0x3c, 0xd3, 0xb9, 0xcc, 0xd7, 0xfd, 0x55, 0xc3, 0xae, 0xe6, 0x23, 0x3f, 0x10, 0x76,
	0x75, 0x27, 0x42, 0x4f, 0xa3, 0x5c, 0xb4, 0xaf, 0x26, 0xbb, 0xb4, 0xde, 0x00, 0xf8, 0xec, 0xcb,
	0xef, 0xff, 0xdf, 0x6a, 0x8f, 0x49, 0x9b, 0xed, 0x8d, 0xcc, 0x1a, 0x59, 0xe9, 0xef, 0x61, 0x9f,
	0xf4, 0xea, 0x3f, 0x93, 0xef, 0x08, 0xdf, 0xd4, 0x11, 0xbd, 0x24, 0xa9, 0xc4, 0x3b, 0x5a, 0x65,
	0x25, 0xde, 0xf1, 0x66, 0xfc, 0xb6, 0xc6, 0x7b, 0x40, 0x5a, 0x15, 0x78, 0xfd, 0x57, 0xab, 0x8d,
	0x87, 0xd6, 0x1b, 0x0f, 0xfd, 0xdb, 0x78, 0xe8, 0xeb, 0xd6, 0x73, 0xd6, 0x5b, 0xcf, 0xf9, 0xb3,
	0xf5, 0x9c, 0xb7, 0xdd, 0x28, 0x56, 0xe3, 0x79, 0x48, 0x47, 0x62, 0x5a, 0x0e, 0x31, 0xe5, 0xb3,
	0x5d, 0xd6, 0x62, 0x97, 0xb6, 0xcc, 0x40, 0x86, 0x37, 0xf4, 0xab, 0x78, 0x7e, 0x16, 0x00, 0x00,
	0xff, 0xff, 0x1c, 0x9f, 0x50, 0x48, 0xb7, 0x03, 0x00, 0x00,
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
	// Queries a token by index.
	Token(ctx context.Context, in *QueryGetTokenRequest, opts ...grpc.CallOption) (*QueryGetTokenResponse, error)
	// Queries a list of token items.
	TokenAll(ctx context.Context, in *QueryAllTokenRequest, opts ...grpc.CallOption) (*QueryAllTokenResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Token(ctx context.Context, in *QueryGetTokenRequest, opts ...grpc.CallOption) (*QueryGetTokenResponse, error) {
	out := new(QueryGetTokenResponse)
	err := c.cc.Invoke(ctx, "/realiotech.network.asset.Query/Token", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) TokenAll(ctx context.Context, in *QueryAllTokenRequest, opts ...grpc.CallOption) (*QueryAllTokenResponse, error) {
	out := new(QueryAllTokenResponse)
	err := c.cc.Invoke(ctx, "/realiotech.network.asset.Query/TokenAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Queries a token by index.
	Token(context.Context, *QueryGetTokenRequest) (*QueryGetTokenResponse, error)
	// Queries a list of token items.
	TokenAll(context.Context, *QueryAllTokenRequest) (*QueryAllTokenResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Token(ctx context.Context, req *QueryGetTokenRequest) (*QueryGetTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Token not implemented")
}
func (*UnimplementedQueryServer) TokenAll(ctx context.Context, req *QueryAllTokenRequest) (*QueryAllTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TokenAll not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Token_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Token(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realiotech.network.asset.Query/Token",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Token(ctx, req.(*QueryGetTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_TokenAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TokenAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realiotech.network.asset.Query/TokenAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TokenAll(ctx, req.(*QueryAllTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "realiotech.network.asset.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Token",
			Handler:    _Query_Token_Handler,
		},
		{
			MethodName: "TokenAll",
			Handler:    _Query_TokenAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "asset/query.proto",
}

func (m *QueryGetTokenRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetTokenRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetTokenRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetTokenResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetTokenResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetTokenResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Token.MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryAllTokenRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllTokenRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllTokenRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *QueryAllTokenResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllTokenResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllTokenResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
	if len(m.Token) > 0 {
		for iNdEx := len(m.Token) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Token[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
func (m *QueryGetTokenRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetTokenResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Token.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryAllTokenRequest) Size() (n int) {
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

func (m *QueryAllTokenResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Token) > 0 {
		for _, e := range m.Token {
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
func (m *QueryGetTokenRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGetTokenRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetTokenRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
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
			m.Index = string(dAtA[iNdEx:postIndex])
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
func (m *QueryGetTokenResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGetTokenResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetTokenResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
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
			if err := m.Token.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryAllTokenRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAllTokenRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllTokenRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryAllTokenResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAllTokenResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllTokenResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
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
			m.Token = append(m.Token, Token{})
			if err := m.Token[len(m.Token)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
