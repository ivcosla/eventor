package producer

import (
	"encoding/json"
	"eventor/store"

	"github.com/Shopify/sarama"
)

type Producer struct {
	emmitter store.Emmitter
}

func New(brokerList []string) Producer {
	return Producer{
		emmitter: store.NewEmmitter(brokerList),
	}
}

func (p *Producer) Emmit(businessEntity string, event string, data interface{}) {
	payload, err := json.Marshal(data)
	if err != nil {
		panic("event data cannot be marshaled into json")
	}

	p.emmitter.Emmitter().Input() <- &sarama.ProducerMessage{
		Topic: businessEntity,
		Key:   sarama.StringEncoder(event),
		Value: sarama.StringEncoder(payload),
	}
}
