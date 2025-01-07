package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// Initialize the database
func initDB() {
	var err error
	db, err = sql.Open("sqlite3", "./resources.db")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v\n", err)
	}


	// Ensure table exists
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS resources (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL
		);`)
	if err != nil {
		log.Fatalf("Failed to create table: %v\n", err)
	}
	log.Println("Database connection successfully initialized.")
}

// GET: Retrieve all resources
func GetResourcesHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name FROM resources")
	if err != nil {
		http.Error(w, "Failed to fetch resources", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var resources []map[string]interface{}
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			http.Error(w, "Failed to parse resources", http.StatusInternalServerError)
			return
		}
		resources = append(resources, map[string]interface{}{"id": id, "name": name})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resources)
}

// POST: Add a new resource
func AddResourceHandler(w http.ResponseWriter, r *http.Request) {
	var newResource struct {
		Name string `json:"name"`
	}

	if db == nil {
		log.Println("Database connection is nil.")
		http.Error(w, "Database connection failed", http.StatusInternalServerError)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&newResource); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	result, err := db.Exec("INSERT INTO resources (name) VALUES (?)", newResource.Name)
	if err != nil {
		http.Error(w, "Failed to add resource", http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id": id,
		"name": newResource.Name,
	})
}

// PUT: Update an existing resource
func UpdateResourceHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var updatedResource struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&updatedResource); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	_, err := db.Exec("UPDATE resources SET name = ? WHERE id = ?", updatedResource.Name, id)
	if err != nil {
		http.Error(w, "Failed to update resource", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DELETE: Remove a resource
func DeleteResourceHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	_, err := db.Exec("DELETE FROM resources WHERE id = ?", id)
	if err != nil {
		http.Error(w, "Failed to delete resource", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func main() {
	// Initialize database
	initDB()
	defer db.Close()

	http.HandleFunc("/resources", GetResourcesHandler)
	http.HandleFunc("/resources/add", AddResourceHandler)
	http.HandleFunc("/resources/update", UpdateResourceHandler)
	http.HandleFunc("/resources/delete", DeleteResourceHandler)

	addr := ":8080"
	fmt.Printf("Starting server on %s...\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Printf("Server failed: %s", err)
	}
}
