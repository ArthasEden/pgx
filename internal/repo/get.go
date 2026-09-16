package repo

import (
	"context"
	"errors"
	"fmt"
	"pgxPractice/internal/service"

	"github.com/jackc/pgx/v5/pgconn"
)

func (r *repo) Get(
	ctx context.Context,
) ([]service.User, error) {
	users := make([]service.User, 0)

	rows, err := r.pool.Query(ctx,
		`select 
		id, 
		name, 
		age,
		phone_number,
		is_active,
		created_at,
		balance from users`)
	if err != nil {
		var pgxErr *pgconn.PgError

		if errors.As(err, &pgxErr) {
			fmt.Println(pgxErr.Code)
		}

		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		u := service.User{}
		if err := rows.Scan(
			&u.ID,
			&u.Name,
			&u.Age,
			&u.PhoneNumber,
			&u.IsActive,
			&u.CreatedAt,
			&u.Balance); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
