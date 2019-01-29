package repositories

import (
	"fmt"
	"log"
	"recify/db"
	"recify/models"
)

const RecipeTableName = "recipe"
const RecipeToCategoryTableName = "recipe_category"
const AmountIngredientTableName = "recipe_ingredient"
const RecipeStepTable = "step"

var CreateRecipe = func(
	recipe *models.Recipe,
	categories []*models.Category,
	amountIngredients []*models.RecipeIngredient,
	steps []*models.Step) (uint, error) {

	conn := db.GetConnection()

	t, err := conn.Begin()

	if err != nil {
		log.Fatal()
	}

	query := fmt.Sprintf(`
		INSERT INTO %s (title, description, created_at) VALUES ($1, $2, current_timestamp) RETURNING id
	`, RecipeTableName)

	res := t.QueryRow(query, recipe.Title, recipe.Description)

	if err != nil {
		_ = t.Rollback()
		log.Fatalln(err)
		return 0, err
	}

	var recipeId uint
	err = res.Scan(&recipeId)

	if err != nil {
		_ = t.Rollback()
		log.Fatal(err)
		return 0, err
	}

	var catQuery = fmt.Sprintf(`INSERT INTO %s VALUES `, RecipeToCategoryTableName)

	for i := 0; i < len(categories); i++ {
		c := categories[i]
		catQuery += fmt.Sprintf("(%d, %d)", recipeId, c.ID)
		if i >= 0 && i < (len(categories)-1) {
			catQuery += ","
		}
	}

	_, err = t.Exec(catQuery)

	if err != nil {
		_ = t.Rollback()
		log.Fatalln(err)
		return 0, err
	}

	var amountIngredientsQuery = fmt.Sprintf(`
		INSERT INTO %s (recipe_id, name, amount, measure) VALUES
	`, AmountIngredientTableName)

	for i := 0; i < len(amountIngredients); i++ {
		ing := amountIngredients[i]
		amountIngredientsQuery += fmt.Sprintf(
			"(%d, %s, %f, '%s')", recipeId, ing.Name, ing.Amount, ing.Measure)
		if i >= 0 && i < (len(amountIngredients)-1) {
			amountIngredientsQuery += ","
		}
	}

	_, err = t.Exec(amountIngredientsQuery)

	if err != nil {
		_ = t.Rollback()
		log.Fatalln(err)
		return 0, err
	}

	var stepsQuery = fmt.Sprintf(`INSERT INTO %s (recipe_id, position, description) VALUES `, RecipeStepTable)

	for i := 0; i < len(steps); i++ {
		st := steps[i]
		stepsQuery += fmt.Sprintf("(%d, %d, '%s')", recipeId, st.Position, st.Description)
		if i >= 0 && i < (len(steps)-1) {
			stepsQuery += ","
		}
	}

	_, err = t.Exec(stepsQuery)

	if err != nil {
		_ = t.Rollback()
		log.Fatalln(err)
		return 0, err
	}

	err = t.Commit()

	if err != nil {
		log.Fatalln(err)
		return 0, err
	}

	return recipeId, nil
}
