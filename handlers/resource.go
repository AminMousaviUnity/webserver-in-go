package handlers

import (
	"encoding/json"
	"net/http"
	"webserver-in-go/db"
	"webserver-in-go/models"
)

// GET: Retrieve all resources
func GetResourcesHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT id, name FROM resources")
	if err != nil {
		http.Error(w, "Failed to fetch resources", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var resources []models.Resource
	for rows.Next() {
		var resource models.Resource
		if err := rows.Scan(&resource.ID, &resource.Name); err != nil {
			http.Error(w, "Failed to parse resources", http.StatusInternalServerError)
			return
		}
		resources = append(resources, resource)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resources)
}

// POST: Add a new resource
func AddResourceHandler(w http.ResponseWriter, r *http.Request) {
	var resource models.Resource
	if err := json.NewDecoder(r.Body).Decode(&resource); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	result, err := db.DB.Exec("INSERT INTO resources (name) VALUES (?)", resource.Name)
	if err != nil {
		http.Error(w, "Failed to add resource", http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	resource.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resource)
}

// PUT: Update an existing resource
func UpdateResourceHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var resource models.Resource
	if err := json.NewDecoder(r.Body).Decode(&resource); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	_, err := db.DB.Exec("UPDATE resources SET name = ? WHERE id = ?", resource.Name, id)
	if err != nil {
		http.Error(w, "Failed to update resource", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DELETE: Remove a resource
func DeleteResourceHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	_, err := db.DB.Exec("DELETE FROM resources WHERE id = ?", id)
	if err != nil {
		http.Error(w, "Failed to delete resource", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}