package handler

import (
	"github.com/casiomacasio/notes-platform/services/auth/internal/events"
	"github.com/casiomacasio/notes-platform/services/auth/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service  *service.Service
	eventBus events.EventBus
}

func NewHandler(service *service.Service, eventBus events.EventBus) *Handler {
	return &Handler{
		service:  service,
		eventBus: eventBus,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.POST("/register", h.register)
	router.POST("/sign-in", h.signIn)
	router.POST("/refresh", h.refresh)
	router.POST("/logout", h.logout)

	return router
}
