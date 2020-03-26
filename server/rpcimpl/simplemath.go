package rpcimpl

import (
	"io"
	"log"
	"strconv"

	pb "github.com/yjiang-dev/simplemath/api"

	"golang.org/x/net/context"
)

// SimpleMathServer a struct type
type SimpleMathServer struct{}

// GreatCommonDivisor GCD
func (sms *SimpleMathServer) GreatCommonDivisor(ctx context.Context, in *pb.GCDRequest) (*pb.GCDResponse, error) {
	first := in.First
	second := in.Second
	for second != 0 {
		first, second = second, first%second
	}
	return &pb.GCDResponse{Result: first}, nil
}

// GetFibonacci Fibonacci
func (sms *SimpleMathServer) GetFibonacci(in *pb.FibonacciRequest, stream pb.SimpleMath_GetFibonacciServer) error {
	a, b := 0, 1
	for i := 0; i < int(in.Count); i++ {
		stream.Send(&pb.FibonacciResponse{Result: int32(a)})
		a, b = b, a+b
	}
	return nil
}

// Statistics method
func (sms *SimpleMathServer) Statistics(stream pb.SimpleMath_StatisticsServer) error {
	var count, maximum, minimum int32
	minimum = int32((^uint32(0)) >> 1)
	maximum = -minimum - 1
	var average, sum float32
	// receive the requests
	for {
		num, err := stream.Recv()
		if err == io.EOF {
			average = sum / float32(count)
			return stream.SendAndClose(&pb.StatisticsResponse{
				Count:   count,
				Maximum: maximum,
				Minimum: minimum,
				Average: average,
			})
		}
		if err != nil {
			log.Fatalf("failed to recv: %v", err)
			return err
		}
		count++
		if maximum < num.Number {
			maximum = num.Number
		}
		if minimum > num.Number {
			minimum = num.Number
		}
		sum += float32(num.Number)
	}
}

// PrimeFactorization method
func (sms *SimpleMathServer) PrimeFactorization(stream pb.SimpleMath_PrimeFactorizationServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("failed to recv: %v", err)
			return err
		}
		stream.Send(&pb.PrimeFactorizationResponse{Result: primeFactorization(int(in.Number))})
	}
	return nil
}

// primeFactorization method
func primeFactorization(num int) string {
	if num <= 2 {
		return strconv.Itoa(num)
	}
	n := num
	prefix := ""
	result := ""
	for i := 2; i <= n; i++ {
		for n != i {
			if n%i == 0 {
				result += prefix + strconv.Itoa(i)
				prefix = " * "
				n /= i
			} else {
				break
			}
		}
	}
	if result == "" {
		result = "1"
	}
	result = " = " + result + " * " + strconv.Itoa(n)
	return strconv.Itoa(num) + result
}
