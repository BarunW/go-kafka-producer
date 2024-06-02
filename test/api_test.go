package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func TestProducer(t *testing.T) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "userInteractionDataGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	if err := c.SubscribeTopics([]string{"user-interaction-data"}, nil); err != nil {
		fmt.Println(err)
		t.Fail()
	}

	// A signal handler or similar could be used to set this to false to break the loop.
	run := true
	for run {
		msg, err := c.ReadMessage(time.Second)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else if !err.(kafka.Error).IsTimeout() {
			// The client will automatically try to recover from all errors.
			// Timeout is not considered an error because it is raised by
			// ReadMessage in absence of messages.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
	c.Close()
}
