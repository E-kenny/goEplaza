package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func App() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Put("/createUser", createUser)
	r.Get("/user/{userID}", getUser)
	http.ListenAndServe(":3000", r)
}
