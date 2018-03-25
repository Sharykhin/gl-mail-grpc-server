package main

import (
	"log"
	"os"

	"github.com/Sharykhin/gl-mail-grpc-server/handler"
)

func main() {
	serverSource := os.Getenv("SERVER_SOURCE")
	log.Fatal(handler.ListenAndServe(serverSource))
}
