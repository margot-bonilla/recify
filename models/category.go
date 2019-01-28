package models

import (
	"fmt"
	"log"
	database "recify/db"

	"github.com/lib/pq"
)

const CategoryTableName = "recipe_category"

type Category struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func GetCategoryById(categoryID uint) *Category {
	category := &Category{}

	conn := database.GetConnection()

	if conn == nil {
		return category
	}

	query := fmt.Sprintf(`SELECT id, name FROM %s WHERE id = $1`, CategoryTableName)

	rows, err := conn.Query(query, categoryID)

	if err != nil {
		log.Fatal(err)
		return category
	}

	for rows.Next() {
		c := new(Category)
		if err := rows.Scan(&c.ID, &c.Name); err != nil {
			log.Fatalln(err)
		}
		category = c
	}

	return category
}

func GetCategoriesByIds(ids []uint) []*Category {

	categories := make([]*Category, 0)

	conn := database.GetConnection()

	if conn == nil {
		return categories
	}

	query := fmt.Sprintf(`SELECT id, name FROM %s WHERE id = ANY ($1)`, CategoryTableName)

	rows, err := conn.Query(query, pq.Array(ids))

	if err != nil {
		log.Fatal(err)
		return categories
	}

	for rows.Next() {
		c := new(Category)
		if err := rows.Scan(&c.ID, &c.Name); err != nil {
			log.Fatalln(err)
		}
		categories = append(categories, c)
	}

	return categories
}
