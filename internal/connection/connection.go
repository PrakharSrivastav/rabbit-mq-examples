package connection

import (
	"github.com/PrakharSrivastav/rabbit-mq-examples/internal/helper"
	"github.com/streadway/amqp"
)

func GetConnection() *amqp.Connection {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	helper.FailOnError(err, "Fialed to connect to RabbitMQ")
	return conn
}

func OpenChannel(conn *amqp.Connection) *amqp.Channel {
	ch, err := conn.Channel()
	helper.FailOnError(err, "Fialed to create a channel")
	return ch
}

func DeclareQueue(channel *amqp.Channel) amqp.Queue {
	q, err := channel.QueueDeclare("local", false, false, false, false, nil)
	helper.FailOnError(err, "Failed to create a queue")
	return q
}

func Consume(ch *amqp.Channel, q amqp.Queue) <-chan amqp.Delivery {

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	helper.FailOnError(err, "Failed consuming message")
	return msgs
}
