package router

import (
	"github.com/go-chi/chi/v5"
)

func NewRouter(fn func(r chi.Router)) *chi.Mux {
	router := chi.NewRouter()

	router.Route("/api/v2/todos", fn)

	return router
}
