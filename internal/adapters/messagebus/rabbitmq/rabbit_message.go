package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMessage struct {
	delivery amqp.Delivery
}

func (rm *RabbitMessage) Body() []byte {
	return rm.delivery.Body
}

func (rm *RabbitMessage) Ack() error {
	return rm.delivery.Ack(false)
}