package models

import (
	u "recify/utils"
	"time"

	"github.com/jinzhu/gorm"
)

const StepTableName = "step"

type Step struct {
	gorm.Model
	Position    uint      `json:"position"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	// @TODO create constraint
	RecipeID uint `json:"recipe_id"`
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
