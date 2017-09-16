package store

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/Shopify/sarama"
)

type Emmitter struct {
	signals  chan os.Signal
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
	e.signals = make(chan os.Signal, 1)
	signal.Notify(e.signals, os.Interrupt)
	go e.listenToChannels()
}

func (e *Emmitter) configure() {
	e.config = sarama.NewConfig()
	e.config.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	e.config.Producer.Return.Successes = false
	e.config.Producer.Return.Errors = true
}

func (e *Emmitter) listenToChannels() {
	for {
		select {
		case err := <-e.emmitter.Errors():
			panic(fmt.Sprint("Got error queuing message", err))
		case <-e.signals:
			e.emmitter.AsyncClose()
			os.Exit(0)
		}
	}
}
