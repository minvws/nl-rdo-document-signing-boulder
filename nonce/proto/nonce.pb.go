// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nonce/proto/nonce.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	proto1 "github.com/letsencrypt/boulder/core/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type NonceMessage struct {
	Nonce                string   `protobuf:"bytes,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NonceMessage) Reset()         { *m = NonceMessage{} }
func (m *NonceMessage) String() string { return proto.CompactTextString(m) }
func (*NonceMessage) ProtoMessage()    {}
func (*NonceMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_9197b76ef104b424, []int{0}
}

func (m *NonceMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NonceMessage.Unmarshal(m, b)
}
func (m *NonceMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NonceMessage.Marshal(b, m, deterministic)
}
func (m *NonceMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NonceMessage.Merge(m, src)
}
func (m *NonceMessage) XXX_Size() int {
	return xxx_messageInfo_NonceMessage.Size(m)
}
func (m *NonceMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_NonceMessage.DiscardUnknown(m)
}

var xxx_messageInfo_NonceMessage proto.InternalMessageInfo

func (m *NonceMessage) GetNonce() string {
	if m != nil {
		return m.Nonce
	}
	return ""
}

type ValidMessage struct {
	Valid                bool     `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidMessage) Reset()         { *m = ValidMessage{} }
func (m *ValidMessage) String() string { return proto.CompactTextString(m) }
func (*ValidMessage) ProtoMessage()    {}
func (*ValidMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_9197b76ef104b424, []int{1}
}

func (m *ValidMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidMessage.Unmarshal(m, b)
}
func (m *ValidMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidMessage.Marshal(b, m, deterministic)
}
func (m *ValidMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidMessage.Merge(m, src)
}
func (m *ValidMessage) XXX_Size() int {
	return xxx_messageInfo_ValidMessage.Size(m)
}
func (m *ValidMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ValidMessage proto.InternalMessageInfo

func (m *ValidMessage) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func init() {
	proto.RegisterType((*NonceMessage)(nil), "nonce.NonceMessage")
	proto.RegisterType((*ValidMessage)(nil), "nonce.ValidMessage")
}

func init() { proto.RegisterFile("nonce/proto/nonce.proto", fileDescriptor_9197b76ef104b424) }

var fileDescriptor_9197b76ef104b424 = []byte{
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcf, 0xcb, 0xcf, 0x4b,
	0x4e, 0xd5, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x07, 0xb3, 0xf5, 0xc0, 0x6c, 0x21, 0x56, 0x30,
	0x47, 0x4a, 0x34, 0x39, 0xbf, 0x08, 0x26, 0x0d, 0x62, 0x42, 0x64, 0x95, 0x54, 0xb8, 0x78, 0xfc,
	0x40, 0xf2, 0xbe, 0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0xa9, 0x42, 0x22, 0x5c, 0x10, 0xf5, 0x12, 0x8c,
	0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x10, 0x0e, 0x48, 0x55, 0x58, 0x62, 0x4e, 0x66, 0x0a, 0x92, 0xaa,
	0x32, 0x10, 0x1f, 0xac, 0x8a, 0x23, 0x08, 0xc2, 0x31, 0x2a, 0x84, 0x9a, 0x15, 0x9c, 0x5a, 0x54,
	0x96, 0x99, 0x9c, 0x2a, 0xa4, 0xcd, 0xc5, 0x0a, 0xe6, 0x0b, 0x71, 0xeb, 0x81, 0x6d, 0x74, 0xcd,
	0x2d, 0x28, 0xa9, 0x94, 0x12, 0xd6, 0x83, 0xb8, 0x0e, 0xd9, 0x5a, 0x25, 0x06, 0x21, 0x13, 0x2e,
	0xb6, 0xa0, 0xd4, 0x94, 0xd4, 0xd4, 0x5c, 0x21, 0x6c, 0x0a, 0xe0, 0xba, 0x90, 0x9d, 0xa1, 0xc4,
	0xe0, 0xc4, 0x1e, 0xc5, 0x0a, 0xf6, 0x47, 0x12, 0x1b, 0x98, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x9a, 0xba, 0xd9, 0x37, 0x07, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NonceServiceClient is the client API for NonceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NonceServiceClient interface {
	Nonce(ctx context.Context, in *proto1.Empty, opts ...grpc.CallOption) (*NonceMessage, error)
	Redeem(ctx context.Context, in *NonceMessage, opts ...grpc.CallOption) (*ValidMessage, error)
}

type nonceServiceClient struct {
	cc *grpc.ClientConn
}

func NewNonceServiceClient(cc *grpc.ClientConn) NonceServiceClient {
	return &nonceServiceClient{cc}
}

func (c *nonceServiceClient) Nonce(ctx context.Context, in *proto1.Empty, opts ...grpc.CallOption) (*NonceMessage, error) {
	out := new(NonceMessage)
	err := c.cc.Invoke(ctx, "/nonce.NonceService/Nonce", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nonceServiceClient) Redeem(ctx context.Context, in *NonceMessage, opts ...grpc.CallOption) (*ValidMessage, error) {
	out := new(ValidMessage)
	err := c.cc.Invoke(ctx, "/nonce.NonceService/Redeem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NonceServiceServer is the server API for NonceService service.
type NonceServiceServer interface {
	Nonce(context.Context, *proto1.Empty) (*NonceMessage, error)
	Redeem(context.Context, *NonceMessage) (*ValidMessage, error)
}

// UnimplementedNonceServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNonceServiceServer struct {
}

func (*UnimplementedNonceServiceServer) Nonce(ctx context.Context, req *proto1.Empty) (*NonceMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Nonce not implemented")
}
func (*UnimplementedNonceServiceServer) Redeem(ctx context.Context, req *NonceMessage) (*ValidMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Redeem not implemented")
}

func RegisterNonceServiceServer(s *grpc.Server, srv NonceServiceServer) {
	s.RegisterService(&_NonceService_serviceDesc, srv)
}

func _NonceService_Nonce_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NonceServiceServer).Nonce(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nonce.NonceService/Nonce",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NonceServiceServer).Nonce(ctx, req.(*proto1.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _NonceService_Redeem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NonceMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NonceServiceServer).Redeem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nonce.NonceService/Redeem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NonceServiceServer).Redeem(ctx, req.(*NonceMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _NonceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nonce.NonceService",
	HandlerType: (*NonceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Nonce",
			Handler:    _NonceService_Nonce_Handler,
		},
		{
			MethodName: "Redeem",
			Handler:    _NonceService_Redeem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nonce/proto/nonce.proto",
}
