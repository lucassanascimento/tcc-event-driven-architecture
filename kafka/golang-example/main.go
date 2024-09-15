package main

import (
    "context"
    "fmt"
    "log"
    "github.com/segmentio/kafka-go"
)

func main(){
	go producer()
	go consumer()
}

func producer() {
    // Configuração do writer do Kafka
    w := kafka.NewWriter(kafka.WriterConfig{
        Brokers: []string{"localhost:9092"},
        Topic:   "example-topic",
        Balancer: &kafka.LeastBytes{},
    })
    defer w.Close()

    // Enviar uma mensagem
    err := w.WriteMessages(context.Background(),
        kafka.Message{
            Key:   []byte("Key-A"),
            Value: []byte("Hello Kafka"),
        },
    )
    if err != nil {
        log.Fatalf("could not write message %v", err)
    }
    fmt.Println("Message sent successfully")
}


func consumer() {
	// Configuração do reader do Kafka
	r := kafka.NewReader(kafka.ReaderConfig{
			Brokers: []string{"localhost:9092"},
			Topic:   "example-topic",
			GroupID: "example-group",
	})
	defer r.Close()

	for {
			m, err := r.ReadMessage(context.Background())
			if err != nil {
					log.Fatalf("could not read message %v", err)
			}
			fmt.Printf("Received message: %s = %s\n", string(m.Key), string(m.Value))
	}
}
