package http

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func App() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/signIn",signIn)
	r.Post("/createUser", createUser)
	r.Get("/user/{userID}", getUser)
	r.Get("/users", getAllUsers)
	r.Patch("/user/{userID}", updateUser)
	r.Delete("/user/{userID}", deleteUser)
	http.ListenAndServe(":"+os.Getenv("PORT"), r)
}
