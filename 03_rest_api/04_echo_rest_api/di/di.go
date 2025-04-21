package di

import (
	"echo/rest_api/internal/handlers"
	"echo/rest_api/internal/repository"
	"echo/rest_api/internal/server"
	"echo/rest_api/internal/usecase"
)

func NewContainer() {
	r := repository.NewRepository()

	u := usecase.NewTodosUsecase(*r)

	h := handlers.NewTodosHandlers(*u)

	s := server.NewServer(*h)

	s.Run()
}
