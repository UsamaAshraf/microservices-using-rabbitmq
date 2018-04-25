package main

import (
	"fmt"
	"log"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		fmt.Sprintf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func main() {
	// 'rabbitmq-server' is the network reference we have to the broker, thanks to Docker Compose.
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

	err = ch.ExchangeDeclare(
			exchangeName, 			// name
			"topic",  				// type
			true,         			// durable
			false,        			// auto-deleted
			false,        			// internal
			false,        			// no-wait
			nil,          			// arguments
	)
	failOnError(err, "Error creating the exchange")
	

	q, err := ch.QueueDeclare(
		"",    // name
		true,  // durable
		false, // delete when usused
		false,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Error creating the queue")


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
		true,   // auto ack
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
		}
	}()
	
	fmt.Println("Listening for events...")
	<-forever
}
