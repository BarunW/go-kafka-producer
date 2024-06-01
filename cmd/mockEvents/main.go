package mockevents

import (
	"context"
	"fmt"
	"log/slog"
	"math/rand"
	"source/cmd/eventSource"
	"time"
)

func GenerateEvents(ctx context.Context, eventChan chan<-eventSource.UserInteractionData){
    ticker := time.NewTicker(13 * time.Millisecond) 
    var resetTicker = func() { 
        t := rand.Intn(60)
        ticker.Reset(time.Duration(t) * time.Second) 
        fmt.Printf("\n------Next event in %d second-----\n", t)
    }
    
    outer:
    for {
        select {
        case <-ticker.C:
            data, err := eventSource.NewUserInteractionData() 
            if err != nil{
                slog.Error("Unable to souce the event", "Details", err.Error())
                resetTicker()
                continue
            }
            eventChan <- *data
            resetTicker()
        case <-ctx.Done():
            break outer
        }
    }
    ticker.Stop()
    close(eventChan)
    slog.Info("Generating Events is Done")

}
