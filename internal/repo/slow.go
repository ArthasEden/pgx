package repo

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func (r *repo) Slow(ctx context.Context) error {
	query := `select pg_sleep(5)`

	tag, err := r.pool.Exec(ctx, query)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}
