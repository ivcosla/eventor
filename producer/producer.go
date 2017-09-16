package producer

import (
	"encoding/json"
	"eventor/store"

	"github.com/Shopify/sarama"
)

type Producer struct {
	businessEntity string
	emmitter       store.Emmitter
}

func New(businessEntity string, emmitter store.Emmitter) Producer {
	return Producer{
		businessEntity: businessEntity,
		emmitter:       emmitter,
	}
}

func (p *Producer) Emmit(event string, data interface{}) {
	payload, err := json.Marshal(data)
	if err != nil {
		panic("event data cannot be marshaled into json")
	}

	p.emmitter.Emmitter().Input() <- &sarama.ProducerMessage{
		Topic: p.businessEntity,
		Key:   sarama.StringEncoder(event),
		Value: sarama.ByteEncoder(payload),
	}
}
