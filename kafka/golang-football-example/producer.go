package main

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/segmentio/kafka-go"
)

func Producer(ctx context.Context, logger *slog.Logger, event Event, broker, topic string) error {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{broker},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
	defer w.Close()

	eventBytes, e := json.Marshal(event)
	if e != nil {
		logger.Error("failure to marshal event", slog.String("error", e.Error()))
		return e
	}
	err := w.WriteMessages(ctx,
		kafka.Message{
			Key:   []byte("football-example"),
			Value: eventBytes,
		},
	)
	if err != nil {
		logger.Error("could not write message:", slog.String("error", err.Error()))
		return err
	}
	return nil
}
