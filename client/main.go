package main

import (
	"log"
	"os"

	pb "github.com/rephus/grpc-gateway-example/template"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:10000"
	defaultName = "rephus"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	r, err := c.SendGet(context.Background(), &pb.TemplateRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Client GET: %s", r.Message)

	r, err = c.SendPost(context.Background(), &pb.TemplateRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Client POST: %s", r.Message)
}
