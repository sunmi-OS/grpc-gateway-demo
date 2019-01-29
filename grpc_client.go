package main

import (
	"os"
	"log"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "grpc-gateway-demo/gateway"
)

const (
	address     = "localhost:9192"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGatewayClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	go func() {
		r, err := c.Echo(context.Background(), &pb.StringMessage{Value: name})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r.Value)
	}()
	time.Sleep(10 * time.Second)
}
