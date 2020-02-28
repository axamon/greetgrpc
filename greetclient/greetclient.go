package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/axamon/greet/greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding/gzip"
)

var version = "0.1.2"

func main() {

	ctx := context.Background()
	var addr = flag.String("addr", "localhost:50051", "Address of the gRPC server")
	var name = flag.String("name", "Gringo", "Firstname to greet")
	flag.Parse()

	fmt.Printf("Client version %s\n", version)

	cc, err := grpc.Dial(
		*addr,
		grpc.WithCompressor(grpc.NewGZIPCompressor()),
		grpc.WithInsecure(),
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
		log.Fatalf("response in error: %v\n", resp)
	}
	log.Printf("response from greet: %v", resp.Result)
}
