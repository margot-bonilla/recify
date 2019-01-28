package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"recify/controllers/v2/inputs"
	"recify/models"
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

	recipe := &models.Recipe{}
	recipe = output.Recipe

	categories := make([]*models.Category, 0)
	categories = output.Categories

	ingredients := make([]*models.RecipeIngredient, 0)
	ingredients = output.AmountIngredients

	steps := make([]*models.Step, 0)
	steps = output.Steps

	recipeID, err := repositories.CreateRecipe(recipe, categories, ingredients, steps)

	var resp map[string]interface{}

	if err != nil {
		resp = u.Message(false, "Something went wrong")
		log.Fatal(err)
	} else {
		resp = u.Message(true, "success")
		resp["recipe_id"] = recipeID
	}

	u.Respond(w, resp)
}
