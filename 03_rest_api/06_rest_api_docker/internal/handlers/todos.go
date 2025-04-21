package handlers

import (
	"base/rest_api/internal/usecase"
	"base/rest_api/internal/validation"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlersInterface interface {
	GetTodos() gin.HandlerFunc
	CreateTodo() gin.HandlerFunc
	UpdateTodo() gin.HandlerFunc
	DeleteTodo() gin.HandlerFunc
}

type Handlers struct {
	usecase usecase.UsecaseInterface
}

func NewTodosHandlers(uc usecase.UsecaseInterface) *Handlers {
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

		todo, err := h.usecase.UpdateTodo(id, body)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
		}

		ctx.JSON(http.StatusOK, todo)
	}
}

func (h *Handlers) DeleteTodo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := validation.ValidateParamDeleteTodo(ctx)

		todo, err := h.usecase.DeleteTodo(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
		}

		ctx.JSON(http.StatusOK, todo)
	}
}
