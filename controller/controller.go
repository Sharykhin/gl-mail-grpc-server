package controller

import (
	"context"

	"github.com/Sharykhin/gl-mail-grpc"
	"github.com/Sharykhin/gl-mail-grpc-server/database"
)

// FailedMailCtrl is a reference to a private struct that provides all necessary controller methods
var FailedMailCtrl failedMailController

type FailedMailProvider interface {
	Create(ctx context.Context, fmr *api.FailMailRequest) (*api.FailMailResponse, error)
}

type failedMailController struct {
	storage FailedMailProvider
}

func (c failedMailController) Create(ctx context.Context, fmr *api.FailMailRequest) (*api.FailMailResponse, error) {
	return c.storage.Create(ctx, fmr)
}

func init() {
	FailedMailCtrl.storage = database.Storage
}
