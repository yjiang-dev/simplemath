package main

import (
	"log"
	"net"

	"github.com/yjiang-dev/simplemath/server/rpcimpl"

	pb "github.com/yjiang-dev/simplemath/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// create the TLS credentials from files
	creds, err := credentials.NewServerTLSFromFile("../cert/server.crt", "../cert/server.key")
	if err != nil {
		log.Fatalf("could not load TLS keys: %s", err)
	}
	// create a gRPC option array with the credentials
	opts := []grpc.ServerOption{grpc.Creds(creds)}
	// create a gRPC server object with server options(opts)
	s := grpc.NewServer(opts...)
	pb.RegisterSimpleMathServer(s, &rpcimpl.SimpleMathServer{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
