package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	queryCreateTable = `create table if not exists users(
    						id serial primary key,
    						name text not null
						)`
	queryInsert = `insert into users (id, name)
					values($1, $2)`
	queryUpdate = `update users
					set name = $1
					where id = $2`
)

func main() {

	ctx := context.Background()
	strConn := "postgres://postgres:pass@localhost:5432/postgres"

	// Создаём пул соедиенений
	pool, err := CreatePool(ctx, strConn)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	defer pool.Close()

	// Создаём таблицу
	_, err = pool.Exec(ctx, queryCreateTable)
	if err != nil {
		fmt.Println("Can't create table:", err)
		return
	}

	// Добавляем пользователя
	if err := AddUser(ctx, pool, 123, "Arthas"); err != nil {
		fmt.Println("Can't add user:", err)
	}

	// Обновляем пользователя
	if err := UpdateUser(ctx, pool, 123, "Sergey"); err != nil {
		fmt.Println("Can't update user:", err)
	}
}

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
		return nil, fmt.Errorf("Can't ping connection: %w", err)

	}
	fmt.Println("Successfully connected to DB!")
	return pool, err
}

func AddUser(
	ctx context.Context,
	pool *pgxpool.Pool,
	id int,
	name string,
) error {
	tag, err := pool.Exec(ctx, queryInsert, id, name)
	if err != nil {
		return err
	}
	fmt.Println("Rows affected:", tag.RowsAffected())
	return nil
}

func UpdateUser(
	ctx context.Context,
	pool *pgxpool.Pool,
	id int,
	name string,
) error {
	tag, err := pool.Exec(ctx, queryUpdate, name, id)
	if err != nil {
		return err
	}
	fmt.Println("Rows affected:", tag.RowsAffected())
	return nil
}
