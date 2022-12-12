package calculus

import (
	"github.com/woodwo/calculus/grpc/proto"
	"context"
)

type GRPCServer struct{
	FibonacciImplementation func() int
}

// Add method for return next Fibonacci
func (s *GRPCServer) Fibonacci(ctx context.Context, req *proto.Empty) (*proto.Value, error) {
	return &proto.Value{
		Value: int32(s.FibonacciImplementation()),
	}, nil
}