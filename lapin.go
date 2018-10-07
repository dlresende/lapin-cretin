package main

import (
	"log"
	"time"

	rmqclient "github.com/dlresende/lapin-cretin/rmqclient"
)

func CreateGhostChannels(uri string, nbOfConnections, nbOfChannelsPerConnection int, forever chan interface{}, rampUpInSec time.Duration) {
	for x := 0; x < nbOfConnections; x++ {
		time.Sleep(rampUpInSec * time.Second)

		go createNChannels(uri, nbOfChannelsPerConnection, forever)
	}

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func createNChannels(uri string, nbOfChannelsPerConnection int, forever chan interface{}) {
	connGhost := rmqclient.OpenConnection(uri)
	defer failOnError(connGhost.Close, "Fail to close connection")

	for x := 0; x < nbOfChannelsPerConnection; x++ {
		chGhost := rmqclient.CreateChannel(connGhost)
		defer failOnError(chGhost.Close, "Fail to close channel")
	}

	<-forever
}

func failOnError(f func() error, msg string) {
	if err := f(); err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
