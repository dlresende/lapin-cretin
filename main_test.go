package main_test

import (
	lapin "github.com/dlresende/lapin-cretin"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	Context("calling with parameters", func() {
		It("should throw error when number of parameters is wrong", func() {
			_, _, err := lapin.ParseArgs([]string{"lapin"})

			Expect(err).To(HaveOccurred())
		})

		It("should throw error when number of connections is not an integer", func() {
			_, _, err := lapin.ParseArgs([]string{"lapin", "notAnInt", "200"})

			Expect(err).To(HaveOccurred())
		})

		It("should throw error when number of connections per channel is not an integer", func() {
			_, _, err := lapin.ParseArgs([]string{"lapin", "10", "notAnInt"})

			Expect(err).To(HaveOccurred())
		})

		It("should convert parameters to integer", func() {
			nbOfConn, nbOfChPerConn, err := lapin.ParseArgs([]string{"lapin", "10", "200"})

			Expect(nbOfConn).To(Equal(10))
			Expect(nbOfChPerConn).To(Equal(200))
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
