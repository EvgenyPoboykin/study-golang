package validation

import (
	"echo/rest_api/internal/models"

	"github.com/labstack/echo/v4"
)

func ValidateBodyCreateTodo(c echo.Context) (*models.CreateTodo, error) {
	body := models.CreateTodo{}
	if err := c.Bind(&body); err != nil {
		return nil, err
	}

	return &body, nil
}

func ValidateArgsUpdateTodo(c echo.Context) (string, *models.UpdateTodo, error) {
	body := models.UpdateTodo{}

	id := c.Param("id")
	if err := c.Bind(&body); err != nil {
		return id, nil, err
	}

	return id, &body, nil
}

func ValidateParamDeleteTodo(c echo.Context) string {
	id := c.Param("id")

	return id
}
