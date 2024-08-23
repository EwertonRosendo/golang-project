package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"io"
	"api/src/models"
	
)

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
func FilterGoogleBooks(bookList []byte) []models.Book {
	var booksResponse models.GoogleBooksResponse
	err := json.Unmarshal(bookList, &booksResponse)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return nil
	}

	var books []models.Book
	for _, item := range booksResponse.Items {
		book := models.Book{
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