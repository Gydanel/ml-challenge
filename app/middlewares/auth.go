package middlewares

import (
	"github.com/gin-gonic/gin"
	"ml-challenge/config"
	"strings"
)

func BasicAuth(c *gin.Context) {
	cfg := config.GetConfig()
	u, s, ok := c.Request.BasicAuth()
	if !ok || len(strings.TrimSpace(u)) < 1 || len(strings.TrimSpace(s)) < 1 {
		c.AbortWithStatus(401)
	}
	var user string
	var secret string
	if user = cfg.GetString("auth.user"); len(strings.TrimSpace(user)) == 0 {
		c.AbortWithStatus(401)
		return
	}
	if secret = cfg.GetString("auth.secret"); len(strings.TrimSpace(secret)) == 0 {
		c.AbortWithStatus(401)
		return
	}
	if user != u || secret != s {
		c.AbortWithStatus(401)
		return
	}
	c.Next()
}

func AuthMiddleware() gin.HandlerFunc {
	return BasicAuth
}
