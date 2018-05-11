# rabbit-mq-examples
This project shows various ways to 
- Setup RabbitMQ
- Connect to RabbitMQ
- Subscribe and consume messages
- Publish messages to queue

## Basics
- RabbitMQ is an open source message broker that implements AMQP protocol.
- The integrated services are loosly coupled and hence more scalable.
- Producers and Consumers communicate via Messages send to queue.
- Good practice: Never publish message to queues directly. Instead Publish to the Exchange which routes it to queue.
- 