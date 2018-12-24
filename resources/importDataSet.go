/**
 * command: go run importDataSet.go full_format_recipes.json
 */
package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"recify/models"
	"strconv"
)
import "fmt"
func main() {
	fileName := os.Args[1]

	// Open our jsonFile
	jsonFile, err := os.Open(fileName)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened " + fileName)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()


	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var recipes []models.Recipe

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &recipes)

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	for i := 0; i < len(recipes); i++ {
		fmt.Println(strconv.Itoa(i) + ") Recipe Type: " + recipes[i].Title)
	}
}

