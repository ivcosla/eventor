package store

import (
	"log"

	"github.com/Shopify/sarama"
)

type Emmitter struct {
	config   *sarama.Config
	emmitter sarama.AsyncProducer
}

func NewEmmitter(brokerList []string) Emmitter {
	emmitter := Emmitter{}
	emmitter.new(brokerList)
	return emmitter
}

func (e *Emmitter) Emmitter() sarama.AsyncProducer {
	return e.emmitter
}

func (e *Emmitter) new(brokerList []string) {
	var err error
	e.configure()
	e.emmitter, err = sarama.NewAsyncProducer(brokerList, e.config)
	if err != nil {
		log.Fatalln("Cannot connect to kafka broker list: ", brokerList)
	}
}

func (e *Emmitter) configure() {
	e.config = sarama.NewConfig()
	e.config.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	e.config.Producer.Return.Successes = false
	e.config.Producer.Return.Errors = true
}
