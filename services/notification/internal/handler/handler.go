package handler

import (
	"github.com/casiomacasio/notes-platform/services/notification/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	notifications := router.Group("/notifications")
	notifications.Use(h.userIdentity)
	{
		notifications.GET("/", h.getAllNotifications)
		notifications.GET("/:id", h.getAllNotificationsByUserId)
	}

	return router
}