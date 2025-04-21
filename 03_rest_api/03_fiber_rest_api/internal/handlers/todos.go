package handlers

import (
	"fiber/rest_api/internal/usecase"
	"fiber/rest_api/internal/validation"

	"github.com/gofiber/fiber/v3"
)

type Handlers struct {
	usecase usecase.Usecase
}

func NewTodosHandlers(uc usecase.Usecase) *Handlers {
	return &Handlers{
		usecase: uc,
	}
}

func (h *Handlers) GetTodos() fiber.Handler {
	return func(c fiber.Ctx) error {
		todos := h.usecase.GetTodos()

		return c.JSON(todos)
	}
}

func (h *Handlers) CreateTodo() fiber.Handler {
	return func(c fiber.Ctx) error {
		args, err := validation.ValidateBodyCreateTodo(c)
		if err != nil {
			return c.JSON(err)
		}

		todo := h.usecase.CreateTodo(args)

		return c.JSON(todo)
	}
}

func (h *Handlers) UpdateTodo() fiber.Handler {
	return func(c fiber.Ctx) error {
		id, body, err := validation.ValidateArgsUpdateTodo(c)
		if err != nil {
			return c.JSON(err)
		}

		todo := h.usecase.UpdateTodo(id, body)

		return c.JSON(todo)
	}
}

func (h *Handlers) DeleteTodo() fiber.Handler {
	return func(c fiber.Ctx) error {
		id := validation.ValidateParamDeleteTodo(c)

		todo := h.usecase.DeleteTodo(id)

		return c.JSON(todo)
	}
}
