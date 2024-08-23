package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"api/src/services"
)

// Handle the /googlebooks endpoint
func SearchGoogleBooks(w http.ResponseWriter, r *http.Request) {
	responseBody, err := services.GoogleBooksRequest("")
	if err != nil {
		http.Error(w, "Failed to fetch books", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseBody)
}

// Handle the /googlebooks/{title} endpoint
func SearchGoogleBooksByTitle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]

	responseBody, err := services.GoogleBooksRequest(title)
	if err != nil {
		http.Error(w, "Failed to fetch books", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseBody)
}
