package main

import (
	"log"
	"net/http"

	"github.com/AbhilashDorle/AlgoTrack/firebase"
	"github.com/AbhilashDorle/AlgoTrack/internal/db"
	"github.com/AbhilashDorle/AlgoTrack/internal/routes"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {

	firebase.InitFirebase()

	r := mux.NewRouter()
	routes.RegisterCardRoutes(r)

	db.Connect()
	defer db.DB.Close()

	// CORS setup (allow all origins for dev)
	corsObj := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)

	log.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", corsObj(r))
}
