package handler

import (
	"net/http"
	"strconv"

	"github.com/casiomacasio/notes-platform/services/user/internal/model"
	"github.com/gin-gonic/gin"
)

type getUserByIdResponses struct {
	Data model.User `json:"data"`
}

func (h *Handler) getMe(c *gin.Context) {
	userId, err := getUserID(c)
	if err != nil {
		return
	}
	user, err := h.service.User.GetUser(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getUserByIdResponses{Data: user})
}

func (h *Handler) getByID(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}
	user, err := h.service.User.GetUser(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getUserByIdResponses{Data: user})
}

func (h *Handler) updateMe(c *gin.Context) {
	userId, err := getUserID(c)
	if err != nil {
		return
	}
	var input model.UpdateUserInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadGateway, err.Error())
	}
	err = h.service.User.UpdateUser(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]string{
		"status": "updated",
	})
}