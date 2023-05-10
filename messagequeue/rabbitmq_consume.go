package messagequeue

import (
	"fmt"
	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://myRabbit:qwer1234@localhost:5672/")
	if err != nil {
		fmt.Printf("%s: %s", "Failed to connect to RabbitMQ", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Printf("%s: %s", "Failed to open a channel", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"test-queue01", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		fmt.Printf("%s: %s", "Failed to declare a queue", err)
	}
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		fmt.Printf("%s: %s", "Failed to register a consumer", err)
	}
	forever := make(chan bool)

	go func() {
		fmt.Printf("Receive start!")
		for d := range msgs {
			fmt.Println("Received a message:", string(d.Body[:]), d.UserId, d.Headers)
		}
	}()

	fmt.Printf("Waiting for messages...")
	<-forever
}
