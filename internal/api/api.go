package api

import (
	"context"
	"pgxPractice/internal/service"

	"github.com/google/uuid"
)

type Service interface {
	Create(ctx context.Context, in service.UserInput) error
	Delete(ctx context.Context, id uuid.UUID) error
	Get(ctx context.Context) ([]service.User, error)
}

type api struct {
	service Service
}

func NewAPI(service Service) *api {
	return &api{
		service: service,
	}
}
