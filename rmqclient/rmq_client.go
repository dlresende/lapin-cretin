package rmqclient

import (
	"fmt"
	"log"

	amqp "github.com/streadway/amqp"
)

func OpenConnection(url string) *amqp.Connection {
	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to connect to RabbitMQ")
	return conn
}

func CreateChannel(conn *amqp.Connection) *amqp.Channel {
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	return ch
}

func DeclareNonDurableNonAutoDeletedQueue(ch *amqp.Channel, queueName string) amqp.Queue {
	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")
	return q
}

func Publish(ch *amqp.Channel, routingKey, message string) {
	err := ch.Publish(
		"",         // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	failOnError(err, "Failed to publish a message")
}

func ConsumeAll(ch *amqp.Channel, queue, consumer string) []string {
	msgs, err := ch.Consume(
		queue,    // queue
		consumer, // consumer
		true,     // auto-ack
		false,    // exclusive
		false,    // no-local
		false,    // no-wait
		nil,      // args
	)
	failOnError(err, "Failed to register a consumer")

	contents := make([]string, 0)
	go func() {
		for d := range msgs {
			contents = append(contents, string(d.Body))
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	err = ch.Cancel(consumer, false)
	failOnError(err, "Failed to cancel consumer "+consumer)
	return contents
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
