package main

import (
	"context"
	"errors"
	"fmt"
	"pgxPractice/internal"

	"github.com/jackc/pgx/v5"
)

func main() {
	//names := []string{"Bob", "Jack", "Alice"}
	ctx := context.Background()
	strConn := "postgres://postgres:pass@localhost:5432/postgres"

	// Создаём пул соедиенений
	pool, err := internal.CreatePool(ctx, strConn)
	if err != nil {
		fmt.Println("Create connection:", err)
		return
	}
	defer pool.Close()

	// Создаём таблицу
	_, err = pool.Exec(ctx, internal.QCreateTable)
	if err != nil {
		fmt.Println("Create table:", err)
		return
	}

	// Добавляем пользователей
	// for i, v := range names {
	// 	if err := internal.Exec(ctx, pool, internal.QAdd, i+1, v); err != nil {
	// 		fmt.Println("Add user:", err)
	//		return
	// 	}
	// }

	// Получаем пользователей
	// if err := internal.Query(ctx, pool, internal.QGetAll); err != nil {
	// 		fmt.Println("Get all users:", err)
	//		return
	// }

	// Получаем пользователя
	if err := internal.QueryRow(ctx, pool, 2); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			fmt.Println("No rows:", err)
			return
		}
		fmt.Println("Error:", err)
		return
	}
}
