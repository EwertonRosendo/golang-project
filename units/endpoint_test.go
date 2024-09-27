package integration

import (
	"net/http"
	"testing"
)

func TestEndpoint(t *testing.T) {
	request, err := http.NewRequest("GET", "http://localhost:5000/books", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		t.Fatalf("failed to perform request: %v", err)
	}
	defer response.Body.Close()

	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}
	expectedStatusCode := 200

	if response.StatusCode != expectedStatusCode {
		t.Errorf("Expected status code %v, but got %v", expectedStatusCode, response.StatusCode)
	}
}
