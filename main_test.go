package main

import (
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMainPageHandler(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate the main page handler
		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		tmpl.Execute(w, nil)
	}))
	defer server.Close()

	// Make a request to the test server
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	// Check status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	// Check that the response contains expected HTML elements
	bodyStr := string(body)
	expectedElements := []string{
		"<title>ksabanty.com</title>",
		"<h1>Welcome to ksabanty.com</h1>",
		"htmx.org@1.9.10",
		"hx-get=\"/api/message\"",
		"hx-target=\"#msg\"",
	}

	for _, element := range expectedElements {
		if !strings.Contains(bodyStr, element) {
			t.Errorf("Expected response to contain %q, but it didn't", element)
		}
	}
}

func TestMessageHandler(t *testing.T) {
	// Create a test request
	req := httptest.NewRequest("GET", "/api/message", nil)

	// Create a response recorder
	w := httptest.NewRecorder()

	// Call the message handler
	messageHandler(w, req)

	// Check status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Check content type
	contentType := w.Header().Get("Content-Type")
	if contentType != "text/html" {
		t.Errorf("Expected Content-Type %q, got %q", "text/html", contentType)
	}

	// Check response body
	expectedBody := "<div><p>Hello, this is a message from the server!</p></div>"
	if w.Body.String() != expectedBody {
		t.Errorf("Expected body %q, got %q", expectedBody, w.Body.String())
	}
}

func TestMessageHandlerWithDifferentMethods(t *testing.T) {
	methods := []string{"POST", "PUT", "DELETE"}

	for _, method := range methods {
		t.Run(method, func(t *testing.T) {
			req := httptest.NewRequest(method, "/api/message", nil)
			w := httptest.NewRecorder()

			messageHandler(w, req)

			// The handler should still work regardless of method
			if w.Code != http.StatusOK {
				t.Errorf("Expected status code %d for %s method, got %d", http.StatusOK, method, w.Code)
			}
		})
	}
}

func TestMessageHandlerResponseStructure(t *testing.T) {
	req := httptest.NewRequest("GET", "/api/message", nil)
	w := httptest.NewRecorder()

	messageHandler(w, req)

	body := w.Body.String()

	// Check that the response contains expected HTML structure
	if !strings.Contains(body, "<div>") {
		t.Error("Expected response to contain <div> tag")
	}

	if !strings.Contains(body, "<p>") {
		t.Error("Expected response to contain <p> tag")
	}

	if !strings.Contains(body, "Hello, this is a message from the server!") {
		t.Error("Expected response to contain the message text")
	}
}

// Integration test that simulates the full flow
func TestHTMXIntegration(t *testing.T) {
	// Create a test server with both handlers
	mux := http.NewServeMux()

	// Register the main page handler
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		tmpl.Execute(w, nil)
	})

	// Register the message handler
	mux.HandleFunc("/api/message", messageHandler)

	server := httptest.NewServer(mux)
	defer server.Close()

	// Test the main page
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatalf("Failed to get main page: %v", err)
	}
	resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Main page returned status %d, expected %d", resp.StatusCode, http.StatusOK)
	}

	// Test the API endpoint
	resp, err = http.Get(server.URL + "/api/message")
	if err != nil {
		t.Fatalf("Failed to get API message: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("API message returned status %d, expected %d", resp.StatusCode, http.StatusOK)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read API response: %v", err)
	}

	if !strings.Contains(string(body), "Hello, this is a message from the server!") {
		t.Error("API response doesn't contain expected message")
	}
}
