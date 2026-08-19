package service

import (
	"context"
	"fmt"
)

func (s *service) Transfer(
	ctx context.Context,
	in InputTransfer,
) error {
	if err := s.repo.Transfer(ctx, in.To, in.From, in.Amount); err != nil {
		return fmt.Errorf("update balance: %w", err)
	}

	return nil
}
