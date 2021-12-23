package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/sandance/GRPC-GO-COURSE/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello, I am client")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)
	//fmt.Printf("Created client: %f", c)
	//doUnary(c)
	doServerStreaming(c)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a server Streaming RPC...")
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Nazmul",
			LastName:  "Islam",
		},
	}
	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Printf("error while calling GreetManyTimes RPC: %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}
		log.Printf("Resppnse from GreetManyTimes: %v", msg.GetResult())

	}
}

/*

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do unary RPC....")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Nazmul",
			LastName:  "Islam",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Printf("Error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from greet: %v", res.Result)

}*/
