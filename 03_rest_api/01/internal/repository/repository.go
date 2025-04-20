package repository

import (
	"example/api/internal/models"

	"github.com/google/uuid"
)

type Repository struct {
	todos *[]models.Todo
}

func NewRepository() *Repository {
	todos := &[]models.Todo{}

	return &Repository{
		todos,
	}
}

func (r *Repository) Todos() *[]models.Todo {
	return r.todos
}

func (r *Repository) Create(args *models.CreateTodo) *models.Todo {
	todo := &models.Todo{
		Id:          uuid.New().String(),
		Description: args.Description,
		IsActive:    true,
	}

	*r.todos = append(*r.todos, *todo)

	return todo
}

func (r *Repository) Update(id string, args *models.UpdateTodo) *models.Todo {
	todo := &models.Todo{
		Id:          id,
		Description: args.Description,
		IsActive:    args.IsActive,
	}

	for _, v := range *r.todos {
		if v.Id == id {
			v.Description = args.Description
			v.IsActive = args.IsActive
		}
	}
	return todo
}

func (r *Repository) Delete(id string) *models.Todo {
	todo := &models.Todo{}

	for _, v := range *r.todos {
		if v.Id == id {
			todo.Id = v.Id
			todo.Description = v.Description
			todo.IsActive = v.IsActive
		}
	}
	return todo
}
