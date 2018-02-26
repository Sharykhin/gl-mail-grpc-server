package controller

import (
	"context"

	"github.com/Sharykhin/gl-mail-grpc"
	"github.com/Sharykhin/gl-mail-grpc-server/database"
	"github.com/Sharykhin/gl-mail-grpc-server/entity"
)

// FailedMailCtrl is a reference to a private struct that provides all necessary controller methods
var FailedMailCtrl failedMailController

type FailedMailProvider interface {
	Create(ctx context.Context, fmr *api.FailMailRequest) (*entity.FailMail, error)
	GetList(ctx context.Context, limit, offset int64) ([]api.FailMailResponse, error)
}

type failedMailController struct {
	storage FailedMailProvider
}

func (c failedMailController) Create(ctx context.Context, fmr *api.FailMailRequest) (*entity.FailMail, error) {
	return c.storage.Create(ctx, fmr)
}

func (c failedMailController) GetList(ctx context.Context, limit, offset int64) ([]api.FailMailResponse, error) {
	return c.storage.GetList(ctx, limit, offset)
}

func init() {
	FailedMailCtrl.storage = database.Storage
}
