package handlers

import (
	"chi/rest_api/internal/usecase"
	"chi/rest_api/internal/utils"
	"chi/rest_api/internal/validation"
	"net/http"
)

type Handlers struct {
	usecase usecase.Usecase
}

func NewTodosHandlers(uc usecase.Usecase) *Handlers {
	return &Handlers{
		usecase: uc,
	}
}

func (h *Handlers) GetTodos() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		todos := h.usecase.GetTodos()

		res, err := utils.Marshal(todos)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))

		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

func (h *Handlers) CreateTodo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		args, err := validation.ValidateBodyCreateTodo(r)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))

		}

		todo := h.usecase.CreateTodo(args)

		res, err := utils.Marshal(todo)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))

		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

func (h *Handlers) UpdateTodo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, body, err := validation.ValidateArgsUpdateTodo(r)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))

		}

		todo := h.usecase.UpdateTodo(id, body)

		res, err := utils.Marshal(todo)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))

		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

func (h *Handlers) DeleteTodo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := validation.ValidateParamDeleteTodo(r)

		todo := h.usecase.DeleteTodo(id)

		res, err := utils.Marshal(todo)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))

		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}
