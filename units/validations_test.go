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
	// Step 1: Create a token for the request
	token, err := authentication.CreateToken(1)
	if err != nil {
		t.Error("Error when tried to create a token: ", err)
	}

	// Step 2: Set up a body for the request (e.g., user data in JSON format)
	userData := map[string]interface{}{
		"user_id": 1,
		"email": "exampleUser",
		"password": "examplePassword",
	}

	// Marshal the user data into a JSON byte array
	bodyBytes, err := json.Marshal(userData)
	if err != nil {
		t.Error("Error when marshaling user data: ", err)
	}

	// Step 3: Create the HTTP request
	req, err := http.NewRequest("POST", "/reviews/1/user?user_id=1", bytes.NewBuffer(bodyBytes)) // POST is typical for auth requests
	if err != nil {
		t.Error("Error creating request: ", err)
	}

	// Step 4: Set the Authorization header with the token
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	// Step 5: Set the Content-Type header to JSON
	req.Header.Set("Content-Type", "application/json")

	// Step 6: Call your validation service
	err = services.ValidateUser(req) // Assuming ValidateUser accepts *http.Request

	// Step 7: Check for errors
	if err != nil {
		t.Error("Error validating user: ", err)
	}
}


