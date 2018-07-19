package main

import (
	"log"
	"time"

	. "github.com/dlresende/lapin-cretin/cfenv"
	. "github.com/dlresende/lapin-cretin/rmqclient"
	amqp "github.com/streadway/amqp"
)

func main() {
	url := GetAmqPUri()

	var conn *amqp.Connection
	conn = OpenConnection(url)
	defer conn.Close()

	var ch *amqp.Channel
	ch = CreateChannel(conn)
	defer ch.Close()

	q := DeclareNonDurableNonAutoDeletedQueue(ch, "test")

	forever := make(chan bool)

	go publishAMessageEveryNSec(ch, q.Name, 1)
	go consumeAllEveryNSec(ch, q.Name, 10)

	go func(nbOfConnections, nbOfChannelsPerConnection int) {
		for x := 0; x < nbOfConnections; x++ {
			var connGhost *amqp.Connection
			connGhost = OpenConnection(url)
			defer connGhost.Close()

			for x := 0; x < nbOfChannelsPerConnection; x++ {
				var chGhost *amqp.Channel
				chGhost = CreateChannel(connGhost)
				defer chGhost.Close()
			}
		}
	}(150, 500)

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func publishAMessageEveryNSec(ch *amqp.Channel, queue string, interval time.Duration) {
	for {
		Publish(ch, queue, time.Now().Local().String())
		time.Sleep(interval * time.Second)
	}
}

func consumeAllEveryNSec(ch *amqp.Channel, queue string, interval time.Duration) {
	for {
		ConsumeAll(ch, queue, "testConsumer")
		time.Sleep(interval * time.Second)
	}
}
