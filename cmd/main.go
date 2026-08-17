package main

import (
	"context"
	"errors"
	"fmt"
	"pgxPractice/internal/repo"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var errUserNotFound = "user not found"

func main() {
	ctx := context.Background()
	strConn := "postgres://postgres:pass@localhost:5432/postgres"

	// Создаём пул соедиенений
	pool, err := pgxpool.New(ctx, strConn)
	if err != nil {
		fmt.Println("Can't connect:", err)
		return
	}

	// Пингуем соединение
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		fmt.Println("Can't ping connection:", err)
		return
	}
	defer pool.Close()

	// Создаём репозиторий
	repo := repo.NewRepo(pool)

	// Создаём пользователя Joe
	if err := repo.Create(ctx, 4, "Joe"); err != nil {
		fmt.Println("Create user:", err)
		return
	}
	fmt.Println()

	// Получаем пользователя
	if err := repo.GetUser(ctx, 4); err != nil {
		fmt.Println("Get user:", err)
		return
	}
	fmt.Println()

	// Обновляем пользователя: меняем имя на Mike
	if err := repo.Update(ctx, 4, "Mike"); err != nil {
		fmt.Println("Update user:", err)
		return
	}
	fmt.Println()

	// Получаем пользователя: ожидаем обновленное имя
	if err := repo.GetUser(ctx, 4); err != nil {
		fmt.Println("Get user:", err)
		return
	}
	fmt.Println()

	// Получаем всех пользователей: ожидаем четверых
	if err := repo.Get(ctx); err != nil {
		fmt.Println("Get users:", err)
		return
	}
	fmt.Println()

	// Удаляем пользователя
	if err := repo.Delete(ctx, 4); err != nil {
		fmt.Println("Delete user:", err)
		return
	}
	fmt.Println()

	// Получаем пользователя: ожидаем ошибку
	if err := repo.GetUser(ctx, 4); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			fmt.Println("Get user:", errUserNotFound)
		} else {
			fmt.Println("Get user:", err)
			return
		}
	}
	fmt.Println()

	// Получаем всех пользователей: ожидаем троих
	if err := repo.Get(ctx); err != nil {
		fmt.Println("Get users:", err)
		return
	}
}
