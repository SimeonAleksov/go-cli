package consumer

import (
	"context"
	"encoding/json"
	"github.com/SimeonAleksov/go-cli/config"
	"time"

	"github.com/SimeonAleksov/go-cli/domain/tracking"
	"github.com/SimeonAleksov/go-cli/log"
	"github.com/segmentio/kafka-go"
)

type KafkaConsumer struct {
	conn  *kafka.Conn
	topic string
}

func NewKafkaConsumer(broker, topic string) *KafkaConsumer {
	log.Infof("Initiating new kafka consumer - topic(%v).", topic)
	conn, err := kafka.DialLeader(context.Background(), "tcp", broker, topic, 0)

	if err != nil {
		log.Fatalf("Failed to connect to Kafka broker: %v", err)
	}

	return &KafkaConsumer{
		conn:  conn,
		topic: topic,
	}
}

func (kc *KafkaConsumer) Start(messages chan<- tracking.Account) {
	defer kc.Close()

	maxRetries := 5
	retryDelay := 5 * time.Second

	for retry := 0; retry < maxRetries; retry++ {
		log.Printf("Attempt #%d\n", retry+1)
		cfg := config.LoadConfigProvider("KAFKA")
		kafkaUrl := cfg.GetString("BROKER")

		reader := kafka.NewReader(kafka.ReaderConfig{
			Brokers:         []string{kafkaUrl},
			Topic:           kc.topic,
			MinBytes:        10e3,
			MaxBytes:        10e6,
			MaxWait:         time.Second,
			StartOffset:     kafka.FirstOffset,
			ReadLagInterval: -1,
		})

		for {
			message, err := reader.ReadMessage(context.Background())
			if err != nil {
				log.Infof("Failed to connect to Kafka broker: %v", err)
				break
			}
			var account tracking.Account

			if err := json.Unmarshal(message.Value, &account); err != nil {
				log.Printf("Error parsing JSON: %v\n", err)
				continue
			}
			messages <- account
		}
		time.Sleep(retryDelay)
		if err := reader.Close(); err != nil {
			log.Printf("Error closing Kafka reader: %v\n", err)
			return
		}
	}
}

func (kc *KafkaConsumer) Close() {
	if kc.conn != nil {
		err := kc.conn.Close()
		if err != nil {
			return
		}
	}
}
