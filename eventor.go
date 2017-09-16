package eventor

import (
	"eventor/consumer"
	"eventor/producer"
	"eventor/store"
)

func NewConsumer(businessEntity string, listener store.Listener) consumer.Consumer {
	return consumer.New(businessEntity, listener)
}

func NewProducer(businessEntity string, emmiter store.Emmitter) producer.Producer {
	return producer.New(businessEntity, emmiter)
}
