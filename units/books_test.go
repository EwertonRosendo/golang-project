package integration

import (
	"api/src/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"testing"
)

var book_test models.Book

func TestAddBook(t *testing.T) {
	book := models.Book{
		Title:       "testing books repository",
		Subtitle:    "adding subtitles for this book",
		Description: "a very good language to learn and work with",
		Thumbnail:   "https://bajaj.com.br/wp-content/themes/twentytwentyone/img/home/dominar-250.png",
		Authors:     "ewerton rosendo",
	}

	marshalled_book, err := json.Marshal(book)
	if err != nil {
		log.Fatalf("impossible to marshal book: %s", err)
	}

	request, err := http.NewRequest("POST", "http://localhost:5000/books/add", bytes.NewReader(marshalled_book))
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		t.Fatalf("failed to perform request: %v", err)
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}

	if err = json.Unmarshal(responseBody, &book); err != nil {
		t.Fatalf("error reading response body: %v", err)
	}

	book_test.ID = book.ID

	expectedStatusCode := 201
	if response.StatusCode != expectedStatusCode {
		t.Errorf("Expected status code %v, but got %v", expectedStatusCode, response.StatusCode)
	}
}

func TestUpdateBook(t *testing.T) {
	if book_test.ID == 0 {
		t.Fatalf("TestUpdateBook: book_test.ID is not set, ensure TestAddBookRepository runs first")
	}

	book := models.Book{
		Title:    "testing update books repository",
		Subtitle: "updating subtitles for this book",
		Authors:  "ewerton rosendo da silva",
	}

	marshalledBook, err := json.Marshal(book)
	if err != nil {
		t.Fatalf("impossible to marshal book: %s", err)
	}

	id := strconv.FormatUint(uint64(book_test.ID), 10)
	request, err := http.NewRequest("PUT", fmt.Sprintf("http://localhost:5000/books/%s", id), bytes.NewReader(marshalledBook))
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		t.Fatalf("failed to perform request: %v", err)
	}

	expectedStatusCode := 204
	if response.StatusCode != expectedStatusCode {
		t.Errorf("Expected status code %v, but got %v", expectedStatusCode, response.StatusCode)
	}
}

func TestGetBooks(t *testing.T) {

	var id string = strconv.FormatUint(uint64(book_test.ID), 10)
	request, err := http.NewRequest("GET", fmt.Sprintf("http://localhost:5000/books/%s", id), nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		t.Fatalf("failed to perform request: %v", err)
	}
	defer response.Body.Close()

	expectedStatusCode := 200
	if response.StatusCode != expectedStatusCode {
		t.Errorf("Expected status code %v, but got %v", expectedStatusCode, response.StatusCode)
	}
}

func TestGetBooksById(t *testing.T) {

	var id string = strconv.FormatUint(uint64(book_test.ID), 10)
	request, err := http.NewRequest("GET", fmt.Sprintf("http://localhost:5000/books/%s", id), nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		t.Fatalf("failed to perform request: %v", err)
	}
	defer response.Body.Close()

	expectedStatusCode := 200
	if response.StatusCode != expectedStatusCode {
		t.Errorf("Expected status code %v, but got %v", expectedStatusCode, response.StatusCode)
	}
}

func TestDeleteBooksById(t *testing.T) {

	var id string = strconv.FormatUint(uint64(book_test.ID), 10)
	request, err := http.NewRequest("DELETE", fmt.Sprintf("http://localhost:5000/books/%s", id), nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		t.Fatalf("failed to perform request: %v", err)
	}
	defer response.Body.Close()

	expectedStatusCode := 204
	if response.StatusCode != expectedStatusCode {
		t.Errorf("Expected status code %v, but got %v", expectedStatusCode, response.StatusCode)
	}
}
