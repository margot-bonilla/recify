package inputs

import "recify/models"

type RecipeInput struct {
	Title       string
	Description string
	Categories  []uint
	Ingredients []AmountIngredientInput
	Steps       []StepInput
}

type AmountIngredientInput struct {
	ID      uint    `json:"id"`
	Amount  float32 `json:"amount"`
	Measure string  `json:"measure"`
}

type StepInput struct {
	Position    uint
	Description string
}

type RecipeOutput struct {
	Recipe            *models.Recipe
	Categories        []*models.Category
	AmountIngredients []*models.RecipeIngredient
	Steps             []*models.Step
}

func (input *RecipeInput) Input2Models() (*RecipeOutput, error) {

	output := &RecipeOutput{}

	// recipe
	recipe := &models.Recipe{}
	recipe.Title = input.Title
	recipe.Description = input.Description

	output.Recipe = recipe

	// categories
	// todo move this checking to service
	categories := models.GetCategoriesByIds(input.Categories)

	output.Categories = categories

	// ingredients
	ingredients := make([]*models.RecipeIngredient, 0)
	ingredientsInput := input.Ingredients

	for i := 0; i < len(ingredientsInput); i++ {
		ing := new(models.RecipeIngredient)
		ingInput := ingredientsInput[i]

		ing.Amount = ingInput.Amount
		ing.Measure = ingInput.Measure
		ing.IngredientID = ingInput.ID

		ingredients = append(ingredients, ing)
	}
	output.AmountIngredients = ingredients

	// steps
	steps := make([]*models.Step, 0)
	stepsInput := input.Steps

	for i := 0; i < len(stepsInput); i++ {
		step := &models.Step{}
		stepInput := stepsInput[i]

		step.Position = stepInput.Position
		step.Description = stepInput.Description

		steps = append(steps, step)
	}
	output.Steps = steps

	return output, nil
}
