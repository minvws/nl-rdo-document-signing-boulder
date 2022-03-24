// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             v3.18.1
// source: ca.proto

package proto

import (
	context "context"
	proto "github.com/letsencrypt/boulder/core/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CertificateAuthorityClient is the client API for CertificateAuthority service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CertificateAuthorityClient interface {
	IssuePrecertificate(ctx context.Context, in *IssueCertificateRequest, opts ...grpc.CallOption) (*IssuePrecertificateResponse, error)
	IssueCertificateForPrecertificate(ctx context.Context, in *IssueCertificateForPrecertificateRequest, opts ...grpc.CallOption) (*proto.Certificate, error)
	GenerateOCSP(ctx context.Context, in *GenerateOCSPRequest, opts ...grpc.CallOption) (*OCSPResponse, error)
}

type certificateAuthorityClient struct {
	cc grpc.ClientConnInterface
}

func NewCertificateAuthorityClient(cc grpc.ClientConnInterface) CertificateAuthorityClient {
	return &certificateAuthorityClient{cc}
}

func (c *certificateAuthorityClient) IssuePrecertificate(ctx context.Context, in *IssueCertificateRequest, opts ...grpc.CallOption) (*IssuePrecertificateResponse, error) {
	out := new(IssuePrecertificateResponse)
	err := c.cc.Invoke(ctx, "/ca.CertificateAuthority/IssuePrecertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAuthorityClient) IssueCertificateForPrecertificate(ctx context.Context, in *IssueCertificateForPrecertificateRequest, opts ...grpc.CallOption) (*proto.Certificate, error) {
	out := new(proto.Certificate)
	err := c.cc.Invoke(ctx, "/ca.CertificateAuthority/IssueCertificateForPrecertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAuthorityClient) GenerateOCSP(ctx context.Context, in *GenerateOCSPRequest, opts ...grpc.CallOption) (*OCSPResponse, error) {
	out := new(OCSPResponse)
	err := c.cc.Invoke(ctx, "/ca.CertificateAuthority/GenerateOCSP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CertificateAuthorityServer is the server API for CertificateAuthority service.
// All implementations must embed UnimplementedCertificateAuthorityServer
// for forward compatibility
type CertificateAuthorityServer interface {
	IssuePrecertificate(context.Context, *IssueCertificateRequest) (*IssuePrecertificateResponse, error)
	IssueCertificateForPrecertificate(context.Context, *IssueCertificateForPrecertificateRequest) (*proto.Certificate, error)
	GenerateOCSP(context.Context, *GenerateOCSPRequest) (*OCSPResponse, error)
	mustEmbedUnimplementedCertificateAuthorityServer()
}

// UnimplementedCertificateAuthorityServer must be embedded to have forward compatible implementations.
type UnimplementedCertificateAuthorityServer struct {
}

func (UnimplementedCertificateAuthorityServer) IssuePrecertificate(context.Context, *IssueCertificateRequest) (*IssuePrecertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IssuePrecertificate not implemented")
}
func (UnimplementedCertificateAuthorityServer) IssueCertificateForPrecertificate(context.Context, *IssueCertificateForPrecertificateRequest) (*proto.Certificate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IssueCertificateForPrecertificate not implemented")
}
func (UnimplementedCertificateAuthorityServer) GenerateOCSP(context.Context, *GenerateOCSPRequest) (*OCSPResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateOCSP not implemented")
}
func (UnimplementedCertificateAuthorityServer) mustEmbedUnimplementedCertificateAuthorityServer() {}

// UnsafeCertificateAuthorityServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CertificateAuthorityServer will
// result in compilation errors.
type UnsafeCertificateAuthorityServer interface {
	mustEmbedUnimplementedCertificateAuthorityServer()
}

func RegisterCertificateAuthorityServer(s grpc.ServiceRegistrar, srv CertificateAuthorityServer) {
	s.RegisterService(&CertificateAuthority_ServiceDesc, srv)
}

func _CertificateAuthority_IssuePrecertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IssueCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateAuthorityServer).IssuePrecertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ca.CertificateAuthority/IssuePrecertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAuthorityServer).IssuePrecertificate(ctx, req.(*IssueCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateAuthority_IssueCertificateForPrecertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IssueCertificateForPrecertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateAuthorityServer).IssueCertificateForPrecertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ca.CertificateAuthority/IssueCertificateForPrecertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAuthorityServer).IssueCertificateForPrecertificate(ctx, req.(*IssueCertificateForPrecertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateAuthority_GenerateOCSP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateOCSPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateAuthorityServer).GenerateOCSP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ca.CertificateAuthority/GenerateOCSP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAuthorityServer).GenerateOCSP(ctx, req.(*GenerateOCSPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CertificateAuthority_ServiceDesc is the grpc.ServiceDesc for CertificateAuthority service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CertificateAuthority_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ca.CertificateAuthority",
	HandlerType: (*CertificateAuthorityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IssuePrecertificate",
			Handler:    _CertificateAuthority_IssuePrecertificate_Handler,
		},
		{
			MethodName: "IssueCertificateForPrecertificate",
			Handler:    _CertificateAuthority_IssueCertificateForPrecertificate_Handler,
		},
		{
			MethodName: "GenerateOCSP",
			Handler:    _CertificateAuthority_GenerateOCSP_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ca.proto",
}

// OCSPGeneratorClient is the client API for OCSPGenerator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OCSPGeneratorClient interface {
	GenerateOCSP(ctx context.Context, in *GenerateOCSPRequest, opts ...grpc.CallOption) (*OCSPResponse, error)
}

type oCSPGeneratorClient struct {
	cc grpc.ClientConnInterface
}

func NewOCSPGeneratorClient(cc grpc.ClientConnInterface) OCSPGeneratorClient {
	return &oCSPGeneratorClient{cc}
}

func (c *oCSPGeneratorClient) GenerateOCSP(ctx context.Context, in *GenerateOCSPRequest, opts ...grpc.CallOption) (*OCSPResponse, error) {
	out := new(OCSPResponse)
	err := c.cc.Invoke(ctx, "/ca.OCSPGenerator/GenerateOCSP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OCSPGeneratorServer is the server API for OCSPGenerator service.
// All implementations must embed UnimplementedOCSPGeneratorServer
// for forward compatibility
type OCSPGeneratorServer interface {
	GenerateOCSP(context.Context, *GenerateOCSPRequest) (*OCSPResponse, error)
	mustEmbedUnimplementedOCSPGeneratorServer()
}

// UnimplementedOCSPGeneratorServer must be embedded to have forward compatible implementations.
type UnimplementedOCSPGeneratorServer struct {
}

func (UnimplementedOCSPGeneratorServer) GenerateOCSP(context.Context, *GenerateOCSPRequest) (*OCSPResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateOCSP not implemented")
}
func (UnimplementedOCSPGeneratorServer) mustEmbedUnimplementedOCSPGeneratorServer() {}

// UnsafeOCSPGeneratorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OCSPGeneratorServer will
// result in compilation errors.
type UnsafeOCSPGeneratorServer interface {
	mustEmbedUnimplementedOCSPGeneratorServer()
}

func RegisterOCSPGeneratorServer(s grpc.ServiceRegistrar, srv OCSPGeneratorServer) {
	s.RegisterService(&OCSPGenerator_ServiceDesc, srv)
}

func _OCSPGenerator_GenerateOCSP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateOCSPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OCSPGeneratorServer).GenerateOCSP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ca.OCSPGenerator/GenerateOCSP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OCSPGeneratorServer).GenerateOCSP(ctx, req.(*GenerateOCSPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OCSPGenerator_ServiceDesc is the grpc.ServiceDesc for OCSPGenerator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OCSPGenerator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ca.OCSPGenerator",
	HandlerType: (*OCSPGeneratorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateOCSP",
			Handler:    _OCSPGenerator_GenerateOCSP_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ca.proto",
}
