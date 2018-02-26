package controller

import (
	"context"
	"testing"

	"encoding/json"
	"time"

	"github.com/Sharykhin/gl-mail-grpc"
	"github.com/Sharykhin/gl-mail-grpc-server/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockStorage struct {
	mock.Mock
}

func (m mockStorage) Create(ctx context.Context, fmr *api.FailMailRequest) (*entity.FailMail, error) {
	ret := m.Called(ctx, fmr)
	fm, err := ret.Get(0), ret.Get(1)
	if err != nil {
		return nil, err.(error)
	}

	return fm.(*entity.FailMail), nil
}

func (m mockStorage) GetList(ctx context.Context, limit, offset int64) ([]entity.FailMail, error) {
	ret := m.Called(ctx, limit, offset)
	fm, err := ret.Get(0), ret.Get(1)
	if err != nil {
		return nil, err.(error)
	}

	return fm.([]entity.FailMail), nil
}

func TestCreate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctx := context.Background()
		fmr := api.FailMailRequest{
			Action:  "test action",
			Payload: json.RawMessage(`{"to":"test@test.com"}`),
			Reason:  "test reason",
		}

		fm := entity.FailMail{
			ID:        21,
			Action:    "test action",
			Payload:   entity.Payload(`{"to":"test@test.com"}`),
			Reason:    "test reason",
			CreatedAt: entity.JSONTime(time.Now()),
		}

		m := new(mockStorage)
		m.On("Create", ctx, &fmr).Return(&fm, nil).Once()

		tt := &failedMailController{
			storage: m,
		}

		fml, err := tt.Create(ctx, &fmr)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		m.AssertExpectations(t)

		assert.NotNil(t, fml)
		assert.Equal(t, "test action", fml.Action)
		assert.Equal(t, "test reason", fml.Reason)
	})
}
