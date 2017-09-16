package store

import (
	"log"
	"os"
	"os/signal"

	"github.com/Shopify/sarama"
)

type Listener struct {
	signals  chan os.Signal
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
	l.signals = make(chan os.Signal, 1)
	signal.Notify(l.signals, os.Interrupt)
	go l.listenToChannels()
}

func (l *Listener) configure() {
	l.config = sarama.NewConfig()
	l.config.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	l.config.Producer.Return.Successes = false
	l.config.Producer.Return.Errors = true
}

func (l *Listener) Listener() sarama.Consumer {
	return l.listener
}

func (l *Listener) listenToChannels() {
	<-l.signals
	l.listener.Close()
	os.Exit(0)
}
