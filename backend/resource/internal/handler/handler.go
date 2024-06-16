package handler

import (
	"golang-coursework/backend/resource/config"
	v1 "golang-coursework/backend/resource/internal/handler/http/v1"
	"golang-coursework/backend/resource/internal/repository"
	"golang-coursework/backend/resource/pkg/logger"

	"github.com/gorilla/mux"
)

type Handler struct {
	v1  *v1.Handler
	cfg *config.Config
}

func NewHandler(repositories *repository.Repositories, log *logger.Logger, cfg *config.Config) *Handler {
	return &Handler{
		v1:  v1.NewHandler(repositories, log, cfg),
		cfg: cfg,
	}
}

func (handler *Handler) GetRouter(router *mux.Router) {
	v1Router := router.PathPrefix(handler.cfg.Server.ApiServer.ApiPrefix).Subrouter()
	handler.v1.GetHandler(v1Router)
}
