package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

const userCtx = "userId"

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrTokenExpired = errors.New("token expired")
)

func AuthMiddleware(parseToken func(accessToken string) (string, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("access_token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing access token cookie"})
			c.Abort()
			return
		}

		userID, err := parseToken(token)
		if err != nil {
			if errors.Is(err, ErrTokenExpired) {
				c.Header("Token-Expired", "true")
				c.JSON(http.StatusUnauthorized, gin.H{"error": "token expired"})
				c.Abort()
				return
			}
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Request.Header.Set("X-User-ID", userID)

		c.Next()
	}
}
