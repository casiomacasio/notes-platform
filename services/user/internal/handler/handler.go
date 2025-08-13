package handler

import (
    "github.com/casiomacasio/notes-platform/services/user/internal/service"
    "github.com/casiomacasio/notes-platform/services/user/internal/events"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
    eventBus events.EventBus
}

func NewHandler(service *service.Service, eventBus events.EventBus) *Handler {
	return &Handler{
		service: service,
        eventBus: eventBus,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()	
	users := router.Group("/users")
	users.Use(h.userIdentity)
	{
		users.GET("/me", h.getMe)
		users.PUT("/me", h.updateMe)
		users.GET("/:id", h.getByID)
	}
	return router
}