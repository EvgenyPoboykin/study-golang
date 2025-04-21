package di

import (
	"makefile/rest_api/internal/handlers"
	"makefile/rest_api/internal/repository"
	"makefile/rest_api/internal/server"
	"makefile/rest_api/internal/usecase"
)

func NewContainer() {
	r := repository.NewRepository()

	u := usecase.NewTodosUsecase(*r)

	h := handlers.NewTodosHandlers(*u)

	s := server.NewServer(*h)

	s.Run()
}
