package api

import (
	"net/http"

	"github.com/google/uuid"
)

func (a *api) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	uuID, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "can't parse id", http.StatusInternalServerError)
		return
	}

	if err := a.service.Delete(r.Context(), uuID); err != nil {
		http.Error(w, "can't delete user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
