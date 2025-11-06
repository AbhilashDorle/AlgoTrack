package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/AbhilashDorle/AlgoTrack/internal/utils"
	"github.com/gorilla/mux"
)

func GetSuggestions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	userID := vars["user_id"]

	fmt.Println("Inside GetSuggestions for user:", userID)

	// Fetch cards from DB
	cards, err := FetchCardsByUser(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	if len(cards) == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "No cards found"})
		return
	}

	// Ensure user_id is set for all cards before sending to AI
	for i := range cards {
		cards[i].UserID = userID
	}

	// Call AI to get suggestions
	suggestionStr, err := utils.CallAI_ForSuggestions(cards)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintf("AI call failed: %s", err.Error())})
		return
	}

	// Debug print
	fmt.Println("Raw AI response:", suggestionStr)

	suggestionStr = strings.TrimSpace(suggestionStr)
	suggestionStr = strings.Trim(suggestionStr, "`")

	start := strings.Index(suggestionStr, "[")
	end := strings.LastIndex(suggestionStr, "]")
	if start != -1 && end != -1 && start < end {
		suggestionStr = suggestionStr[start : end+1]
	}

	// Unmarshal AI response into []Card
	var aiCards []Card
	err = json.Unmarshal([]byte(suggestionStr), &aiCards)
	if err != nil {
		// If unmarshalling fails, log the AI output for debugging
		fmt.Println("Failed to unmarshal AI response:", err)
		fmt.Println("AI response string:", suggestionStr)

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error":       fmt.Sprintf("Failed to parse AI response: %s", err.Error()),
			"ai_response": suggestionStr,
		})
		return
	}

	// Ensure user_id is set in all AI cards (in case AI overwrote it)
	for i := range aiCards {
		aiCards[i].UserID = userID
	}

	// Final debug print
	fmt.Println("AI suggestions:", aiCards)

	// Send response to frontend
	json.NewEncoder(w).Encode(aiCards)
}
