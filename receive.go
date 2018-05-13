package main

import (
	"log"

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

	q, err := ch.QueueDeclare("local", false, false, false, false, nil)
	helper.FailOnError(err, "Failed to declare queue")

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	helper.FailOnError(err, "Failed consuming message")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Recieved message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
