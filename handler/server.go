package handler

import (
	"context"

	"log"

	"github.com/Sharykhin/gl-mail-grpc"
	"github.com/Sharykhin/gl-mail-grpc-server/controller"
)

type server struct {
}

func (s server) CreateFailMail(ctx context.Context, fmr *api.FailMailRequest) (*api.FailMailResponse, error) {
	log.Println("Create a new failed mail message")
	fm, err := controller.FailedMailCtrl.Create(ctx, fmr)
	if err != nil {
		log.Printf("could not create a new failed mail: %v", err)
		return nil, err
	}
	return fm, err
}
