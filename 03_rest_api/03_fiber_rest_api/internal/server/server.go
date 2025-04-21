package server

import (
	"fiber/rest_api/internal/handlers"

	"github.com/gofiber/fiber/v3"
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
	router := fiber.New()

	r := router.Group("/api/v3")

	r.Get("/todos", s.handlers.GetTodos())
	r.Put("/todos", s.handlers.CreateTodo())
	r.Patch("/todos/:id", s.handlers.UpdateTodo())
	r.Delete("/todos/:id", s.handlers.DeleteTodo())

	router.Listen(":4321")
}
