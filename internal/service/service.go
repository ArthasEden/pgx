package service

import (
	"context"

	"github.com/google/uuid"
)

type Repo interface {
	Create(ctx context.Context, u User) error
	Delete(ctx context.Context, id uuid.UUID) error
	Get(ctx context.Context) ([]User, error)
}

type service struct {
	repo Repo
}

func NewService(repo Repo) *service {
	return &service{
		repo: repo,
	}
}
