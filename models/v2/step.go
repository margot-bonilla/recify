package v2

type Step struct {
	RecipeId    uint   `json:"recipe_id"`
	Position    uint   `json:"position"`
	Description string `json:"description"`
}
