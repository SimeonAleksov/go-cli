package tui

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/SimeonAleksov/go-cli/consumer"
	"github.com/SimeonAleksov/go-cli/log"
)

type TUI struct {
	messages  chan string
	messagesM sync.Mutex
	consumer  *consumer.KafkaConsumer
}

func NewTUI(kafkaConsumer *consumer.KafkaConsumer) *TUI {
	tui := &TUI{
		consumer: kafkaConsumer,
	}

	return tui
}

func (t *TUI) Start() {
  messages := make(chan string, 100)
	go t.consumer.Start(messages)

	t.messages = messages
  go t.logKafkaMessages()
	go t.listenForTermination()
}

func (t *TUI) Stop() {
	t.consumer.Close()
}

func (t *TUI) listenForTermination() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	t.Stop()
}

func (t *TUI) logKafkaMessages() {
	for msg := range t.messages {
		log.Infoln("Received Kafka message:", msg)
	}
}

