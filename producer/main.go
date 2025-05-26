package main

import (
	"log"

	"github.com/streadway/amqp"
)

func main() {
	// connect to rabbitmq server
	// url structure "amqp://<username>:<password>@<host>:<port>/"
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ:", err)
	}
	defer conn.Close()

	// create channel to send message
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Failed to open channel:", err)
	}
	defer ch.Close()

	// create a queue message
	q, err := ch.QueueDeclare(
		"task_queue", // name: the queue name
		true,         // durable: survives RabbitMQ restarts
		false,        // autoDelete: delete when unused? (no)
		false,        // exclusive: used by one connection only? (no)
		false,        // noWait: wait for server response? (no)
		nil,          // arguments: optional settings (none)
	)
	if err != nil {
		log.Fatal("Failed to declare queue:", err)
	}

	body := "this is body"
	err = ch.Publish(
		"",     // exchange: "" = default direct exchange
		q.Name, // routing key: queue name (must match)
		false,  // mandatory: return message if no queue matches? (no)
		false,  // immediate: deliver immediately to consumer? (no)
		amqp.Publishing{
			DeliveryMode: amqp.Persistent, // make message persistent
			ContentType:  "text/plain",    // MIME type
			Body:         []byte(body),    // actual message body
		})
	if err != nil {
		log.Fatal("Failed to publish message:", err)
	}

	log.Printf(" [x] Sent: %s", body)
}
