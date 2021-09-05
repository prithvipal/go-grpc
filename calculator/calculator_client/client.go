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
	doPrimeNumberDecomposition(c)
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
