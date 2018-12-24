package models

import (
	u "recify/utils"
)

const RecipeIngredientTableName = "recipe_ingredient"

type RecipeIngredient struct {
	Recipe *Recipe `json:"recipe"`
	RecipeID uint `json:"recipe_id"`
	Ingredient Ingredient `json:"ingredient"`
	IngredientId uint `json:"ingredient_id"`
	Amount int `json:"amount"`
}

func (*RecipeIngredient) Table() string {
	return RecipeIngredientTableName
}

func (recipeIngredient *RecipeIngredient) Create() map[string] interface{} {

	//if resp, ok := recipeIngredient.Validate(); !ok {
	//	return resp
	//}

	GetDB().Create(recipeIngredient)

	resp := u.Message(true, "success")
	resp[RecipeIngredientTableName] = recipeIngredient
	return resp
}