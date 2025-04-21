package server

import (
	"makefile/rest_api/internal/handlers"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router   *gin.Engine
	handlers handlers.Handlers
}

func NewServer(handlers handlers.Handlers) *Server {
	server := &Server{}
	router := gin.Default()

	server.router = router
	server.handlers = handlers

	return server
}

func (s *Server) Run() {
	baseUrl := s.router.Group("/api/v1/todos")

	baseUrl.GET("/", s.handlers.GetTodos())
	baseUrl.PUT("/", s.handlers.CreateTodo())
	baseUrl.PATCH("/:id", s.handlers.UpdateTodo())
	baseUrl.DELETE("/:id", s.handlers.DeleteTodo())

	s.router.Run("localhost:4321")
}
