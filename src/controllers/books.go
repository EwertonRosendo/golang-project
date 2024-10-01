package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"fmt"
	"os"
	"encoding/json"
	"github.com/gorilla/mux"
	"strings"
	"io"
	"net/http"
	"regexp"
	"strconv"
)

// Handle the /googlebooks endpoint
func SearchBooks(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repository := repositories.NewBookRepository(db)
	books, err := repository.SearchBooks()
	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, books)
}

// Handle the /googlebooks/{title} endpoint
func SearchBooksByTitle(w http.ResponseWriter, r *http.Request) {

}

func FindBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	bookID, err := strconv.ParseUint(params["book_id"], 10, 64)

	if err != nil {
		responses.ERR(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewBookRepository(db)
	book, err := repository.FindBookById(bookID)

	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookID, err := strconv.ParseUint(params["book_id"], 10, 64)

	if err != nil {
		responses.ERR(w, http.StatusBadRequest, err)
	}

	db, err := database.Connect()
	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewBookRepository(db)
	if err = repository.Delete(bookID); err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	bookID, err := strconv.ParseUint(params["book_id"], 10, 64)
	if err != nil {
		responses.ERR(w, http.StatusBadRequest, err)
		return
	}

	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ERR(w, http.StatusBadRequest, err)
		return
	}

	var book models.Book
	if err = json.Unmarshal(bodyRequest, &book); err != nil {
		responses.ERR(w, http.StatusBadRequest, err)
		return
	}

	book.FormatBook()

	db, err := database.Connect()
	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewBookRepository(db)
	
	if err, _ = repository.Update(bookID, book); err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}

	

	responses.JSON(w, http.StatusNoContent, nil)
}

func AddBookWithFile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Reached upload\n")
	
	r.ParseMultipartForm(200 << 20) // Maximum of 200MB file allowed
	
	file, handler, err := r.FormFile("form-id")

	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}

	r.ParseMultipartForm(0)
	
	defer file.Close()

	r.ParseMultipartForm(0)
	var book models.Book

	title := r.FormValue("title")
	author := r.FormValue("author")
	subtitle := r.FormValue("subtitle")
	publisher := r.FormValue("publisher")
	publishedAt := r.FormValue("published_at")
	description := r.FormValue("description")

	// Print the form data
	fmt.Printf("Title: %s\n", title)
	fmt.Printf("Author: %s\n", author)
	fmt.Printf("Subtitle: %s\n", subtitle)
	fmt.Printf("Publisher: %s\n", publisher)
	fmt.Printf("Published At: %s\n", publishedAt)
	fmt.Printf("Description: %s\n", description)

	book.Title = r.FormValue("title")
	book.Subtitle = r.FormValue("subtitle")
	book.Description = r.FormValue("description")
	book.Authors = r.FormValue("author")
	book.Publisher = r.FormValue("publisher")
	book.Published_at = r.FormValue("published_at")
	book.Thumbnail = regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(book.Title, " ")
	book.Thumbnail = strings.TrimSpace(book.Thumbnail)
	book.FormatBook()

	db, err := database.Connect()
	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewBookRepository(db)
	book.ID, err = repository.Create(book)

	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}

	dst, err := os.Create(fmt.Sprintf("static/%s", (book.Thumbnail + ".jpg")))
	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}
	defer dst.Close()

	// 4. Copy the uploaded file to the destination file
	if _, err := io.Copy(dst, file); err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return

	}
	responses.JSON(w, http.StatusCreated, book)
	fmt.Fprintf(w, "File uploaded successfully: %s\n", handler.Filename)
}

func UpdateBookWithFile(w http.ResponseWriter, r *http.Request) {
	// Parse the multipart form (limit 200MB)
	err := r.ParseMultipartForm(200 << 20) 
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, "Unable to parse form")
		return
	}
	
	// Parse book ID
	params := mux.Vars(r)
	bookID, err := strconv.ParseUint(params["book_id"], 10, 64)
	if err != nil {
		responses.ERR(w, http.StatusBadRequest, err)
		return
	}

	// Populate book struct
	var book models.Book
	book.Title = r.FormValue("title")
	book.Description = r.FormValue("description")
	book.Authors = r.FormValue("authors")
	book.Publisher = r.FormValue("publisher")
	book.Published_at = r.FormValue("published_at")
	book.FormatBook()

	// Connect to the database
	db, err := database.Connect()
	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	// Update book in the database
	var thumbnail string
	repository := repositories.NewBookRepository(db)
	if err, thumbnail = repository.Update(bookID, book); err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}

	// Handle file upload
	file, _, err := r.FormFile("form-id")
	if err != nil {
		responses.JSON(w, http.StatusOK, "File uploaded and saved successfully")
		return
	}
	defer file.Close()

	// Remove old thumbnail file from folder
	if file != nil {
		err = os.Remove(fmt.Sprintf("static/%s", thumbnail))
		if err != nil {
			responses.ERR(w, http.StatusInternalServerError, err)
			return
		}
	}

	// Save new file in the same folder
	dst, err := os.Create(fmt.Sprintf("static/%s", thumbnail))
	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}
	defer dst.Close()

	// Copy the uploaded file to the destination
	if _, err := io.Copy(dst, file); err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}

	// Respond with success
	responses.JSON(w, http.StatusOK, "File uploaded and saved successfully")
}

