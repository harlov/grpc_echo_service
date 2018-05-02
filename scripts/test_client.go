package main

import (
	"log"
	"os"
	"time"

	pb "github.com/harlov/grpc_echo_service/pkg/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewEchoClient(conn)

	// Contact the server and print out its response.

	name := os.Args[1]

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Say(ctx, &pb.EchoRequest{Msg: name, Mid: 10})
	if err != nil {
		log.Fatalf("could not echo: %v", err)
	}
	log.Printf("Echo reponse: %s", r.Msg)
}
