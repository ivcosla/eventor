package consumer

import (
	"encoding/json"
	"eventor/store"

	"github.com/Shopify/sarama"
)

type consumer struct {
	hooks                  map[string]func(b []byte)
	listener               store.Listener
	businessEntity         string
	businessEntityListener sarama.PartitionConsumer
}

func New(businessEntity string, listener store.Listener) consumer {
	return consumer{
		hooks:          make(map[string]func(b []byte)),
		listener:       listener,
		businessEntity: businessEntity,
	}
}

func (c *consumer) Register(key string, callback func(b []byte)) {
	c.hooks[key] = callback
}

func (c *consumer) Listen() {
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

func (c *consumer) ingestEvents() {
	var key string

	for {
		select {
		case event := <-c.businessEntityListener.Messages():
			json.Unmarshal(event.Key, &key)
			c.fire(key, event.Value)
		}
	}
}

func (c *consumer) fire(key string, b []byte) {
	c.hooks[key](b)
}
