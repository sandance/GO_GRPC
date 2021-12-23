package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/sandance/GRPC-GO-COURSE/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct{}


func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Received Sum RPC: %v", req)
	firstNumber := req.FirstNumber
	secondNumber := req.SecondNumber

	sum := firstNumber + secondNumber
	result := &calculatorpb.SumResponse{
		SumResult: sum,
	}
	return result, nil
}
func main() {
	fmt.Println("Calculator Server")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatal("Failed to Listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)

	}
}
