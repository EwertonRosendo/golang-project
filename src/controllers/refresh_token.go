package controllers

import (
	"api/src/authentication"
	"api/src/models"
	"api/src/responses"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func RefreshToken(w http.ResponseWriter, r *http.Request) {
    // Accessing the headers
    authHeader := r.Header.Get("Authorization")
    fmt.Println("Authorization Header:", authHeader)
    
    // Example of accessing a custom header, e.g., "X-Custom-Header"
    customHeader := r.Header.Get("X-Custom-Header")
    fmt.Println("Custom Header:", customHeader)

    params := mux.Vars(r)
    userID, err := strconv.ParseUint(params["user_id"], 10, 64)
    fmt.Println("User ID:", userID)

    if err != nil {
        responses.ERR(w, http.StatusInternalServerError, err)
        return
    }

    refresh_token, err := authentication.RefreshToken(r, userID)

    if err != nil {
        responses.ERR(w, http.StatusInternalServerError, err)
        return
    }

    refresh_token_obj := models.Token{RefreshToken: refresh_token}

    responses.JSON(w, http.StatusAccepted, refresh_token_obj)
}

func Teste(w http.ResponseWriter, r *http.Request) {
    responses.JSON(w, http.StatusAccepted, models.User{Nick: "TESTE"} )
}

