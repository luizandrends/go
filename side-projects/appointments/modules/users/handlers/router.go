package handlers

import (
	"github.com/go-chi/chi/v5"
)

type Password string

func CreateUserRouter() chi.Router {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Post("/create", handleCreateUser())
	})

	return r
}
