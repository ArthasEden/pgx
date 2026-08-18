package repo

import (
	"context"
	"pgxPractice/internal/service"

	"github.com/jackc/pgx/v5"
)

func (r *repo) Create(
	ctx context.Context,
	u service.User,
) error {
	query := `
	insert into users (
	id,
	name,
	age,
	phone_number,
	is_active,
	created_at
	)
	values($1, $2, $3, $4, $5, $6)`

	tag, err := r.pool.Exec(
		ctx, query,
		u.ID,
		u.Name,
		u.Age,
		u.PhoneNumber,
		u.IsActive,
		u.CreatedAt,
	)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}
