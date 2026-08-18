package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os/signal"
	"pgxPractice/internal/api"
	"pgxPractice/internal/repo"
	"pgxPractice/internal/service"
	"syscall"
	"time"

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

	repo := repo.NewRepo(pool)
	svc := service.NewService(repo)
	api := api.NewAPI(svc)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: api.NewHandler(),
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	fmt.Println("Starting HTTP-Server")
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Println("Can't starting server", err)
			cancel()
		}
	}()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	fmt.Println("Stopping HTTP-Server...")
	if err := server.Shutdown(shutdownCtx); err != nil {
		fmt.Println("Can't shutdown server", err)
		return
	}
}
