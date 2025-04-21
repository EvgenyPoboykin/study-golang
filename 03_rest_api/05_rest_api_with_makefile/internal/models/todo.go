package models

type Todo struct {
	Id          string `json:"id" db:"id"`
	IsActive    bool   `json:"isActive" db:"is_active"`
	Description string `json:"description" db:"description"`
}

type CreateTodo struct {
	Description string `json:"description" db:"description"`
}

type UpdateTodo struct {
	IsActive    bool   `json:"isActive" db:"is_active"`
	Description string `json:"description" db:"description"`
}
