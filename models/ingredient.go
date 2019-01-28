package models

import (
	u "recify/utils"
	"time"

	"github.com/jinzhu/gorm"
)

const IngredientTableName = "ingredient"

type Ingredient struct {
	gorm.Model
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Measure   string    `json:"measure"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// @TODO
	//RecipeIngredient *RecipeIngredient `json:"recipe_ingredient"`
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
