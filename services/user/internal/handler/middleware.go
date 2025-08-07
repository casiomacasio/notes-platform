package handler

import (
	"errors"
	"net/http"
	"github.com/gin-gonic/gin"
)

const userCtx = "userId"

var (
	ErrInvalidToken  = errors.New("invalid token")
	ErrTokenExpired  = errors.New("token expired")
)

func (h Handler) userIdentity(c *gin.Context) {
	token, err := c.Cookie("access_token")
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "missing access token cookie")
		return
	}

	userID, err := h.service.Authorization.ParseToken(token)
	if err != nil {
		if errors.Is(err, ErrTokenExpired) {
			c.Header("Token-Expired", "true")
			newErrorResponse(c, http.StatusUnauthorized, "token expired")
			return
		}
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userID)
}

func getUserID(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id not found")
	}
	userID, ok := id.(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id invalid")
		return 0, errors.New("user id invalid")
	}
	return userID, nil
}