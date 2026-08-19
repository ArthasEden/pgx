package api

import "net/http"

func (a *api) NewHandler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /create", a.Create)
	mux.HandleFunc("DELETE /delete/{id}", a.Delete)
	mux.HandleFunc("GET /get", a.Get)
	mux.HandleFunc("POST /transfer", a.Transfer)

	return mux
}
