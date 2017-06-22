package main

import (
	"log"
	"net"

	pb "github.com/rephus/grpc-gateway-example/template"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":10000"
)

type server struct{}

func (s *server) SendGet(ctx context.Context, in *pb.TemplateRequest) (*pb.TemplateResponse, error) {
	return &pb.TemplateResponse{Message: "Received GET method " + in.Name}, nil
}

func (s *server) SendPost(ctx context.Context, in *pb.TemplateRequest) (*pb.TemplateResponse, error) {
	return &pb.TemplateResponse{Message: "Received POST method " + in.Name}, nil
}
func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
