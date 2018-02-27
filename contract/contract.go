package contract

import (
	"context"

	"github.com/Sharykhin/gl-mail-grpc"
	"github.com/Sharykhin/gl-mail-grpc-server/entity"
)

type FailedMailProvider interface {
	Create(ctx context.Context, fmr *api.FailMailRequest) (*entity.FailMail, error)
	GetList(ctx context.Context, limit, offset int64) ([]entity.FailMail, error)
}
