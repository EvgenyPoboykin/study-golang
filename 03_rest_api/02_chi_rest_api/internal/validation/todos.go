package validation

import (
	"chi/rest_api/internal/models"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func ValidateBodyCreateTodo(r *http.Request) (*models.CreateTodo, error) {
	body := models.CreateTodo{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return nil, err
	}

	return &body, nil
}

func ValidateArgsUpdateTodo(r *http.Request) (string, *models.UpdateTodo, error) {
	body := models.UpdateTodo{}

	id := chi.URLParam(r, "id")
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return id, nil, err
	}

	return id, &body, nil
}

func ValidateParamDeleteTodo(r *http.Request) string {
	id := chi.URLParam(r, "id")

	return id
}
