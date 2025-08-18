package handler

import (
	"github.com/casiomacasio/notes-platform/services/note/internal/events"
	"github.com/casiomacasio/notes-platform/services/note/internal/service"
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

	router.POST("/", h.createNote)
	router.GET("/", h.getAllNotes)
	router.GET("/:id", h.getNoteByID)
	router.PUT("/:id", h.updateNote)
	router.DELETE("/:id", h.deleteNote)

	return router
}
