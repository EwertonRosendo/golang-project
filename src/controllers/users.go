package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil{
		responses.ERR(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(bodyRequest, &user); err != nil{
		responses.ERR(w, http.StatusBadRequest, err)
		return
	}
	
	if err = user.Prepare("signup"); err != nil{
		responses.ERR(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil{
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	user.ID, err = repository.Create(user)

	if err != nil{
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, user)
}
func FindUsers(w http.ResponseWriter, r *http.Request) {

	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	db, err := database.Connect()
	 if err != nil{
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	 }
	 defer db.Close()
	 repository := repositories.NewUserRepository(db)
	 users, err := repository.SearchUsers(nameOrNick)
	 if err != nil{
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	 }
	 responses.JSON(w, http.StatusOK, users)
}
func FindUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["user_id"], 10, 64)
	if err != nil {
		responses.ERR(w, http.StatusBadRequest, err)
		return
	}
	db, err := database.Connect()
	 if err != nil{
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	 }
	 defer db.Close()

	 repository := repositories.NewUserRepository(db)
	 user, err := repository.FindUserById(userID)

	 if err != nil{
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, user)
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["user_id"], 10, 64)
	if err != nil {
		responses.ERR(w, http.StatusBadRequest,  err)
		return
	}

	tokenUserID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.ERR(w, http.StatusUnauthorized,  err)
		return
	}
	if userID != tokenUserID {
		responses.ERR(w, http.StatusForbidden, errors.New("you can not update another user"))
		return
	}

	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil{
		responses.ERR(w, http.StatusBadRequest, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(bodyRequest, &user); err!= nil {
		responses.ERR(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare("update"); err != nil {
		responses.ERR(w, http.StatusBadRequest, err)
		return
	}
	db, err := database.Connect()
	 if err != nil{
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	 }
	 defer db.Close()

	 repository := repositories.NewUserRepository(db)
	 if err = repository.Update(userID, user); err != nil{
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	 }

	 responses.JSON(w, http.StatusNoContent, nil)
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["user_id"], 10, 64)

	if err != nil{
		responses.ERR(w, http.StatusBadRequest, err)
	}

	tokenUserID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.ERR(w, http.StatusUnauthorized,  err)
		return
	}
	if userID != tokenUserID {
		responses.ERR(w, http.StatusForbidden, errors.New("you can not delete another user"))
		return
	}

	db, err := database.Connect()
	 if err != nil{
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	 }
	 defer db.Close()

	 repository := repositories.NewUserRepository(db)
	 if err = repository.Delete(userID); err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	 }

	 responses.JSON(w, http.StatusNoContent, nil)

}