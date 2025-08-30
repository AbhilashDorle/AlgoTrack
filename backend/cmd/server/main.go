package main

import (
	"log"
	"net/http"

	"github.com/AbhilashDorle/AlgoTrack/internal/routes"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// func handler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Access-Control-Allow-Origin", "*") // allow CORS for testing
// 	fmt.Fprint(w, `{"message": "Request Received"}`)
// }

func main() {
	r := mux.NewRouter()
	routes.RegisterCardRoutes(r)

	// CORS setup (allow all origins for dev)
	corsObj := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)

	log.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", corsObj(r))
}
