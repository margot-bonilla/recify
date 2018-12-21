package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"recify/controllers"
	"strings"
)

func sayHello(res http.ResponseWriter, req *http.Request) {
	message := req.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message
	res.Write([]byte(message))
}

//func main() {
//
//	router := mux.NewRouter()
//	//router.Use(app.JwtAuthentication) //attach JWT auth middleware
//
//	router.HandleFunc("/", sayHello)
//
//	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
//	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
//
//	port := os.Getenv("PORT") //Get port from .env file, we did not specify any port so this should return an empty string when tested locally
//	if port == "" {
//		port = "8000" //localhost
//	}
//
//	fmt.Println(port)
//
//	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
//	if err != nil {
//		fmt.Print(err)
//	}
//}

func main() {
		router := mux.NewRouter()
		//router.Use(app.JwtAuthentication) //attach JWT auth middleware

		router.HandleFunc("/", sayHello)

		router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
		router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")

		port := os.Getenv("PORT") //Get port from .env file, we did not specify any port so this should return an empty string when tested locally
		if port == "" {
			port = "8000" //localhost
		}

		fmt.Println(port)

		err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
		if err != nil {
			fmt.Print(err)
		}
}