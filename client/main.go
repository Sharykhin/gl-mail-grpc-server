package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"encoding/json"
	"fmt"

	"github.com/Sharykhin/gl-mail-grpc"
	"google.golang.org/grpc/credentials"
)

const (
	address = "127.0.0.1:50051"
)

var cert = "server.crt"

func main() {

	cred, err := credentials.NewClientTLSFromFile(cert, "")
	if err != nil {
		log.Fatalf("could not load tls cert: %s", err)
	}

	// Set up a connection to the gRPC server.
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(cred))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	// Creates a new CustomerClient
	client := api.NewFailMailClient(conn)

	f := &api.FailMailRequest{
		Action:  "test action",
		Payload: json.RawMessage(`{"to": "chapal@inbox.ru"}`),
		Reason:  "test reason",
	}

	resp, err := client.CreateFailMail(context.Background(), f)
	if err != nil {
		log.Fatalf("Could not create a failed message: %v", err)
	}
	fmt.Println(resp)
}
