package main

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"os"
	"os/signal"
)

const rabbitConnString = "amqp://guest:guest@localhost:5672/"

func main() {
	conn, err := amqp.Dial(rabbitConnString)
	if err != nil {
		fmt.Printf("new server: %v", err)
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Printf("Connected to RabbitMQ")

	// wait for ctrl+c
	signalChanel := make(chan os.Signal, 1)
	signal.Notify(signalChanel, os.Interrupt)
	<-signalChanel
}
