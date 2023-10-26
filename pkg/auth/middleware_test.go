package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var authMiddleware = NewMiddleware(TYPE_API_KEY)
var nextHandler = func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
var authHandler = authMiddleware.AuthHandler(nextHandler)

func TestAuthenticated(t *testing.T) {
	req, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("X-API-KEY", authMiddleware.apiKey)

	rr := httptest.NewRecorder()
	authHandler(rr, req)

	expectedStatus := http.StatusOK
	if rr.Code != expectedStatus {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, expectedStatus)
	}
}

func TestWrongApiKey(t *testing.T) {
	req, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("X-API-KEY", "wrong key")

	rr := httptest.NewRecorder()
	authHandler(rr, req)

	expectedStatus := http.StatusUnauthorized
	if rr.Code != expectedStatus {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, expectedStatus)
	}
}

func TestMissingHeader(t *testing.T) {
	req, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	authHandler(rr, req)

	expectedStatus := http.StatusUnauthorized
	if rr.Code != expectedStatus {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, expectedStatus)
	}
}
