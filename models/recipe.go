package models

import (
	"github.com/jinzhu/gorm"
	"log"
	u "recify/utils"
)

const RecipeTableName = "recipe"

type Recipe struct {
	gorm.Model
	Title            string            `json:"title"`
	Steps            []Step            `json:directions;gorm:"foreignKey:RecipeId"`
	Description      string            `json:"desc"`
	UserId           uint              `json:"user_id"`
	Ingredients      []Ingredient      `json:"ingredients"`
	Categories       []Category        `json:"categories"`
	RecipeIngredient *RecipeIngredient `json:"recipe_ingredient"`
	Rating           float64           `json:"rating"`
	Calories         float64           `json:"calories"`
	Fat              int               `json:"fat"`
	Date             string            `json:"date"`
	Protein          float64           `json:"protein"`
	Sodium           int               `json:"sodium"`
}

func (*Recipe) TableName() string {
	return RecipeTableName
}

/*
 This struct function validate the required parameters sent through the http request body
returns message and true if the requirement is met
*/
func (recipe *Recipe) Validate() (map[string]interface{}, bool) {

	// @TODO check if the user exists
	if recipe.UserId <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	temp := &Recipe{}

	err := GetDB().Table(CategoryTableName).Where("name = ?", recipe.Title).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}
	if temp.Title != "" {
		return u.Message(false, "Recipe already in db."), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
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
