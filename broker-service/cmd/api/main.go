package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const webPort = "80"

type Config struct {
	Rabbit *amqp.Connection
}

func main() {

	conn, err := connect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer conn.Close()
	log.Println("connected to rabbit")

	app := &Config{
		Rabbit: conn,
	}

	log.Printf("Starting broker service at port: %s\n", webPort)

	// create a http setver
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	// start the server
	err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
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
