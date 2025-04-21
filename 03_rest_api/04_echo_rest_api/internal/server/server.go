package server

import (
	"echo/rest_api/internal/handlers"

	"github.com/labstack/echo/v4"
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
	router := echo.New()

	r := router.Group("/api/v4")

	r.GET("/todos", s.handlers.GetTodos())
	r.PUT("/todos", s.handlers.CreateTodo())
	r.PATCH("/todos/:id", s.handlers.UpdateTodo())
	r.DELETE("/todos/:id", s.handlers.DeleteTodo())

	router.Start(":4321")
}
