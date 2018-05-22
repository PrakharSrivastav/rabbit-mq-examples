package main

import (
	"log"

	"github.com/PrakharSrivastav/rabbit-mq-examples/internal/connection"
)

func main() {
	conn := connection.GetConnection()
	defer conn.Close()

	channel := connection.OpenChannel(conn)
	defer channel.Close()

	msgs := connection.Consume(channel, connection.DeclareQueue(channel, "local-queue"))

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Recieved message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
