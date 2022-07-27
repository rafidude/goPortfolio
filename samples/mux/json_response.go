package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/now", NowHandler)

	log.Println("Listening ...")
	http.ListenAndServe(":8080", r)
}

func NowHandler(resp http.ResponseWriter, _ *http.Request) {
	now := time.Now()

	payload := make(map[string]string)
	payload["now"] = now.Format(time.ANSIC)

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)

	json.NewEncoder(resp).Encode(payload)
}
