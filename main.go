package main

import (
	"time"

	. "github.com/dlresende/lapin-cretin/rmqclient"
	amqp "github.com/streadway/amqp"
)

func main() {
	url := "amqp://guest:guest@localhost:5672"
	var conn *amqp.Connection
	for x := 0; x < 390; x++ {
		conn = OpenConnection(url)
	}
	// defer conn.Close()

	var ch *amqp.Channel
	for x := 0; x < 55000; x++ {
		ch = CreateChannel(conn)
	}

	// defer ch.Close()
	q := DeclareNonDurableNonAutoDeletedQueue(ch, "test")
	start := time.Now()

	Publish(ch, q.Name, start.Local().String())
	ConsumeAll(ch, q.Name, "testConsumer")
}
