// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.9.0
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

// DemoClient is the client API for Demo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DemoClient interface {
	Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type demoClient struct {
	cc grpc.ClientConnInterface
}

func NewDemoClient(cc grpc.ClientConnInterface) DemoClient {
	return &demoClient{cc}
}

func (c *demoClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/demo.Demo/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DemoServer is the server API for Demo service.
// All implementations must embed UnimplementedDemoServer
// for forward compatibility
type DemoServer interface {
	Ping(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedDemoServer()
}

// UnimplementedDemoServer must be embedded to have forward compatible implementations.
type UnimplementedDemoServer struct {
}

func (UnimplementedDemoServer) Ping(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedDemoServer) mustEmbedUnimplementedDemoServer() {}

// UnsafeDemoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DemoServer will
// result in compilation errors.
type UnsafeDemoServer interface {
	mustEmbedUnimplementedDemoServer()
}

func RegisterDemoServer(s grpc.ServiceRegistrar, srv DemoServer) {
	s.RegisterService(&Demo_ServiceDesc, srv)
}

func _Demo_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemoServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo.Demo/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemoServer).Ping(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Demo_ServiceDesc is the grpc.ServiceDesc for Demo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Demo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "demo.Demo",
	HandlerType: (*DemoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Demo_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "demo.proto",
}
