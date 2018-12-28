package v2

import (
	"fmt"
	"github.com/lib/pq"
	"log"
	"recify/db"
)

const CategoryTableName = "recipe_category"

type Category struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

func GetCategoryById(categoryId uint) *Category {
	category := &Category{}

	conn := db.GetConnection()

	if conn == nil {
		return category
	}

	query := fmt.Sprintf(`SELECT id, name FROM %s WHERE id = $1`, CategoryTableName)

	rows, err := conn.Query(query, categoryId)

	if err != nil {
		log.Fatal(err)
		return category
	}

	for rows.Next() {
		c := new(Category)
		if err := rows.Scan(&c.Id, &c.Name); err != nil {
			log.Fatalln(err)
		}
		category = c
	}

	return category
}

func GetCategoriesByIds(ids []uint) []*Category {

	categories := make([]*Category, 0)

	conn := db.GetConnection()

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
		if err := rows.Scan(&c.Id, &c.Name); err != nil {
			log.Fatalln(err)
		}
		categories = append(categories, c)
	}

	return categories
}
