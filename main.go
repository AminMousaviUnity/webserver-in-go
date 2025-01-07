package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a custom handler called MyHandler!")
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
	http.ListenAndServe(":8080", nil)
}

