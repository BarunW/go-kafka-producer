package test

import (
	"context"
	"fmt"
	"source/cmd/eventSource"
	mockevents "source/cmd/mockEvents"
	"testing"
	"time"
)

func TestGenerateEvents(t *testing.T){
    var (
        dur time.Duration = 10
        unit time.Duration = time.Second
    )
    fmt.Printf("This will test take %v + %v\n",dur, unit)

    ctx, cancel := context.WithCancel(context.Background()) 
    evenChan := make(chan eventSource.UserInteractionData, 8)
    
    timer := time.NewTimer(dur * unit)    
    go mockevents.GenerateEvents(ctx, evenChan)
   
    outer:
    for{
        select{
        case <-timer.C:
            break outer
        case data := <-evenChan : 
            fmt.Println(data)
        }
    }
    cancel()
    
}
