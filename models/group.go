package models

import (
	u "recify/utils"
	"time"

	"github.com/jinzhu/gorm"
)

const GroupTableName string = "group"

type Group struct {
	gorm.Model
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (*Group) TableName() string {
	return GroupTableName
}

// Validate incoming user details...
func (group *Group) Validate() (map[string]interface{}, bool) {

	// @todo validate group

	return u.Message(false, "Requirement passed"), true
}

func (group *Group) Create() map[string]interface{} {

	if resp, ok := group.Validate(); !ok {
		return resp
	}

	GetDB().Create(group)

	response := u.Message(true, "User has been created")
	response[GroupTableName] = group

	return response
}
