// +build main1

package main

import (
	"bytes"
	"log"
	"time"

	"github.com/PrakharSrivastav/rabbit-mq-examples/internal/connection"
	"github.com/PrakharSrivastav/rabbit-mq-examples/internal/helper"
)

func main() {
	c := connection.GetConnection()
	defer c.Close()

	ch := connection.OpenChannel(c)
	defer ch.Close()

	q := connection.DeclareQueue(ch, "task_queue")

	msgs, err := ch.Consume(
		q.Name,
		"",
		false, // auto-ack
		false,
		false,
		false,
		nil)

	helper.FailOnError(err, "Error consuming message from channel")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Recieved message : %s", d.Body)
			dot_count := bytes.Count(d.Body, []byte("."))
			t := time.Duration(dot_count)
			time.Sleep(t * time.Second)
			log.Println("Done")
			d.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
