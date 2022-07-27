package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// curl http://localhost:8080/hello\?name\=rafi
func main() {
	r := mux.NewRouter()

	r.HandleFunc("/hello", func(resp http.ResponseWriter, req *http.Request) {
		name := req.URL.Query().Get("name")
		if name == "" {
			name = "guest"
		}
		fmt.Fprintf(resp, "Hello %s!", name)
	})

	log.Println("Listening ...")
	http.ListenAndServe(":8080", r)
}
