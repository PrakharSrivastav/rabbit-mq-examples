// +build main1

package main

import (
	"github.com/PrakharSrivastav/rabbit-mq-examples/internal/helper"
	"github.com/streadway/amqp"
)

func main() {
	// create an amqp connection
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	helper.FailOnError(err, "Fialed to connect to RabbitMQ")
	defer conn.Close()

	// open a channel from connection
	ch, err := conn.Channel()
	helper.FailOnError(err, "Fialed to create a channel")
	defer ch.Close()

	// declare a queue
	q, err := ch.QueueDeclare("local", false, false, false, false, nil)
	helper.FailOnError(err, "Failed to create a queue")

	body := "This is the message body"

	err = ch.Publish("", q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	})
	helper.FailOnError(err, "Failed to publish message")
}
