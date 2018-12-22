package controllers
import (
	"net/http"
	"recify/models"
	"encoding/json"
	u "recify/utils"
	"strconv"
	"github.com/gorilla/mux"
)

var CreateRecipe = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user") . (uint) //Grab the id of the user that send the request
	recipe := &models.Recipe{}

	err := json.NewDecoder(r.Body).Decode(recipe)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	recipe.UserId = user
	resp := recipe.Create()
	u.Respond(w, resp)
}

var GetRecipesFor = func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["user_id"])
	if err != nil {
		//The passed path parameter is not an integer
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	data := models.GetRecipes(uint(id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}