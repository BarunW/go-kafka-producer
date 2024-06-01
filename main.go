package main

import (
	"context"
	"fmt"
	"log"
	"source/cmd/api"

	"github.com/segmentio/kafka-go"
)


func main(){
    pd := api.NewProducers()
    
    topic := "userInteractionData"
    topicConfigs := []kafka.TopicConfig{
        {
            Topic:             topic,
            NumPartitions:     1,
            ReplicationFactor: 1,
        },
    }
 
    if err := pd.NewTopic(topicConfigs); err != nil{   
      log.Fatal(err) 
    } 
    
    fmt.Println(pd.Conn.RemoteAddr())

    ctx, _ := context.WithCancel(context.Background())
    api.PushUserInteractionData(topic, ctx, pd) 
    defer pd.Conn.Close()
     
}
