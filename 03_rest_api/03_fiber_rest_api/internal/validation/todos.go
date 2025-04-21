package validation

import (
	"encoding/json"
	"fiber/rest_api/internal/models"

	"github.com/gofiber/fiber/v3"
)

func ValidateBodyCreateTodo(c fiber.Ctx) (*models.CreateTodo, error) {
	body := models.CreateTodo{}
	if err := json.Unmarshal(c.Body(), &body); err != nil {
		return nil, err
	}

	return &body, nil
}

func ValidateArgsUpdateTodo(c fiber.Ctx) (string, *models.UpdateTodo, error) {
	body := models.UpdateTodo{}

	id := c.Params("id")
	if err := json.Unmarshal(c.Body(), &body); err != nil {
		return id, nil, err
	}

	return id, &body, nil
}

func ValidateParamDeleteTodo(c fiber.Ctx) string {
	id := c.Params("id")

	return id
}
