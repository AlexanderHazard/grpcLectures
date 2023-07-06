package main

import (
	"context"
	pb "grpcLectures/greet/proto"
	"log"
)

func doSum(client pb.SumServiceClient) {

	res, err := client.Sum(context.Background(), &pb.SumRequest{
		First:  1,
		Second: 2,
	})

	if err != nil {
		log.Fatal("Sum didn't calculate %v\n", err)
	}

	log.Printf("Result=%v\n", res)
}
