package v2

import (
	"database/sql"
	"log"
	"recify/db"
	"time"
)

type Recipe struct {
	Id uint `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (recipe *Recipe) Create() (sql.Result, error){

	connection := db.GetConnection()

	query := `INSERT INTO recipe(title, description, created_at) VALUES ($1, $2, current_timestamp)`

	stmt, err := connection.Prepare(query)

	if err != nil {
		log.Fatalln(err)
		panic(err)
	}

	return stmt.Exec(recipe.Title, recipe.Description)
}