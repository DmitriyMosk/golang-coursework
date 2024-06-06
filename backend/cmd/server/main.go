package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"golang-coursework/cmd/config"
	"golang-coursework/cmd/jira"
	"golang-coursework/database"
	"golang-coursework/routes"
	"golang-coursework/utils"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func setupLogger() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	file, err := os.OpenFile("logs/logs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logrus.SetOutput(file)
	} else {
		logrus.SetOutput(os.Stdout)
	}

	errorFile, err := os.OpenFile("logs/err_logs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		mw := io.MultiWriter(os.Stderr, errorFile)
		logrus.SetOutput(mw)
	} else {
		logrus.SetOutput(os.Stderr)
	}

	logrus.SetLevel(logrus.InfoLevel)
}

func main() {
	var err error

	config.GConfig, err = config.InitConfig("D:/go_cursejob/golang-coursework/backend/config")

	if err != nil {
		logrus.Fatalf("Failed to load config: %v", err)
		fmt.Println("Can not read CFG")
	}

	setupLogger()

	DB, err := database.InitDB(config.GConfig)

	jiraClient := jira.NewJiraClient(config.GConfig.ProgramSettings.JiraURL)

	r := mux.NewRouter()
	routes.SetupRoutes(r, DB, jiraClient, &config.GConfig)

	// Применяем GzipMiddleware ко всем маршрутам
	r.Use(utils.GzipMiddleware)

	port := viper.GetString("port")
	if port == "" {
		port = "8000"
	}

	srv := &http.Server{
		Handler:      r,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	http.Handle("/", r)
	logrus.Infof("Server is running on port %d", config.GConfig.Port)
	log.Fatal(srv.ListenAndServe(), r)
}
