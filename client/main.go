package main

import (
	"log"

	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"flag"
	"io"

	"encoding/json"

	"github.com/Sharykhin/gl-mail-grpc"
	"google.golang.org/grpc/credentials"
	"os"
)

const (
	address = "localhost:50051"
)

var cert = "../server.crt"

func main() {

	action := flag.String("action", "create", "type of action")
	flag.Parse()

	// Set up a connection to the gRPC server.
	conn := createConn()
	defer conn.Close() // nolint: errcheck

	// Creates a new CustomerClient
	client := api.NewFailMailClient(conn)
	switch *action {
	case "create":
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
	case "list":
		filter := &api.FailMailFilter{Limit: 5, Offset: 0}
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		stream, err := client.GetFailMails(ctx, filter)
		if err != nil {
			log.Fatalf("Could not stream fail mails: %v", err)
		}
		for {
			m, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("%v.GetFailMails(_) = _, %v", client, err)
			}
			log.Printf("fail mail: %v", m)
		}
	case "count":
		ctx := context.Background()
		res, err := client.CountFailMails(ctx, &api.Empty{})
		if err != nil {
			log.Fatalf("Could not count: %v", err)
		}
		log.Println(res)
	}
}

func createConn() *grpc.ClientConn {
	env := os.Getenv("APP_ENV")
	if env == "prod" {
		cred, err := credentials.NewClientTLSFromFile(cert, "")

		if err != nil {
			log.Fatalf("Could not load tls cert: %s", err)
		}

		// Set up a connection to the gRPC server.
		conn, err := grpc.Dial(address, grpc.WithTransportCredentials(cred))
		if err != nil {
			log.Fatalf("Could not connet to a grpc server: %v", err)
		}
		return conn
	} else {
		// Set up a connection to the gRPC server.
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Could not connet to a grpc server: %v", err)
		}
		return conn
	}
}
