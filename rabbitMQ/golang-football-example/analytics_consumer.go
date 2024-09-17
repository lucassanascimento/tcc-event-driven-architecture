package main

import (
	"context"
	"encoding/json"
	"log"
	"log/slog"

	"github.com/streadway/amqp"
)

func AnalyticsConsumer(ctx context.Context, logger *slog.Logger, conn *amqp.Connection, queueName string) error {
	// Opens a new channel for the consumer
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel for consumer: %v", err)
	}
	defer ch.Close()
	logger.Info("AnalyticsConsumer started. Waiting for events...")

	// Declare the queue
	q, err := ch.QueueDeclare(
		queueName, // name da fila
		true,      // durable
		false,     // auto-delete
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		logger.Error("Failed to declare queue", slog.String("error", err.Error()))
		return err
	}

	// Link the queue to the exchange
	err = ch.QueueBind(
		q.Name,   // nome da fila
		"",       // routing key
		"events", // exchange
		false,
		nil,
	)
	if err != nil {
		logger.Error("Failed to bind queue", slog.String("error", err.Error()))
		return err
	}

	// Start consuming messages
	msgs, err := ch.Consume(
		q.Name,      // nome da fila
		"analytics", // consumer name
		true,        // auto-ack
		false,       // exclusive
		false,       // no-local
		false,       // no-wait
		nil,         // arguments
	)
	if err != nil {
		logger.Error("Failed to register a consumer", slog.String("error", err.Error()))
		return err
	}

	var totalScore, teamATotalScore, teamBTotalScore int
	// Function to update counters
	updateScores := func(event Event) {
		totalScore += event.ScoreTeamA + event.ScoreTeamB // Total Goals in the Match
		teamATotalScore += event.ScoreTeamA               // Total goals Team A
		teamBTotalScore += event.ScoreTeamB               // Total goals Team B

		logger.Info("AnalyticsConsumer:",
			"Match", event.MatchID,
			"Total Goals", totalScore,
			"Total Goals: "+event.TeamA, teamATotalScore,
			"Total Goals: "+event.TeamB, teamBTotalScore)
	}

	// Message Processing Loop
	for {
		select {
		case <-ctx.Done(): // Checks if the context has been canceled
			logger.Info("Context canceled, stopping consumer...")
			return ctx.Err() // Returns the error associated with the context cancellation

		case m := <-msgs: // Processes received messages
			// Converts the event from JSON to a struct in Go
			var event Event
			if err := json.Unmarshal(m.Body, &event); err != nil {
				logger.Error("unmarshal", slog.String("error", err.Error()))
				return err
			}

			// Updates the counters upon receiving a new event
			updateScores(event)
		}
	}
}
