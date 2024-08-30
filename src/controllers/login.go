package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/validations"
	"api/src/authentication"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request){

	bodyRequest , err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERR(w, http.StatusUnprocessableEntity, err)
		return
	}
	var user models.User
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		responses.ERR(w, http.StatusBadRequest, err)
	}
	
	db, err := database.Connect()
	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repositories.NewUserRepository(db)

	userFromDataBase, err := repository.FindByEmail(user.Email)

	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}
	
	if err = validations.ValidatePassword(userFromDataBase.Password, user.Password);err != nil{
		responses.ERR(w, http.StatusUnauthorized, err)
		return
	}
	authentication.CreateAccessTokenCookie(w, r, userFromDataBase.ID)
	
	token, err := authentication.CreateToken(userFromDataBase.ID)
	if err != nil{
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}
	var user_refresh_token models.Token
	user_refresh_token.RefreshToken = token

	responses.JSON(w, http.StatusAccepted, user_refresh_token)

}