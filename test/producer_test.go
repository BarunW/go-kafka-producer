package test

import (
	"context"
	"log"
	"testing"

	"github.com/segmentio/kafka-go"
)

func TestProducer(t *testing.T) {
    // Define Kafka reader configuration
    reader := kafka.NewReader(kafka.ReaderConfig{
        Brokers: []string{"127.0.0.1:9092"},
        Topic:   "userInteractionData",
    })
    defer reader.Close()

    for {
        // Read a message
        message, err := reader.ReadMessage(context.Background())
        if err != nil {
            t.Fatal("failed to read message:", err)
        }

        log.Printf("Received message: key=%s, value=%s", string(message.Key), string(message.Value))
    }
}

