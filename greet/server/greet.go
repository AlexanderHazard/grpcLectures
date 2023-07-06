package main

import (
	"context"
	"fmt"
	pb "grpcLectures/greet/proto"
	"log"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function was invoved with %v\n", in)
	return &pb.GreetResponse{
		Result: fmt.Sprintf("Hello %s", in.FirstName),
	}, nil
}

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	sum := in.First + in.Second
	log.Printf("Sum invoved, with first=%v second=%v\n", in.First, in.Second)
	return &pb.SumResponse{
		Sum: sum,
	}, nil
}
