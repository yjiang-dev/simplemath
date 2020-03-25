package rpcimpl

import (
	pb "github.com/yjiang-dev/simplemath/api"

	"golang.org/x/net/context"
)

type SimpleMathServer struct{}

func (sms *SimpleMathServer) GreatCommonDivisor(ctx context.Context, in *pb.GCDRequest) (*pb.GCDResponse, error) {
	first := in.First
	second := in.Second
	for second != 0 {
		first, second = second, first%second
	}
	return &pb.GCDResponse{Result: first}, nil
}
