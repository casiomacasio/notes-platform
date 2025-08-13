package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getAllNotifications(c *gin.Context) {
	notitifications, err := h.service.Notification.GetAllNotifications()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": notitifications,
	})
}

func (h *Handler) getAllNotificationsByUserId(c *gin.Context) {
	userId, err := getUserID(c)
	if err != nil {
		return
	}

	notifications, err := h.service.Notification.GetAllNotificationsByUserId(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": notifications,
	})
}
