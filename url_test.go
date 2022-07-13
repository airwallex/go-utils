package utils

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test URL", func() {
	It("EncodeURL should work", func() {
		encoded, err := EncodeURL(`https://www.airwallex.com?components=["search"]`)
		Expect(err).To(BeNil())
		Expect(encoded).To(Equal("https://www.airwallex.com?components=%5B%22search%22%5D"))
	})

	It("GetURLHost should work", func() {
		host, err := GetURLHost(`https://www.airwallex.com?components=["search"]`)
		Expect(err).To(BeNil())
		Expect(host).To(Equal("www.airwallex.com"))
	})
})
