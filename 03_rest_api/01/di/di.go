package di

import (
	"example/api/internal/handlers"
	"example/api/internal/repository"
	"example/api/internal/server"
	"example/api/internal/usecase"
)

func NewContainer() {
	r := repository.NewRepository()

	u := usecase.NewTodosUsecase(*r)

	h := handlers.NewTodosHandlers(*u)

	s := server.NewServer(*h)

	s.Run()
}
