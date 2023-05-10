package messagequeue

import (
	"fmt"
	"github.com/streadway/amqp"
	"time"
)

func main() {
	conn, err := amqp.Dial("amqp://myRabbit:qwer1234@localhost:5672/red-Rabbit")
	if err != nil {
		fmt.Println("Failed to connect to RabbitMQ", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println("Failed to open a channel", err)
	}
	defer ch.Close()

	body := "my red rabbit"
	err = ch.Publish(
		"exch",        // exchange
		"routing-key", // routing key
		false,         // mandatory
		false,         // immediate
		amqp.Publishing{
			//DeliveryMode: amqp.Persistent,
			//ContentType:  "application/json", // 조절
			Body:      []byte(body),
			Timestamp: time.Now(),
		})
	if err != nil {
		fmt.Println("Failed to publish a message", err)
	}
	fmt.Println(body)
}
