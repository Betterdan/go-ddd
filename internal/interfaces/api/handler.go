package api

import "demo/internal/interfaces/api/handler"

type HandlerList struct {
	UserHandler *handler.UserHandler
	// 可以添加更多的 handler
}

func NewHandlerList(
	userHandler *handler.UserHandler,
) *HandlerList {
	return &HandlerList{
		UserHandler: userHandler,
	}
}
