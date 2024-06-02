package main

import (
	"context"
	"source/cmd/api"
)

func main(){
    p := api.NewProducersApi()
   
    topic := "userInteractionData"
    
    ctx, _ := context.WithCancel(context.Background())
    api.PushUserInteractionData(topic, ctx, p.Producer) 
    defer p.Close()
     
}
