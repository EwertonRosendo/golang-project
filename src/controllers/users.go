package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil{
		responses.ERR(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(bodyRequest, &user); err != nil{
		responses.ERR(w, http.StatusBadRequest, err)
		return
	}
	
	if err = user.Prepare(); err != nil{
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
	w.Write([]byte("searching User by id!"))
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating User!"))
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("deleting User!"))
}