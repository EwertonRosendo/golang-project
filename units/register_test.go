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

var user_test models.User // Renaming to 'user_test' to store the created user for further tests
var token models.Token

// TestRegisterUser handles the registration of a user and stores the result in 'user_test'.
func TestRegisterUser(t *testing.T) {
	// Define a new user
	user := models.User{
		Name:  "ewerton rosendo da silva",
		Email: "ewertonrosendodasilva@gmail.com",
		Nick:  "RosendoSilva",
		Password: "sgfsd3232",
	}

	// Marshal the user struct into JSON
	marshalled, err := json.Marshal(user)
	if err != nil {
		log.Fatalf("impossible to marshal user: %s", err)
	}

	// Create a new POST request to register the user
	request, err := http.NewRequest("POST", "http://localhost:5000/users", bytes.NewReader(marshalled))
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	// Send the request using the default HTTP client
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		t.Fatalf("failed to perform request: %v", err)
	}
	defer response.Body.Close()

	// Read the response body
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}

	// Unmarshal the response into 'user_test'
	if err = json.Unmarshal(responseBody, &user_test); err != nil {
		t.Fatalf("error reading response body: %v", err)
	}

	// Print the user for debugging purposes
	fmt.Println("Created User:", user_test)

	// Validate the status code (201 Created is expected)
	expectedStatusCode := 201
	if response.StatusCode != expectedStatusCode {
		t.Errorf("Expected status code %v, but got %v", expectedStatusCode, response.StatusCode)
	}
}

func TestLoginUser(t *testing.T) {
	// Define a new user
	user := models.User{
		Email: "ewertonrosendodasilva@gmail.com",
		Password: "sgfsd3232",
	}

	// Marshal the user struct into JSON
	marshalled, err := json.Marshal(user)
	if err != nil {
		log.Fatalf("impossible to marshal user: %s", err)
	}

	// Create a new POST request to register the user
	request, err := http.NewRequest("POST", "http://localhost:5000/login", bytes.NewReader(marshalled))
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	// Send the request using the default HTTP client
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		t.Fatalf("failed to perform request: %v", err)
	}
	defer response.Body.Close()

	// Read the response body
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}

	// Unmarshal the response into 'user_test'
	if err = json.Unmarshal(responseBody, &token); err != nil {
		t.Fatalf("error reading response body: %v", err)
	}

	// Print the user for debugging purposes
	fmt.Println("token from the user:", token.RefreshToken)

	// Validate the status code (201 Created is expected)
	expectedStatusCode := 202
	if response.StatusCode != expectedStatusCode {
		t.Errorf("Expected status code %v, but got %v", expectedStatusCode, response.StatusCode)
	}
}

/*
// TestUpdateUser updates the user based on the result from TestRegisterUser
func TestUpdateUser(t *testing.T) {
	// Example Bearer token (replace this with a valid token)

	user := models.User{
		Name:  "ewerton rosendo da silva1",
		Email: "ewertonrosendodasilva1@gmail.com",
		Nick:  "RosendoSilva1",
	}

	// Marshal the user struct into JSON
	marshalled, err := json.Marshal(user)
	if err != nil {
		log.Fatalf("impossible to marshal user: %s", err)
	}

	// Create a new PUT request to update the user
	request, err := http.NewRequest("PUT", fmt.Sprintf("http://localhost:5000/users/%d", user_test.ID), bytes.NewReader(marshalled))
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	// Add the Authorization header with the Bearer token
	request.Header.Set("Authorization", "Bearer "+token.RefreshToken)

	// Optionally, set the content type if the API expects JSON
	request.Header.Set("Content-Type", "application/json")

	// Send the request using the default HTTP client
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		t.Fatalf("failed to perform request: %v", err)
	}
	defer response.Body.Close()

	// Read and print the response body
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}
	fmt.Println("Update Response:", string(responseBody))

	// Validate the status code (200 OK is expected)
	expectedStatusCode := 200
	if response.StatusCode != expectedStatusCode {
		t.Errorf("Expected status code %v, but got %v", expectedStatusCode, response.StatusCode)
	}
}


func TestDeleteUser(t *testing.T) {
	// Example Bearer token (replace this with a valid token)
	

	// Ensure that the user_test contains the registered user
	if user_test.ID == 0 {
		t.Fatalf("no user found to delete, please check TestRegisterUser")
	}

	// Create a new DELETE request to remove the user
	request, err := http.NewRequest("DELETE", fmt.Sprintf("http://localhost:5000/users/%d", user_test.ID), nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	// Add the Authorization header with the Bearer token
	request.Header.Set("Authorization", "Bearer "+token.RefreshToken)

	// Send the request using the default HTTP client
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		t.Fatalf("failed to perform request: %v", err)
	}
	defer response.Body.Close()

	// Read and print the response body
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}
	fmt.Println("Delete Response:", string(responseBody))

	// Validate the status code (200 OK or 204 No Content is expected)
	expectedStatusCode := 200 // Could also be 204 depending on API design
	if response.StatusCode != expectedStatusCode {
		t.Errorf("Expected status code %v, but got %v", expectedStatusCode, response.StatusCode)
	}
}
*/

func TestCloseTests(t *testing.T) {


	
	// Create a new POST request to register the user
	request, err := http.NewRequest("GET", "http://localhost:5000/clean_database", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	// Send the request using the default HTTP client
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		t.Fatalf("failed to perform request: %v", err)
	}
	defer response.Body.Close()

	expectedStatusCode := 204
	if response.StatusCode != expectedStatusCode {
		t.Errorf("Expected status code %v, but got %v", expectedStatusCode, response.StatusCode)
	}
}