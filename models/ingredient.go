package models

import (
	"github.com/jinzhu/gorm"
	u "recify/utils"
)

const IngredientTableName = "ingredient"

type Ingredient struct {
	gorm.Model
	Name             string            `json:"name"`
	Measure          string            `json:"measure"`
	RecipeIngredient *RecipeIngredient `json:"recipe_ingredient"`
}

func (*Ingredient) TableName() string {
	return IngredientTableName
}

func (ingredient *Ingredient) Create() map[string]interface{} {

	//if resp, ok := recipe.Validate(); !ok {
	//	return resp
	//}

	GetDB().Create(ingredient)

	resp := u.Message(true, "success")
	resp[IngredientTableName] = ingredient
	return resp
}
