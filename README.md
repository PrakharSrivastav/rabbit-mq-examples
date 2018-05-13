# rabbit-mq-examples
This project shows various ways to 
- Setup RabbitMQ
- Connect to RabbitMQ
- Subscribe and consume messages
- Publish messages to queue

## Basics
- RabbitMQ is an open source message broker that implements AMQP protocol.
- The integrated services are loosly coupled and hence more scalable.
- **Producers** and **Consumers** communicate via **Messages** sent to a **Queue**.
- Good practice: Never publish message to queues directly. Instead Publish to an **Exchange** that forwards it to **Queue**.
- **Exchange** connects to a queue via a **Binding** and a **Binding Key**.
- An **Exchange** can communicate to multiple **Queues** and route messages based on **Binding** and the **Binding Key**
- The **Producer** publishes the message to the **Exchange** with a **Routing Key**. 
- The **Consumer** can acknowledge the message consumption or reject it so that it returns to the queue.



### Exchange types : Routing techniques
- **Fanout** : When a message sent to an exchange is broadcasted to all the queues. Exchange ignores all the routing rules and sends the message to all the queues it knows about.
- **Direct** : Forwards the message to the queue where **Routing Key = Binding Key**.
- **Topic** : Allows partial matches of keys. Uses wildcards to match the Routing Key with binding Keys.
- **Header** : Uses message **headers** to route instead of Routing key.
- **Default**: Matches **routing key** with **queue name** instead of binding key. RabbitMQ creates this Exchange type for each queue automatically.

### Rabbit MQ setup.
- Instead of installing a full blown RabbitMQ server, we just use a docker image.
  ```
  sudo docker run --hostname rabbit-mq-example -p 8080:15672 -p 5672:5672 rabbitmq:3.7.5-management-alpine
  ```
- RabbitMQ provides capability to programatically create Exchanges
