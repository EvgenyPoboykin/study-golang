package server

import (
	"example/api/internal/handlers"

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
	baseUrl := s.router.Group("/api/v1")

	baseUrl.GET("/todos", s.handlers.GetTodos())
	baseUrl.PUT("/todos", s.handlers.CreateTodo())
	baseUrl.PATCH("/todos/:id", s.handlers.UpdateTodo())
	baseUrl.DELETE("/todos/:id", s.handlers.DeleteTodo())
}
