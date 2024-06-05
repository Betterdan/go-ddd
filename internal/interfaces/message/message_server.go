package message

import (
	"demo/internal/infrastructure/config"
	"demo/internal/infrastructure/mq"
	"go.uber.org/fx"
	"log"
)

type MessageServer struct {
	handlers map[string]mq.MessageHandler
}

func NewMessageServer(cfg *config.Config, factory *MessageHandlerFactory) *MessageServer {

	handlers := make(map[string]mq.MessageHandler)
	for _, handlerConfig := range cfg.KafkaConfig.KafkaTopics {
		handler, err := factory.CreateHandler(handlerConfig)
		if err != nil {
			log.Fatalf("failed to create handler: %v", err)
		}
		if handler == nil {
			continue
		}
		handlers[handlerConfig.Name] = handler

	}
	return &MessageServer{
		handlers: handlers,
	}
}

func (s *MessageServer) Start() {
	for name, handler := range s.handlers {
		go func(h mq.MessageHandler, topic string) {
			client := mq.GetKafkaClient(topic)
			if client != nil {
				err := client.Subscribe(topic, h.HandleMessage)
				if err != nil {
					log.Fatalf("failed to subscribe to topic %s: %v", topic, err)
				}
			}
		}(handler, name)
	}
}

var Module = fx.Options(
	fx.Provide(NewMessageHandlerFactory),
	fx.Provide(NewMessageServer),
)
