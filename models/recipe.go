package models

import (
	"github.com/jinzhu/gorm"
	"log"
	u "recify/utils"
)

type Recipe struct {
	gorm.Model
	Title string `json:"title"`
	Steps string `json:"steps"`
	Description string `json:"description"`
	UserId uint `json:"user_id"` //The user that this contact belongs to
	//IngredientList []Ingredient `gorm:foreignKey:RecipeRefer`
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
	resp["recipe"] = recipe
	return resp
}

//func GetRecipe(id uint) *Recipe {
//	recipe := &Recipe{}
//	err := GetDB().Table("recipe").Where("id = ?", id).First(recipe).Error
//
//	if err != nil {
//		log.Println(err)
//		return nil
//	}
//
//	return recipe
//}

func GetRecipes(id uint) []*Recipe {
	recipes := make([]*Recipe, 0)
	err := GetDB().Table("recipe").Where("user_id = ?", id).Find(&recipes).Error

	if err != nil {
		log.Println(err)
		return nil
	}

	return recipes
}