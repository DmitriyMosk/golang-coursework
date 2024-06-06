package routes

import (
	"golang-coursework/handlers"
	"golang-coursework/middleware"

	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router) {
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/resources", handlers.GetResources).Methods("GET")
	api.HandleFunc("/resources/{id}", handlers.GetResource).Methods("GET")
	api.HandleFunc("/resources", handlers.CreateResource).Methods("POST")
	api.HandleFunc("/resources/{id}", handlers.UpdateResource).Methods("PUT")
	api.HandleFunc("/resources/{id}", handlers.DeleteResource).Methods("DELETE")

	// Применение middleware для аутентификации
	//api.Use(middleware.JwtAuthentication)

	// Применение middleware для обработки таймаутов
	api.Use(middleware.TimeoutMiddleware)
}

//golang-coursework
