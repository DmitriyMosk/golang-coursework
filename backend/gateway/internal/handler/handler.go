package handler

import (
	"golang-coursework/backend/gateway/config"
	v1 "golang-coursework/backend/gateway/internal/handler/http/v1"
	"golang-coursework/backend/gateway/pkg/logger"
	"net/http"
)

type Handler struct {
	v1  *v1.Proxy
	cfg *config.Config
}

func NewHandler(log *logger.Logger, cfg *config.Config) *Handler {
	return &Handler{
		v1:  v1.NewProxy(log, cfg),
		cfg: cfg,
	}
}

func (handler *Handler) GetProxy(serveMux *http.ServeMux) {
	handler.v1.NewProxy(serveMux, handler.cfg.Server.ApiServer.ApiPrefix)
}
