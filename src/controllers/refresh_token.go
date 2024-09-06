package controllers

import (
	"api/src/authentication"
	"api/src/models"
	"api/src/responses"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
)

func RefreshToken(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["user_id"], 10, 64)

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