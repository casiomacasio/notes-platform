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

	router.GET("/", h.userIdentity, h.getAllNotifications)
	router.GET("/:id", h.userIdentity, h.getAllNotificationsByUserId)

	return router
}
