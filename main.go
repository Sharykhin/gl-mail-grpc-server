package main

import (
	"github.com/Sharykhin/gl-mail-grpc-server/handler"
	"log"
	"fmt"
	"os"
)

func main() {
	serverSource := os.Getenv("SERVER_SOURCE")
	fmt.Printf("Start listening on %s \n", serverSource)
	log.Fatal(handler.ListenAndServe(serverSource))
}
