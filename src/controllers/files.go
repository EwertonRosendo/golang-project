package controllers

import (
	"net/http"
	"github.com/gorilla/mux"
)

func ServeStaticFiles(w http.ResponseWriter, r *http.Request) {
	// Use mux.Vars to get the filename
	vars := mux.Vars(r)
	file := vars["file"]

	// Serve the file from the ./static directory
	http.ServeFile(w, r, "./static/"+file)
}