package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"fmt"

	//"fmt"
	//"os"
	//"api/src/services"
	"encoding/json"
	"io"
	"net/http"
	//"github.com/gorilla/mux"
)

// Handle the /googlebooks endpoint
func SearchBooks(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	 if err != nil{
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	 }
	 defer db.Close()
	 repository := repositories.NewBookRepository(db)
	 books, err := repository.SearchBooks()
	 if err != nil{
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	 }
	 responses.JSON(w, http.StatusOK, books)
}

// Handle the /googlebooks/{title} endpoint
func SearchBooksByTitle(w http.ResponseWriter, r *http.Request) {
	
}

// add a book from google books api to our database
func AddBook(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := io.ReadAll(r.Body)

	fmt.Println("params: ", string(bodyRequest))

	if err != nil{
		responses.ERR(w, http.StatusUnprocessableEntity, err)
		return
	}
	var book models.Book

	if err = json.Unmarshal(bodyRequest, &book); err != nil{
		responses.ERR(w, http.StatusBadRequest, err)
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