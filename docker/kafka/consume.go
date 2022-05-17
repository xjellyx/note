package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"sync"
)

var (
	wg sync.WaitGroup
)

func main() {
	consumer, err := sarama.NewConsumer([]string{"localhost:9092", "localhost:9093"}, nil)

	if err != nil {
		panic(err)
	}

	partitionList, _err := consumer.Partitions("test")

	if _err != nil {
		panic(_err)
	}

	for partition := range partitionList {
		pc, err := consumer.ConsumePartition("test", int32(partition), 1)
		if err != nil {
			panic(err)
		}
		ch := make(chan bool)
		wg.Add(1)
		go func(sarama.PartitionConsumer) {
			defer wg.Done()
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d, Offset:%d, Key:%s, Value:%s\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			}
			ch <- true
		}(pc)
		<-ch
		pc.AsyncClose()
		consumer.Close()
	}
	wg.Wait()
}
