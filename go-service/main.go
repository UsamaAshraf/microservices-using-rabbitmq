package main

import (
	"fmt"
	"log"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		fmt.Sprintf("%s: %s", msg, err)
		// Exit the program.
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func main() {
	// 'rabbitmq-server' is the network reference we have to the broker, 
	// thanks to Docker Compose.
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq-server:5672/")
	failOnError(err, "Error connecting to the broker")
	// Make sure we close the connection whenever the program is about to exit.
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	// Make sure we close the channel whenever the program is about to exit.
	defer ch.Close()
	
	exchangeName := "user_updates"
	bindingKey := "user.profile.*"

	// Create the exchange if it doesn't already exist.
	err = ch.ExchangeDeclare(
			exchangeName, 	// name
			"topic",  		// type
			true,         	// durable
			false,        	// auto-deleted
			false,        	// internal
			false,        	// no-wait
			nil,          	// arguments
	)
	failOnError(err, "Error creating the exchange")
	
	// Create the exchange if it doesn't already exist.
	// This does not need to be done in the publisher since the
	// queue is only relevant to the consumer, which subscribes to it.
	q, err := ch.QueueDeclare(
			"",    // name
			true,  // durable
			false, // delete when unused
			false, // exclusive
			false, // no-wait
			nil,   // arguments
	)
	failOnError(err, "Error creating the queue")

	// Bind the exchange to the queue based on a string pattern (binding key).
	err = ch.QueueBind(
			q.Name,       // queue name
			bindingKey,   // binding key
			exchangeName, // exchange
			false,
			nil,
	)
	failOnError(err, "Error binding the queue")

	
	msgs, err := ch.Consume(
			q.Name, // queue
			"",     // consumer
			false,  // auto ack
			false,  // exclusive
			false,  // no local
			false,  // no wait
			nil,    // args
	)
	failOnError(err, "Failed to register as a consumer")


	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received: %s", d.Body)

			// Update the user data on the service's associated datastore...

			// The 'false' indicates the success of a single delivery, 'true' would mean that
			// this delivery and all prior unacknowledged deliveries on this channel will be acknowledged,
			// which I find no reason to do in this example.
			d.Ack(false)
		}
	}()
	
	fmt.Println("Service listening for events...")
	
	// Block until 'forever' receives a value, which will never happen.
	<-forever
}
