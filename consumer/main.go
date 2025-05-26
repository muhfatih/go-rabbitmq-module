package main

import (
	"log"
	"time"

	"github.com/streadway/amqp"
)

func main() {
	// Connect to the RabbitMQ server running locally
	// Format: "amqp://<user>:<password>@<host>:<port>/"
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ:", err)
	}
	defer conn.Close()

	// Open a channel for communication
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Failed to open a channel:", err)
	}
	defer ch.Close()

	// Declare the same queue as the producer to make sure it exists
	q, err := ch.QueueDeclare(
		"task_queue", // queue name
		true,         // durable: survives broker restarts
		false,        // autoDelete: don't delete when last consumer disconnects
		false,        // exclusive: not limited to this connection
		false,        // noWait: wait for server confirmation
		nil,          // arguments: none for now
	)
	if err != nil {
		log.Fatal("Failed to declare a queue:", err)
	}

	// Register as a consumer to receive messages from the queue
	msgs, err := ch.Consume(
		q.Name, // queue name
		"",     // consumer tag: empty = auto-generated
		true,   // autoAck: messages are considered "acknowledged" as soon as delivered
		false,  // exclusive: allow other consumers
		false,  // noLocal: not used by RabbitMQ (for other AMQP impls)
		false,  // noWait: wait for server response
		nil,    // arguments: none
	)
	if err != nil {
		log.Fatal("Failed to register a consumer:", err)
	}

	log.Println(" [*] Waiting for messages. To exit press CTRL+C")

	// Consume messages in a loop
	for d := range msgs {
		log.Printf(" [x] Received: %s", d.Body)

		// Simulate some work by sleeping (e.g. sending email, processing job, etc.)
		time.Sleep(2 * time.Second)

		log.Println(" [x] Done")
	}
}
