package message

import (
	"demo/internal/application/service"
	"demo/internal/infrastructure/config"
	"demo/internal/infrastructure/mq"
	"demo/internal/interfaces/message/handler"
	"fmt"
)

type MessageHandlerFactory struct {
	userService *service.UserService
	// 其他全局依赖
}

func NewMessageHandlerFactory(userService *service.UserService) *MessageHandlerFactory {
	return &MessageHandlerFactory{
		userService: userService,
		// 初始化其他全局依赖
	}
}

func (f *MessageHandlerFactory) CreateHandler(cfg config.TopicConfig) (mq.MessageHandler, error) {
	switch cfg.Name {
	case "user":
		return handler.NewUserMessageHandler(f.userService), nil
	// 添加其他处理器类型
	default:
		return nil, fmt.Errorf("unknown handler type: %s", cfg.Handler)
	}
}
