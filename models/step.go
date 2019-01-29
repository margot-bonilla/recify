package models

const stepTableName = "step"

// Step model
type Step struct {
	RecipeID    uint   `json:"recipe_id"`
	Position    uint   `json:"position"`
	Description string `json:"description"`
}

// TableName get name of the table
func (*Step) TableName() string {
	return stepTableName
}
