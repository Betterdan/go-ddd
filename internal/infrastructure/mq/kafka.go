package mq

import (
	"context"
	"github.com/segmentio/kafka-go"
)

type KafkaMQ struct {
	writer *kafka.Writer
	reader *kafka.Reader
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
