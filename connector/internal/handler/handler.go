package handler

import (
	"golang-coursework/connector/config"
	v1 "golang-coursework/connector/internal/handler/http/v1"
	"golang-coursework/connector/internal/repository"
	"golang-coursework/connector/internal/service"
	"golang-coursework/connector/pkg/logger"

	"github.com/gorilla/mux"
)

type Handler struct {
	v1  *v1.Handler
	cfg *config.Config
}

func NewHandler(services *service.Services, repositories *repository.Repositories, log *logger.Logger, cfg *config.Config) *Handler {
	return &Handler{
		v1:  v1.NewHandler(services, repositories, log, cfg),
		cfg: cfg,
	}
}

func (handler *Handler) GetRouter(router *mux.Router) {
	v1Router := router.PathPrefix(handler.cfg.Server.ApiServer.ApiPrefix).Subrouter()
	handler.v1.GetHandler(v1Router)
}
