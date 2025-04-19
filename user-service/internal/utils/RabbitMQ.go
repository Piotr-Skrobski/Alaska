package utils

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func SetUpQueue(channel *amqp.Channel) amqp.Queue {
	q, err := channel.QueueDeclare(
		"User.Deleted",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatalf("failed to declare queue: %v", err)
	}

	return q
}
