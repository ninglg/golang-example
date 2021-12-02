package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/Shopify/sarama"
)

// Sarama配置选项
var (
	brokers = "127.0.0.1:9092,127.0.0.1:9092,127.0.0.1:9092"
	version = "0.10.2.0"  //V0_10_2_0及以上版本才支持消费者组
	group   = "consumer_monkey_coursedata_statistics"  //消费者组
	topics  = "topic_monkey_coursedata_statistics"    //支持逗号分隔的多个topic
	channelBufferSize = 2
)

func main() {
	sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)

	version, err := sarama.ParseKafkaVersion(version)
	if err != nil {
		log.Panicf("Error parsing Kafka version: %v", err)
	}

	/**
	 * Construct a new Sarama configuration.
	 * The Kafka cluster version has to be defined before the consumer/producer is initialized.
	 */
	config := sarama.NewConfig()
	config.Version = version

	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	config.Consumer.MaxProcessingTime = 10 * time.Second
	config.ChannelBufferSize = channelBufferSize

	// 自定义一个消费者标识字符串
	config.ClientID = "sarama_kafka_consumer"

	/**
	 * 设置一个新的Sarama消费者组
	 */
	consumer := Consumer{
		ready: make(chan bool),
	}

	ctx, cancel := context.WithCancel(context.Background())
	client, err := sarama.NewConsumerGroup(strings.Split(brokers, ","), group, config)
	if err != nil {
		log.Panicf("Error creating consumer group client: %v", err)
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			// 必须在死循环中调用consume
			// 如果服务侧产生重平衡, the consumer session will need to be recreated to get the new claims
			if err := client.Consume(ctx, strings.Split(topics, ","), &consumer); err != nil {
				log.Panicf("Error from consumer: %v", err)
			}
			// check if context was cancelled, signaling that the consumer should stop
			if ctx.Err() != nil {
				log.Println(ctx.Err())
				return
			}
			consumer.ready = make(chan bool)
		}
	}()

	<-consumer.ready // 等待消费者设置完毕
	log.Println("Sarama consumer up and running!...")

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-ctx.Done():
		log.Println("terminating: context cancelled")
	case <-sigterm:
		log.Println("terminating: via signal")
	}
	cancel()
	wg.Wait()
	if err = client.Close(); err != nil {
		log.Panicf("Error closing client: %v", err)
	}
}

// Consumer 定义一个Sarama消费者组方式的消费者
type Consumer struct {
	ready chan bool
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(consumer.ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/Shopify/sarama/blob/master/consumer_group.go#L27-L29
	for message := range claim.Messages() {
		// 这里处理真实的消息负载
		err := PrintKafkaMessage(message)
		if err != nil {
			log.Printf("Error claim message: %v", message)
			// 消息后续会重试消费
			continue
		}

		session.MarkMessage(message, "")
	}

	return nil
}

func PrintKafkaMessage(message *sarama.ConsumerMessage) error {
	log.Printf("Message claimed: value = %s, timestamp = %v, topic = %s", string(message.Value), message.Timestamp, message.Topic)

	return nil
}
