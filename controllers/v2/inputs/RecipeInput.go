package inputs

type RecipeInput struct {
	Title string
	Description string
	Categories []uint
	Ingredients []IngredientInput
	Steps []StepInput
}

type IngredientInput struct {
	Id uint `json:"id"`
	Amount float32 `json:"amount"`
	Measure string `json:"measure"`

}

type StepInput struct {
	position uint
	description string
}