package v2

type AmountIngredient struct {
	RecipeId     uint    `json:"recipe_id"`
	IngredientId uint    `json:"ingredient_id"`
	Amount       float32 `json:"amount"`
	Measure      string  `json:"measure"`
}
