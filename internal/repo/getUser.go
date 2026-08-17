package repo

import (
	"context"
	"fmt"
)

func (r *repo) GetUser(
	ctx context.Context,
	id int,
) error {
	var (
		userID int
		name   string
		q      = `select id, name from users where id = $1`
	)

	if err := r.pool.QueryRow(ctx, q, id).Scan(&userID, &name); err != nil {
		return err
	}
	fmt.Printf("Get user: %d = %s\n", id, name)

	return nil
}
