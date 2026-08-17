package internal

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreatePool(
	ctx context.Context,
	strConn string,
) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, strConn)
	if err != nil {
		return nil, fmt.Errorf("Can't connect: %w", err)
	}

	// Пингуем соединение
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("Can't ping connection: %w", err)

	}
	fmt.Println("Successfully connected to DB!")
	return pool, nil
}

func Exec(
	ctx context.Context,
	pool *pgxpool.Pool,
	query string,
	id int,
	name string,
) error {
	tag, err := pool.Exec(ctx, query, id, name)
	if err != nil {
		return err
	}
	fmt.Println("Rows affected:", tag.RowsAffected())
	return nil
}

func Query(
	ctx context.Context,
	pool *pgxpool.Pool,
	query string,
) error {
	rows, err := pool.Query(ctx, query)
	if err != nil {
		return err
	}
	defer rows.Close()

	var id int
	var name string
	for rows.Next() {
		if err := rows.Scan(&id, &name); err != nil {
			return err
		}
		fmt.Printf("id:%d, name:%s\n", id, name)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}

func QueryRow(
	ctx context.Context,
	pool *pgxpool.Pool,
	id int,
) error {
	var (
		userID int
		name   string
		q      = `select id, name from users where id = $1`
	)
	if err := pool.QueryRow(ctx, q, id).Scan(&userID, &name); err != nil {
		return err
	}
	fmt.Println("Get user:", userID, name)

	return nil
}
