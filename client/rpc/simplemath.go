package rpc

import (
	"log"
	"strconv"
	"time"

	pb "github.com/yjiang-dev/simplemath/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	address = "localhost:50051"
)

func GreatCommonDivisor(first, second string) {
	// create the client TLS credentials
	creds, err := credentials.NewClientTLSFromFile("../cert/server.crt", "")
	// initiate a connection with the server using creds
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSimpleMathClient(conn)
	a, _ := strconv.ParseInt(first, 10, 32)
	b, _ := strconv.ParseInt(second, 10, 32)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GreatCommonDivisor(ctx, &pb.GCDRequest{First: int32(a), Second: int32(b)})
	if err != nil {
		log.Fatalf("could not compute: %v", err)
	}
	log.Printf("The Greatest Common Divisor of %d and %d is %d", a, b, r.Result)
}
