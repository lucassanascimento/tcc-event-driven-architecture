package main

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/streadway/amqp"
)

func Producer(ctx context.Context, logger *slog.Logger, ch *amqp.Channel, event Event) error {
	body, err := json.Marshal(event)
	if err != nil {
		logger.Error("Failed to marshal event", slog.String("error", err.Error()))
		return err
	}

	// Publish the message on the "events" exchange
	err = ch.Publish(
		"events", // exchange
		"",       // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		logger.Error("Failed to publish message", slog.String("error", err.Error()))
		return err
	}
	return nil
}
