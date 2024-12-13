package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"product-worker/api"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Job struct {
	Action string
	Market string
	Data any
}

func main() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id": "product-event",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		log.Fatal(err)
	}

	err = c.SubscribeTopics([]string{"test-topic-2"}, nil);

	if err != nil {
		log.Fatal(err)
	}

	client := api.NewClient()

	run := true
	for run {
		event := c.Poll(100)
		if event == nil {
			continue
		}
		switch e := event.(type) {
		case *kafka.Message:
			var job Job
			err := json.Unmarshal(e.Value, &job)

			if err != nil {
				log.Printf("Message parsing error: %s\n", err)
				break
			}

			api, exists := client.Api(job.Market)

			if !exists {
				log.Printf("Unsupported market\n")
				break
			}

			fmt.Printf("action: %s\n", job.Action)
			switch job.Action {
			case "stop":
				err = api.Stop(job.Data)
			default:
				err = errors.New("unavailable action")
			}

			if err != nil {
				log.Printf("send retry-product-event")
				// 재시도 토픽으로 메세지 전송
				break
			}

			// DB 저장
		case kafka.Error:
			log.Printf("Error: %v: %v\n", e.Code(), e)
			run = false
		default:
			log.Printf("Ignored %v\n", e)
		}
	}

	log.Printf("Closing consumer\n")
	c.Close()
}