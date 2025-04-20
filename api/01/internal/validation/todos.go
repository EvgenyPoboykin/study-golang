package validation

import (
	"example/api/internal/models"

	"github.com/gin-gonic/gin"
)

func ValidateBodyCreateTodo(c *gin.Context) *models.CreateTodo {
	body := models.CreateTodo{}
	c.ShouldBindJSON(&body)

	return &body
}

func ValidateArgsUpdateTodo(c *gin.Context) (string, *models.UpdateTodo) {
	id := c.Param("id")

	body := models.UpdateTodo{}
	c.ShouldBindJSON(&body)

	return id, &body
}

func ValidateParamDeleteTodo(c *gin.Context) string {
	id := c.Param("id")

	return id
}
