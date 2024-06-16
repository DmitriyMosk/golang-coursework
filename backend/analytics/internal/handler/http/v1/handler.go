package v1

import (
	"golang-coursework/backend/analytics/config"
	"golang-coursework/backend/analytics/internal/repository"
	"golang-coursework/backend/analytics/pkg/logger"

	"github.com/gorilla/mux"
)

type Handler struct {
	analyticsHandler *AnalyticsHandler
	cfg              *config.Config
}

func NewHandler(repositories *repository.Repositories, log *logger.Logger, cfg *config.Config) *Handler {
	return &Handler{
		analyticsHandler: NewAnalyticsHandler(repositories, log, cfg),
		cfg:              cfg,
	}
}

func (handler *Handler) GetHandler(router *mux.Router) {
	analyticsRouter := router.PathPrefix(handler.cfg.Server.AnalyticsHTTP.AnalyticsPrefix).Subrouter()
	handler.analyticsHandler.GetAnalyticsHandler(analyticsRouter)
}
