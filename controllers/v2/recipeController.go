package v2

import (
	"encoding/json"
	"log"
	"net/http"
	"recify/controllers/v2/inputs"
	"recify/models/v2"
	"recify/repositories"
	u "recify/utils"
)

var CreateRecipe = func(w http.ResponseWriter, r *http.Request) {
	// userId := r.Context().Value("user") . (uint)
	recipeInput := &inputs.RecipeInput{}

	err := json.NewDecoder(r.Body).Decode(recipeInput)

	if err != nil {
		u.Respond(w, u.Message(false, "Error decoding request body"))
		return
	}

	output, err := recipeInput.Input2Models()

	if err != nil {
		u.Respond(w, u.Message(false, "Error parsing request body"))
		return
	}

	recipe := &v2.Recipe{}
	recipe = output.Recipe

	categories := make([]*v2.Category, 0)
	categories = output.Categories

	ingredients := make([]*v2.AmountIngredient, 0)
	ingredients = output.AmountIngredients

	steps := make([]*v2.Step, 0)
	steps = output.Steps

	recipeId, err := repositories.CreateRecipe(recipe, categories, ingredients, steps)

	var resp map[string]interface{}

	if err != nil {
		resp = u.Message(false, "Something went wrong")
		log.Fatal(err)
	} else {
		resp = u.Message(true, "success")
		resp["recipe_id"] = recipeId
	}

	u.Respond(w, resp)
}
