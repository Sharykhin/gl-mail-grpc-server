package handler

import (
	"fmt"
	"net"

	"github.com/Sharykhin/gl-mail-grpc"
	"github.com/Sharykhin/gl-mail-grpc-server/database"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	crt = "server.crt"
	key = "server.key"
)

// ListenAndServe creates grps server and start listening income connections
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
	storage := database.Storage
	api.RegisterFailMailServer(s, &server{storage: storage})
	return s.Serve(lis)
}
