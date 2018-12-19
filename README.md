# kafka-test
[![Build Status](https://travis-ci.org/nicovogelaar/kafka-test.svg?branch=master)](https://travis-ci.org/nicovogelaar/kafka-test)
```
BenchmarkConfluentProducer_Produce-2          	 5000000	      6235 ns/op	     339 B/op	      11 allocs/op
--- BENCH: BenchmarkConfluentProducer_Produce-2
    confluent_test.go:44: 1
    confluent_test.go:44: 30
    confluent_test.go:44: 1000
    confluent_test.go:44: 30000
    confluent_test.go:44: 1000000
    confluent_test.go:44: 3000000
    confluent_test.go:44: 5000000
BenchmarkConfluentProducer_ProduceChannel-2   	 5000000	      6105 ns/op	     501 B/op	      12 allocs/op
--- BENCH: BenchmarkConfluentProducer_ProduceChannel-2
    confluent_test.go:91: 1
    confluent_test.go:91: 100
    confluent_test.go:91: 10000
    confluent_test.go:91: 1000000
    confluent_test.go:91: 5000000
BenchmarkSaramaProducer_SyncProducer-2        	  100000	    226690 ns/op	    3506 B/op	      65 allocs/op
--- BENCH: BenchmarkSaramaProducer_SyncProducer-2
    sarama_test.go:46: 1
    sarama_test.go:46: 100
    sarama_test.go:46: 5000
    sarama_test.go:46: 100000
BenchmarkSaramaProducer_AsyncProducer-2       	10000000	      4093 ns/op	     504 B/op	      10 allocs/op
--- BENCH: BenchmarkSaramaProducer_AsyncProducer-2
    sarama_test.go:83: 1
    sarama_test.go:83: 100
    sarama_test.go:83: 10000
    sarama_test.go:83: 1000000
    sarama_test.go:83: 10000000
PASS
ok  	github.com/nicovogelaar/kafka-test	174.574s
```
