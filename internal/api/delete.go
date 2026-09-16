package api

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (a *api) Delete(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	id := r.PathValue("id")
	uuID, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "can't parse id", http.StatusInternalServerError)
		return
	}

	if err := a.service.Delete(ctx, uuID); err != nil {
		http.Error(w, "can't delete user", getHTTPStatus(err))
		return
	}

	w.WriteHeader(http.StatusOK)
}
