// +build main1

package main

import (
	"github.com/PrakharSrivastav/rabbit-mq-examples/internal/connection"
	"github.com/PrakharSrivastav/rabbit-mq-examples/internal/helper"
	"github.com/streadway/amqp"
)

func main() {

	conn := connection.GetConnection()
	defer conn.Close()

	channel := connection.OpenChannel(conn)
	defer channel.Close()

	q := connection.DeclareQueue(channel)

	body := "This is the message body3"
	err := channel.Publish("", q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	})
	helper.FailOnError(err, "Failed to publish message")
}
