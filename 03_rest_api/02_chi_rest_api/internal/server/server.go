package server

import (
	"chi/rest_api/internal/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	handlers handlers.Handlers
}

func NewServer(handlers handlers.Handlers) *Server {
	server := &Server{}

	server.handlers = handlers

	return server
}

func (s *Server) Run() {
	router := chi.NewRouter()

	router.Route("/api/v2", func(r chi.Router) {
		r.Get("/todos", s.handlers.GetTodos())
		r.Put("/todos", s.handlers.CreateTodo())
		r.Patch("/todos/{id}", s.handlers.UpdateTodo())
		r.Delete("/todos/{id}", s.handlers.DeleteTodo())
	})

	http.ListenAndServe("localhost:4321", router)
}
