package integration

import (
	"api/src/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"testing"
)

var user_test models.User
var token models.Token

func TestRegisterUser(t *testing.T) {
	user := models.User{
		Name:     "ewerton rosendo da silva",
		Email:    "ewertonrosendodasilva@gmail.com",
		Nick:     "RosendoSilva",
		Password: "sgfsd3232",
	}

	marshalled, err := json.Marshal(user)
	if err != nil {
		log.Fatalf("impossible to marshal user: %s", err)
	}

	request, err := http.NewRequest("POST", "http://localhost:5000/users", bytes.NewReader(marshalled))
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		t.Fatalf("failed to perform request: %v", err)
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}

	if err = json.Unmarshal(responseBody, &user_test); err != nil {
		t.Fatalf("error reading response body: %v", err)
	}

	fmt.Println("Created User:", user_test)

	expectedStatusCode := 201
	if response.StatusCode != expectedStatusCode {
		t.Errorf("Expected status code %v, but got %v", expectedStatusCode, response.StatusCode)
	}
}

func TestLoginUser(t *testing.T) {
	user := models.User{
		Email:    "ewertonrosendodasilva@gmail.com",
		Password: "sgfsd3232",
	}

	marshalled, err := json.Marshal(user)
	if err != nil {
		log.Fatalf("impossible to marshal user: %s", err)
	}

	request, err := http.NewRequest("POST", "http://localhost:5000/login", bytes.NewReader(marshalled))
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		t.Fatalf("failed to perform request: %v", err)
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}

	if err = json.Unmarshal(responseBody, &token); err != nil {
		t.Fatalf("error reading response body: %v", err)
	}

	fmt.Println("token from the user:", token.RefreshToken)

	expectedStatusCode := 202
	if response.StatusCode != expectedStatusCode {
		t.Errorf("Expected status code %v, but got %v", expectedStatusCode, response.StatusCode)
	}
}
