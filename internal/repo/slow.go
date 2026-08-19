package repo

import "context"

func (r *repo) Slow(ctx context.Context) error {
	_, err := r.pool.Exec(ctx, `SELECT pg_sleep(5)`)
	return err
}
