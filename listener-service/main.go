package main

import (
	"listener/event"
	"log"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	// connect to rabbitmq
	conn, err := connect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer conn.Close()
	log.Println("connected to rabbit")
	// start listening for the messages
	log.Println("listening for mq messages..")
	// create a consumer
	consumer, err := event.NewConsumer(conn)
	if err != nil {
		panic(err)
	}
	// watch the queue and consume events
	err = consumer.Listen([]string{"log.INFO", "log.WARNING", "log.ERROR"})
	if err != nil {
		panic(err)
	}
}

func connect() (*amqp.Connection, error) {
	var count int32
	var backOff = 1 * time.Second
	var connection *amqp.Connection

	// do not continue untill rabbit is ready
	for {
		c, err := amqp.Dial("amqp://guest:guest@rabbitmq")
		if err != nil {
			log.Println("rabbitmq not yet ready")
			count++
		} else {
			connection = c
			break
		}
		if count > 5 {
			log.Println(err)
			return nil, err
		}
		backOff += time.Second * 2
		log.Println("backing off..")
		time.Sleep(backOff)
	}

	return connection, nil
}
