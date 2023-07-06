package main

import (
	"context"
	pb "grpcLectures/greet/proto"
	"log"
)

func doGreet(client pb.GreetServiceClient) {
	log.Println("doGreet was invoked")

	res, err := client.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Sasha",
	})

	if err != nil {
		log.Fatalf("Couldn't greet: %v\n", err)
	}

	log.Printf("Greeting: %s\n", res.Result)
}
