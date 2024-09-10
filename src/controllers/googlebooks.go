package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/services"
	"encoding/json"
	"os"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

// Handle the /googlebooks endpoint
func SearchGoogleBooks(w http.ResponseWriter, r *http.Request) {
	
	responseBody, err := services.GoogleBooksRequest("GOOGLE")
	if err != nil {
		http.Error(w, "Failed to fetch books", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseBody)
}

// Handle the /googlebooks/{title} endpoint
func SearchGoogleBooksByTitle(w http.ResponseWriter, r *http.Request) {
	vars 	:= mux.Vars(r)
	title 	:= vars["title"]

	responseBody, err := services.GoogleBooksRequest(title)
	if 	err != nil {
		http.Error(w, "Failed to fetch books", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseBody)
}

// add a book from google books api to our database
func AddGoogleBook(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := io.ReadAll(r.Body)

	if err != nil{
		responses.ERR(w, http.StatusUnprocessableEntity, err)
		return
	}
	var book models.Book
	
	
	if err = json.Unmarshal(bodyRequest, &book); err != nil{
		responses.ERR(w, http.StatusBadRequest, err)
		return
	}

	out, err := os.Create("static/"+book.Title+".jpg")
	if err != nil  {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}
	defer out.Close()

	resp, err := http.Get(book.Thumbnail)
	if err != nil {
		responses.ERR(w, http.StatusBadRequest, err)
		return 
	}
  	defer resp.Body.Close()
	  _, err = io.Copy(out, resp.Body)
	  if err != nil  {
		responses.ERR(w, http.StatusInternalServerError, err)
		return 
	  }
	
	book.FormatBook()

	db, err := database.Connect()
	if err != nil{
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewBookRepository(db)
	book.ID, err = repository.Create(book)
	
	if err != nil{
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, book)
}