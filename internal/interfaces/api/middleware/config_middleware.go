package middleware

import (
	"demo/internal/infrastructure/config"
	"github.com/gin-gonic/gin"
)

func ConfigMiddleware(config *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("config", config)
		c.Next()
	}
}
