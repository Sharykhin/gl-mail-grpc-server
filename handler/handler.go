package handler

import (
	"fmt"
	"net"

	"os"

	"log"

	"github.com/Sharykhin/gl-mail-grpc"
	"github.com/Sharykhin/gl-mail-grpc-server/database"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	crt string
	key string
	env string
)

// ListenAndServe creates grps server and start listening income connections
func ListenAndServe(serverSource string) error {
	lis, err := net.Listen("tcp", serverSource)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	var s *grpc.Server
	if env == "prod" {
		// Create the TLS credentials
		cred, err := credentials.NewServerTLSFromFile(crt, key)
		if err != nil {
			return fmt.Errorf("could not load TLS keys: %v", err)
		}

		// Creates a new gRPC server
		s = grpc.NewServer(grpc.Creds(cred))
	} else {
		// Creates a new gRPC server
		s = grpc.NewServer()
	}

	storage := database.Storage
	api.RegisterFailMailServer(s, &server{storage: storage})
	fmt.Printf("Start listening on %s. Env: %s \n", serverSource, env)
	return s.Serve(lis)
}

func init() {
	env = os.Getenv("APP_ENV")
	if env == "prod" {
		crt = os.Getenv("KEY_SERVER_CRT")
		if crt == "" {
			log.Fatal("Env variable KEY_SERVER_CRT is not specified")
		}
		key = os.Getenv("KEY_SERVER_KEY")
		if key == "" {
			log.Fatal("Env variable KEY_SERVER_KEY is not specified")
		}
	}
}
