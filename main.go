package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Mock in-memory data
var resources = []map[string]interface{} {
	{"id": 1, "name": "Resource 1"},
	{"id": 2, "name": "Resource 2"},
}

// POST: Add a new resource
// Testing with: `curl -X POST -H "Content-Type: application/json" -d '{"name":"New Resource"}' http://localhost:8080/resources/add`
func AddResourceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var newResource map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&newResource); err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}

	// Add the new resource with an incremented ID
	newResource["id"] = len(resources) + 1
	resources = append(resources, newResource)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newResource)
}

// PUT: Update an existing resource
// Testing with: `curl -X PUT -H "Content-Type: application/json" -d '{"name":"Updated Resource"}' "http://localhost:8080/resources/update?id=1"`
func UpdateResourceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 || id > len(resources) {
		http.Error(w, "Invalid resource ID", http.StatusBadRequest)
		return
	}

	var updatedResource map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updatedResource); err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}

	// Update the resource in memory
	updatedResource["id"] = id
	resources[id-1] = updatedResource

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedResource)
}


// DELETE: Remove a resource
// Testing with: `curl -X DELETE "http://localhost:8080/resources/delete?id=2"`
func DeleteResourceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 || id > len(resources) {
		http.Error(w, "Invalid resource ID", http.StatusBadRequest)
		return
	}

	// Remove the resource from the slice
	resources = append(resources[:id-1], resources[id:]...)

	w.WriteHeader(http.StatusNoContent) // No Content response
}

// Get: Retrieve all resources
// Testing with: "curl -X GET http://localhost:8080/resources"
func GetResourcesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resources)
}

func main() {
	// GET: Retrieve all resources
	http.HandleFunc("/resources", GetResourcesHandler)

	// POST: Add a new resource
	http.HandleFunc("/resources/add", AddResourceHandler)

	// PUT: Update an existing resource
	http.HandleFunc("/resources/update", UpdateResourceHandler)

	// DELETE: Remove a resource
	http.HandleFunc("/resources/delete", DeleteResourceHandler)

	// Start the server
	addr := ":8080"
	fmt.Printf("Starting server on %s...\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Printf("Server failed: %s", err)
	}
}
