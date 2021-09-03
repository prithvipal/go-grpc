package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Prithvipal/go-grpc/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function was involked with %v \n", req)
	result := "Hello " + req.Greeting.GetFirstName()
	return &greetpb.GreetResponse{Result: result}, nil
}

func main() {
	fmt.Println("Hello, I'm a server")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
