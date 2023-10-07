// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: demo.proto

package demo

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

// DemonstratorClient is the client API for Demonstrator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DemonstratorClient interface {
	DoSomething(ctx context.Context, in *DoRequest, opts ...grpc.CallOption) (*DoReply, error)
}

type demonstratorClient struct {
	cc grpc.ClientConnInterface
}

func NewDemonstratorClient(cc grpc.ClientConnInterface) DemonstratorClient {
	return &demonstratorClient{cc}
}

func (c *demonstratorClient) DoSomething(ctx context.Context, in *DoRequest, opts ...grpc.CallOption) (*DoReply, error) {
	out := new(DoReply)
	err := c.cc.Invoke(ctx, "/demo.Demonstrator/DoSomething", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DemonstratorServer is the server API for Demonstrator service.
// All implementations must embed UnimplementedDemonstratorServer
// for forward compatibility
type DemonstratorServer interface {
	DoSomething(context.Context, *DoRequest) (*DoReply, error)
	mustEmbedUnimplementedDemonstratorServer()
}

// UnimplementedDemonstratorServer must be embedded to have forward compatible implementations.
type UnimplementedDemonstratorServer struct {
}

func (UnimplementedDemonstratorServer) DoSomething(context.Context, *DoRequest) (*DoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoSomething not implemented")
}
func (UnimplementedDemonstratorServer) mustEmbedUnimplementedDemonstratorServer() {}

// UnsafeDemonstratorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DemonstratorServer will
// result in compilation errors.
type UnsafeDemonstratorServer interface {
	mustEmbedUnimplementedDemonstratorServer()
}

func RegisterDemonstratorServer(s grpc.ServiceRegistrar, srv DemonstratorServer) {
	s.RegisterService(&Demonstrator_ServiceDesc, srv)
}

func _Demonstrator_DoSomething_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemonstratorServer).DoSomething(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo.Demonstrator/DoSomething",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemonstratorServer).DoSomething(ctx, req.(*DoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Demonstrator_ServiceDesc is the grpc.ServiceDesc for Demonstrator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Demonstrator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "demo.Demonstrator",
	HandlerType: (*DemonstratorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DoSomething",
			Handler:    _Demonstrator_DoSomething_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "demo.proto",
}