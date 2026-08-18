package service

import (
	"context"
	"fmt"
)

func (s *service) Create(ctx context.Context, in UserInput) error {
	user := NewUser(in)

	if err := s.repo.Create(ctx, user); err != nil {
		return fmt.Errorf("Create user %w", err)
	}

	return nil
}
