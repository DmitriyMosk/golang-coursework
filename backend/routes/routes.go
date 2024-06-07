package routes

/*

"golang-coursework/cmd/config"
	"golang-coursework/cmd/etl"
	"golang-coursework/cmd/jira"

*/

import (
	"golang-coursework/cmd/config"
	"golang-coursework/cmd/jira"
	"golang-coursework/handlers"
	"golang-coursework/middleware"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type PageInfo struct {
	CurrentPage   int `json:"currentPage"`
	ProjectsCount int `json:"projectsCount"`
	PageCount     int `json:"pageCount"`
}

func SetupRoutes(r *mux.Router, db *gorm.DB, jiraClient *jira.JiraClient, config *config.Config) {
	api := r.PathPrefix("").Subrouter()
	api.HandleFunc("/api/v1/resources", handlers.GetResources).Methods("GET")
	api.HandleFunc("/api/v1/resources/{id}", handlers.GetResource).Methods("GET")
	api.HandleFunc("/api/v1/resources", handlers.CreateResource).Methods("POST")
	api.HandleFunc("/api/v1/resources/{id}", handlers.UpdateResource).Methods("PUT")
	api.HandleFunc("/api/v1/resources/{id}", handlers.DeleteResource).Methods("DELETE")

	// Применение middleware для аутентификации
	api.Use(middleware.JwtAuthentication)

	// Применение middleware для обработки таймаутов
	api.Use(middleware.TimeoutMiddleware)

	apiV2 := r.PathPrefix("").Subrouter()
	apiV2.HandleFunc("/api/v2/updateProject", handlers.UpdateProjectV2).Methods("GET")
	apiV2.HandleFunc("/api/v2/projects", handlers.GetProjectsV2).Methods("GET")

	apiV2.Use(middleware.TimeoutMiddleware)

}
