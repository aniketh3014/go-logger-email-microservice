package event

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Emmiter struct {
	connection *amqp.Connection
}

func NewEmmiter(conn *amqp.Connection) (Emmiter, error) {
	emmiter := Emmiter{
		conn,
	}

	err := emmiter.setup()
	if err != nil {
		return Emmiter{}, err
	}

	return emmiter, nil
}

func (e *Emmiter) setup() error {
	channel, err := e.connection.Channel()
	if err != nil {
		return err
	}

	defer channel.Close()

	return declareEnchange(channel)
}

func (e *Emmiter) Push(event string, severity string) error {
	channel, err := e.connection.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	log.Println("pushing to mq")

	err = channel.Publish(
		"logs_topic",
		severity,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(event),
		},
	)

	if err != nil {
		return err
	}

	return nil
}
