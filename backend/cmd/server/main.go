package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Request Received")
}

func main() {
	http.HandleFunc("/add/playlist", handler)
	http.ListenAndServe(":8080", nil)
}
