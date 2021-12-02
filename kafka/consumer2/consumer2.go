package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":     "localhost",
		"broker.address.family": "v4",
		"group.id":              "consumer_group_monkey_courseapi_enweeklyreport",
		"auto.offset.reset":     "earliest",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{"topic_monkey_courseapi_enweeklyreport"}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

	c.Close()
}
