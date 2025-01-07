package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// MyHandler is a custom HTTP handler that implements the http.Handler interface
type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a custom handler called MyHandler!")
}

// RequestInfoHandler handles requests and displays request information
func RequestInfoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)

	// Display request details
	fmt.Fprintf(w, "Method: %s\n", r.Method)
	fmt.Fprintf(w, "URL: %s\n", r.URL)
	fmt.Fprintf(w, "Header: %v\n", r.Header)
	fmt.Fprintf(w, "Content-Type: %s\n", r.Header.Get("Content-Type"))

	// Read and safely close the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	fmt.Fprintf(w, "Body: %s\n", body)

	// Parse and display form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Form: %v\n", r.Form)
	fmt.Fprintf(w, "Form value 'name': %s\n", r.Form.Get("name"))
}

// ResponseInfoHandler sends a JSON response
func ResponseInfoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Prepare a response map and encode it as JSON
	response := map[string]interface{} {
		"status": "success",
		"message": "Hello, world!",
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

func main() {
	// Root route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})

	// Custom handler
	handler := &MyHandler{}
	http.Handle("/customHandler", handler)
	
	// Request info handler
	http.HandleFunc("/req-info", RequestInfoHandler)

	// Response info handler
	http.HandleFunc("/res-info", ResponseInfoHandler)

	// Static file server
	staticDir := "./static"
	if _, err := os.Stat(staticDir); os.IsNotExist(err) {
		fmt.Printf("Static directory '%s' does not exist. Please create it and add files.\n", staticDir)
		return
	}
	fs := http.FileServer(http.Dir(staticDir))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	// Start the server
	addr := ":8080"
	fmt.Printf("Starting server on %s...\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Printf("Server failed: %s", err)
	}
}

