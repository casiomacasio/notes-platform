package handler

import (
	"github.com/casiomacasio/notes-platform/services/note/internal/service"
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

	notes := router.Group("/notes")
	notes.Use(h.userIdentity)
	{
		notes.POST("/", h.createNote)
		notes.GET("/", h.getAllNotes)
		notes.GET("/:id", h.getNoteByID)
		notes.PUT("/:id", h.updateNote)
		notes.DELETE("/:id", h.deleteNote)
	}

	return router
}