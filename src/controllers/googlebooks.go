package controllers

import (
	"net/http"
	"net/url"
	"fmt"
	"github.com/gorilla/mux"
	"os"
	"io"
)
// func to handle the /googlebooks endpoint
func SearchGoogleBooks(w http.ResponseWriter, r *http.Request) {
	responseBody, err := GoogleBooksRequest("")
	if err != nil{
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}

// func to handle the /googlebooks/{title} endpoint
func SearchGoogleBooksByTitle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title_request := vars["title"]
	responseBody, err := GoogleBooksRequest(title_request)
	if err != nil{
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}

//  func to make the request to the google books api
func GoogleBooksRequest(title string) ([]byte, error) {
	// Use a different variable name to avoid shadowing the url package
	apiURL := fmt.Sprintf("https://www.googleapis.com/books/v1/volumes?q=intitle:%s&maxResults=30&key=%s", url.QueryEscape(title), os.Getenv("GOOGLE_BOOKS_API_KEY"))

	request, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	request.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("failed to perform request: %v", err)
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	return responseBody, nil
}