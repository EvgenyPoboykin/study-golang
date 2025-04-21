package usecase

import (
	"base/rest_api/internal/models"
	"base/rest_api/internal/repository"
)

type UsecaseInterface interface {
	GetTodos() *[]models.Todo
	CreateTodo(args *models.CreateTodo) *models.Todo
	UpdateTodo(id string, args *models.UpdateTodo) (*models.Todo, error)
	DeleteTodo(id string) (*models.Todo, error)
}

type Usecase struct {
	repository repository.RepositoryInterface
}

func NewTodosUsecase(repository repository.RepositoryInterface) *Usecase {
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
