package handler

import (
	"demo/internal/application/service"
	"demo/internal/infrastructure/logger"
	"demo/internal/infrastructure/mq"
	"go.uber.org/zap"
)

type UserMessageHandler struct {
	userService *service.UserService
}

func NewUserMessageHandler(userService *service.UserService) mq.MessageHandler {
	return &UserMessageHandler{userService: userService}
}

func (h *UserMessageHandler) HandleMessage(msg mq.Message) {
	logger.Info("Received message: key=%s, value=%s", zap.String(msg.Key, string(msg.Value)))
}
