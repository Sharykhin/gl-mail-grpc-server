package handler

import (
	"fmt"
	"net"

	"github.com/Sharykhin/gl-mail-grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	crt = "server.crt"
	key = "server.key"
)

func ListenAndServe(serverSource string) error {
	lis, err := net.Listen("tcp", serverSource)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	// Create the TLS credentials
	cred, err := credentials.NewServerTLSFromFile(crt, key)
	if err != nil {
		return fmt.Errorf("could not load TLS keys: %v", err)
	}

	// Creates a new gRPC server
	s := grpc.NewServer(grpc.Creds(cred))
	api.RegisterFailMailServer(s, &server{})
	return s.Serve(lis)
}
