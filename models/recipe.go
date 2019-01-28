package models

import "time"

const RecipeTableName = "recipe"

type Recipe struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// TableName get name of the table
func (*Recipe) TableName() string {
	return RecipeTableName
}
