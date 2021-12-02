package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Shopify/sarama"
)

var producer sarama.SyncProducer //同步生产者
//var producer sarama.AsyncProducer  //异步生产者

type MqConfig struct {
	Topics     []string `json:"topics"`
	Servers    []string `json:"servers"`
	ConsumerId string   `json:"consumerGroup"`
}

var cfg = &MqConfig{
	Topics:     []string{"topic_monkey_coursedata_statistics"},
	//Topics: []string{"topic_monkey_courseapi_enweeklyreport"},
	Servers:    []string{"127.0.0.1:9092", "127.0.0.1:9092", "127.0.0.1:9092"},
	ConsumerId: "consumer_monkey_coursedata_statistics",
	//ConsumerId: "consumer_group_monkey_courseapi_enweeklyreport",
}

func init() {
	var err error

	mqConfig := sarama.NewConfig()
	mqConfig.Net.DialTimeout = 1 * time.Second
	mqConfig.Net.WriteTimeout = 1 * time.Second
	mqConfig.Net.ReadTimeout = 1 * time.Second

	mqConfig.Producer.Return.Successes = true
	mqConfig.Producer.Return.Errors = true
	mqConfig.Version = sarama.V0_10_2_0

	producer, err = sarama.NewSyncProducer(cfg.Servers, mqConfig)
	if err != nil {
		msg := fmt.Sprintf("Kafak producer create fail. err: %v", err)
		fmt.Println(msg)
		panic(msg)
	}
}

func produce(topic string, key string, content string) error {
	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Key:       sarama.StringEncoder(key),
		//Value:     sarama.StringEncoder(content + time.Now().Format("2006-01-02 15:04:05")),
		Value:     sarama.StringEncoder(content),
		Timestamp: time.Now(),
	}

	_, _, err := producer.SendMessage(msg)
	if err != nil {
		msg := fmt.Sprintf("Send Error topic: %v. key: %v. content: %v", topic, key, content)
		fmt.Println(msg)
		return err
	}
	fmt.Printf("Send OK\ntopic:%s key:%s value:%s\n", topic, key, content)

	return nil
}

func main() {
	for {
		//the key of the kafka messages
		//do not set the same the key for all messages, it may cause partition im-balance
		key := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
		//value = `{"trace_id":"\"hmk_2_2518903_4054b4ae-100b-44e5-8fad-b5912591bcf5\"","value":{"user_id":2518903,"lesson_id":5,"section_id":25,"num":1,"type":2,"subject_type":1,"remark":"","now":"2021-05-14T16:00:07.228223238+08:00"}}`
		value := `{"trace_id":"\"` + key + `\"","value":{"user_id":2518903,"lesson_id":5,"section_id":25,"num":1,"type":2,"subject_type":1,"remark":"","now":"2021-05-14T16:00:07.228223238+08:00"}}`


		_ = produce(cfg.Topics[0], key, value)
		time.Sleep(100 * time.Millisecond)
	}
}
