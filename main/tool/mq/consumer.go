package mq

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var Ch chan bool

var topics = make(map[string]*func(kafka.Message))

//var consumer = make(map[string]*kafka.Consumer)

func Register(_topic string, f func(message kafka.Message)) {
	topics[_topic] = &f
}

func init() {
	Ch = make(chan bool)
	//go handle()
}

func handle() {
	select {
	case <-Ch:
		c, err := kafka.NewConsumer(&kafka.ConfigMap{
			"bootstrap.servers": "121.37.87.181:9092",
			"group.id":          "server",
			"auto.offset.reset": "earliest",
		})
		if err != nil {
			panic(err)
		}
		keys := make([]string, len(topics))
		for key := range topics {
			if key == "" {
				continue
			}
			keys = append(keys, key)
		}
		c.SubscribeTopics(keys, nil)
		for {
			message, err1 := c.ReadMessage(-1)
			if err1 != nil {
				fmt.Println(err1)
				continue
			}
			fmt.Println(string(message.Value) + ":" + *message.TopicPartition.Topic)
		}

	}
}
