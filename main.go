package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()
	connStr := "postgres://postgres:pass@localhost:5432/postgres"

	pool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		fmt.Println("Can't connect:", err)
		return
	}
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		fmt.Println("Can't ping connection:", err)
		return
	}
	fmt.Println("Successfully connected to DB!")

	var res int
	if err := pool.QueryRow(ctx, "SELECT 1").Scan(&res); err != nil {
		fmt.Println("Can't do query:", err)
		return
	}

	fmt.Println("Query result:", res)
}
