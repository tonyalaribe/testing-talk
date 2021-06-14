// +build  e2e

package main

import (
	"testing"

	"github.com/imroc/req"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const rootURL = "http://localhost:8080"

func TestE2E(t *testing.T) {
	defer GinkgoRecover()
	RegisterFailHandler(Fail)
	RunSpecs(t, "Geolocator Suite")
}

var _ = Describe("Here we describe the test", func() {
	It("should fail with a wrong auth", func() {
		r, err := req.Get(rootURL + "/submit?auth=WRONG&age=19&name=tony&email=abc@xyz.com")
		Expect(err).To(BeNil())
		Expect(r.Response().StatusCode).To(Equal(401))
	})

	It("should succeed", func() {
		r, err := req.Get(rootURL + "/submit?auth=ABC&age=19&name=tony&email=abc@xyz.com")
		Expect(err).To(BeNil())
		Expect(r.Response().StatusCode).To(Equal(200))
		Expect(r.String()).To(Equal("Success in Handling  request"))
	})
})
