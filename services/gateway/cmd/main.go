package main

import (
	"io"
	"net/http"

	"github.com/casiomacasio/notes-platform/services/gateway/internal/auth"
	"github.com/casiomacasio/notes-platform/services/gateway/internal/handler"
	"github.com/casiomacasio/notes-platform/services/gateway/internal/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/auth/register", handler.Register)
	r.POST("/auth/login", handler.Login)
	r.POST("/auth/refresh", handler.Refresh)
	r.POST("/auth/logout", handler.Logout)

	authorized := r.Group("/")
	authorized.Use(middleware.AuthMiddleware(auth.ParseToken))

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
