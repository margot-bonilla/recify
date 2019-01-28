/**
 * command: go run resources/importDataSet.go resources/full_format_recipes.json
 */
package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"recify/models"
)
import "fmt"

type RecipeDataSet struct {
	Title       string
	Directions  []string
	Fat         float32
	Date        string
	Categories  []string
	Calories    float32
	Desc        string
	Protein     float32
	Rating      float64
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

	// we initialize our Users array
	var recipes []RecipeDataSet

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &recipes)

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	for i := 0; i < len(recipes); i++ {
		fmt.Println(recipes[i].Title)
		var recipe models.Recipe
		recipe.Title = recipes[i].Title
		recipe.Rating = recipes[i].Rating
		recipe.UserId = 1
		recipe.Create()

		if recipe.ID != 0 {
			for j := 0; j < len(recipes[i].Directions); j++ {
				var step models.Step
				step.Position = j
				step.Description = recipes[i].Directions[j]
				step.RecipeId = recipe.ID
				step.Create()
			}
			for j := 0; j < len(recipes[i].Categories); j++ {
				var cat models.Category
				cat.Name = recipes[i].Categories[j]
				cat.Create()

				// @TODO relation recipe-category
			}

			for j := 0; j < len(recipes[i].Ingredients); j++ {
				// @TODO create a rule to split the text into amount - measure - ingredient

				var ing models.Ingredient
				ing.Name = recipes[i].Ingredients[j]
				ing.Create()

				var recipeIngredient models.RecipeIngredient
				recipeIngredient.IngredientId = ing.ID
				recipeIngredient.RecipeID = recipe.ID

				recipeIngredient.Create()
			}

		}
	}
}
