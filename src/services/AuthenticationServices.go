package services

import (
	"api/src/authentication"
	
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ValidateUser(r *http.Request) (error) {

	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["user_id"], 10, 64)
	if err != nil {
		return err
	}
	
	tokenUserID, err := authentication.ExtractUserID(r)
	if err != nil {
		return err
	}
	
	if userID != tokenUserID {
		return err
	}
	return nil
}
