package api

import (
	"log/slog"
	"net"
	"strconv"

	"github.com/segmentio/kafka-go"
)

type producers struct{
    Conn *kafka.Conn
}

func NewProducers() *producers{
    conn := initKafkaConn() 
    pr := &producers{
       Conn: conn,
    }
    return pr
}

func initKafkaConn() *kafka.Conn{
    conn, err := kafka.Dial("tcp", "localhost:9092")
    if err != nil {
        panic(err.Error())
    }
    // connection close should be handle by the caller    
    return conn
}

func(p *producers) NewTopic(topicConfigs []kafka.TopicConfig) error{
    // to create topics when auto.create.topics.enable='false'
    controller, err := p.Conn.Controller()
    if err != nil {
        slog.Error(err.Error())
        return err
    }
    var controllerConn *kafka.Conn
    controllerConn, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
    if err != nil {
        slog.Error(err.Error())
        return err
    }
    defer controllerConn.Close()

    err = controllerConn.CreateTopics(topicConfigs...)
    if err != nil {
        slog.Error(err.Error())
        return err
    }
    return nil
}

