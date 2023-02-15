package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "example-go-grpc/examplepb/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreetingServiceServer
}

func (s *server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	name := in.GetName()
	message := fmt.Sprintf("Hello, %s!", name)
	return &pb.GreetResponse{Message: message}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreetingServiceServer(s, &server{})
	log.Println("server started")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
