package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func init() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	fmt.Println(dbURI)

	conn, err := gorm.Open("postgres", dbURI) // os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	db.SingularTable(true)
	//db.Debug().AutoMigrate(&User{})
	//db.Debug().AutoMigrate(&User{}, &Recipe{}, &Step{}, &Ingredient{}, &Category{}, &RecipeIngredient{})
}

// GetDB return db connection
func GetDB() *gorm.DB {
	return db
}
