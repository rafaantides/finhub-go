package messagebus

type Message interface {
	Body() []byte
	Ack() error
}

type MessageBus interface {
	SendMessage(queueName string, body []byte) error
	ConsumeMessages(queueName string) (<-chan Message, error)
	Close()
}
