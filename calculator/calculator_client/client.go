package main

import (
	"context"
	"fmt"
	"io"
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

	// doSum(c)
	doComputeAverage(c)
}

func doSum(c calculatorpb.CalculatorServiceClient) {
	req := &calculatorpb.CalculateRequest{FirstNum: 10, SecondNum: 20}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while making sum service request %v", err)
	}
	fmt.Println("Result", res.Result)
}

func doPrimeNumberDecomposition(c calculatorpb.CalculatorServiceClient) {
	req := &calculatorpb.PrimeNumberDecompositionRequest{
		Num: 120,
	}
	reqStream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling PrimeNumberDecomposition rpc... %v", err)
	}
	for {
		msg, err := reqStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading streaming: %v", err)
		}
		log.Println("Result: ", msg.GetResult())
	}
}

func doComputeAverage(c calculatorpb.CalculatorServiceClient) {

	requests := []*calculatorpb.ComputeAverageRequest{
		{
			Num: 1,
		},
		{
			Num: 2,
		},
		{
			Num: 3,
		},
		{
			Num: 4,
		},
	}
	stream, err := c.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalf("Error while calling ComputeAverage rpc: %v", err)
	}
	for _, req := range requests {
		log.Printf("sending request: %v", req)
		err := stream.Send(req)
		if err != nil {
			log.Fatalf("error while sending request to ComputeAverage rpc: %v", err)
		}
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while reading ComputeAverage response: %v", err)
	}

	fmt.Println("Result:", res.GetResult())
}
