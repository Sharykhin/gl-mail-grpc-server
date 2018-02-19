package main

import (
	"log"
	"net"
	"google.golang.org/grpc"
	api "github.com/Sharykhin/gl-mail-grpc"
	"context"
)

const (
	port = ":50051"
)

type Server struct {

}

func (s Server) Create(ctx context.Context, fmr api.FailMailRequest) (api.FailMailResponse, error) {
	return api.FailMailResponse{
		ID: 19,
		Action: "register",
		Payload: []byte(`{}`),
		Reason: "reason",
		CreatedAt: "sad",
		DeletedAt: "",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Creates a new gRPC server
	s := grpc.NewServer()
	s.Serve(lis)
}