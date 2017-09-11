package consumer

import (
	"encoding/json"
	"eventor/store"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type entity struct {
	id int
}

type updateListener struct {
	Increment int
}

var _ = Describe("When I register a handler", func() {
	var listener store.Listener
	var consumer consumer

	BeforeEach(func() {
		listener = store.NewListener([]string{"localhost:9092"})
		consumer = NewConsumer("business_entity", listener)
	})

	Context("And I call it", func() {
		It("Should update the business entity", func() {
			domainObject := entity{
				id: 3,
			}
			consumer.Register("update listener", func(b []byte) {
				var payload updateListener
				json.Unmarshal(b, &payload)

				domainObject.id += payload.Increment
			})
			producerPayload := []byte(`{"increment":2}`)

			consumer.fire("update listener", producerPayload)

			Expect(domainObject.id).To(Equal(5))
		})
	})
})
