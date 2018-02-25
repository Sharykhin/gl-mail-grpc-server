package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"encoding/json"
	"fmt"

	"github.com/Sharykhin/gl-mail-grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the gRPC server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
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
