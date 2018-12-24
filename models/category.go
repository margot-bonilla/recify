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

func (category *Category) Validate() (map[string] interface{}, bool) {

	temp := &Category{}

	err := GetDB().Table(CategoryTableName).Where("name = ?", category.Name).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}
	if temp.Name != "" {
		return u.Message(false, "Category already in db."), false
	}

	return u.Message(false, "Requirement passed"), true
}


func (category *Category) Create() map[string] interface{} {

	if resp, ok := category.Validate(); !ok {
		return resp
	}

	GetDB().Create(category)

	resp := u.Message(true, "success")
	resp[CategoryTableName] = category
	return resp
}