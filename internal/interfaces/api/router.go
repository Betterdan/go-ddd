package api

import (
	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	Handler *Handler
}

func NewRouter(config *RouterConfig) *gin.Engine {
	router := gin.Default()

	userHandler := config.Handler.UserHandler

	// 注册用户路由
	userGroup := router.Group("/users")
	{
		userGroup.GET("/:id", userHandler.GetUser)
	}

	return router
}

func NewRouterConfig(handler *Handler) *RouterConfig {
	return &RouterConfig{Handler: handler}
}
