package app

import (
	"database/sql"
	"fmt"
	"golang-coursework/backend/analytics/config"
	"golang-coursework/backend/analytics/internal/handler"
	"golang-coursework/backend/analytics/internal/repository"
	"golang-coursework/backend/analytics/pkg/logger"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	log    *logger.Logger
	cfg    *config.Config
	db     *sql.DB
	server *http.Server
}

func NewApp(cfg *config.Config, log *logger.Logger) (*App, error) {

	var db *sql.DB
	sqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DB.HostDB,
		cfg.DB.PortDB,
		cfg.DB.UserDB,
		cfg.DB.PasswordDB,
		cfg.DB.NameDB)
	db, err := sql.Open("postgres", sqlInfo)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	repositories := repository.NewRepositories(db)

	analyticsHandlers := handler.NewHandler(repositories, log, cfg)

	analyticsRouter := mux.NewRouter()

	analyticsHandlers.GetRouter(analyticsRouter)

	analyticsServer := &http.Server{
		Addr:    cfg.Server.AnalyticsHTTP.AnalyticsHost + ":" + cfg.Server.AnalyticsHTTP.AnalyticsPort,
		Handler: analyticsRouter,
	}

	return &App{
		log:    log,
		cfg:    cfg,
		db:     db,
		server: analyticsServer,
	}, nil
}

func (app *App) Run() error {
	err := app.server.ListenAndServe()
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	return nil
}

func (app *App) Close() error {

	if err := app.server.Close(); err != nil {
		return fmt.Errorf(err.Error())
	}
	return app.db.Close()
}
