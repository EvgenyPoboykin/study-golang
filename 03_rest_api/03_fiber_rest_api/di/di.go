package di

import (
	"fiber/rest_api/internal/handlers"
	"fiber/rest_api/internal/repository"
	"fiber/rest_api/internal/server"
	"fiber/rest_api/internal/usecase"
)

func NewContainer() {
	r := repository.NewRepository()

	u := usecase.NewTodosUsecase(*r)

	h := handlers.NewTodosHandlers(*u)

	s := server.NewServer(*h)

	s.Run()
}
