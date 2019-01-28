package models

import (
	"time"
)

const GroupTableName string = "group"

type Group struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// TableName get name of the table
func (*Group) TableName() string {
	return GroupTableName
}
