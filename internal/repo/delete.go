package repo

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *repo) Delete(
	ctx context.Context,
	id uuid.UUID,
) error {
	query := `delete from users where id = $1`

	tag, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		var pgxErr *pgconn.PgError

		if errors.As(err, &pgxErr) {
			fmt.Println(pgxErr.Code)
		}

		return err
	}

	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}
