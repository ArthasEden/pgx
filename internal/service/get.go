package service

import (
	"context"
	"fmt"
)

func (s *service) Get(ctx context.Context) ([]User, error) {
	users, err := s.repo.Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("Get user %w", err)
	}
	return users, nil

}
