package rpc

import (
	"log"
	"strconv"
	"time"

	pb "github.com/yjiang-dev/simplemath/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func GreatCommonDivisor(first, second string) {
	// get a connection
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	// create a client
	c := pb.NewSimpleMathClient(conn)
	a, _ := strconv.ParseInt(first, 10, 32)
	b, _ := strconv.ParseInt(second, 10, 32)
	// create a ctx
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// remote call
	r, err := c.GreatCommonDivisor(ctx, &pb.GCDRequest{First: int32(a), Second: int32(b)})
	if err != nil {
		log.Fatalf("could not compute: %v", err)
	}
	log.Printf("The Greatest Common Divisor of %d and %d is %d", a, b, r.Result)
}
