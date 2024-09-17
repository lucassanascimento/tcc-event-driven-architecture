package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/streadway/amqp"
)

func NotifyConsumer(ctx context.Context, logger *slog.Logger, conn *amqp.Connection, queueName string) error {
	// Opens a new channel for the consumer
	ch, err := conn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open a channel for consumer: %w", err)
	}
	defer ch.Close()
	logger.Info("NotifyConsumer started. Waiting for events...")

	// Declare the queue
	q, err := ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // autoDelete
		false,     // exclusive
		false,     // noWait
		nil,       // args
	)
	if err != nil {
		return fmt.Errorf("failed to declare queue: %w", err)
	}

	// Link the queue to the exchange
	if err := ch.QueueBind(
		q.Name,   // queue name
		"",       // routing key
		"events", // exchange
		false,    // noWait
		nil,      // args
	); err != nil {
		return fmt.Errorf("failed to bind queue: %w", err)
	}

	// Start consuming messages
	msgs, err := ch.Consume(
		q.Name,     // queue
		"notifier", // consumer tag
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		return fmt.Errorf("failed to register a consumer: %w", err)
	}

	// Message Processing Loop
	for {
		select {
		case <-ctx.Done(): // Checks if the context has been canceled
			logger.Info("Context canceled, stopping NotifyConsumer...")
			return ctx.Err()

		case msg := <-msgs: // Processes received messages
			var event Event
			if err := json.Unmarshal(msg.Body, &event); err != nil {
				logger.Error("Failed to unmarshal message", slog.String("error", err.Error()))
				continue // Logs the error, but continues processing the next messages
			}

			logger.Info(
				"NotifyConsumer", "Match", event.MatchID,
				"Score"+event.TeamA, event.ScoreTeamA,
				"Score"+event.TeamB, event.ScoreTeamB,
			)
		}
	}
}
