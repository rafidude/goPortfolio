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
	router.HandleFunc("/greet", handlers.Greet).Methods(http.MethodGet)
	router.HandleFunc("/customers", handlers.GetCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", handlers.GetCustomer).Methods(http.MethodGet)

	// Starting the server
	log.Fatal(http.ListenAndServe(":8080", router))
}
