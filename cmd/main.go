package main

import (
	server "github.com/harlov/grpc_echo_service/cmd/server"
	"log"
	"net"
)

const (
	port = ":50051"
)

func main() {
	s := server.NewServer()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
