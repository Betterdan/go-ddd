package api

import (
	"demo/internal/infrastructure/config"
	"demo/internal/interfaces/api/middleware"
	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	handlerList *HandlerList
	config      *config.Config
}

func NewRouterConfig(handlerList *HandlerList, config *config.Config) *RouterConfig {
	return &RouterConfig{handlerList: handlerList, config: config}
}

func NewRouter(rc *RouterConfig) *gin.Engine {
	router := gin.Default()

	//加载中间件
	router.Use(middleware.ConfigMiddleware(rc.config))

	userHandler := rc.handlerList.UserHandler
	// 注册用户路由
	userGroup := router.Group("/users")
	{
		userGroup.GET("/:id", userHandler.GetUser)
		userGroup.GET("/test", userHandler.Test)
	}

	return router
}
