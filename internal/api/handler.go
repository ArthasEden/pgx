package api

import "net/http"

func (a *api) NewHandler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /create", a.Create)
	mux.HandleFunc("DELETE /delete/{id}", a.Delete)
	mux.HandleFunc("GET /get", a.Get)

	return mux
}
