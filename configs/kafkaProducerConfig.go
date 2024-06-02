package configs

import "github.com/confluentinc/confluent-kafka-go/v2/kafka"

func (c configs) KafkaConfigMap() kafka.ConfigMap {
	return kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		//        "sasl.username":     "<CLUSTER API KEY>",
		//        "sasl.password":     "<CLUSTER API SECRET>",
		//
		//        // Fixed properties
		//        "security.protocol": "SASL_SSL",
		//        "sasl.mechanisms":   "PLAIN",
		//        "acks":              "all"}
	}
}
