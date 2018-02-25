package handler

import (
	"log"
	"net"

	"fmt"

	"os"

	"github.com/Sharykhin/gl-mail-grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var serverSource string

var (
	crt = "server.crt"
	key = "server.key"
)

func init() {
	serverSource = os.Getenv("SERVER_SOURCE")
}

func ListenAndServe() {
	fmt.Printf("Start listening on %s \n", serverSource)
	lis, err := net.Listen("tcp", serverSource)
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
