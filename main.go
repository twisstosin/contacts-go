package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/twisstosin/contacts-go/app"
	"github.com/twisstosin/contacts-go/controllers"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/user/register", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/contacts/new", controllers.CreateContact).Methods("POST")
	router.HandleFunc("/api/me/contacts", controllers.GetContactsFor).Methods("GET") // user/id/contacts

	router.Use(app.JwtAuthentication) // attach jwt middleware
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000" // localhost default
	}

	fmt.Println("Port ", port)

	err := http.ListenAndServe(":"+port, router) // launch the app, visit localhost:8000/api

	if err != nil {
		fmt.Print(err)
	}
}
