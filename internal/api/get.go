package api

import (
	"encoding/json"
	"net/http"
)

func (a *api) Get(w http.ResponseWriter, r *http.Request) {
	users, err := a.service.Get(r.Context())
	if err != nil {
		http.Error(w, "can't get users", http.StatusInternalServerError)
		return
	}

	dtoResp := DTOServiceToApi(users)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(dtoResp); err != nil {
		return
	}
}
