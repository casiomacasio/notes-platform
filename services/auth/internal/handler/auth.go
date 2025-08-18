package handler

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/casiomacasio/notes-platform/services/auth/internal/model"
	"github.com/casiomacasio/notes-platform/services/auth/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var isProd = os.Getenv("APP_ENV") == "production"

func setAuthCookies(c *gin.Context, accessToken, refreshToken, refreshTokenId string) {
	sameSite := http.SameSiteLaxMode
	secure := false

	if isProd {
		sameSite = http.SameSiteNoneMode
		secure = true
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   secure,
		SameSite: sameSite,
		MaxAge:   15 * 60,
	})

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   secure,
		SameSite: sameSite,
		MaxAge:   30 * 24 * 60 * 60,
	})

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "refresh_token_id",
		Value:    refreshTokenId,
		Path:     "/",
		HttpOnly: true,
		Secure:   secure,
		SameSite: sameSite,
		MaxAge:   30 * 24 * 60 * 60,
	})
}

func (h *Handler) register(c *gin.Context) {
	var input model.CreateUserRequest
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid body")
		return
	}
	id, err := h.service.Authorization.CreateUser(input)
	if err != nil {
		if errors.Is(err, repository.ErrUsernameExists) {
			newErrorResponse(c, http.StatusConflict, "Username already in use")
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	userEvent := model.UserCreatedEvent{
		UserId: id,
		Name:   input.Name,
		Email:  input.Email,
	}
	h.eventBus.Publish("user_events", userEvent)

	n := model.Notification{
		UserId:    id,
		Type:      "user_signed_up",
		Title:     "Welcome to the platform!",
		Message:   "Hello, you have successfully signed up.",
		Status:    "unread",
		CreatedAt: time.Now(),
	}
	h.eventBus.Publish("notifications", n)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.service.GetUser(input.Email, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	tokenID, refreshToken, err := h.service.GenerateRefreshToken(user.Id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	accessToken, err := h.service.GenerateToken(user.Id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	n := model.Notification{
		UserId:    user.Id,
		Type:      "user_signed_in",
		Title:     "Welcome back!",
		Message:   "Hello, you have successfully signed in.",
		Status:    "unread",
		CreatedAt: time.Now(),
	}
	setAuthCookies(c, accessToken, refreshToken, tokenID)
	h.eventBus.Publish("notifications", n)
	c.JSON(http.StatusOK, map[string]string{
		"message": "logged in successfully",
	})
}

func (h *Handler) refresh(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "no refresh token cookie")
		return
	}
	refreshTokenId, err := c.Cookie("refresh_token_id")
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "refresh_token_id is missed")
		return
	}
	refreshTokenUUID, err := uuid.Parse(refreshTokenId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid refresh token format")
		return
	}
	userId, err := h.service.GetUserByRefreshTokenAndRefreshTokenId(refreshToken, refreshTokenUUID)
	if err != nil {
		if errors.Is(err, repository.ErrRefreshTokenExpired) {
			c.Header("RefreshToken-Expired", "true")
			newErrorResponse(c, http.StatusUnauthorized, "refresh_token expired, must re-login")
			return
		}
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	newAccessToken, err := h.service.GenerateToken(userId)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	newTokenId, newRefreshToken, err := h.service.Authorization.GenerateRefreshToken(userId)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "failed to generate new refresh token")
		return
	}

	setAuthCookies(c, newAccessToken, newRefreshToken, newTokenId)

	c.JSON(http.StatusOK, map[string]string{
		"message": "token refreshed",
	})
}

func (h *Handler) logout(c *gin.Context) {
	refreshTokenId, err := c.Cookie("refresh_token_id")
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "refresh_token_id is missed")
		return
	}
	refreshTokenUUID, err := uuid.Parse(refreshTokenId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid refresh token format")
		return
	}
	userId, err := h.service.Authorization.GetUserByRefreshTokenId(refreshTokenUUID)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("couldn't get the user: %s", err.Error()))
		return
	}
	if err := h.service.Authorization.RevokeRefreshToken(refreshTokenUUID); err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("couldn't revoke refresh token: %s", err.Error()))
	}

	deleteCookie := func(name string) {
		http.SetCookie(c.Writer, &http.Cookie{
			Name:     name,
			Value:    "",
			Path:     "/",
			HttpOnly: true,
			Secure:   isProd,
			SameSite: http.SameSiteLaxMode,
			MaxAge:   -1,
		})
	}

	deleteCookie("access_token")
	deleteCookie("refresh_token")
	deleteCookie("refresh_token_id")
	n := model.Notification{
		UserId:    userId,
		Type:      "user_logged_out",
		Title:     "Goodbye!",
		Message:   "You successfully logged out",
		Status:    "unread",
		CreatedAt: time.Now(),
	}
	h.eventBus.Publish("notifications", n)
	c.JSON(http.StatusOK, gin.H{"message": "logged out successfully"})
}
