package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/AbhilashDorle/AlgoTrack/internal/db"
	"github.com/gorilla/mux"
)

type Card struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UserID      string `json:"user_id"`
}

// AddCard inserts a new card into the DB
func AddCard(w http.ResponseWriter, r *http.Request) {
	var card Card

	// Decode JSON from request body
	if err := json.NewDecoder(r.Body).Decode(&card); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	/*id, title description, user_id*/

	err := db.DB.QueryRow(
		"INSERT INTO cards (user_id, title, description) VALUES ($1, $2, $3) RETURNING id",
		card.UserID, card.Title, card.Description,
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

// UpdateCard updates an existing card
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

	// Validate ID and UserID
	if card.ID <= 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid card ID"})
		return
	}

	// Update DB - ensure only the owner can update
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
		json.NewEncoder(w).Encode(map[string]string{"error": "No card found with given ID for this user"})
		return
	}

	// Success response
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(card)
}

// DeleteCard handles DELETE /cards/{id}
func DeleteCard(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idStr := vars["id"]

	// Convert string ID to int
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid card ID"})
		return
	}

	// Delete from DB - ensure only owner can delete
	res, err := db.DB.Exec(
		"DELETE FROM cards WHERE id = $1",
		id,
	)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	} else {
		fmt.Println(err)
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Card not found for this user"})
		return
	}

	// Success response
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(map[string]string{"message": "Card deleted successfully"})
}

func GetCards(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	userID := vars["user_id"]
	cards, err := FetchCardsByUser(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	// If no cards found
	if len(cards) == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "No cards found for this user"})
		return
	}

	// Success
	json.NewEncoder(w).Encode(cards)
}

func FetchCardsByUser(userID string) ([]Card, error) {
	rows, err := db.DB.Query("SELECT id, title, description FROM cards WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cards := []Card{}
	for rows.Next() {
		var card Card
		if err := rows.Scan(&card.ID, &card.Title, &card.Description); err != nil {
			return nil, err
		}
		cards = append(cards, card)
	}

	return cards, nil
}
