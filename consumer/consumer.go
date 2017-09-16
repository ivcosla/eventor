package consumer

import (
	"eventor/store"

	"github.com/Shopify/sarama"
)

type Consumer struct {
	hooks                  map[string]func(b []byte)
	listener               store.Listener
	businessEntity         string
	businessEntityListener sarama.PartitionConsumer
}

func New(businessEntity string, listener store.Listener) Consumer {
	return Consumer{
		hooks:          make(map[string]func(b []byte)),
		listener:       listener,
		businessEntity: businessEntity,
	}
}

func (c *Consumer) Register(key string, callback func(b []byte)) {
	c.hooks[key] = callback
}

func (c *Consumer) Listen() {
	var err error

	eventStore := c.listener.Listener()
	c.businessEntityListener, err = eventStore.ConsumePartition(
		c.businessEntity,
		0,
		sarama.OffsetNewest,
	)

	if err != nil {
		panic(err)
	}

	go c.ingestEvents()
}

func (c *Consumer) ingestEvents() {
	for {
		select {
		case event := <-c.businessEntityListener.Messages():
			c.fire(string(event.Key), event.Value)
		}
	}
}

func (c *Consumer) fire(key string, b []byte) {
	c.hooks[key](b)
}
