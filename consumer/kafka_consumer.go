package consumer

import (
	"context"
  "time"

	"github.com/segmentio/kafka-go"
  "github.com/SimeonAleksov/go-cli/log"
)


type KafkaConsumer struct {
    conn *kafka.Conn
    topic string
}


func NewKafkaConsumer(broker, topic string) *KafkaConsumer {
    log.Infof("Initiating new kafka consumer - topic(%v).", topic)
    conn, err := kafka.DialLeader(context.Background(), "tcp", broker, topic, 0)

    if err != nil {
      log.Fatalf("Failed to connect to Kafka broker: %v", err)            
    }

    return &KafkaConsumer{
      conn: conn,
      topic: topic,
    }
}


func (kc *KafkaConsumer) Start(messages chan<- string) {
    defer kc.Close()

    reader := kafka.NewReader(kafka.ReaderConfig{
      Brokers: []string{"localhost:9092"},
      Topic: kc.topic,
      MinBytes: 10e3,
      MaxBytes: 10e6,
      MaxWait:         time.Second,
      StartOffset:     kafka.FirstOffset,
      ReadLagInterval: -1,
    })

     for {
       message, err := reader.ReadMessage(context.Background()) 
       if err != nil {
          log.Fatalf("Failed to connect to Kafka broker: %v", err)
          continue
      }

      messages <- string(message.Value)
    }

}

func (kc *KafkaConsumer) Close() {
    if kc.conn != nil {
        kc.conn.Close()
  }
}

