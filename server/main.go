package main

import (
	"log"
	"net"

	"github.com/yjiang-dev/simplemath/server/rpcimpl"

	pb "github.com/yjiang-dev/simplemath/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// the address to bind
const (
	port = ":50051"
)

// main start a gRPC server and waits for connection
func main() {
	// create a listener on TCP port 50051
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// create a gRPC server object
	s := grpc.NewServer()
	// attach the server instance to ther gRPC server
	pb.RegisterSimpleMathServer(s, &rpcimpl.SimpleMathServer{})
	reflection.Register(s)
	// start the server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
