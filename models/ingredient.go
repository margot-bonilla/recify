package models

import (
	"time"
)

const IngredientTableName = "ingredient"

type Ingredient struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName get name of the table
func (*Ingredient) TableName() string {
	return IngredientTableName
}
