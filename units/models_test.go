package integration

import (
	//"api/src/config"
	"api/src/database"
	"api/src/validations"
	"api/src/models"
	"api/src/repositories"
	//"bytes"
	//"database/sql"
	//"encoding/json"
	//"fmt"
	//"io"
	//"log"
	//"net/http"
	//"os"
	//"strconv"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var user_id uint64
func TestModelCreateUser(t *testing.T) {
	
	db, err := database.Connect()
	if err != nil {
		t.Error("Error while connecting with the database: ", err)
		
	}
	defer db.Close()

	var user_body models.User 
	user_body.Name = "test_data_base"
	user_body.Nick = "testing_database"
	user_body.Email = "testing_database@gmail.com"
	hashPassword, err := validations.Hash("user.Password")
	user_body.Password = string(hashPassword)

	repository := repositories.NewUserRepository(db)
	if user_id, err = repository.Create(user_body); err != nil {
		t.Error("Error while creating an user: ", err)
	}
}

func TestModelDeleteUser(t *testing.T) {
	
	db, err := database.Connect()
	if err != nil {
		t.Error("Error while connecting with the database: ", err)
		
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	if err = repository.Delete(user_id); err != nil {
		t.Error("Error while creating an user: ", err)
	}
}