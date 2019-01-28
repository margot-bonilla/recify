package db

import (
	"database/sql"
	"fmt"
	"os"
)

func GetConnection() *sql.DB {

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)

	database, err := sql.Open("postgres", dbUri)

	if err != nil {
		fmt.Print(err)
	}

	return database
}
