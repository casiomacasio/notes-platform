package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrTokenExpired = errors.New("token expired")
)

func getUserID(c *gin.Context) (int, error) {
	userIDStr := c.GetHeader("X-User-ID")
	if userIDStr == "" {
		newErrorResponse(c, http.StatusUnauthorized, "user id not found")
		return 0, errors.New("user id not found")
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "user id invalid")
		return 0, errors.New("user id invalid")
	}
	return userID, nil
}
