package models

import (
	"log"
	u "recify/utils"

	"github.com/jinzhu/gorm"
)

const RecipeTableName = "recipe"

type Recipe struct {
	gorm.Model
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`

	Rating   float64 `json:"rating"`
	Calories float64 `json:"calories"`
	Fat      int     `json:"fat"`
	Date     string  `json:"date"`
	Protein  float64 `json:"protein"`
	Sodium   int     `json:"sodium"`

	//Ingredients      []Ingredient      `json:"ingredients"`
	//Categories       []Category        `json:"categories"`
	//RecipeIngredient *RecipeIngredient `json:"recipe_ingredient"`
}

func (*Recipe) TableName() string {
	return RecipeTableName
}

func (recipe *Recipe) Validate() (map[string]interface{}, bool) {

	if recipe.Title == "" {
		return u.Message(false, "The title cannot be empty"), false
	}

	return u.Message(false, "Requirement passed"), true
}

func (recipe *Recipe) Create() map[string]interface{} {

	if resp, ok := recipe.Validate(); !ok {
		return resp
	}

	GetDB().Create(recipe)

	resp := u.Message(true, "success")
	resp[RecipeTableName] = recipe
	return resp
}

// @TODO get group recipes

func GetRecipes() []*Recipe {
	recipes := make([]*Recipe, 0)
	err := GetDB().Table(RecipeTableName).Find(&recipes).Error

	if err != nil {
		log.Println(err)
		return nil
	}

	return recipes
}
