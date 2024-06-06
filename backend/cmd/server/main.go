package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"golang-coursework/routes"
	"golang-coursework/utils"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./config/")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	requiredConfigs := []string{"port", "resourseTimeout", "analyticsTimeout", "dbUser", "dbPassword", "dbHost", "dbPort", "dbName", "jwtSecret"}
	for _, config := range requiredConfigs {
		if !viper.IsSet(config) {
			log.Fatalf("The %s is not configured", config)
		}
	}
}

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
	initConfig()
	setupLogger()

	r := mux.NewRouter()
	routes.SetupRoutes(r)

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

	logrus.Infof("Starting server on port %s", port)
	log.Fatal(srv.ListenAndServe())
}
