package models

// UserLikesRecipeTableName table name
const UserLikesRecipeTableName = "user_likes_recipe"

// UserLikesRecipe relation table
type UserLikesRecipe struct {
	UserID   uint `json:"user_id"`
	RecipeID uint `json:"recipe_id"`
}

// TableName get name of the table
func (*UserLikesRecipe) TableName() string {
	return UserLikesRecipeTableName
}
