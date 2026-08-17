package repo

import (
	"context"
	"fmt"
)

func (r *repo) Create(
	ctx context.Context,
	id int,
	name string,
) error {
	var q = `insert into users (id, name)
				values($1, $2)`

	_, err := r.pool.Exec(ctx, q, id, name)
	if err != nil {
		return err
	}

	fmt.Printf("Create user: %d = %s\n", id, name)
	return nil
}
