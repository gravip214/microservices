// main.go
package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"microservice/handlers"
	"microservice/services"
	"net/http"
	//"os"
)

func main() {
	// Load environment variables
	err := services.LoadEnv()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/top-track", handlers.GetTopTrackInfo).Methods("GET")
	//r.HandleFunc("/users/{id:[0-9]+}", getUserHandler).Methods("GET")

	// Start the server
	http.Handle("/", r)
	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
