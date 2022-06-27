package handlers

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
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

func Greet(w http.ResponseWriter, r *http.Request) {
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

func GetCustomers(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Accept") == "application/xml" {
		getCustomersXML(w, r)
	} else {
		getCustomersJSON(w, r)
	}
}

//get a single customer
func GetCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]
	fmt.Fprintf(w, "Customer ID is: %s\n", id)
}
