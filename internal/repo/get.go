package repo

import (
	"context"
	"pgxPractice/internal/service"
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
		created_at from users`)
	if err != nil {
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
			&u.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
