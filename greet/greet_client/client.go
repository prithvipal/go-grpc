package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/Prithvipal/go-grpc/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm a client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect %v", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	// fmt.Printf("Created client %v", c)
	// doUnary(c)
	doServerStreaming(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to Unary RPC")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Prithvipal",
			LastName:  "Singh",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", res.Result)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to server streaming RPC")
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Prithvipal",
			LastName:  "Singh",
		},
	}
	reqStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes RPC: %v", err)
	}
	for {
		msg, err := reqStream.Recv()
		if err == io.EOF {
			// we have reached end of stream
			break
		}
		if err != nil {
			log.Fatalf("Error while reading stream %v", err)
		}
		log.Println("Response from GreetManyTimes:", msg.GetResult())
	}
}
