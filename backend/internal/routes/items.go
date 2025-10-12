package routes

import (
	"fmt"
	"net/http"

	"github.com/AbhilashDorle/AlgoTrack/internal/handlers"
	"github.com/AbhilashDorle/AlgoTrack/internal/middleware"
	"github.com/gorilla/mux"
)

func RegisterCardRoutes(r *mux.Router) {
	fmt.Println("Request received in router")
	r.Handle("/add/card", middleware.AuthMiddleware(http.HandlerFunc(handlers.AddCard))).Methods("POST")
	r.Handle("/update/card", middleware.AuthMiddleware(http.HandlerFunc(handlers.UpdateCard))).Methods("PUT")
	r.Handle("/card/{id}", middleware.AuthMiddleware(http.HandlerFunc(handlers.DeleteCard))).Methods("DELETE")

	// r.HandleFunc("/add/card", handlers.AddCard).Methods("POST")
	// r.HandleFunc("/update/card", handlers.UpdateCard).Methods("PUT")
	// r.HandleFunc("/cards/{id}", handlers.DeleteCard).Methods("DELETE")
}
