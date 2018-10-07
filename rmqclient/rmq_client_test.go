package rmqclient_test

import (
	"log"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/dlresende/lapin-cretin/rmqclient"
)

var _ = Describe("Channels", func() {

	It("should publish and consume a message", func() {
		url := "amqp://guest:guest@localhost:5672"
		conn := OpenConnection(url)
		defer conn.Close()
		ch := CreateChannel(conn)
		defer ch.Close()
		q := DeclareNonDurableNonAutoDeletedQueue(ch, "test")
		start := time.Now()
		ConsumeAll(ch, q.Name, "testConsumer")

		Publish(ch, q.Name, start.Local().String())
		msgs := ConsumeAll(ch, q.Name, "testConsumer")

		Expect(msgs[0]).To(Equal(start.Local().String()))
	})
})

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
