package handler

import (
	"log"
	"net"

	"fmt"

	"github.com/Sharykhin/gl-mail-grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	port = ":50051"
)

var (
	crt = "server.crt"
	key = "server.key"
)

func ListenAndServe() {
	fmt.Printf("Start listening on tcp %s \n", port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create the TLS credentials
	cred, err := credentials.NewServerTLSFromFile(crt, key)
	if err != nil {
		log.Fatalf("could not load TLS keys: %v", err)
	}

	// Creates a new gRPC server
	s := grpc.NewServer(grpc.Creds(cred))
	api.RegisterFailMailServer(s, &server{})
	log.Fatal(s.Serve(lis))
}
