package routes

import (
	"fmt"

	"github.com/AbhilashDorle/AlgoTrack/internal/handlers"
	"github.com/gorilla/mux"
)

func RegisterCardRoutes(r *mux.Router) {
	fmt.Println("Request received in router")
	r.HandleFunc("/add/card", handlers.AddCard).Methods("POST")
	r.HandleFunc("/update/card", handlers.UpdateCard).Methods("PUT")
	r.HandleFunc("/cards/{id}", handlers.DeleteCard).Methods("DELETE")
}
