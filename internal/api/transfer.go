package api

import (
	"context"
	"encoding/json"
	"net/http"
	"pgxPractice/internal/service"
	"time"
)

func (a *api) Transfer(w http.ResponseWriter, r *http.Request) {
	req := DTOreqTransfer{}
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "can't parse request", http.StatusInternalServerError)
		return
	}

	in := service.InputTransfer{
		To:     req.To,
		From:   req.From,
		Amount: req.Amount,
	}

	if err := a.service.Transfer(ctx, in); err != nil {
		http.Error(w, "can't update balance", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
