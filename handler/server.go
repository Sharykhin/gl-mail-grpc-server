package handler

import (
	"context"

	"log"

	"fmt"

	"github.com/Sharykhin/gl-mail-grpc"
	"github.com/Sharykhin/gl-mail-grpc-server/controller"
)

type server struct {
}

func (s server) CreateFailMail(ctx context.Context, fmr *api.FailMailRequest) (*api.FailMailResponse, error) {
	log.Printf("Create a new failed mail message: %s \n", fmr)
	fm, err := controller.FailedMailCtrl.Create(ctx, fmr)
	if err != nil {
		log.Printf("could not create a new failed mail: %v", err)
		return nil, err
	}
	return fm, err
}

func (s server) GetFailMails(filter *api.FailMailFilter, stream api.FailMail_GetFailMailsServer) error {
	log.Printf("Get list of failed messages: %s \n", filter)
	mm, err := controller.FailedMailCtrl.GetList(context.Background(), filter.Limit, filter.Offset)
	if err != nil {
		return fmt.Errorf("could not get list: %v", err)
	}
	for _, m := range mm {
		if err := stream.Send(&m); err != nil {
			return fmt.Errorf("could not send entity into a stream: %v", err)
		}
	}
	return nil
}
