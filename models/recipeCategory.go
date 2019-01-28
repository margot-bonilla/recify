package models

// RecipeCategoryTableName table name
const RecipeCategoryTableName = "recipe_category"

// RecipeCategory relation table with amount attribute
type RecipeCategory struct {
	RecipeID   uint `json:"recipe_id"`
	CategoryID uint `json:"category_id"`
}

// TableName get name of the table
func (*RecipeCategory) TableName() string {
	return RecipeCategoryTableName
}
