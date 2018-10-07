package cfenv_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCfenv(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CF env Suite")
}
