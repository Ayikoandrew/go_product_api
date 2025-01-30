package test // use your actual package name

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"main.go/middleware"
)

func TestLogReport(t *testing.T) {
	// Create a buffer to capture log output
	buf := &bytes.Buffer{}
	logger := log.New(buf, "", 0)

	// Create a test handler that we'll wrap with our middleware
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Create our middleware
	middleware := middleware.LogReport(logger, nil) // passing nil for handlers.Product since it's not used

	// Create a test server with our middleware wrapped around our test handler
	ts := httptest.NewServer(middleware(testHandler))
	defer ts.Close()

	// Make a test request
	_, err := http.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	// Check if the log contains expected information
	logOutput := buf.String()
	if len(logOutput) == 0 {
		t.Error("Expected log output but got none")
	}
}
