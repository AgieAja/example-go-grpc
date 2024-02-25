package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/AgieAja/example-go-grpc/examplepb/proto"
	pbProduct "github.com/AgieAja/example-go-grpc/productpb/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreetingServiceServer
	pbProduct.UnimplementedProductServiceServer
}

func (s *server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	name := in.GetName()
	message := fmt.Sprintf("Hello, %s!", name)
	return &pb.GreetResponse{Message: message}, nil
}

func (s *server) GetProduct(ctx context.Context, in *pbProduct.Empty) (*pbProduct.GetProductsResponse, error) {
	response := &pbProduct.GetProductsResponse{
		Products: []*pbProduct.GetProductResponse{
			{
				Id:   1,
				Name: "Product 1",
			},
			{
				Id:   2,
				Name: "Product 2",
			},
			{
				Id:   3,
				Name: "Product 3",
			},
		},
	}
	return response, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreetingServiceServer(s, &server{})
	pbProduct.RegisterProductServiceServer(s, &server{})
	log.Println("server started")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
