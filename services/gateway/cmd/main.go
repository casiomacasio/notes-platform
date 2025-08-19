package main

import (
	"io"
	"net/http"
	"os"
	"time"

	"github.com/casiomacasio/notes-platform/services/gateway/internal/auth"
	"github.com/casiomacasio/notes-platform/services/gateway/internal/handler"
	"github.com/casiomacasio/notes-platform/services/gateway/internal/middleware"
	"github.com/casiomacasio/notes-platform/services/gateway/internal/redis"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading .env file: %s", err.Error())
	}

	rdb, err := redis.NewRedisClient(redis.RedisConfig{
		Host:     viper.GetString("redis.host"),
		Port:     viper.GetString("redis.port"),
		Password: os.Getenv("REDIS_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("failed to connect to Redis: %v", err)
	}
	if rdb == nil {
		logrus.Fatalf("redis client has nil value")
	}
	defer rdb.Close()
	r := gin.Default()
	r.Use(middleware.GlobalRateLimitMiddleware(rdb, 1000, time.Minute))

	authRoutes := r.Group("/auth")
	authRoutes.Use(middleware.RateLimitIPMiddleware(rdb, 10, time.Minute))

	{
		authRoutes.POST("/register", handler.Register)
		authRoutes.POST("/login", handler.Login)
		authRoutes.POST("/refresh", handler.Refresh)
		authRoutes.POST("/logout", handler.Logout)
	}

	authorized := r.Group("/")
	authorized.Use(middleware.AuthMiddleware(auth.ParseToken))
	authorized.Use(middleware.RateLimitMiddleware(rdb, 20, time.Minute))

	authorized.Any("/note/*proxyPath", func(c *gin.Context) {
		target := "http://note:8082" + c.Param("proxyPath")
		proxy(c, target)
	})

	authorized.Any("/user/*proxyPath", func(c *gin.Context) {
		target := "http://user:8083" + c.Param("proxyPath")
		proxy(c, target)
	})

	authorized.Any("/notification/*proxyPath", func(c *gin.Context) {
		target := "http://notification:8084" + c.Param("proxyPath")
		proxy(c, target)
	})

	r.Run(":8080")
}

func proxy(c *gin.Context, target string) {
	req, err := http.NewRequest(c.Request.Method, target, c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	req.Header = c.Request.Header

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	for k, v := range resp.Header {
		for _, vv := range v {
			c.Writer.Header().Add(k, vv)
		}
	}

	c.Status(resp.StatusCode)
	io.Copy(c.Writer, resp.Body)
}

func initConfig() error {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
