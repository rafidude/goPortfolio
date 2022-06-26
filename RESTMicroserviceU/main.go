package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	Name string `json:"name" xml:"name"`
	City string `json:"city" xml:"city"`
	Zip  string `json:"zip_code" xml:"zip_code"`
}

//declare a slice of customers
var customers = []Customer{
	{Name: "John", City: "New York", Zip: "10001"},
	{Name: "Jane", City: "New York", Zip: "10002"},
}

func main() {
	router := mux.NewRouter()

	// Setting up routes
	router.HandleFunc("/greet", greet)
	router.HandleFunc("/customersjson", getCustomersJSON)
	router.HandleFunc("/customersxml", getCustomersXML)
	router.HandleFunc("/customers", getCustomers)

	// Starting the server
	log.Fatal(http.ListenAndServe(":8080", router))
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q\n", r.URL.Path)
}

func getCustomersJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

//send XML data
func getCustomersXML(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/xml")
	xml.NewEncoder(w).Encode(customers)
}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Accept") == "application/xml" {
		getCustomersXML(w, r)
	} else {
		getCustomersJSON(w, r)
	}
}
