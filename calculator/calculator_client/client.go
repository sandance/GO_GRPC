package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/sandance/GRPC-GO-COURSE/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello, I am client")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	c := calculatorpb.NewCalculatorServiceClient(conn)
	//fmt.Printf("Created client: %f", c)
	//doUnary(c)
	doStreamPrimary(c)

}

func doStreamPrimary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Staring streaming RPC for Prime...")
	req := &calculatorpb.PrimeNumberDecompositionRequest{
		Number: 120,
	}
	stream, err := c.PrimerNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling PrimeDecomposition RPC: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Something happen", err)
		}
		fmt.Println(res.GetPrimeFactor())
	}
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do unary RPC....")
	req := &calculatorpb.SumRequest{
		FirstNumber:  5,
		SecondNumber: 40,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Sum RPC: %v", err)
	}
	log.Printf("Response from calculator: %v", res.SumResult)

}
