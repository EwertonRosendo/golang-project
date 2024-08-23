package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"io"
	"github.com/gorilla/mux"
)

// Define the response structure for Google Books API
type GoogleBooksResponse struct {
	Items []struct {
		VolumeInfo struct {
			Title         string   `json:"title"`
			Subtitle      string   `json:"subtitle"`
			Description   string   `json:"description"`
			Authors       []string `json:"authors"`
			PublishedDate string   `json:"publishedDate"`
			Publisher     string   `json:"publisher"`
			ImageLinks    struct {
				Thumbnail string `json:"thumbnail"`
			} `json:"imageLinks"`
		} `json:"volumeInfo"`
	} `json:"items"`
}

// Define the Book struct (optional, if you need to manipulate data)
type Book struct {
	Title         string   `json:"title,omitempty"`
	Subtitle      string   `json:"subtitle,omitempty"`
	Description   string   `json:"description,omitempty"`
	Authors       []string `json:"authors,omitempty"`
	PublishedDate string   `json:"publishedDate,omitempty"`
	Publisher     string   `json:"publisher,omitempty"`
	Thumbnail     string   `json:"thumbnail,omitempty"`
}

// Handle the /googlebooks endpoint
func SearchGoogleBooks(w http.ResponseWriter, r *http.Request) {
	responseBody, err := GoogleBooksRequest("")
	if err != nil {
		http.Error(w, "Failed to fetch books", http.StatusInternalServerError)
		return
	}

	// Optional: Filter books
	filteredBooks := FilterGoogleBooks(responseBody)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filteredBooks)
}

// Handle the /googlebooks/{title} endpoint
func SearchGoogleBooksByTitle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]

	responseBody, err := GoogleBooksRequest(title)
	if err != nil {
		http.Error(w, "Failed to fetch books", http.StatusInternalServerError)
		return
	}

	// Optional: Filter books
	filteredBooks := FilterGoogleBooks(responseBody)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filteredBooks)
}

// Make the request to the Google Books API
func GoogleBooksRequest(title string) ([]byte, error) {
	apiURL := fmt.Sprintf("https://www.googleapis.com/books/v1/volumes?q=intitle:%s&maxResults=10&key=%s", url.QueryEscape(title), os.Getenv("GOOGLE_BOOKS_API_KEY"))

	request, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

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

// Optional: Filter and manipulate the book data
func FilterGoogleBooks(bookList []byte) []Book {
	var booksResponse GoogleBooksResponse
	err := json.Unmarshal(bookList, &booksResponse)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return nil
	}

	var books []Book
	for _, item := range booksResponse.Items {
		book := Book{
			Title:         item.VolumeInfo.Title,
			Subtitle:      item.VolumeInfo.Subtitle,
			Description:   item.VolumeInfo.Description,
			Authors:       item.VolumeInfo.Authors,
			PublishedDate: item.VolumeInfo.PublishedDate,
			Publisher:     item.VolumeInfo.Publisher,
			Thumbnail:     item.VolumeInfo.ImageLinks.Thumbnail,
		}
		books = append(books, book)
	}
	return books
}
