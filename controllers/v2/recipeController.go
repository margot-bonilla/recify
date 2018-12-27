package v2

import (
	"encoding/json"
	"log"
	"net/http"
	"recify/controllers/v2/inputs"
	"recify/models/v2"
	"recify/services"
	u "recify/utils"
	"time"
)

var CreateRecipe = func(w http.ResponseWriter, r *http.Request) {
	// userId := r.Context().Value("user") . (uint)
	recipeInput := &inputs.RecipeInput{}

	err:= json.NewDecoder(r.Body).Decode(recipeInput)

	if err != nil {
		u.Respond(w, u.Message(false, "Error decoding request body"))
		return
	}

	// assign to categories :=
	categories := v2.GetCategoriesByIds(recipeInput.Categories)


	// create amountIngredient
	//for i:= 0; i < len(recipeInput.Ingredients); i++ {
	//
	//}

	recipe := &v2.Recipe{}
	recipe.Title = recipeInput.Title
	recipe.Description = recipeInput.Description
	recipe.CreatedAt = time.Now().UTC()

	services.CreateRecipe(recipe, categories)

	// _ , err = recipe.Create()

	var resp map[string] interface{}

	if err != nil {
		resp = u.Message(false, "Something went wrong")
		log.Fatal(err)
	} else {
		resp = u.Message(true, "success")
	}

	u.Respond(w, resp)
}