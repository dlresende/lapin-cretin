package main_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	lapin "github.com/dlresende/lapin-cretin"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	AMQP_URI        = "amqp://guest:guest@localhost:5672/"
	RMQ_MGT_API_URI = "http://guest:guest@localhost:15672/api/overview"
	_10             = 10
	_20             = 20
)

var _ = Describe("Main", func() {
	Context("Create ghost channels", func() {
		It("should create N connections and M channels per connection", func(done Done) {
			forever := make(chan interface{})

			go lapin.CreateGhostChannels(AMQP_URI, _10, _20, forever, 0)

			Eventually(func() int {
				overview := getOverview(RMQ_MGT_API_URI)
				return overview.ObjectTotals.Connections
			}, 10, 1).Should(Equal(_10))

			Eventually(func() int {
				overview := getOverview(RMQ_MGT_API_URI)
				return overview.ObjectTotals.Channels
			}, 10, 1).Should(Equal(_10 * _20))

			close(done)
		}, 10)
	})
})

type Overview struct {
	ObjectTotals ObjectTotals `json:"object_totals"`
}

type ObjectTotals struct {
	Channels    int `json:"channels"`
	Connections int `json:"connections"`
	Consumers   int `json:"consumers"`
	Exchanges   int `json:"exchanges"`
	Queues      int `json:"queues"`
}

func getOverview(uri string) Overview {
	resp, err := http.Get(uri)
	Expect(err).NotTo(HaveOccurred())
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	var overview Overview
	json.Unmarshal(body, &overview)

	return overview
}
