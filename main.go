package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"source/cmd/api"
	"source/configs"
)

func main() {
	conf := configs.NewConfigs()
	p := api.NewProducersApi(conf.KafkaConfigMap())

	topic := "user-interaction-data"

	ctx, cancel := context.WithCancel(context.Background())
	go graceFullShutDown(cancel)
	api.PushUserInteractionData(topic, ctx, p.Producer)
	defer p.Close()

	slog.Info("Producer Successfully Stop")
}

func graceFullShutDown(cancel context.CancelFunc) {
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Kill)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan
	cancel()

}
