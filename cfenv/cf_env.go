package cfenv

import (
	"encoding/json"
	"log"
	"os"
)

type VcapServices struct {
	PRabbitMQ []PRabbitMQ `json:"p-rabbitmq"`
}

type PRabbitMQ struct {
	Credentials Credentials `json:"credentials"`
}

type Credentials struct {
	Protocols Protocols `json:"protocols"`
}

type Protocols struct {
	AMQP AMQP `json:"amqp"`
}

type AMQP struct {
	URI string `json:"uri"`
}

func GetAmqPUri() string {
	var vcapServices VcapServices
	err := json.Unmarshal([]byte(os.Getenv("VCAP_SERVICES")), &vcapServices)
	failOnError(err, "Could not parse VCAP_SERVICES")
	return vcapServices.PRabbitMQ[0].Credentials.Protocols.AMQP.URI
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
