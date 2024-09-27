package integration

import (
	"api/src/authentication"
	"api/src/validations"
	"testing"
	"net/http"
	"api/src/services"
	"fmt"
	"bytes"
	"encoding/json"
)

func TestCreateHash(t *testing.T) {

	_, err := validations.Hash("passwordpassword")
	if err != nil {
		t.Error("Error when trying to hash a password: ", err)
	}
	
}

func TestValidatePassword(t *testing.T) {
	password, err := validations.Hash("passwordpassword")
	if err != nil {
		t.Error("Error when trying to hash a password: ", err)
	}
	err = validations.ValidatePassword(string(password), "passwordpassword")
	if err != nil {
		t.Error("Error when tried to validate a password with a hash: ", err)
	}
}

func TestUserAuthenticy(t *testing.T) {
	token, err := authentication.CreateToken(1)
	if err != nil {
		t.Error("Error when tried to create a token: ", err)
	}

	userData := map[string]interface{}{
		"user_id": 1,
		"email": "exampleUser",
		"password": "examplePassword",
	}

	bodyBytes, err := json.Marshal(userData)
	if err != nil {
		t.Error("Error when marshaling user data: ", err)
	}

	req, err := http.NewRequest("POST", "/reviews/1/user?user_id=1", bytes.NewBuffer(bodyBytes)) // POST is typical for auth requests
	if err != nil {
		t.Error("Error creating request: ", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	req.Header.Set("Content-Type", "application/json")

	err = services.ValidateUser(req)

	if err != nil {
		t.Error("Error validating user: ", err)
	}
}
