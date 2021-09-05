package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/Prithvipal/go-grpc/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct {
}

func (*server) Sum(ctx context.Context, req *calculatorpb.CalculateRequest) (*calculatorpb.CalculateResponse, error) {
	result := req.FirstNum + req.SecondNum
	return &calculatorpb.CalculateResponse{Result: result}, nil
}

func (*server) PrimeNumberDecomposition(req *calculatorpb.PrimeNumberDecompositionRequest, stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error {
	num := req.Num
	var k int32 = 2
	for num > 1 {
		if num%k == 0 {
			res := &calculatorpb.PrimeNumberDecompositionResponse{
				Result: k,
			}
			num = num / k
			stream.Send(res)
			time.Sleep(1 * time.Second)
		} else {
			k = k + 1
		}
	}
	return nil
}

func main() {
	fmt.Println("Hello from calculator server")

	lis, err := net.Listen("tcp", "0.0.0.0:50041")
	if err != nil {
		log.Fatalf("Error while listening the port %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error while serving %v", err)
	}
}
