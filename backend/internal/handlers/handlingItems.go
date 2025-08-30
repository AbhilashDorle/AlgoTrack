package handlers

import (
	"encoding/json"
	"net/http"
)

type Card struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func AddCard(w http.ResponseWriter, r *http.Request) {
	var card Card

	// Decode JSON from request body
	if err := json.NewDecoder(r.Body).Decode(&card); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// (Here you’d insert into DB — for now, just echo back)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(card)
}
