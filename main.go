package main

import (
	"fmt"
	"log"
	"net/http"
	"webserver-in-go/db"
	"webserver-in-go/handlers"
)

func main() {
	// Initialize the database
	err := db.InitDB("./data/resource.db")
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.CloseDB()

	// Register HTTP handlers
	http.HandleFunc("/resources", handlers.GetResourcesHandler)
	http.HandleFunc("/resources/add", handlers.AddResourceHandler)
	http.HandleFunc("/resources/update", handlers.UpdateResourceHandler)
	http.HandleFunc("/resources/delete", handlers.DeleteResourceHandler)

	addr := ":8080"
	fmt.Printf("Starting server on %s...", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Printf("Server failed: %s", err)
	}
}
