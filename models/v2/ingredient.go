package v2

import "time"

type Ingredient struct {
	Id uint
	Name string
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}