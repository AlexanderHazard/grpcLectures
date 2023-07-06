package main

import (
	"google.golang.org/grpc"
	pb "grpcLectures/greet/proto"
	"log"
	"net"
)

var addr = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
	pb.SumServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	server := grpc.NewServer()
	pb.RegisterSumServiceServer(server, &Server{})
	if err = server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}

}
