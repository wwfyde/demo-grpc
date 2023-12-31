package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/wwfyde/demo-grpc/examples/helloworld/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to ")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()

	// Create a connection to server
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Client pb
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its Response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	start := time.Now()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	latency := time.Since(start)
	fmt.Println("total latency: ", latency)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

}
