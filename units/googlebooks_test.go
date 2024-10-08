package integration

import (
	"api/src"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGoogleDefaultEndpoint(t *testing.T) {
	req, err := http.NewRequest("GET", "/googlebooks", nil)
	if err != nil {
		t.Fatal(err)
	}

	newRecorder := httptest.NewRecorder()

	router.Generate().ServeHTTP(newRecorder, req)

	statusCode := 200
	if newRecorder.Result().StatusCode != statusCode {
		t.Errorf("TestInfoRequest() test returned an unexpected result: got %v want %v", newRecorder.Result().StatusCode, statusCode)
	}
}

func TestGoogleWithParamsEndpoint(t *testing.T) {
	req, err := http.NewRequest("GET", "/googlebooks/ruby", nil)
	if err != nil {
		t.Fatal(err)
	}

	newRecorder := httptest.NewRecorder()

	router.Generate().ServeHTTP(newRecorder, req)

	statusCode := 200
	if newRecorder.Result().StatusCode != statusCode {
		t.Errorf("TestInfoRequest() test returned an unexpected result: got %v want %v", newRecorder.Result().StatusCode, statusCode)
	}
}

func TestGoogleWithWrongUrl(t *testing.T) {
	req, err := http.NewRequest("GET", "/googlebooks/", nil)
	if err != nil {
		t.Fatal(err)
	}

	newRecorder := httptest.NewRecorder()

	router.Generate().ServeHTTP(newRecorder, req)

	statusCode := 404
	if newRecorder.Result().StatusCode != statusCode {
		t.Errorf("TestInfoRequest() test returned an unexpected result: got %v want %v", newRecorder.Result().StatusCode, statusCode)
	}
}
