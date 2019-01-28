package models

// UserGroupTableName table name
const UserGroupTableName = "user_group"

// UserGroup relation table
type UserGroup struct {
	UserID  uint `json:"user_id"`
	GroupID uint `json:"group_id"`
}

// TableName get name of the table
func (*UserGroup) TableName() string {
	return UserGroupTableName
}
