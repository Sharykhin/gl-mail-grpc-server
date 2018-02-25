package handler

import (
	"log"
	"net"

	"github.com/Sharykhin/gl-mail-grpc"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func ListenAndServe() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Creates a new gRPC server
	s := grpc.NewServer()
	api.RegisterFailMailServer(s, &server{})
	s.Serve(lis)
}
