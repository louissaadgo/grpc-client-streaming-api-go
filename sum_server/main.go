package main

import (
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"projects/grpcclientstreamingapi/sumpb/sumpb"
	"strconv"
)

const port = ":50051"

type server struct {
	sumpb.UnimplementedSumServiceServer
}

func (s *server) Sum(stream sumpb.SumService_SumServer) error {
	var sum int32 = 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Println("Sending response")
			stream.SendAndClose(&sumpb.SumResponse{Sum: "The sum is " + strconv.Itoa(int(sum))})
			break
		}
		if err != nil {
			log.Fatalf("Failed to receive dat: %v", err)
		}
		log.Println("Received: ", req.GetNum())
		sum += req.GetNum()
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen to port: %v", err)
	}
	log.Println("Listening to port")
	s := grpc.NewServer()
	sumpb.RegisterSumServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
