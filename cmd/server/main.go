package server

import (
	"errors"
	"fmt"
	"io/ioutil"

	"crypto/tls"
	"crypto/x509"

	pb "github.com/harlov/grpc_echo_service/pkg/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

const (
	crtPath = "configs/keys/out/grpc_echo.crt"
	keyPath = "configs/keys/out/grpc_echo.key"
	caPath  = "configs/keys/out/grpc_echo.crt"
)

type server struct{}

func (s *server) Say(ctx context.Context, in *pb.EchoRequest) (*pb.EchoReply, error) {
	return &pb.EchoReply{Msg: "You say: " + in.Msg, Mid: in.Mid}, nil
}

func NewServer() (s *grpc.Server, err error) {
	certificate, err := tls.LoadX509KeyPair(crtPath, keyPath)
	if err != nil {
		return nil, fmt.Errorf("could not load server key pair: %s", err)
	}
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(caPath)
	if err != nil {
		return nil, fmt.Errorf("couldn read CA certificate: %s", err)
	}

	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		return nil, errors.New("failed to append ca certs")
	}
	creds := credentials.NewTLS(&tls.Config{
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{certificate},
		ClientCAs:    certPool,
	})
	s = grpc.NewServer(grpc.Creds(creds))
	pb.RegisterEchoServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	return s, nil
}
