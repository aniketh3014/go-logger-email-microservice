package event

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func declareEnchange(channel *amqp.Channel) error {
	return channel.ExchangeDeclare(
		"logs_topic", // name
		"topic",      // type
		true,         // durable?
		false,        // auto delete?
		false,        // internal?
		false,        // no wait?
		nil,          // atguments?
	)
}

func declareRandomQueue(channel *amqp.Channel) (amqp.Queue, error) {
	return channel.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete?
		true,  // exclusive?
		false, // no-wait
		nil,   // arguments
	)
}
