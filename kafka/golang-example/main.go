package main

import (
	"context"
	"log/slog"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	logger := slog.Default()

	go consumer(logger)

	for {
		producer(logger)
		time.Sleep(2 * time.Second)
	}
}

func producer(logger *slog.Logger) {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9093"},
		Topic:    "example-topic",
		Balancer: &kafka.LeastBytes{},
	})
	defer w.Close()

	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("Key-A"),
			Value: []byte("Hello Kafka"),
		},
	)
	if err != nil {
		logger.Error("could not write message: %v", err)
		panic(err)
	}
	logger.Info("Message sent successfully")
}

func consumer(logger *slog.Logger) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9093"},
		Topic:   "example-topic",
	})
	defer r.Close()

	logger.Info("Consumer started. Waiting for messages...")

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			logger.Error("could not read message: ", err)
			panic(err)
		}
		logger.Info("Received message:", "key", string(m.Key), "value", string(m.Value))
	}
}
