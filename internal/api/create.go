package api

import (
	"context"
	"encoding/json"
	"net/http"
	"pgxPractice/internal/service"
	"time"
)

func (a *api) Create(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	dtoReq := DTOreq{}

	if err := json.NewDecoder(r.Body).Decode(&dtoReq); err != nil {
		http.Error(w, "can't decode data", http.StatusBadRequest)
		return
	}

	in := service.UserInput{
		Name:        dtoReq.Name,
		Age:         dtoReq.Age,
		PhoneNumber: dtoReq.PhoneNumber,
	}

	if err := a.service.Create(ctx, in); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
