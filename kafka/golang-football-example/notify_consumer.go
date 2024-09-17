package main

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/segmentio/kafka-go"
)

func NotifyConsumer(ctx context.Context, logger *slog.Logger, broker, topic string) error {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker},
		Topic:   topic,
	})
	defer r.Close()

	logger.Info("NotifyConsumer started. Waiting for events...")
	for {
		select {
		case <-ctx.Done():
			logger.Info("Context canceled, stopping consumer...")
			return ctx.Err() // Returns the error associated with the context cancellation

		default:

			m, err := r.ReadMessage(ctx)
			if err != nil {
				logger.Error("could not read message: ", slog.String("error", err.Error()))
				return err
			}
			var event Event
			unErr := json.Unmarshal(m.Value, &event)
			if unErr != nil {
				logger.Error("unmarshal", slog.String("error", unErr.Error()))
				return unErr
			}

			logger.Info("NotifyConsumer", "Match", event.MatchID, "Score"+event.TeamA, event.ScoreTeamA, "Score"+event.TeamB, event.ScoreTeamB)
		}
	}
}
