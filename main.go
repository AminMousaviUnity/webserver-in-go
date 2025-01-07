package main

import (
	"fmt"
	"io"
	"net/http"
)

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a custom handler called MyHandler!")
}

func req_info_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Method: %s\n", r.Method)
	fmt.Fprintf(w, "URL: %s\n", r.URL)
	fmt.Fprintf(w, "Header: %v\n", r.Header)
	fmt.Fprintf(w, "Content-Type: %s\n", r.Header.Get("Content-Type"))

	// Read and close the request body
	body, _ := io.ReadAll(r.Body)
	r.Body.Close()
	fmt.Fprintf(w, "Body: %s\n", body)

	// Parse and access form data
	r.ParseForm()
	fmt.Fprintf(w, "Form: %v\n", r.Form)
	fmt.Fprintf(w, "Form value 'name': %s\n", r.Form.Get("name"))
}

func main() {
	// Using HandleFunc for root address
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})

	// Use a CustomHandler which implements the ServeHTTP interface:
	// type Handler interface {
	// 	ServeHTTP(ResponseWriter, *Request)
	// }
	handler := &MyHandler{}
	http.Handle("/customHandler", handler)
	

	// Request info handler
	http.HandleFunc("/req-info", req_info_handler)

	http.ListenAndServe(":8080", nil)
}

