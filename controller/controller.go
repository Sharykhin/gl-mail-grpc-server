package controller

import (
	"context"

	"github.com/Sharykhin/gl-mail-grpc"
	"github.com/Sharykhin/gl-mail-grpc-server/contract"
	"github.com/Sharykhin/gl-mail-grpc-server/database"
	"github.com/Sharykhin/gl-mail-grpc-server/entity"
)

// FailMail is a reference to a private struct that provides all necessary controller methods
var FailMail failedMail

type failedMail struct {
	storage contract.FailedMailProvider
}

func (c failedMail) Create(ctx context.Context, fmr *api.FailMailRequest) (*entity.FailMail, error) {
	return c.storage.Create(ctx, fmr)
}

func (c failedMail) GetList(ctx context.Context, limit, offset int64) ([]entity.FailMail, error) {
	return c.storage.GetList(ctx, limit, offset)
}

func init() {
	FailMail.storage = database.Storage
}
