package http

import (
	"github.com/a-h/templ"
	"github.com/gorilla/mux"
	"github.com/reztheperson/ith/internal/api"
	"github.com/reztheperson/ith/internal/web/templates"
)

func setupAPIRoutes(r *mux.Router) {
	r.HandleFunc("/api/login", api.Login).Methods("GET")
}

func setupPageRoutes(r *mux.Router) {
	r.Handle("/", templ.Handler(templates.Home())).Methods("GET")
}
