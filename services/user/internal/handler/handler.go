package handler

import (
    "github.com/casiomacasio/notes-platform/services/user/internal/service"
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
	users := router.Group("/users")
	users.Use(h.userIdentity)
	{
		users.GET("/me", h.getMe)
		users.PUT("/me", h.updateMe)
		users.GET("/:id", h.getByID)
	}
	return router
}