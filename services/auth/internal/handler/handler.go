package handler

import (
    "github.com/casiomacasio/notes-platform/services/auth/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service     *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service:     service,
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