package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Prithvipal/go-grpc/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello from client")

	cc, err := grpc.Dial("localhost:50041", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error while dial %v", err)
	}

	defer cc.Close()
	c := calculatorpb.NewCalculatorServiceClient(cc)
	req := &calculatorpb.CalculateRequest{FirstNum: 10, SecondNum: 20}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while making sum service request %v", err)
	}
	fmt.Println("Result", res.Result)
}
