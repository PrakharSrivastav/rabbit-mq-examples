package main

import (
	"log"
	"os"
	"strings"

	"github.com/PrakharSrivastav/rabbit-mq-examples/internal/helper"

	"github.com/streadway/amqp"

	"github.com/PrakharSrivastav/rabbit-mq-examples/internal/connection"
)

func main() {
	conn := connection.GetConnection()
	defer conn.Close()
	channel := connection.OpenChannel(conn)
	defer channel.Close()

	queue := connection.DeclareQueue(channel, "task_queue")

	body := bodyFrom(os.Args)

	err := channel.Publish("", queue.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         []byte(body),
	})

	helper.FailOnError(err, "Error publishing to worker queue")
	log.Printf(" [x] Sent %s", body)
}

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}
