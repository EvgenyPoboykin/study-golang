package server

import (
	"base/rest_api/internal/handlers"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router   *gin.Engine
	handlers handlers.HandlersInterface
}

func NewServer(handlers handlers.HandlersInterface) *Server {
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

	s.router.Run(":4321")
}
