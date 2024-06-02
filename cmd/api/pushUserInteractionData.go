package api

import (
	"context"
	"encoding/json"
	"log/slog"
	mockevents "source/cmd/mockEvents"
	"source/models"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func encode(data models.UserInteractionData) ([]byte, error) {
	dataByt, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return dataByt, nil
}

func writeMsg(topic string, data models.UserInteractionData, p *kafka.Producer) error {
	// Produce messages to topic (asynchronously)
	datByt, err := encode(data)
	if err != nil {
		return err
	}

	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          datByt,
	}, nil)

	return nil
}

func PushUserInteractionData(topic string, ctx context.Context, p *kafka.Producer) {
	eventChan := make(chan models.UserInteractionData, 8)
	go mockevents.GenerateEvents(ctx, eventChan)
    outer:
	for {
		select {
		case data := <-eventChan:
			if err := writeMsg(topic, data, p); err != nil {
				slog.Error("Failed to write msg", "Details", err.Error())
				continue
			}
		case <-ctx.Done():
			break outer
		}
	}
}
