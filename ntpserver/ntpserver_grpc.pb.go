// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ntpserver

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// NtpServiceClient is the client API for NtpService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NtpServiceClient interface {
	GetServer(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type ntpServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNtpServiceClient(cc grpc.ClientConnInterface) NtpServiceClient {
	return &ntpServiceClient{cc}
}

func (c *ntpServiceClient) GetServer(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/ntpserver.NtpService/GetServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NtpServiceServer is the server API for NtpService service.
// All implementations must embed UnimplementedNtpServiceServer
// for forward compatibility
type NtpServiceServer interface {
	GetServer(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedNtpServiceServer()
}

// UnimplementedNtpServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNtpServiceServer struct {
}

func (UnimplementedNtpServiceServer) GetServer(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServer not implemented")
}
func (UnimplementedNtpServiceServer) mustEmbedUnimplementedNtpServiceServer() {}

// UnsafeNtpServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NtpServiceServer will
// result in compilation errors.
type UnsafeNtpServiceServer interface {
	mustEmbedUnimplementedNtpServiceServer()
}

func RegisterNtpServiceServer(s grpc.ServiceRegistrar, srv NtpServiceServer) {
	s.RegisterService(&NtpService_ServiceDesc, srv)
}

func _NtpService_GetServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NtpServiceServer).GetServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ntpserver.NtpService/GetServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NtpServiceServer).GetServer(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// NtpService_ServiceDesc is the grpc.ServiceDesc for NtpService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NtpService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ntpserver.NtpService",
	HandlerType: (*NtpServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetServer",
			Handler:    _NtpService_GetServer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ntpserver/ntpserver.proto",
}
