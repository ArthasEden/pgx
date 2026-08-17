package repo

import (
	"context"
	"fmt"
)

func (r *repo) Get(
	ctx context.Context,
) error {
	var (
		id   int
		name string
		q    = `select id, name from users`
	)

	rows, err := r.pool.Query(ctx, q)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&id, &name); err != nil {
			return err
		}
		fmt.Printf("Get user: %d = %s\n", id, name)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}
