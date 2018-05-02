package server

import (
	"fmt"
	pb "github.com/harlov/grpc_echo_service/pkg/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (s *server) Say(ctx context.Context, in *pb.EchoRequest) (*pb.EchoReply, error) {
	return &pb.EchoReply{Msg: "You say: " + in.Msg, Mid: in.Mid}, nil
}

func NewServer() (s *grpc.Server) {
	s = grpc.NewServer()
	pb.RegisterEchoServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	return s
}
