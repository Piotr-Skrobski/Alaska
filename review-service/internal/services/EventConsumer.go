package services

import (
	"encoding/json"
	"log"

	"github.com/Piotr-Skrobski/Alaska/shared-events/events"
	amqp "github.com/rabbitmq/amqp091-go"
)

type EventConsumer struct {
	channel        *amqp.Channel
	reviewService  *ReviewService
}

func NewEventConsumer(channel *amqp.Channel, reviewService *ReviewService) *EventConsumer {
	return &EventConsumer{
		channel:       channel,
		reviewService: reviewService,
	}
}

func (ec *EventConsumer) ConsumeUserDeletedEvents() {
	queue, err := ec.channel.QueueDeclare(
		"User.Deleted", // queue name
		true,           // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		log.Fatalf("failed to declare queue: %v", err)
	}

	msgs, err := ec.channel.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		log.Fatalf("failed to register consumer: %v", err)
	}

	go func() {
		for msg := range msgs {
			var event events.UserDeleted
			if err := json.Unmarshal(msg.Body, &event); err != nil {
				log.Printf("failed to unmarshal UserDeleted event: %v", err)
				continue
			}

			log.Printf("Received UserDeleted event for user_id: %d", event.UserID)
			if err := ec.reviewService.DeleteReviewsByUserID(event.UserID); err != nil {
				log.Printf("failed to delete reviews for user_id %d: %v", event.UserID, err)
			} else {
				log.Printf("Successfully deleted reviews for user_id: %d", event.UserID)
			}
		}
	}()
}

func (ec *EventConsumer) ConsumeMovieDeletedEvents() {
	queue, err := ec.channel.QueueDeclare(
		"Movie.Deleted", // queue name
		true,            // durable
		false,           // delete when unused
		false,           // exclusive
		false,           // no-wait
		nil,             // arguments
	)
	if err != nil {
		log.Fatalf("failed to declare queue: %v", err)
	}

	msgs, err := ec.channel.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		log.Fatalf("failed to register consumer: %v", err)
	}

	go func() {
		for msg := range msgs {
			var event events.MovieDeletedEvent
			if err := json.Unmarshal(msg.Body, &event); err != nil {
				log.Printf("failed to unmarshal MovieDeletedEvent: %v", err)
				continue
			}

			log.Printf("Received MovieDeleted event for movie_id: %s", event.MovieID)
			if err := ec.reviewService.DeleteReviewsByMovieID(event.MovieID); err != nil {
				log.Printf("failed to delete reviews for movie_id %s: %v", event.MovieID, err)
			} else {
				log.Printf("Successfully deleted reviews for movie_id: %s", event.MovieID)
			}
		}
	}()
}
