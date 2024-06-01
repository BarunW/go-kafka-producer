package api

import (
    "bytes"
    "context"
    "encoding/gob"
    "log/slog"
    es "source/cmd/eventSource"
    mockevents "source/cmd/mockEvents"
    "time"

    "github.com/segmentio/kafka-go"
)

func encode(data es.UserInteractionData) ([]byte, error){
    buffer := bytes.Buffer{}
    e := gob.NewEncoder(&buffer)

    if err := e.Encode(data); err != nil{
       return nil, err 
    }

    return buffer.Bytes(), nil
}


func writeMsg(topic string, data es.UserInteractionData, pd *producers) error {       
    writer := &kafka.Writer{
       Addr: pd.Conn.RemoteAddr(),
       Topic: topic,
       MaxAttempts: 2,
       WriteTimeout: 10 * time.Second,
    }
    databByt, err := encode(data)   
    if err != nil{ 
        slog.Error("Unable to encode the user interaction data", "Deatails", err.Error())
        return err
    }

    err = writer.WriteMessages(
        context.Background(),
        kafka.Message{
            Value: databByt,
        },
    )
    if err != nil{ return err }
    slog.Info("Successfully Written the message")
    return nil
}
    

func PushUserInteractionData(topic string, ctx context.Context, p *producers){
    eventChan := make(chan es.UserInteractionData, 8) 
    go mockevents.GenerateEvents(ctx, eventChan) 
    outer:
    for {
        select{
        case data := <-eventChan:
            if err := writeMsg(topic, data, p); err != nil{ 
                slog.Error("Failed to write msg", "Details", err.Error())
                continue 
            }
        case <-ctx.Done():
            break outer 
        }
    }
}
