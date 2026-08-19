package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (r *repo) Transfer(
	ctx context.Context,
	to uuid.UUID,
	from uuid.UUID,
	amount int,
) error {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	tag, err := tx.Exec(ctx,
		"UPDATE users SET balance = balance + $1 WHERE id = $2",
		amount, to)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	tag, err = tx.Exec(ctx,
		"UPDATE users SET balance = balance - $1 WHERE id = $2",
		amount, from)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}
