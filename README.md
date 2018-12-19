# kafka-test
[![Build Status](https://travis-ci.org/nicovogelaar/kafka-test.svg?branch=master)](https://travis-ci.org/nicovogelaar/kafka-test)
```
BenchmarkConfluentProducer_Produce-2          	 5000000	      5483 ns/op	     324 B/op	       9 allocs/op
--- BENCH: BenchmarkConfluentProducer_Produce-2
    confluent_test.go:48: 1
    confluent_test.go:48: 30
    confluent_test.go:48: 1000
    confluent_test.go:48: 100000
    confluent_test.go:48: 2000000
    confluent_test.go:48: 3000000
    confluent_test.go:48: 5000000
BenchmarkConfluentProducer_ProduceChannel-2   	 5000000	      4501 ns/op	     323 B/op	       9 allocs/op
--- BENCH: BenchmarkConfluentProducer_ProduceChannel-2
    confluent_test.go:97: 1
    confluent_test.go:97: 100
    confluent_test.go:97: 10000
    confluent_test.go:97: 1000000
    confluent_test.go:97: 5000000
BenchmarkSaramaProducer_SyncProducer-2        	  100000	    215920 ns/op	    3322 B/op	      62 allocs/op
--- BENCH: BenchmarkSaramaProducer_SyncProducer-2
    sarama_test.go:51: 1
    sarama_test.go:51: 100
    sarama_test.go:51: 10000
    sarama_test.go:51: 100000
BenchmarkSaramaProducer_AsyncProducer-2       	10000000	      3405 ns/op	     331 B/op	       7 allocs/op
--- BENCH: BenchmarkSaramaProducer_AsyncProducer-2
    sarama_test.go:90: 1
    sarama_test.go:90: 100
    sarama_test.go:90: 10000
    sarama_test.go:90: 1000000
    sarama_test.go:90: 10000000
PASS
ok  	github.com/nicovogelaar/kafka-test	158.059s
```
