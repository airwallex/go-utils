package utils

import (
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test OS", func() {
	It("Getenv should work", func() {
		Expect(Getenv("foo", "bar")).To(Equal("bar"))
		os.Setenv("foo", "baz")
		Expect(Getenv("foo", "bar")).To(Equal("baz"))
	})
})
