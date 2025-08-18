package handler

import (
	"github.com/casiomacasio/notes-platform/services/user/internal/events"
	"github.com/casiomacasio/notes-platform/services/user/internal/service"
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

	router.GET("/me", h.userIdentity, h.getMe)
	router.PUT("/me", h.userIdentity, h.updateMe)
	router.GET("/:id", h.userIdentity, h.getByID)

	return router
}
