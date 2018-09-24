package lib_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/vsanna/go_web/lib"
)

var _ = Describe("SecureRandom", func() {
	It("secure random generate 16length string", func() {
		token, _ := SecureRandom()
		Expect(len(token)).To(Equal(16))
		Ω(len(token)).To(Equal(16))
	})
})
