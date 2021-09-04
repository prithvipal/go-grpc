package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Prithvipal/go-grpc/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct {
}

func (*server) Sum(ctx context.Context, req *calculatorpb.CalculateRequest) (*calculatorpb.CalculateResponse, error) {
	result := req.FirstNum + req.SecondNum
	return &calculatorpb.CalculateResponse{Result: result}, nil
}

func main() {
	fmt.Println("Hello from server")

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
