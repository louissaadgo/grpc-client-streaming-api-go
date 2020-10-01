package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"projects/grpcclientstreamingapi/sumpb/sumpb"
)

var nums = []int32{5, 10, 13}

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()
	c := sumpb.NewSumServiceClient(conn)
	stream, err := c.Sum(context.Background())
	if err != nil {
		log.Fatalf("Something went wrong: %v", err)
	}
	for i := 0; i < 3; i++ {
		stream.Send(&sumpb.SumRequest{Num: nums[i]})
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalln("Something went wrong ", err)
	}
	log.Printf("The Sum Is: %v", res.GetSum())
}
