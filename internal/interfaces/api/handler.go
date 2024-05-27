package api

import "demo/internal/interfaces/api/handler"

type Handler struct {
	UserHandler *handler.UserHandler
	// 可以添加更多的 handler
}

func NewHandler(
	userHandler *handler.UserHandler,
) *Handler {
	return &Handler{
		UserHandler: userHandler,
	}
}
