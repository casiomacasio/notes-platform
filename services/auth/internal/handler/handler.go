package handler

import (
    "github.com/casiomacasio/notes-platform/services/auth/internal/service"
    "github.com/casiomacasio/notes-platform/services/auth/internal/events"
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
	auth := router.Group("/auth")
	{
		auth.POST("/register", h.register)
		auth.POST("/sign-in", h.signIn)
		auth.POST("/refresh", h.refresh)
		auth.POST("/logout", h.logout)
	} 
	return router
}