/**
 * command: go run resources/importDataSet.go resources/full_format_recipes.json
 */
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"recify/models"
	"recify/repositories"
)

type recipeDataSet struct {
	Title       string
	Directions  []string
	Fat         float32
	Date        string
	Categories  []string
	Calories    float32
	Desc        string
	Protein     float32
	Rating      float32
	Ingredients []string
	Sodium      float32
}

func main() {
	//fileName := os.Args[1] //"full_format_recipes.json"

	// Open our jsonFile
	jsonFile, err := os.Open("resources/full_format_recipes.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened ")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our recipes array
	var recipes []recipeDataSet

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'recipes' which we defined above
	json.Unmarshal(byteValue, &recipes)

	// we iterate through every recipe
	for i := 0; i < len(recipes); i++ {
		fmt.Println(recipes[i].Title)
		var recipe models.Recipe
		recipe.Title = recipes[i].Title
		recipe.Rating = recipes[i].Rating

		var steps = make([]models.Step, len(recipes[i].Directions))
		for j := 0; j < len(recipes[i].Directions); j++ {
			var step models.Step
			step.Position = uint(j)
			step.Description = recipes[i].Directions[j]
			step.RecipeID = recipe.ID
			steps[j] = step
		}

		var categories = make([]models.Category, len(recipes[i].Categories))
		for j := 0; j < len(recipes[i].Categories); j++ {
			var cat models.Category
			cat.Name = recipes[i].Categories[j]
			categories[j] = cat
		}

		var ingredients = make([]models.RecipeIngredient, len(recipes[i].Ingredients))
		for j := 0; j < len(recipes[i].Ingredients); j++ {
			// @TODO create a rule to split the text into amount - measure - ingredient

			var ing models.RecipeIngredient
			ing.Name = recipes[i].Ingredients[j]
			ing.Amount = 1
			ing.Measure = ""
			ingredients[j] = ing
		}

		recipeID, err := repositories.CreateRecipe(recipe, categories, ingredients, steps)

		if err != nil {
			log.Fatal()
		}
		log.Printf("Created recipe with id " + string(recipeID))
	}
}
