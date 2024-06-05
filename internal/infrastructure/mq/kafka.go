package mq

import (
	"context"
	"demo/internal/infrastructure/config"
	"github.com/segmentio/kafka-go"
)

type KafkaMQ struct {
	writer *kafka.Writer
	reader *kafka.Reader
}

func GetKafkaClientMap(config *config.Config) map[string]*KafkaMQ {
	mqs := make(map[string]*KafkaMQ)
	for _, handlerConfig := range config.KafkaConfig.KafkaTopics {
		mqs[handlerConfig.Name] = NewKafkaMQ(config.KafkaConfig.KafkaBrokers, handlerConfig.Name)
	}
	return mqs
}

func GetKafkaClient(topicName string) *KafkaMQ {
	if client, exists := MyKafkaClientMap[topicName]; exists {
		return client
	} else {
		return nil
	}
}

func NewKafkaMQ(brokers []string, topic string) *KafkaMQ {
	return &KafkaMQ{
		writer: kafka.NewWriter(kafka.WriterConfig{
			Brokers: brokers,
			Topic:   topic,
		}),
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers: brokers,
			Topic:   topic,
		}),
	}
}

func (k *KafkaMQ) Publish(topic string, msg Message) error {
	return k.writer.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(msg.Key),
		Value: msg.Value,
	})
}

func (k *KafkaMQ) Subscribe(topic string, handler func(Message)) error {
	for {
		m, err := k.reader.ReadMessage(context.Background())
		if err != nil {
			return err
		}
		handler(Message{
			Key:   string(m.Key),
			Value: m.Value,
		})
	}
}

func (k *KafkaMQ) Close() error {
	if err := k.writer.Close(); err != nil {
		return err
	}
	return k.reader.Close()
}
