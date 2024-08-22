package controllers

import (
	"net/http"
	"net/url"
	"fmt"
	"github.com/gorilla/mux"
	"os"
	"io"
	"bytes"
	
	"encoding/json"
)

func SearchGoogleBooks(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("https://www.googleapis.com/books/v1/volumes?q=intitle:golang&maxResults=30&key=%s", os.Getenv("GOOGLE_BOOKS_API_KEY"))
	w.Write([]byte(fmt.Sprintf("Search book default route and the url to google books api is %s", url)))
}
func SearchGoogleBooksByTitle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	url := fmt.Sprintf("https://www.googleapis.com/books/v1/volumes?q=intitle:%s&maxResults=30&key=%s", url.QueryEscape(title), os.Getenv("GOOGLE_BOOKS_API_KEY"))
	request, error := http.NewRequest("GET", url, nil)
	if error != nil {
        fmt.Println(error)
    }
	request.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{}
	response, error := client.Do(request)
	
	if error != nil {
        fmt.Println(error)
    }
	
	responseBody, error := io.ReadAll(response.Body)

    if error != nil {
        fmt.Println(error)
    }
	formattedData := formatJSON(responseBody)

	//responses.JSON(w, http.StatusOK, formattedData )
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonData, err := json.Marshal(formattedData)
	if err != nil {
		http.Error(w, "Error converting response to JSON", http.StatusInternalServerError)
		return
	}
	
	w.Write(jsonData)
	defer response.Body.Close()
}

func formatJSON(data []byte) string {
    var out bytes.Buffer
    err := json.Indent(&out, data, "", "  ")

    if err != nil {
        fmt.Println(err)
    }

    d := out.Bytes()
    return string(d)
}