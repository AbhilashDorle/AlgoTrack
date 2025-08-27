package handlers

import (
	"fmt"
	"net/http"
)

func AddItem(w http.ResponseWriter, r *http.Request) {
	// 	var item Item
	// 	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
	// 		http.Error(w, err.Error(), http.StatusBadRequest)
	// 		return
	// 	}

	// 	if err := db.InsertItem(item); err != nil {
	// 		http.Error(w, "Failed to insert item", http.StatusInternalServerError)
	// 		return
	// 	}

	// 	w.WriteHeader(http.StatusCreated)
	// 	json.NewEncoder(w).Encode(item)
	fmt.Println("Request Received in handler")
}
