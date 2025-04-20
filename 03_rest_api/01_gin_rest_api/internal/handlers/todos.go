package handlers

import (
	"gin/rest_api/internal/usecase"
	"gin/rest_api/internal/validation"
	"net/http"

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
		ctx.JSON(http.StatusOK, h.usecase.GetTodos())
	}
}

func (h *Handlers) CreateTodo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		args := validation.ValidateBodyCreateTodo(ctx)

		todo := h.usecase.CreateTodo(args)

		ctx.JSON(http.StatusOK, todo)
	}
}

func (h *Handlers) UpdateTodo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, body := validation.ValidateArgsUpdateTodo(ctx)

		todo := h.usecase.UpdateTodo(id, body)

		ctx.JSON(http.StatusOK, todo)
	}
}

func (h *Handlers) DeleteTodo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := validation.ValidateParamDeleteTodo(ctx)

		todo := h.usecase.DeleteTodo(id)

		ctx.JSON(http.StatusOK, todo)
	}
}
