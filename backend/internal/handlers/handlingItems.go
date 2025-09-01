package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AbhilashDorle/AlgoTrack/internal/db"
)

type Card struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func AddCard(w http.ResponseWriter, r *http.Request) {
	var card Card
	fmt.Println("Inside Add Card")

	// Decode JSON from request body
	if err := json.NewDecoder(r.Body).Decode(&card); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err := db.DB.QueryRow(
		"INSERT INTO cards (title, description) VALUES ($1, $2) RETURNING id",
		card.Title, card.Description,
	).Scan(&card.ID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// (Here you’d insert into DB — for now, just echo back)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(card)
}

func UpdateCard(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var card Card

	// Decode JSON
	if err := json.NewDecoder(r.Body).Decode(&card); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}

	// Validate ID
	if card.ID <= 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid card ID"})
		return
	}

	// Update DB
	res, err := db.DB.Exec(
		"UPDATE cards SET title = $1, description = $2 WHERE id = $3",
		card.Title, card.Description, card.ID,
	)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil || rowsAffected == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "No card found with given ID"})
		return
	}

	// Success response
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*") // For development only
	json.NewEncoder(w).Encode(card)
}
