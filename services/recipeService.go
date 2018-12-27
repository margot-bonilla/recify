package services

import (
	"fmt"
	"log"
	"recify/db"
	"recify/models/v2"
)

var CreateRecipe = func(recipe* v2.Recipe, categories []*v2.Category)  uint {


	conn := db.GetConnection()

	t, err := conn.Begin()

	if err != nil {
		log.Fatal()
	}

	query := `INSERT INTO recipe(title, description, created_at) VALUES ($1, $2, current_timestamp) RETURNING id`

	res := t.QueryRow(query, recipe.Title, recipe.Description)

	if err != nil {
		t.Rollback()
		log.Fatalln(err)
	}

	var recipeId uint
	err = res.Scan(&recipeId)

	if err != nil {
		log.Fatal()
	}

	var categoriesQuery = "INSERT INTO recipe_to_category VALUES "

	for i:= 0; i < len(categories); i++ {
		c := categories[i]
		categoriesQuery += fmt.Sprintf("(%d, %d)", recipeId, c.Id)

		if i >= 0 && i < (len(categories) - 1) {
			categoriesQuery += ","
		}
	}

	_ , err = t.Exec(categoriesQuery)

	if err != nil {
		t.Rollback()
		log.Fatalln(err)
	}

	err = t.Commit()

	if err != nil {
		log.Fatalln(err)
	}

	return recipeId
}