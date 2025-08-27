package main

import (
	"net/http"

	"github.com/AbhilashDorle/AlgoTrack/internal/routes"
	"github.com/gorilla/mux"
)

// func handler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Access-Control-Allow-Origin", "*") // allow CORS for testing
// 	fmt.Fprint(w, `{"message": "Request Received"}`)
// }

func main() {
	// http.HandleFunc("/add/card", handler)

	r := mux.NewRouter()

	routes.RegisterItemRoutes(r)

	http.ListenAndServe(":8080", r)
}
