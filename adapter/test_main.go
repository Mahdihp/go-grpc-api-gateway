package main

import (
	"context"
	"fmt"
	"github.com/hellokvn/go-grpc-api-gateway/pkg/config"
	kafka "github.com/segmentio/kafka-go"
	"log"
	"time"
)

func main() {

	cfg, err := config.LoadConfig()
	fmt.Println(cfg)
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	conn, err := kafka.DialLeader(context.Background(), "tcp", cfg.Kafka_Server, "test-topic", 0)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	conn.SetReadDeadline(time.Now().Add(20 * time.Second))
	//batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max
	b := make([]byte, 10e3) // 10KB max per message
	// fetch 10KB min, 1MB max

	for {
		n, err := conn.Read(b)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(string(b[:n]))
	}

	//if err := batch.Close(); err != nil {
	//	log.Fatal("failed to close batch:", err)
	//}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close connection:", err)
	}
}
