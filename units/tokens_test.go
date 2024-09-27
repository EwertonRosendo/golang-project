package integration

import (
	"api/src/authentication"
	"testing"
	"net/http"
	"fmt"
)

func TestCreateToken(t *testing.T) {
	_, err := authentication.CreateToken(1)
	if err != nil {
		t.Error("Error when tried create a token: ", err)
	}	
}
func TestValidateToken(t *testing.T) {
	token, err := authentication.CreateToken(1)
	if err != nil {
		t.Error("Error when tried create a token: ", err)
	}
	var r http.Request 
	r.Header = map[string][]string{
		"Authorization": {fmt.Sprintf("Bearer %s", token)},
	}
	err = authentication.TokenValidation(&r)
	if err != nil {
		t.Error("Error when tried to validate a token: ", err)
	}
}

func TestExtractUser(t *testing.T) {
	token, err := authentication.CreateToken(1)
	if err != nil {
		t.Error("Error when tried create a token: ", err)
	}
	var r http.Request 
	r.Header = map[string][]string{
		"Authorization": {fmt.Sprintf("Bearer %s", token)},
	}
	userID, err := authentication.ExtractUserID(&r)
	if err != nil {
		t.Error("Error when tried to validate a token: ", err)
	}
	if userID != 1 {
		t.Error("Error: users' id were not the same: ", err)
	}
}

