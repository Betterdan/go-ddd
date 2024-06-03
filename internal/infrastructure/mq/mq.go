package mq

type Message struct {
	Key   string
	Value []byte
}

type MQ interface {
	Publish(topic string, msg Message) error
	Subscribe(topic string, handler func(Message)) error
	Close() error
}

type MessageHandler interface {
	HandleMessage(msg Message)
}
