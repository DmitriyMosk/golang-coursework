package server

import (
	"golang-coursework/pkg/config"
	"golang-coursework/pkg/handlers"
	"golang-coursework/pkg/logging"
	"net/http"

	"github.com/gorilla/mux"
)

func Run() error {
	cfg, err := config.LoadConfig("configs/config.yml")
	if err != nil {
		return err
	}

	logger := logging.NewLogger()

	r := mux.NewRouter()
	r.HandleFunc("/api/issues", handlers.FetchIssues).Methods("GET")
	r.HandleFunc("/api/projects", handlers.FetchProjects).Methods("GET")

	srv := &http.Server{
		Addr:    ":" + cfg.Connector.ServerPort,
		Handler: r,
	}

	logger.Println("Starting server on port", cfg.Connector.ServerPort)
	return srv.ListenAndServe()
}
