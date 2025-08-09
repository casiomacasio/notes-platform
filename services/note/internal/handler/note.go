package handler

import (
	"net/http"
	"strconv"
	"github.com/casiomacasio/notes-platform/services/note/internal/events"
	"github.com/casiomacasio/notes-platform/services/note/internal/model"
	"github.com/gin-gonic/gin"
)

type getNoteByIdResponse struct {
	Data model.Note `json:"data"`
}

func (h *Handler) CreateNotification(c *gin.Context) {
    var input model.NotificationInput
    if err := c.BindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    h.eventBus.Publish("notifications", events.Event{
        Type: input.Type,
        Data: input.Data,
    })

    c.JSON(http.StatusOK, gin.H{"status": "notification published"})
}

func (h *Handler) createNote(c *gin.Context) {
	userId, err := getUserID(c)
	if err != nil {
		return
	}

	var input model.CreateNoteInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	noteId, err := h.service.Note.CreateNote(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": noteId,
	})
}

func (h *Handler) getAllNotes(c *gin.Context) {
	userId, err := getUserID(c)
	if err != nil {
		return
	}

	notes, err := h.service.Note.GetAllNotes(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": notes,
	})
}

func (h *Handler) getNoteByID(c *gin.Context) {
	userId, err := getUserID(c)
	if err != nil {
		return
	}

	noteId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid note id")
		return
	}

	note, err := h.service.Note.GetNoteByID(userId, noteId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getNoteByIdResponse{Data: note})
}

func (h *Handler) updateNote(c *gin.Context) {
	userId, err := getUserID(c)
	if err != nil {
		return
	}

	noteId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid note id")
		return
	}

	var input model.UpdateNoteInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Note.UpdateNote(userId, noteId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"status": "updated",
	})
}

func (h *Handler) deleteNote(c *gin.Context) {
	userId, err := getUserID(c)
	if err != nil {
		return
	}

	noteId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid note id")
		return
	}

	err = h.service.Note.DeleteNote(userId, noteId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"status": "deleted",
	})
}