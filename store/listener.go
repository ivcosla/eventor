package store

import (
	"log"

	"github.com/Shopify/sarama"
)

type Listener struct {
	config   *sarama.Config
	listener sarama.Consumer
}

func NewListener(brokerList []string) Listener {
	listener := Listener{}
	listener.new(brokerList)
	return listener
}

func (l *Listener) new(brokerList []string) {
	var err error
	l.configure()
	l.listener, err = sarama.NewConsumer(brokerList, l.config)
	if err != nil {
		log.Fatalln("Listener.New: Cannot connect to kafka broker list: ", brokerList)
	}
}

func (l *Listener) configure() {
	l.config = sarama.NewConfig()
	l.config.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	l.config.Producer.Return.Successes = false
	l.config.Producer.Return.Errors = true
}
