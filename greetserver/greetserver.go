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

	"google.golang.org/grpc/credentials"
	_ "google.golang.org/grpc/encoding/gzip" // Install the gzip compressor
)

var version = "0.1.5"

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

	certFile, keyFile := "server.crt", "server.key"
	creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer(grpc.Creds(creds))
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
