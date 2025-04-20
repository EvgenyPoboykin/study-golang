package repository

import (
	"chi/rest_api/internal/models"

	"github.com/google/uuid"
)

type Repository struct {
	todos *[]models.Todo
}

func NewRepository() *Repository {
	todos := &[]models.Todo{
		{
			Id:          "637272b9-ff59-4b46-90fb-087a3b01fc21",
			Description: "Create Example gin RestAPI",
			IsActive:    true,
		},
	}

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
	todos := &[]models.Todo{}

	todo := &models.Todo{
		Id:          id,
		Description: args.Description,
		IsActive:    args.IsActive,
	}

	for _, v := range *r.todos {
		if v.Id == id {
			*todos = append(*todos, *todo)
		} else {
			*todos = append(*todos, v)
		}
	}

	*r.todos = *todos

	return todo
}

func (r *Repository) Delete(id string) *models.Todo {
	todos := &[]models.Todo{}

	todo := &models.Todo{}

	for _, v := range *r.todos {
		if v.Id != id {
			*todos = append(*todos, v)
		} else {
			todo.Description = v.Description
			todo.IsActive = v.IsActive
			todo.Id = v.Id
		}
	}

	*r.todos = *todos

	return todo
}
