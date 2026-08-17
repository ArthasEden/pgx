package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ctx     = context.Background()
	strConn = "postgres://postgres:pass@localhost:5432/postgres"

	queryCreateTable = `create table if not exists users(
    						id serial primary key,
    						name text not null
						)`
	queryInsert = `insert into users (id, name)
					values(21, 'Arthas')`
	queryUpdate = `update users
					set name = 'Sergey'
					where id = 21`
)

func main() {
	// Создаём пул соедиенений
	pool, err := pgxpool.New(ctx, strConn)
	if err != nil {
		fmt.Println("Can't connect:", err)
		return
	}
	defer pool.Close()

	// Пингуем соединение
	if err := pool.Ping(ctx); err != nil {
		fmt.Println("Can't ping connection:", err)
		return
	}
	fmt.Println("Successfully connected to DB!")

	// Используем pool.QueryRow - метод, когда ожидаем от зароса ответ в виде строк
	var res int
	if err := pool.QueryRow(ctx, "SELECT 1").Scan(&res); err != nil {
		fmt.Println("Can't do query:", err)
		return
	}
	fmt.Println("Query result:", res)

	// Создаём таблицу
	_, err = pool.Exec(ctx, queryCreateTable)
	if err != nil {
		fmt.Println("Can't create table:", err)
		return
	}

	// Добавляем пользователя
	tag, err := pool.Exec(ctx, queryInsert)
	if err != nil {
		fmt.Println("Can't add user:", err)
		return
	}
	fmt.Println("Rows affected:", tag.RowsAffected())

	// Обновляем пользователя
	tag, err = pool.Exec(ctx, queryUpdate)
	if err != nil {
		fmt.Println("Can't update user:", err)
		return
	}
	fmt.Println("Rows affected:", tag.RowsAffected())
}
