package services

import (
	"encoding/json"
	"log"

	"github.com/Piotr-Skrobski/Alaska/shared-events/events"
	amqp "github.com/rabbitmq/amqp091-go"
)

type EventPublisher struct {
	channel *amqp.Channel
}

func NewEventPublisher(channel *amqp.Channel) *EventPublisher {
	return &EventPublisher{
		channel: channel,
	}
}

func (ep *EventPublisher) PublishUserDeleted(event events.UserDeleted) error {
	body, err := json.Marshal(event)
	if err != nil {
		return err
	}

	err = ep.channel.Publish(
		"",
		"User.Deleted",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		log.Printf("failed to publish user deleted event: %v", err)
		return err
	}
	return nil
}
