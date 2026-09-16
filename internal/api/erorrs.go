package api

import (
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func getHTTPStatus(err error) int {
	if errors.Is(err, pgx.ErrNoRows) {
		return http.StatusNotFound
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505":
			return http.StatusConflict
		case "23503":
			return http.StatusConflict
		case "23502", "23514":
			return http.StatusBadRequest
		}
	}

	return http.StatusInternalServerError
}
