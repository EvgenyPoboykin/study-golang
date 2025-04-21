package handlers

import (
	"echo/rest_api/internal/usecase"
	"echo/rest_api/internal/validation"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handlers struct {
	usecase usecase.Usecase
}

func NewTodosHandlers(uc usecase.Usecase) *Handlers {
	return &Handlers{
		usecase: uc,
	}
}

func (h *Handlers) GetTodos() echo.HandlerFunc {
	return func(c echo.Context) error {
		todos := h.usecase.GetTodos()

		return c.JSON(http.StatusOK, todos)
	}
}

func (h *Handlers) CreateTodo() echo.HandlerFunc {
	return func(c echo.Context) error {
		args, err := validation.ValidateBodyCreateTodo(c)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		todo := h.usecase.CreateTodo(args)

		return c.JSON(http.StatusOK, todo)
	}
}

func (h *Handlers) UpdateTodo() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, body, err := validation.ValidateArgsUpdateTodo(c)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		todo, err := h.usecase.UpdateTodo(id, body)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusOK, todo)
	}
}

func (h *Handlers) DeleteTodo() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := validation.ValidateParamDeleteTodo(c)

		todo, err := h.usecase.DeleteTodo(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusOK, todo)
	}
}
