package v1

import (
	"golang-coursework/connector/config"
	"golang-coursework/connector/internal/repository"
	"golang-coursework/connector/internal/service"
	"golang-coursework/connector/pkg/logger"

	"github.com/gorilla/mux"
)

type Handler struct {
	connectorHandler *ConnectorHandler
	cfg              *config.Config
}

func NewHandler(services *service.Services, repositories *repository.Repositories, log *logger.Logger, cfg *config.Config) *Handler {
	return &Handler{
		connectorHandler: NewConnectorHandler(services, repositories, log, cfg),
		cfg:              cfg,
	}
}

func (handler *Handler) GetHandler(router *mux.Router) {
	connectorRouter := router.PathPrefix(handler.cfg.Server.ConnectorHTTP.ConnectorPrefix).Subrouter()
	handler.connectorHandler.GetConnectorHandler(connectorRouter)
}
