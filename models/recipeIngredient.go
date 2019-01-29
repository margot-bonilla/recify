package models

import "time"

// RecipeIngredientTableName table name
const RecipeIngredientTableName = "recipe_ingredient"

// RecipeIngredient relation table with amount attribute
type RecipeIngredient struct {
	RecipeID  uint      `json:"recipe_id"`
	Name      string    `json:"name"`
	Amount    float32   `json:"amount"`
	Measure   string    `json:"measure"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName get name of the table
func (*RecipeIngredient) TableName() string {
	return RecipeIngredientTableName
}
