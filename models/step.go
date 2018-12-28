package models

import (
	"github.com/jinzhu/gorm"
	u "recify/utils"
)

const StepTableName = "step"

type Step struct {
	gorm.Model
	Position    int    `json:"position"`
	Description string `json:"description"`
	RecipeId    uint   `json:"recipe_id"`
}

func (*Step) TableName() string {
	return StepTableName
}

func (step *Step) Create() map[string]interface{} {

	//if resp, ok := recipe.Validate(); !ok {
	//	return resp
	//}

	GetDB().Create(step)

	resp := u.Message(true, "success")
	resp[StepTableName] = step
	return resp
}
