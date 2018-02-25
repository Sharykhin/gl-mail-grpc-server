package handler

import (
	"context"

	"fmt"

	"time"

	"github.com/Sharykhin/gl-mail-grpc"
)

type server struct {
}

func (s server) CreateFailMail(ctx context.Context, fmr *api.FailMailRequest) (*api.FailMailResponse, error) {
	fmt.Println("CreateFailMail method is called", fmr)
	return &api.FailMailResponse{
		ID:        19,
		Action:    fmr.Action,
		Payload:   fmr.Payload,
		Reason:    fmr.Reason,
		CreatedAt: time.Now().String(),
		DeletedAt: "",
	}, nil
}
