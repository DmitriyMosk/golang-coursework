package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Resource struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var resources = []Resource{
	{ID: "1", Name: "Resource 1"},
	{ID: "2", Name: "Resource 2"},
}

func GetResources(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"data": resources,
		"_links": map[string]interface{}{
			"self": map[string]string{
				"href": r.RequestURI,
			},
		},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetResource(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range resources {
		if item.ID == params["id"] {
			response := map[string]interface{}{
				"data": item,
				"_links": map[string]interface{}{
					"self": map[string]string{
						"href": r.RequestURI,
					},
					"all_resources": map[string]string{
						"href": "/api/v1/resources",
					},
				},
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Resource not found"})
}

func CreateResource(w http.ResponseWriter, r *http.Request) {
	var resource Resource
	_ = json.NewDecoder(r.Body).Decode(&resource)
	resource.ID = "3" // Assign a new ID in a real application
	resources = append(resources, resource)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resource)
}

func UpdateResource(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range resources {
		if item.ID == params["id"] {
			var resource Resource
			_ = json.NewDecoder(r.Body).Decode(&resource)
			resource.ID = item.ID
			resources[index] = resource
			json.NewEncoder(w).Encode(resource)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Resource not found"})
}

func DeleteResource(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range resources {
		if item.ID == params["id"] {
			resources = append(resources[:index], resources[index+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Resource not found"})
}
