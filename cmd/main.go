package main

import (
	"github.com/woodwo/calculus/grpc/proto"
	"github.com/woodwo/calculus/pkg/server"
	"log"
	"net"

	"google.golang.org/grpc"
)


func main() {
    // Create new gRPC server instance
    s := grpc.NewServer()

    // Put function that return Fibonacci
    srv := &calculus.GRPCServer{
        FibonacciImplementation: calculus.Fibonacci(),
    }

    // Register gRPC server for handle
    proto.RegisterCalculusServer(s, srv)

    // Listen on port 8080
    l, err := net.Listen("tcp", ":8080")
    if err != nil {
        log.Fatal(err)
    }

    // Start gRPC server
    if err := s.Serve(l); err != nil {
        log.Fatal(err)
    }
}
