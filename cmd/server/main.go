package main

import (
	"fmt"
	"github.com/NicholasRodrigues/pub-sub-server/internal/pubsub"
	"github.com/NicholasRodrigues/pub-sub-server/internal/routing"
	amqp "github.com/rabbitmq/amqp091-go"
	"os"
	"os/signal"
)

const rabbitConnString = "amqp://guest:guest@localhost:5672/"

func main() {
	conn, err := amqp.Dial(rabbitConnString)
	if err != nil {
		fmt.Printf("new rabbitmq connection: %v", err)
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Printf("Connected to RabbitMQ")

	connChanel, err := conn.Channel()
	if err != nil {
		fmt.Printf("new rabbitmq channel: %v", err)
		os.Exit(1)
	}

	exchange := routing.ExchangePerilDirect
	routingKey := routing.PauseKey
	dataToSend := routing.PlayingState{
		IsPaused: true,
	}

	if err := pubsub.PublishJSON(connChanel, exchange, routingKey, dataToSend); err != nil {
		fmt.Printf("publishing message: %v", err)
		os.Exit(1)
	}

	// wait for ctrl+c
	signalChanel := make(chan os.Signal, 1)
	signal.Notify(signalChanel, os.Interrupt)
	<-signalChanel
}
