package app

import (
	"log"
	"microservice1/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	// Setting up routes
	router.HandleFunc("/greet", handlers.Greet)
	router.HandleFunc("/customers", handlers.GetCustomers)

	// Starting the server
	log.Fatal(http.ListenAndServe(":8080", router))
}
