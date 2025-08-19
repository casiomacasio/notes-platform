package middleware

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"net/http"
	"strconv"
	"time"
)

var ctx = context.Background()

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

func RateLimitMiddleware(redisClient *redis.Client, limit int, window time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := getUserID(c)
		if err != nil {
			c.Abort()
			return
		}

		key := fmt.Sprintf("rate_limit:%d", userID)
		count, err := redisClient.Incr(ctx, key).Result()
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, "redis error")
			c.Abort()
			return
		}

		if count == 1 {
			redisClient.Expire(ctx, key, window)
		}

		if count > int64(limit) {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "rate limit exceeded",
			})
			return
		}

		c.Next()
	}
}

func GlobalRateLimitMiddleware(redisClient *redis.Client, limit int, window time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := "rate_limit:global"
		count, err := redisClient.Incr(ctx, key).Result()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
			return
		}

		if count == 1 {
			redisClient.Expire(ctx, key, window)
		}

		if count > int64(limit) {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "global rate limit exceeded",
			})
			return
		}

		c.Next()
	}
}

func RateLimitIPMiddleware(redisClient *redis.Client, limit int, window time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		key := fmt.Sprintf("rate_limit:ip:%s", ip)

		count, err := redisClient.Incr(ctx, key).Result()
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, "redis error")
			c.Abort()
			return
		}

		if count == 1 {
			redisClient.Expire(ctx, key, window)
		}

		if count > int64(limit) {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "rate limit exceeded (IP)",
			})
			return
		}

		c.Next()
	}
}
