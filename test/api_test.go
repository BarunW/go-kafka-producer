package test

import (
	"fmt"
	"source/cmd/api"
	"testing"

	"github.com/segmentio/kafka-go"
)

func TestKafkaConnection(t *testing.T){
    p := api.NewProducers()
    NewTopicTest(t, p.Conn, p.NewTopic)
    t.Cleanup(func() {
        p.Conn.Close()
    }) 
}

func NewTopicTest(t *testing.T, conn *kafka.Conn, i interface{}){
    f, ok:= i.(func(config []kafka.TopicConfig) error) 
    if !ok { t.Fail() }

    topic := "testTopic"
    topicConfigs := []kafka.TopicConfig{
        {
            Topic:             topic,
            NumPartitions:     1,
            ReplicationFactor: 1,
        },
    }
    if err := f(topicConfigs); err != nil{
        t.Fail()
    }

    // Delete the testTopic
    err := conn.DeleteTopics(topic)
    if err != nil{
        fmt.Println(err)
        t.Fail()
    } 
}


