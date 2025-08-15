package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // allow CORS for testing
	fmt.Fprint(w, `{"message": "Request Received"}`)
}

func main() {
	http.HandleFunc("/add/card", handler)
	http.ListenAndServe(":8080", nil)
}
