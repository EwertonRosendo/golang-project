package controllers

import (
	"net/http"
	"net/url"
	"fmt"
	"github.com/gorilla/mux"
	"os"
	"io"
)

func SearchGoogleBooks(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("https://www.googleapis.com/books/v1/volumes?q=intitle:golang&maxResults=30&key=%s", os.Getenv("GOOGLE_BOOKS_API_KEY"))
	w.Write([]byte(fmt.Sprintf("Search book default route and the url to google books api is %s", url)))
}
func SearchGoogleBooksByTitle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	url := fmt.Sprintf("https://www.googleapis.com/books/v1/volumes?q=intitle:%s&maxResults=30&key=%s", url.QueryEscape(title), os.Getenv("GOOGLE_BOOKS_API_KEY"))

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		http.Error(w, "Error creating request", http.StatusInternalServerError)
		return
	}

	request.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		http.Error(w, "Error executing request", http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		http.Error(w, "Error reading response", http.StatusInternalServerError)
		return
	}

	// Return the raw JSON response from Google Books API
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}
