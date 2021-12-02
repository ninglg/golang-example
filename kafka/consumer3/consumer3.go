package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		GroupID: "consumer_group_monkey_courseapi_enweeklyreport",
		Topic:   "topic_monkey_courseapi_enweeklyreport",
		MinBytes: 10,    // 10B
		MaxBytes: 10e6, // 10MB
	})

	ctx := context.Background()
	for {
		m, err := r.FetchMessage(ctx)
		if err != nil {
			log.Fatal("failed to fetch message:", err)
		}

		do(m)
		if err := r.CommitMessages(ctx, m); err != nil {
			log.Fatal("failed to commit messages:", err)
		}
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}

func do(m kafka.Message) {
	fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	time.Sleep(500 * time.Millisecond)
}
