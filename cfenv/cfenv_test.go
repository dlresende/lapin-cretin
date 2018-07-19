package cfenv_test

import (
	"io/ioutil"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/dlresende/lapin-cretin/cfenv"
)

var _ = Describe("Cfenv", func() {

	BeforeSuite(func() {
		os.Setenv("VCAP_SERVICES", readFile("testdata/vcap_services.json"))
	})

	It("should read AMQP URI from VCAP_SERVICES", func() {
		uri := GetAmqPUri()

		Expect(uri).To(Equal("amqp://93da6eed-447c-44bf-8089-284db322435d:my-secret-password@10.0.4.50:5672/eb1439d4-3788-401b-bcbf-574f0717c5ed"))
	})

})

func readFile(filepath string) string {
	file, _ := os.Open(filepath)
	defer file.Close()

	content, _ := ioutil.ReadAll(file)
	return string(content)
}
