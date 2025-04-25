package usecase

import (
	"base/rest_api/internal/models"
)

type repository interface {
	Todos() *[]models.Todo
	Create(args *models.CreateTodo) *models.Todo
	Update(id string, args *models.UpdateTodo) (*models.Todo, error)
	Delete(id string) (*models.Todo, error)
}

type Usecase struct {
	repository repository
}

func NewTodosUsecase(repository repository) *Usecase {
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
