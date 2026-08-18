package api

import (
	"context"
	"pgxPractice/internal/service"

	"github.com/google/uuid"
)

type Service interface {
	Create(ctx context.Context, input service.UserInput) error
	Delete(ctx context.Context, uuid uuid.UUID) error
	Get(ctx context.Context) ([]service.User, error)
	UpdateBalance(ctx context.Context, uuid uuid.UUID, amount int) error
}

type api struct {
	service Service
}

func NewAPI(service Service) *api {
	return &api{
		service: service,
	}
}
