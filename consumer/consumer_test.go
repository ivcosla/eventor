package consumer

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type entity struct {
	id int
}

var _ = Describe("When I register a handler", func() {
	Context("And I call it", func() {
		It("Should update the business entity", func() {
			listener := listeners{}
			domainObject := entity{
				id: 3,
			}
			listener.Register("update listener", func() {
				domainObject.id++
			})

			listener.fire("update listener")

			Expect(domainObject.id).To(Equal(4))
		})
	})
})
