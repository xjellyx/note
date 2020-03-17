package main

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	uuid "github.com/satori/go.uuid"
	"strconv"
	"sync"
)

type student struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  string `json:"sex"`
}

var (
	loc = &sync.RWMutex{}
)

func main() {
	var conf = sarama.NewConfig()
	conf.Producer.Return.Successes = true
	conf.Producer.RequiredAcks = sarama.WaitForAll
	conf.Producer.Partitioner = sarama.NewRandomPartitioner
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, conf)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	msg := &sarama.ProducerMessage{Topic: "topic001",
		Partition: int32(-1),
		Key:       sarama.StringEncoder("key"),
	}

	var (
		i  = 2000000
		wg = &sync.WaitGroup{}
	)
	for j := 0; j < i; j++ {

		wg.Add(1)
		go func() {
			defer wg.Done()
			loc.Lock()
			defer loc.Unlock()
			var s = new(student)
			s.Id = uuid.NewV4().String()
			println(s.Id)
			s.Name = strconv.FormatInt(int64(i), 10)
			s.Sex = "boy"
			s.Age = 18
			body, _ := json.Marshal(s)

			msg.Value = sarama.ByteEncoder(body)
			paritition, offset, err := producer.SendMessage(msg)

			if err != nil {
				fmt.Println("Send Message Fail")
			}

			fmt.Printf("Partion = %d, offset = %d\n", paritition, offset)
		}()

	}

	wg.Wait()
}
