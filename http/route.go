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

	// RESTy routes for "articles" resource
	r.Route("/", func(r chi.Router) {
		r.Post("/signIn", signIn)

		r.Route("/users", func(r chi.Router) {
			r.Use(auth)
			r.Post("/", createUser)
			r.Get("/", getAllUsers)

			// Subrouters:
			r.Route("/{userID}", func(r chi.Router) {
				r.Get("/", getUser)
				r.Patch("/", updateUser)
				r.Delete("/", deleteUser)
			})

		})

	})
	http.ListenAndServe(":"+os.Getenv("PORT"), r)
}
