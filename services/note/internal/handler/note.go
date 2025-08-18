package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/casiomacasio/notes-platform/services/note/internal/model"
	"github.com/gin-gonic/gin"
)

type getNoteByIdResponse struct {
	Data model.Note `json:"data"`
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

	n := model.Notification{
		UserId:    userId,
		Type:      "Note Created",
		Title:     "Note was created!",
		Message:   fmt.Sprintf("Hello, you have made a note: %s", input.Title),
		Status:    "unread",
		CreatedAt: time.Now(),
	}
	h.eventBus.Publish("notifications", n)
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

	n := model.Notification{
		UserId:    userId,
		Type:      "Note Updated",
		Title:     "Note was Updated!",
		Message:   fmt.Sprintf("Hello, you have updated a note: %s", *input.Title),
		Status:    "unread",
		CreatedAt: time.Now(),
	}
	h.eventBus.Publish("notifications", n)

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

	n := model.Notification{
		UserId:    userId,
		Type:      "Note Deleted",
		Title:     "Note was deleted!",
		Message:   "Hello, you have deleted a note",
		Status:    "unread",
		CreatedAt: time.Now(),
	}
	h.eventBus.Publish("notifications", n)

	c.JSON(http.StatusOK, map[string]string{
		"status": "deleted",
	})
}
