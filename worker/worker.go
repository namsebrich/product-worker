package worker

import (
	"encoding/json"
	"log"
	"product-worker/client"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Worker struct {
	run      bool
	consumer *kafka.Consumer
	client   *client.Client
}

func New() *Worker {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "product-event",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		log.Fatal(err)
	}

	err = consumer.SubscribeTopics([]string{"product-event"}, nil)

	if err != nil {
		log.Fatal(err)
	}

	client := client.New()

	return &Worker{run: true, consumer: consumer, client: client}
}

func (w *Worker) Run() {
	for w.run {
		event := w.consumer.Poll(100)
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

			err = job.Run(w.client)

			// if err != nil {
			// 	// 재시도 토픽 전송
			// 	log.Printf("send retry-product-event")
			// } else {
			// 	// DB 저장
			// }

		case kafka.Error:
			log.Printf("Error: %v: %v\n", e.Code(), e)
			w.run = false
		default:
			log.Printf("Ignored %v\n", e)
		}
	}
}

func (w *Worker) Close() {
	log.Printf("Closing worker\n")
	w.consumer.Close()
}
