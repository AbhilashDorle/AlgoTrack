package routes

import (
	"fmt"

	"github.com/AbhilashDorle/AlgoTrack/internal/handlers"
	"github.com/gorilla/mux"
)

func RegisterCardRoutes(r *mux.Router) {
	fmt.Println("Request received in router")
	r.HandleFunc("/add/card", handlers.AddCard).Methods("POST")
}
