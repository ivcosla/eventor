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
			consumer := NewConsumer()
			domainObject := entity{
				id: 3,
			}
			consumer.Register("update listener", func() {
				domainObject.id++
			})

			consumer.fire("update listener")

			Expect(domainObject.id).To(Equal(4))
		})
	})
})
