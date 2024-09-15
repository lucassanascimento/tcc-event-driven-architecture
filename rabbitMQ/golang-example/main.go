package main

import (
	"log"
	"log/slog"
	"time"

	"github.com/streadway/amqp"
)

func main() {
	logger := slog.Default()

	conn, err := amqp.Dial("amqp://root:root@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	// Estabelece um canal
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	// Declara uma fila (queue)
	q, err := ch.QueueDeclare(
		"example-queue", // nome da fila
		false,           // não durável
		false,           // auto delete quando não usada
		false,           // exclusiva
		false,           // sem esperar
		nil,             // argumentos adicionais
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	// Iniciar consumidor em uma goroutine
	go consumer(conn, q.Name, logger)

	// Loop do produtor que envia mensagens a cada 2 segundos
	for {
		producer(ch, q.Name, logger)
		time.Sleep(2 * time.Second)
	}
}

func producer(ch *amqp.Channel, queueName string, logger *slog.Logger) {
	// Mensagem que será enviada
	body := "Hello RabbitMQ"
	err := ch.Publish(
		"",        // exchange
		queueName, // routing key (nome da fila)
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		log.Fatalf("Failed to publish a message: %v", err)
	}
	logger.Info("Message sent successfully")
}

func consumer(conn *amqp.Connection, queueName string, logger *slog.Logger) {
	// Abre um novo canal para o consumidor
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel for consumer: %v", err)
	}
	defer ch.Close()

	// Consome mensagens da fila
	msgs, err := ch.Consume(
		queueName, // nome da fila
		"",        // nome do consumidor (deixa vazio para gerar um automaticamente)
		true,      // auto-ack (confirmação automática de recebimento)
		false,     // exclusivo
		false,     // sem espera
		false,     // sem argumentos adicionais
		nil,       // argumentos adicionais
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	logger.Info("Consumer started. Waiting for messages...")
	// Loop para processar mensagens recebidas
	for msg := range msgs {
		logger.Info("Received message", "body", string(msg.Body))
	}
}
