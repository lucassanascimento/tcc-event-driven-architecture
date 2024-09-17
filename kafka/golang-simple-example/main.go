package main

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	ctx := context.Background()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	broker := "localhost:9093"
	topic := "simple-example"

	go consumer(ctx, logger, broker, topic)

	for {
		producer(ctx, logger, broker, topic)
		time.Sleep(2 * time.Second)
	}
}

func producer(ctx context.Context, logger *slog.Logger, broker, topic string) {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{broker},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
	defer w.Close()

	err := w.WriteMessages(ctx,
		kafka.Message{
			Key:   []byte("Key-A"),
			Value: []byte("Hello Kafka"),
		},
	)
	if err != nil {
		logger.Error("could not write message", "erro", slog.AnyValue(err))
		panic(err)
	}
	logger.Info("Message sent successfully")
}

func consumer(ctx context.Context, logger *slog.Logger, broker, topic string) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker},
		Topic:   topic,
	})
	defer r.Close()

	logger.Info("Consumer started. Waiting for messages...")

	for {
		m, err := r.ReadMessage(ctx)
		if err != nil {
			logger.Error("could not read message", "erro", slog.AnyValue(err))
			panic(err)
		}
		logger.Info("Received message:", "key", string(m.Key), "value", string(m.Value))
	}
}
