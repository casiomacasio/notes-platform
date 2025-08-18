package handler

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) { forwardAuth(c, "http://auth:8081/register", true) }
func Login(c *gin.Context)    { forwardAuth(c, "http://auth:8081/login", true) }
func Refresh(c *gin.Context)  { forwardAuth(c, "http://auth:8081/refresh", true) }
func Logout(c *gin.Context)   { forwardAuth(c, "http://auth:8081/logout", true) }

func forwardAuth(c *gin.Context, target string, setCookies bool) {
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

	if setCookies {
		for _, cookie := range resp.Cookies() {
			http.SetCookie(c.Writer, cookie)
		}
	}

	c.Status(resp.StatusCode)
	io.Copy(c.Writer, resp.Body)
}
