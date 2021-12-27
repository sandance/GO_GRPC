package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/sandance/GRPC-GO-COURSE/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) FindMaximum(stream calculatorpb.CalculatorService_FindMaximumServer) error {
	maximum := int32(0)
	fmt.Printf("Recived FindMaximum RPC")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		number := req.GetNumber()
		if number > maximum {
			maximum = number
			sendErr := stream.Send(&calculatorpb.FindMaximumResponse{
				Maximum: maximum,
			})

			if sendErr != nil {
				log.Fatalf("Error while sending data to client: %v", err)
				return err
			}
		}

	}

}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Received Sum RPC: %v\n", req)
	firstNumber := req.FirstNumber
	secondNumber := req.SecondNumber

	sum := firstNumber + secondNumber
	result := &calculatorpb.SumResponse{
		SumResult: sum,
	}
	return result, nil
}

func (*server) PrimerNumberDecomposition(req *calculatorpb.PrimeNumberDecompositionRequest, stream calculatorpb.CalculatorService_PrimerNumberDecompositionServer) error {
	number := req.GetNumber()
	divisor := int64(2)

	for number > 1 {
		if number%divisor == 0 {
			fmt.Printf("Prime Number: %v\n", number)
			res := &calculatorpb.PrimeNumberDecompositionResponse{
				PrimeFactor: divisor,
			}
			stream.Send(res)
			number = number / divisor
		} else {
			divisor++
			fmt.Printf("Divisor has increased: %v\n", divisor)
		}
	}
	return nil

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
