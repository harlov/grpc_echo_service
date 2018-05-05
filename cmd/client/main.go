package client

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	pb "github.com/harlov/grpc_echo_service/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	basePath = "configs/keys/out/"
	caPath   = "configs/keys/out/grpc_echo.crt"
	address  = "localhost:50051"
)

func RunEchoClien(clientName string, msg string) error {

	certificate, err := tls.LoadX509KeyPair(basePath+clientName+".crt",
		basePath+clientName+".key")

	if err != nil {
		return fmt.Errorf("could not load client key pair: %s", err)
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(caPath)

	if err != nil {
		return fmt.Errorf("could not read ca certificate: %s", err)
	}

	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		return errors.New("failed to append ca certs")
	}

	creds := credentials.NewTLS(&tls.Config{
		ServerName:   "grpc_echo", // NOTE: this is required!
		Certificates: []tls.Certificate{certificate},
		RootCAs:      certPool,
	})

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewEchoClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Say(ctx, &pb.EchoRequest{Msg: msg, Mid: 10})
	if err != nil {
		log.Fatalf("could not echo: %v", err)
	}
	log.Printf("Echo reponse: %s", r.Msg)
	return nil
}
