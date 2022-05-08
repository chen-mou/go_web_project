package mq

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"sync"
	"time"
)

var producer *kafka.Producer
var con context.Context
var lock = &sync.Mutex{}

func init() {
	//create()
}

func create() {
	var err error
	var f context.CancelFunc
	producer, err = kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "121.37.87.181:9092",
	})
	if err != nil {
		panic(err)
	}
	con, f = context.WithTimeout(context.Background(), 20*time.Minute)
	go close(con, f)
}

func close(con context.Context, f context.CancelFunc) {
	select {
	case <-con.Done():
		lock.Lock()
		producer.Close()
		producer = nil
		lock.Unlock()
		f()
	}
}

func Send(topic string, message interface{}) {
	lock.Lock()
	if producer == nil {
		create()
	}
	lock.Unlock()
	go func() {
		for e := range producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()
	val, _ := json.Marshal(message)
	producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic: &topic,
		},
		Value: val,
	}, nil)
	producer.Flush(15 * 1000)
}
