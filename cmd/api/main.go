package api

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type producerApi struct {
	Producer *kafka.Producer
}

func NewProducersApi(configMap kafka.ConfigMap) *producerApi {
	p, err := kafka.NewProducer(&configMap)
	if err != nil {
		panic(err)
	}
	pa := &producerApi{
		Producer: p,
	}
	pa.msgReportHandler()
	return pa
}

func (pa *producerApi) msgReportHandler() {
	// Delivery report handler for produced messages
	go func() {
		for e := range pa.Producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()
}

func (pa *producerApi) Close() {
	// Wait for message deliveries before shutting down
	pa.Producer.Flush(15 * 1000)
	pa.Producer.Close()
}
