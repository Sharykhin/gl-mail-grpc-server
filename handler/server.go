package handler

import (
	"context"
	"fmt"
	"log"

	"strings"

	"encoding/json"

	"github.com/Sharykhin/gl-mail-grpc"
	"github.com/Sharykhin/gl-mail-grpc-server/controller"
)

type server struct {
}

// TODO: think how to mock this func
func validate(fmr *api.FailMailRequest) error {
	if strings.Trim(fmr.Action, "") == "" {
		return fmt.Errorf("action is required")
	}

	if fmr.Payload == nil {
		return fmt.Errorf("payload is required")
	}

	var stuff struct{}
	err := json.Unmarshal(fmr.Payload, &stuff)
	if err != nil {
		return fmt.Errorf("payload must be a valid json")
	}

	if strings.Trim(fmr.Reason, " ") == "" {
		return fmt.Errorf("reason is required")
	}

	return nil
}

func (s server) CreateFailMail(ctx context.Context, fmr *api.FailMailRequest) (*api.FailMailResponse, error) {
	log.Printf("CreateFailMail is called with request: %v \n", fmr)
	if err := validate(fmr); err != nil {
		return nil, err
	}
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
