package integration

import (
	//"io"
	"net/http"
	"testing"
	//"fmt"
)

func TestEndpoint(t *testing.T) {
	// Set up a new request.
	request, err := http.NewRequest("GET", "http://localhost:5000/books", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	// Use the default client to send the request.
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		t.Fatalf("failed to perform request: %v", err)
	}
	defer response.Body.Close()

	// Read the response body.
	//responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}

	// Print the response body as a string for debugging purposes.
	//fmt.Println(string(responseBody))

	// Expected status code.
	expectedStatusCode := 200

	// Check if the status code matches.
	if response.StatusCode != expectedStatusCode {
		t.Errorf("Expected status code %v, but got %v", expectedStatusCode, response.StatusCode)
	}
}
