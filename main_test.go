package main

import (
	"encoding/json"
	"hello-world/config"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func init() {
	cfg = config.LoadConfig()
}

func TestSetupServer(t *testing.T) {
	server := setupServer()

	// Test server port configuration
	expectedAddr := ":" + cfg.AppPort
	if server.Addr != expectedAddr {
		t.Errorf("Expected server address %s, got %s", expectedAddr, server.Addr)
	}

	// Test route handlers
	_, ok := server.Handler.(http.Handler)
	if !ok {
		t.Fatal("Expected server handler to implement http.Handler interface")
	}

	// Test routes by making test requests
	testCases := []struct {
		path           string
		expectedStatus int
		expectedBody   string
	}{
		{"/", http.StatusOK, "Welcome"},
		{"/about", http.StatusOK, "About Us"},
		{"/ping", http.StatusOK, "pong"},
		{"/static/css/style.css", http.StatusOK, ""},
		{"/nonexistent", http.StatusNotFound, "404"},
	}

	for _, tc := range testCases {
		req := httptest.NewRequest("GET", tc.path, nil)
		w := httptest.NewRecorder()

		server.Handler.ServeHTTP(w, req)

		resp := w.Result()
		body, _ := io.ReadAll(resp.Body)

		if resp.StatusCode != tc.expectedStatus {
			t.Errorf("Path %s: expected status %d, got %d", tc.path, tc.expectedStatus, resp.StatusCode)
		}

		if tc.expectedBody != "" && !strings.Contains(string(body), tc.expectedBody) {
			t.Errorf("Path %s: expected response to contain '%s'", tc.path, tc.expectedBody)
		}
	}
}

func TestHandleHome(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	handleHome(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK; got %v", resp.StatusCode)
	}
	if !strings.Contains(string(body), "Welcome") {
		t.Error("Expected response to contain 'Welcome'")
	}
	if !strings.Contains(string(body), cfg.AppName) {
		t.Error("Expected response to contain AppName")
	}
}

func TestHandleAbout(t *testing.T) {
	req := httptest.NewRequest("GET", "/about", nil)
	w := httptest.NewRecorder()

	handleAbout(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK; got %v", resp.StatusCode)
	}
	if !strings.Contains(string(body), "About Us") {
		t.Error("Expected response to contain 'About Us'")
	}
}

func TestPing(t *testing.T) {
	req := httptest.NewRequest("GET", "/ping", nil)
	w := httptest.NewRecorder()

	ping(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK; got %v", resp.StatusCode)
	}

	var response struct {
		Message   string `json:"message"`
		Ip        string `json:"ip"`
		UserAgent string `json:"useragent"`
	}

	if err := json.Unmarshal(body, &response); err != nil {
		t.Fatalf("Failed to parse response JSON: %v", err)
	}

	if response.Message != "pong" {
		t.Errorf("Expected message 'pong'; got %v", response.Message)
	}
	if response.Ip == "" {
		t.Error("Expected IP address to be set")
	}
}

func TestLogIncomingRequest(t *testing.T) {
	req := httptest.NewRequest("GET", "/test", nil)
	req.Header.Set("User-Agent", "test-agent")

	logIncomingRequest(req) // Just ensure it doesn't panic
}
