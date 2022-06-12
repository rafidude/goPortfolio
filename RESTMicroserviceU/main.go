package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	Name string `json:"name"`
	City string `json:"city"`
	Zip  string `json:"zip_code"`
}

func main() {
	router := mux.NewRouter()

	// Setting up routes
	router.HandleFunc("/greet", greet)
	router.HandleFunc("/customers", getCustomersJSON)

	// Starting the server
	log.Fatal(http.ListenAndServe(":8080", router))
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q\n", r.URL.Path)
}

func getCustomersJSON(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "John", City: "New York", Zip: "10001"},
		{Name: "Jane", City: "New York", Zip: "10001"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
