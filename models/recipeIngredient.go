package models

// RecipeIngredientTableName table name
const RecipeIngredientTableName = "recipe_ingredient"

// RecipeIngredient relation table with amount attribute
type RecipeIngredient struct {
	RecipeID     uint    `json:"recipe_id"`
	IngredientID uint    `json:"ingredient_id"`
	Amount       float32 `json:"amount"`
	Measure      string  `json:"measure"`
}

// TableName get name of the table
func (*RecipeIngredient) TableName() string {
	return RecipeIngredientTableName
}
