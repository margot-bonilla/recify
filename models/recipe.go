package models

import (
	"github.com/jinzhu/gorm"
	"log"
	u "recify/utils"
)

const RecipeTableName = "recipe"

type Recipe struct {
	gorm.Model
	Title string `json:"title"`
	Steps []Step `gorm:"foreignKey:RecipeId"`
	Description string `json:"description"`
	UserId uint `json:"user_id"`
	RecipeIngredient *RecipeIngredient `json:"recipe_ingredient"`
}

func (*Recipe) TableName() string {
	return RecipeTableName
}

/*
 This struct function validate the required parameters sent through the http request body
returns message and true if the requirement is met
*/
func (recipe *Recipe) Validate() (map[string] interface{}, bool) {

	if recipe.UserId <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}



func (recipe *Recipe) Create() map[string] interface{} {

	if resp, ok := recipe.Validate(); !ok {
		return resp
	}

	GetDB().Create(recipe)

	resp := u.Message(true, "success")
	resp[RecipeTableName] = recipe
	return resp
}

func GetUserRecipes(id uint) []*Recipe {
	recipes := make([]*Recipe, 0)
	err := GetDB().Table(RecipeTableName).Where("user_id = ?", id).Find(&recipes).Error

	if err != nil {
		log.Println(err)
		return nil
	}

	return recipes
}

func GetRecipes() []*Recipe {
	recipes := make([]*Recipe, 0)
	err := GetDB().Table(RecipeTableName).Find(&recipes).Error

	if err != nil {
		log.Println(err)
		return nil
	}

	return recipes
}