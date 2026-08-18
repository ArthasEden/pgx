package repo

import (
	"context"

	"github.com/google/uuid"

	"github.com/jackc/pgx/v5"
)

func (r *repo) Delete(
	ctx context.Context,
	id uuid.UUID,
) error {
	query := `delete from users where id = $1`

	tag, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}
