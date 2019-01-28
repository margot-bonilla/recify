package models

// GroupRecipeTableName table name
const GroupRecipeTableName = "group_recipe"

// GroupRecipe relation table
type GroupRecipe struct {
	GroupID  uint `json:"group_id"`
	RecipeID uint `json:"recipe_id"`
}

// TableName get name of the table
func (*GroupRecipe) TableName() string {
	return GroupRecipeTableName
}
