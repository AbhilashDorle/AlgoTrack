package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AbhilashDorle/AlgoTrack/internal/utils"
	"github.com/gorilla/mux"
)

func GetSuggestions(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["user_id"]

	fmt.Println("InsideGetSuggestions")

	cards, err := FetchCardsByUser(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	fmt.Println("InsideGetSuggestions after fetching cards")
	fmt.Println(cards)
	// Now you can do whatever with the cards, e.g., pick one randomly
	if len(cards) == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "No cards found"})
		return
	}

	suggestion, err := utils.CallAI_ForSuggestions(cards)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	// w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// json.NewEncoder(w).Encode(cards)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(map[string]string{"suggestion": suggestion})
}
