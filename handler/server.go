package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/Sharykhin/gl-mail-grpc"
	"github.com/Sharykhin/gl-mail-grpc-server/contract"
)

type server struct {
	storage contract.FailedMailProvider
}

func (s server) CreateFailMail(ctx context.Context, fmr *api.FailMailRequest) (*api.FailMailResponse, error) {
	log.Printf("CreateFailMail is called with request: %v \n", fmr)
	if err := validate(fmr); err != nil {
		return nil, err
	}
	fm, err := s.storage.Create(ctx, fmr)
	if err != nil {
		log.Printf("could not create a new failed mail: %v", err)
		return nil, err
	}

	return &api.FailMailResponse{
		ID:        fm.ID,
		Action:    fm.Action,
		Payload:   fm.Payload,
		Reason:    fm.Reason,
		CreatedAt: fm.CreatedAt,
	}, err
}

func (s server) GetFailMails(filter *api.FailMailFilter, stream api.FailMail_GetFailMailsServer) error {
	log.Printf("GetFailMails is called with request: %s \n", filter)
	fml, err := s.storage.GetList(context.Background(), filter.Limit, filter.Offset)
	if err != nil {
		return fmt.Errorf("could not get list: %v", err)
	}

	for _, fm := range fml {
		m := api.FailMailEntity{
			ID:        fm.ID,
			Action:    fm.Action,
			Payload:   fm.Payload,
			Reason:    fm.Reason,
			CreatedAt: fm.CreatedAt,
		}
		if fm.DeletedAt != nil {
			m.DeletedAt = fm.DeletedAt.String()
		}
		if err := stream.Send(&m); err != nil {
			return fmt.Errorf("could not send entity into a stream: %v", err)
		}
	}
	return nil
}

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
