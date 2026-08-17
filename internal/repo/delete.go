package repo

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (r *repo) Delete(
	ctx context.Context,
	id int,
) error {
	var q = `delete from users where id = $1`

	tag, err := r.pool.Exec(ctx, q, id)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	fmt.Printf("Delete user: %d \n", id)
	return nil
}
