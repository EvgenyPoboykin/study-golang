package usecase

import (
	"echo/rest_api/internal/models"
	"echo/rest_api/internal/repository"
)

type Usecase struct {
	repository repository.Repository
}

func NewTodosUsecase(repository repository.Repository) *Usecase {
	return &Usecase{
		repository,
	}
}

func (u *Usecase) GetTodos() *[]models.Todo {
	return u.repository.Todos()
}

func (u *Usecase) CreateTodo(args *models.CreateTodo) *models.Todo {
	return u.repository.Create(args)
}

func (u *Usecase) UpdateTodo(id string, args *models.UpdateTodo) (*models.Todo, error) {
	return u.repository.Update(id, args)
}

func (u *Usecase) DeleteTodo(id string) (*models.Todo, error) {
	return u.repository.Delete(id)
}
