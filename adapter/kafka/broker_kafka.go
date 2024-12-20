package kafka

import (
	"context"
	"fmt"
	"github.com/hellokvn/go-grpc-api-gateway/pkg/config"
	kafka "github.com/segmentio/kafka-go"
	"log"
	"time"
)

func KafkaProducer(cfg config.Config, topic string) *kafka.Writer {

	writer := kafka.Writer{
		Addr:     kafka.TCP(cfg.Kafka_Server),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
	_, err := kafka.DialLeader(context.Background(), "tcp", cfg.Kafka_Server, topic, 0)
	if err != nil {
		panic(err.Error())
	}
	return &writer
}

func KafkaConsumer(cfg config.Config, topic string) {
	conn, err := kafka.DialLeader(context.Background(), "tcp", cfg.Kafka_Server, topic, 1)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

	b := make([]byte, 10e3) // 10KB max per message
	for {
		n, err := batch.Read(b)
		if err != nil {
			break
		}
		fmt.Println(string(b[:n]))
	}

	if err := batch.Close(); err != nil {
		log.Fatal("failed to close batch:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close connection:", err)
	}
}
