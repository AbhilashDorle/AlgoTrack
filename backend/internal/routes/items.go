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
	r.Handle("/deletecards/{id}", middleware.AuthMiddleware(http.HandlerFunc(handlers.DeleteCard))).Methods("DELETE")
	r.Handle("/getcards/{user_id}", middleware.AuthMiddleware(http.HandlerFunc(handlers.GetCards))).Methods("GET")

	r.Handle("/suggestions/{user_id}", middleware.AuthMiddleware(http.HandlerFunc(handlers.GetSuggestions))).Methods("GET")

}
