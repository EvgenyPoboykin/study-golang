package di

import (
	"chi/rest_api/internal/handlers"
	"chi/rest_api/internal/repository"
	"chi/rest_api/internal/server"
	"chi/rest_api/internal/usecase"
)

func NewContainer() {
	r := repository.NewRepository()

	u := usecase.NewTodosUsecase(*r)

	h := handlers.NewTodosHandlers(*u)

	s := server.NewServer(*h)

	s.Run()
}
