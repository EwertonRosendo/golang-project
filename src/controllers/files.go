package controllers

import (
	//"api/src/database"
	//"api/src/models"
	//"api/src/repositories"
	"api/src/database"
	"api/src/repositories"
	"api/src/responses"
	"fmt"
	"strings"

	//"os"
	//"api/src/services"
	//"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

// Handle the /googlebooks endpoint
func SaveFile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Reached upload\n")
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
	if err != nil{
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}

	fmt.Println("the thumbnail is for the book with the id of: ", bookID)

	// 1. Param input for multipart file upload
	r.ParseMultipartForm(200 << 20) // Maximum of 200MB file allowed

	// 2. Retrieve file from form-data
	// "form-id" is the form key that the client should use when uploading the file
	file, handler, err := r.FormFile("form-id")
	r.ParseMultipartForm(0)
	title := r.FormValue("title")
	title = strings.TrimSpace(title)
	title = strings.ReplaceAll(title," ", "-")
	fmt.Println("O TITULO É :    -----",title)
	if err != nil {
		errStr := fmt.Sprintf("Error in reading the file: %s\n", err)
		fmt.Println(errStr)
		fmt.Fprintf(w, errStr)
		return
	}
	defer file.Close()

	// Create the images directory if it doesn't exist
	err = os.MkdirAll("images", os.ModePerm)
	if err != nil {
		errStr := fmt.Sprintf("Error creating directory: %s\n", err)
		fmt.Println(errStr)
		fmt.Fprintf(w, errStr)
		return
	}
	fmt.Println("tentando remover o arquivo ", book.Thumbnail)
	err = os.Remove(fmt.Sprintf("static/%s", book.Thumbnail))
    if err != nil { 
        responses.ERR(w, http.StatusInternalServerError, err)
		return
    } 
	

	// 3. Create a new file in the images directory
	//dst, err := os.Create(fmt.Sprintf("images/%s", handler.Filename))
	dst, err := os.Create(fmt.Sprintf("static/%s", (book.Title+".jpg")))
	if err != nil {
		errStr := fmt.Sprintf("Error creating file: %s\n", err)
		fmt.Println(errStr)
		fmt.Fprintf(w, errStr)
		return
	}
	defer dst.Close()

	// 4. Copy the uploaded file to the destination file
	if _, err := io.Copy(dst, file); err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
		
	}

	// 5. Respond to the client that the file was saved successfully
	fmt.Println("change the cover image name")
	err = repository.UpdateThumbnail(book.ID, book)
	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}
	
	fmt.Fprintf(w, "File uploaded successfully: %s\n", handler.Filename)
	fmt.Println("File saved successfully")
}

func ServeStaticFiles(w http.ResponseWriter, r *http.Request) {
	// Use mux.Vars to get the filename
	vars := mux.Vars(r)
	file := vars["file"]

	// Serve the file from the ./static directory
	http.ServeFile(w, r, "./static/"+file)
}


