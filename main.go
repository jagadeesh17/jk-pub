package main

import (
	"app/middlewares"
	"fmt"
	"net/http"

	"app/api"

	"github.com/gorilla/mux"
)

func main() {
	// Create a new router
	router := mux.NewRouter()
	// Apply JWT middleware to protect routes
	router.Use(middlewares.JWTMiddleware)

	// Define routes and their corresponding handlers
	router.HandleFunc("/contracts", api.CreateNewContract).Methods("POST")
	router.HandleFunc("/subscriptions", api.ConfigureSubscription).Methods("POST")

	// Start the server on port 8080
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", router)
}
