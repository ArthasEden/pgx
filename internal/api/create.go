package api

import (
	"encoding/json"
	"net/http"
	"pgxPractice/internal/service"
)

func (a *api) Create(w http.ResponseWriter, r *http.Request) {
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

	if err := a.service.Create(r.Context(), in); err != nil {
		http.Error(w, "can't create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
