package di

import (
	"gin/rest_api/internal/handlers"
	"gin/rest_api/internal/repository"
	"gin/rest_api/internal/server"
	"gin/rest_api/internal/usecase"
)

func NewContainer() {
	r := repository.NewRepository()

	u := usecase.NewTodosUsecase(*r)

	h := handlers.NewTodosHandlers(*u)

	s := server.NewServer(*h)

	s.Run()
}
