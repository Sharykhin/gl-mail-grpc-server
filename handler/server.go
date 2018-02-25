package handler

import (
	"context"

	"github.com/Sharykhin/gl-mail-grpc"
)

type server struct {
}

func (s server) CreateFailMail(ctx context.Context, fmr *api.FailMailRequest) (*api.FailMailResponse, error) {
	return &api.FailMailResponse{
		ID:        19,
		Action:    "register",
		Payload:   []byte(`{}`),
		Reason:    "reason",
		CreatedAt: "sad",
		DeletedAt: "",
	}, nil
}
