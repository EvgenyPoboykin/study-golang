package handlers

import (
	"encoding/json"

	"example/api/internal/usecase"
	"example/api/internal/validation"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	usecase usecase.Usecase
}

func NewTodosHandlers(uc usecase.Usecase) *Handlers {
	return &Handlers{
		usecase: uc,
	}
}

func (h *Handlers) GetTodos() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		todos, err := json.Marshal(h.usecase.GetTodos())
		if err != nil {
			ctx.JSON(400, err)
		}

		ctx.JSON(200, todos)
	}
}

func (h *Handlers) CreateTodo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		args := validation.ValidateBodyCreateTodo(ctx)

		todo, err := json.Marshal(h.usecase.CreateTodo(args))
		if err != nil {
			ctx.JSON(400, err)
		}

		ctx.JSON(200, todo)
	}
}

func (h *Handlers) UpdateTodo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, body := validation.ValidateArgsUpdateTodo(ctx)

		todo, err := json.Marshal(h.usecase.UpdateTodo(id, body))
		if err != nil {
			ctx.JSON(400, err)
		}

		ctx.JSON(200, todo)
	}
}

func (h *Handlers) DeleteTodo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := validation.ValidateParamDeleteTodo(ctx)

		todo, err := json.Marshal(h.usecase.DeleteTodo(id))
		if err != nil {
			ctx.JSON(400, err)
		}

		ctx.JSON(200, todo)
	}
}
