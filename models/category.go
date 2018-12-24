package models

import (
	"github.com/jinzhu/gorm"
	u "recify/utils"
)

const CategoryTableName = "category"

type Category struct {
	gorm.Model
	Name string `json:"name"`
}

func (*Category) TableName() string {
	return CategoryTableName
}

func (category *Category) Create() map[string] interface{} {

	//if resp, ok := recipe.Validate(); !ok {
	//	return resp
	//}

	GetDB().Create(category)

	resp := u.Message(true, "success")
	resp[CategoryTableName] = category
	return resp
}