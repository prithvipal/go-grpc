package main

import (
	"context"
	"fmt"
	"io"
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

func (*server) ComputeAverage(stream calculatorpb.CalculatorService_ComputeAverageServer) error {
	fmt.Println("ComputeAverage function was involked with stream request ")
	sum := int32(0)
	count := 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			result := float64(sum) / float64(count)
			return stream.SendAndClose(&calculatorpb.ComputeAverageResponse{
				Result: float64(result),
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}
		sum += req.GetNum()
		count++
	}

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
