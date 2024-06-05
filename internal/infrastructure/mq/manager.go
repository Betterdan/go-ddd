package mq

import "demo/internal/infrastructure/config"

var (
	MyKafkaClientMap  map[string]*KafkaMQ
	MyPulsarClientMap map[string]*MQ
)

func InitMqClient(config *config.Config) {
	MyKafkaClientMap = GetKafkaClientMap(config)
}
