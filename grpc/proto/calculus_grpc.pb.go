// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: grpc/proto/calculus.proto

package proto

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

// CalculusClient is the client API for Calculus service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalculusClient interface {
	Fibonacci(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Value, error)
}

type calculusClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculusClient(cc grpc.ClientConnInterface) CalculusClient {
	return &calculusClient{cc}
}

func (c *calculusClient) Fibonacci(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Value, error) {
	out := new(Value)
	err := c.cc.Invoke(ctx, "/calculus.calculus/Fibonacci", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculusServer is the server API for Calculus service.
// All implementations should embed UnimplementedCalculusServer
// for forward compatibility
type CalculusServer interface {
	Fibonacci(context.Context, *Empty) (*Value, error)
}

// UnimplementedCalculusServer should be embedded to have forward compatible implementations.
type UnimplementedCalculusServer struct {
}

func (UnimplementedCalculusServer) Fibonacci(context.Context, *Empty) (*Value, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fibonacci not implemented")
}

// UnsafeCalculusServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalculusServer will
// result in compilation errors.
type UnsafeCalculusServer interface {
	mustEmbedUnimplementedCalculusServer()
}

func RegisterCalculusServer(s grpc.ServiceRegistrar, srv CalculusServer) {
	s.RegisterService(&Calculus_ServiceDesc, srv)
}

func _Calculus_Fibonacci_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculusServer).Fibonacci(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculus.calculus/Fibonacci",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculusServer).Fibonacci(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Calculus_ServiceDesc is the grpc.ServiceDesc for Calculus service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Calculus_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "calculus.calculus",
	HandlerType: (*CalculusServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Fibonacci",
			Handler:    _Calculus_Fibonacci_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/proto/calculus.proto",
}
