package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	// Kafka broker address from environment variable
	kafkaBroker := "kafka:9092"

	// Add a delay to wait for Kafka to be ready
	time.Sleep(10 * time.Second)

	// Kafka writer configuration
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{kafkaBroker},
		Topic:   "example-topic",
	})

	// Write a message
	err := writer.WriteMessages(context.Background(), kafka.Message{
		Value: []byte("Hello Kafka"),
	})
	if err != nil {
		log.Fatal("could not write message:", err)
	}

	// Kafka reader configuration
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{kafkaBroker},
		GroupID: "go-consumer-group",
		Topic:   "example-topic",
	})

	// Read messages
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Printf("could not read message %v\n", err)
			continue
		}
		fmt.Printf("received: %s\n", string(msg.Value))
	}
}
