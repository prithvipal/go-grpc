package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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
	// doServerStreaming(c)
	doClientStreaming(c)
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

func doClientStreaming(c greetpb.GreetServiceClient) {
	requests := []*greetpb.LongGreetRequest{
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Prithvi",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Parth",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Rajani",
			},
		},
	}
	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("error while calling LongGreet rpc: %v", err)
	}
	for _, req := range requests {
		fmt.Println("sending request", req)
		err := stream.Send(req)
		if err != nil {
			log.Fatalf("error while sending message to stream: %v", err)
		}
		time.Sleep(1 * time.Second)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response from LongGreet: %v", err)
	}
	fmt.Println("LongGreet Response:", res.Result)

}
