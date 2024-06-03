package message

import (
	"demo/internal/infrastructure/config"
	"demo/internal/infrastructure/mq"
	"go.uber.org/fx"
	"log"
)

type MessageServer struct {
	handlers map[string]mq.MessageHandler
	mqs      map[string]mq.MQ
}

func NewMessageServer(cfg *config.Config, factory *MessageHandlerFactory) *MessageServer {

	handlers := make(map[string]mq.MessageHandler)
	mqs := make(map[string]mq.MQ)
	for _, handlerConfig := range cfg.KafkaConfig.KafkaTopics {
		handler, err := factory.CreateHandler(handlerConfig)
		if err != nil {
			log.Fatalf("failed to create handler: %v", err)
		}
		handlers[handlerConfig.Name] = handler
		mqs[handlerConfig.Name] = mq.NewKafkaMQ(cfg.KafkaConfig.KafkaBrokers, handlerConfig.Name)
	}
	return &MessageServer{
		handlers: handlers,
		mqs:      mqs,
	}
}

func (s *MessageServer) Start() {
	for name, handler := range s.handlers {
		mqImpl := s.mqs[name]
		go func(h mq.MessageHandler, topic string, mq mq.MQ) {
			err := mqImpl.Subscribe(topic, h.HandleMessage)
			if err != nil {
				log.Fatalf("failed to subscribe to topic %s: %v", topic, err)
			}
		}(handler, name, mqImpl)
	}
}

var Module = fx.Options(
	fx.Provide(NewMessageHandlerFactory),
	fx.Provide(NewMessageServer),
)
