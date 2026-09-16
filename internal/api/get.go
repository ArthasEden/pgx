package api

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

func (a *api) Get(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	users, err := a.service.Get(ctx)
	if err != nil {
		http.Error(w, "can't get users", getHTTPStatus(err))
		return
	}

	dtoResp := DTOServiceToApi(users)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(dtoResp); err != nil {
		return
	}
}
