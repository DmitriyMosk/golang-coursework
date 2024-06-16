package v1

import (
	"golang-coursework/backend/resource/config"
	"golang-coursework/backend/resource/internal/repository"
	"golang-coursework/backend/resource/pkg/logger"

	"github.com/gorilla/mux"
)

type Handler struct {
	resourceHandler *ResourceHandler
	cfg             *config.Config
}

func NewHandler(repositories *repository.Repositories, log *logger.Logger, cfg *config.Config) *Handler {
	return &Handler{
		resourceHandler: NewResourceHandler(repositories, log, cfg),
		cfg:             cfg,
	}
}

func (handler *Handler) GetHandler(router *mux.Router) {
	resourceRouter := router.PathPrefix(handler.cfg.Server.ResourceHTTP.ResourcePrefix).Subrouter()
	handler.resourceHandler.GetResourceHandler(resourceRouter)
}
