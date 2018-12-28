package inputs

import (
	"recify/models/v2"
)

type RecipeInput struct {
	Title       string
	Description string
	Categories  []uint
	Ingredients []AmountIngredientInput
	Steps       []StepInput
}

type AmountIngredientInput struct {
	Id      uint    `json:"id"`
	Amount  float32 `json:"amount"`
	Measure string  `json:"measure"`
}

type StepInput struct {
	Position    uint
	Description string
}

type RecipeOutput struct {
	Recipe            *v2.Recipe
	Categories        []*v2.Category
	AmountIngredients []*v2.AmountIngredient
	Steps             []*v2.Step
}

func (input *RecipeInput) Input2Models() (*RecipeOutput, error) {

	output := &RecipeOutput{}

	// recipe
	recipe := &v2.Recipe{}
	recipe.Title = input.Title
	recipe.Description = input.Description

	output.Recipe = recipe

	// categories
	// todo move this checking to service
	categories := v2.GetCategoriesByIds(input.Categories)

	output.Categories = categories

	// ingredients
	ingredients := make([]*v2.AmountIngredient, 0)
	ingredientsInput := input.Ingredients

	for i := 0; i < len(ingredientsInput); i++ {
		ing := new(v2.AmountIngredient)
		ingInput := ingredientsInput[i]

		ing.Amount = ingInput.Amount
		ing.Measure = ingInput.Measure
		ing.IngredientId = ingInput.Id

		ingredients = append(ingredients, ing)
	}
	output.AmountIngredients = ingredients

	// steps
	steps := make([]*v2.Step, 0)
	stepsInput := input.Steps

	for i := 0; i < len(stepsInput); i++ {
		step := &v2.Step{}
		stepInput := stepsInput[i]

		step.Position = stepInput.Position
		step.Description = stepInput.Description

		steps = append(steps, step)
	}
	output.Steps = steps

	return output, nil
}
