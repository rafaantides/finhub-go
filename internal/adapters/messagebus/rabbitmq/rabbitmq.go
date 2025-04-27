package rabbitmq

import (
	"context"
	"fmt"
	"finhub-go/internal/core/ports/outbound/messagebus"
	"finhub-go/internal/utils/logger"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	log     *logger.Logger
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewRabbitMQ(user, password, host, port string) (messagebus.MessageBus, error) {
	log := logger.NewLogger("RabbitMQ")

	amqpURI := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, password, host, port)

	conn, err := amqp.Dial(amqpURI)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	log.Start("Host: %s:%s | User: %s", host, port, user)

	return &RabbitMQ{
		log:     log,
		conn:    conn,
		channel: ch,
	}, nil
}

func (r *RabbitMQ) ensureQueueExists(queueName string) error {
	if queueName == "" {
		return fmt.Errorf("queue name cannot be empty")
	}

	_, err := r.channel.QueueDeclare(
		queueName,
		true,  // durable
		false, // autoDelete
		false, // exclusive
		false, // noWait
		nil,
	)
	if err != nil {
		r.log.Error("Failed to declare queue '%s': %v", queueName, err)
	}
	return err
}

func (r *RabbitMQ) SendMessage(queueName string, body []byte) error {
	if err := r.ensureQueueExists(queueName); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.channel.PublishWithContext(
		ctx,
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	if err != nil {
		r.log.Error("Failed to send message to queue '%s': %v\nPayload: %s", queueName, err, string(body))
		return err
	}

	r.log.Info("Message sent to queue '%s': %s", queueName, string(body))
	return nil
}

func (r *RabbitMQ) ConsumeMessages(queueName string) (<-chan messagebus.Message, error) {
	if err := r.ensureQueueExists(queueName); err != nil {
		return nil, err
	}

	deliveries, err := r.channel.Consume(
		queueName,
		"",
		false, // autoAck
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		r.log.Error("Failed to consume from queue '%s': %v", queueName, err)
		return nil, err
	}

	msgChan := make(chan messagebus.Message)

	go func() {
		for d := range deliveries {
			msgChan <- &RabbitMessage{delivery: d}
		}
		close(msgChan)
	}()

	return msgChan, nil
}

func (r *RabbitMQ) DeleteQueue(queue string) error {
	_, err := r.channel.QueueDelete(queue, false, false, false)
	return err
}

func (r *RabbitMQ) Close() {
	if err := r.channel.Close(); err != nil {
		r.log.Error("Failed to close channel: %v", err)
	}
	if err := r.conn.Close(); err != nil {
		r.log.Error("Failed to close connection: %v", err)
	}
}
