package repo

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (r *repo) Update(
	ctx context.Context,
	id int,
	name string,
) error {
	var q = `update users
				set name = $2
				where id = $1`

	tag, err := r.pool.Exec(ctx, q, id, name)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	fmt.Printf("Update user: %d = %s\n", id, name)
	return nil
}
