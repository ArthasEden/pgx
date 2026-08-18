package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (s *service) Delete(
	ctx context.Context,
	id uuid.UUID,
) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("Delete user %w", err)
	}

	return nil
}
