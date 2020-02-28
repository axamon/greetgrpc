package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/axamon/greet/greetpb"
	"google.golang.org/grpc"

	_ "google.golang.org/grpc/encoding/gzip" // Install the gzip compressor
)

var version = "0.1.2"

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {

	log.Printf("Greet function invoked with request: %v\n", req)

	firstname := strings.Title(req.GetGreeting().GetFirstname())
	result := "Ciao " + firstname

	return &greetpb.GreetResponse{Result: result}, nil
}

func main() {

	var addr = flag.String("addr", "0.0.0.0:50051", "Server address")

	flag.Parse()

	fmt.Printf("Server version: %s\n", version)

	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
