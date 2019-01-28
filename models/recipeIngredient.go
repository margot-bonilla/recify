package models

import (
	u "recify/utils"
	"time"

	"github.com/jinzhu/gorm"
)

// RecipeIngredientTableName table name
const RecipeIngredientTableName = "recipe_ingredient"

// RecipeIngredient relation table with amount attribute
type RecipeIngredient struct {
	gorm.Model
	RecipeID     uint      `json:"recipe_id"`
	IngredientID uint      `json:"ingredient_id"`
	Amount       float32   `json:"amount"`
	Measure      string    `json:"measure"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	// @TODO
	//RecipeIngredient *RecipeIngredient `json:"recipe_ingredient"`
}

// TableName get name of the table
func (*RecipeIngredient) TableName() string {
	return RecipeIngredientTableName
}

// Create create row in the database from the object form
func (ingredient *RecipeIngredient) Create() map[string]interface{} {

	//if resp, ok := recipe.Validate(); !ok {
	//	return resp
	//}

	GetDB().Create(ingredient)

	resp := u.Message(true, "success")
	resp[IngredientTableName] = ingredient
	return resp
}
