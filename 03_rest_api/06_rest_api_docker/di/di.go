package di

import (
	"base/rest_api/internal/handlers"
	"base/rest_api/internal/repository"
	"base/rest_api/internal/server"
	"base/rest_api/internal/usecase"
)

func NewContainer() {
	r := repository.NewRepository()

	u := usecase.NewTodosUsecase(*r)

	h := handlers.NewTodosHandlers(*u)

	s := server.NewServer(*h)

	s.Run()
}
