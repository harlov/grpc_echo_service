package main

import (
	"log"
	"net"
	"os"

	"github.com/harlov/grpc_echo_service/cmd/client"
	server "github.com/harlov/grpc_echo_service/cmd/server"
)

const (
	port = ":50051"
)

func main() {
	if os.Args[1] == "client" {
		client.RunEchoClien(os.Args[2], os.Args[3])
		return
	}

	s, err := server.NewServer()
	if err != nil {
		panic("error create server")
	}
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
