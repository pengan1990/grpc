package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "github.com/binlake/grpc_test/protos"
)

var (
	address     = "localhost:9000"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewIUserServiceClient(conn)

	// Contact the server and print out its response.
	r, err := c.Get(context.Background(), &pb.UserRequest{Id:1})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s %s", r.Id, r.Name)
}
