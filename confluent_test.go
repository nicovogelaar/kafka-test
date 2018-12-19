package kafka

import (
	"strconv"
	"testing"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/myles-mcdonnell/blondie"
)

func newConfluentProducer(broker string) (*kafka.Producer, error) {
	return kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
		"go.produce.channel.size": 100000,
	})
}

func BenchmarkConfluentProducer_Produce(b *testing.B) {
	opts := blondie.DefaultOptions()
	opts.QuietMode = true

	blondie.WaitForDeps([]blondie.DepCheck{blondie.NewTcpCheck("kafka-broker", 9092, 30*time.Second)}, opts)

	producer, err := newConfluentProducer("kafka-broker:9092")
	if err != nil {
		b.Fatalf("failed to create confluent producer: %v", err)
	}

	go func() {
		for e := range producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					b.Logf("delivery failed: %v", ev.TopicPartition)
				}
			}
		}
	}()

	topic := "benchmark_topic"
	topicPtr := &topic

	b.Logf("%v", b.N)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		err = producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: topicPtr, Partition: kafka.PartitionAny},
			Value:          []byte(strconv.Itoa(n)),
		}, nil)
		if err != nil {
			b.Fatalf("failed to produce message: %v", err)
		}
		if n % 100000 == 0 {
			producer.Flush(15000)
		}
	}

	go func () {
		producer.Flush(15000)
		producer.Close()
	}()
}

func BenchmarkConfluentProducer_ProduceChannel(b *testing.B) {
	opts := blondie.DefaultOptions()
	opts.QuietMode = true

	blondie.WaitForDeps([]blondie.DepCheck{blondie.NewTcpCheck("kafka-broker", 9092, 30*time.Second)}, opts)

	producer, err := newConfluentProducer("kafka-broker:9092")
	if err != nil {
		b.Fatalf("failed to create confluent producer: %v", err)
	}

	go func() {
		for e := range producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					b.Logf("delivery failed: %v", ev.TopicPartition)
				}
			}
		}
	}()

	topic := "benchmark_topic"
	topicPtr := &topic

	b.Logf("%v", b.N)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		producer.ProduceChannel() <- &kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: topicPtr, Partition: kafka.PartitionAny},
			Value:          []byte(strconv.Itoa(n)),
		}
	}

	go func () {
		producer.Flush(15000)
		producer.Close()
	}()
}
