package app

import (
	"database/sql"
	"fmt"
	"golang-coursework/backend/resource/config"
	resourceHandler "golang-coursework/backend/resource/internal/handler"
	"golang-coursework/backend/resource/internal/repository"
	"golang-coursework/backend/resource/pkg/logger"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	log *logger.Logger
	cfg *config.Config

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

	handlers := resourceHandler.NewHandler(repositories, log, cfg)

	router := mux.NewRouter()

	handlers.GetRouter(router)

	resourceServer := &http.Server{
		Addr:    cfg.Server.ResourceHTTP.ResourceHost + ":" + cfg.Server.ResourceHTTP.ResourcePort,
		Handler: router,
	}

	return &App{
		log:    log,
		cfg:    cfg,
		db:     db,
		server: resourceServer,
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
