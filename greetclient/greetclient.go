package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/axamon/greetgrpc/greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/encoding/gzip"
)

var version = "0.1.5"

func main() {

	ctx := context.Background()
	var addr = flag.String("addr", "localhost:50051", "Address of the gRPC server")
	var name = flag.String("name", "Gringo", "Firstname to greet")
	flag.Parse()

	fmt.Printf("Client version %s\n", version)

	certFile := "server.crt"
	creds, err := credentials.NewClientTLSFromFile(certFile, "")
	if err != nil {
		log.Fatal(err)
	}

	cc, err := grpc.Dial(
		*addr,
		// grpc.WithCompressor(grpc.NewGZIPCompressor()),
		// grpc.WithInsecure(),
		grpc.WithTransportCredentials(creds),
	)
	if err != nil {
		log.Fatalf("could not connect to server: %v", err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)

	doUnary(ctx, c, name)
}

func doUnary(ctx context.Context, c greetpb.GreetServiceClient, name *string) {

	// fmt.Printf("Client creato: %f\n", c)
	req := &greetpb.GreetRequest{Greeting: &greetpb.Greeting{Firstname: *name}}

	resp, err := c.Greet(ctx, req, grpc.UseCompressor(gzip.Name))
	if err != nil {
		log.Fatalf("response %v in error: %v\n", resp, err)
	}
	log.Printf("response from greet: %v", resp.Result)
}
