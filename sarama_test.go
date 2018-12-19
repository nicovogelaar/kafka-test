package kafka

import (
	"testing"
	"time"

	"github.com/Shopify/sarama"
	"github.com/myles-mcdonnell/blondie"
)

func newSaramaSyncProducer(broker string) (sarama.SyncProducer, error) {
	conf := sarama.NewConfig()
	conf.Producer.RequiredAcks = sarama.WaitForAll
	conf.Producer.Retry.Max = 10
	conf.Producer.Return.Successes = true
	conf.Producer.Return.Errors = true
	conf.Producer.MaxMessageBytes = 100 << 20
	conf.Producer.Flush.MaxMessages = 10000
	return sarama.NewSyncProducer([]string{broker}, conf)
}

func newSaramaAsyncProducer(broker string) (sarama.AsyncProducer, error) {
	conf := sarama.NewConfig()
	conf.Producer.RequiredAcks = sarama.WaitForAll
	conf.Producer.Retry.Max = 10
	conf.Producer.Return.Successes = false
	conf.Producer.Return.Errors = true
	conf.Producer.MaxMessageBytes = 100 << 20
	conf.Producer.Flush.MaxMessages = 10000
	return sarama.NewAsyncProducer([]string{broker}, conf)
}

func BenchmarkSaramaProducer_SyncProducer(b *testing.B) {
	opts := blondie.DefaultOptions()
	opts.QuietMode = true

	blondie.WaitForDeps([]blondie.DepCheck{blondie.NewTcpCheck("kafka-broker", 9092, 30*time.Second)}, opts)

	producer, err := newSaramaSyncProducer("kafka-broker:9092")
	if err != nil {
		b.Fatalf("failed to create sarama producer: %v", err)
	}

	topic := "benchmark_topic"

	b.Logf("%v", b.N)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _, err = producer.SendMessage(&sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.StringEncoder(n),
		})
		if err != nil {
			b.Fatalf("failed to produce message: %v", err)
		}
	}

	if err = producer.Close(); err != nil {
		b.Fatalf("failed to close producer: %v", err)
	}
}

func BenchmarkSaramaProducer_AsyncProducer(b *testing.B) {
	opts := blondie.DefaultOptions()
	opts.QuietMode = true

	blondie.WaitForDeps([]blondie.DepCheck{blondie.NewTcpCheck("kafka-broker", 9092, 30*time.Second)}, opts)

	producer, err := newSaramaAsyncProducer("kafka-broker:9092")
	if err != nil {
		b.Fatalf("failed to create sarama producer: %v", err)
	}

	go func () {
		for err := range producer.Errors() {
			b.Logf("producer error: %v", err)
		}
	}()

	topic := "benchmark_topic"

	b.Logf("%v", b.N)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		producer.Input() <- &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.StringEncoder(n),
		}
	}

	if err = producer.Close(); err != nil {
		b.Fatalf("failed to close producer: %v", err)
	}
}
