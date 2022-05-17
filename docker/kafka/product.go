package main

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"strconv"
	"sync"
	"time"
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
	producer, err := sarama.NewSyncProducer([]string{"192.168.3.85:9092", "192.168.3.85:9093"}, conf)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	msg := &sarama.ProducerMessage{Topic: "test",
		Partition: int32(-1),
		Key:       sarama.StringEncoder("key"),
	}

	var (
		i  = 200
		wg = &sync.WaitGroup{}
	)
	for j := 0; j < i; j++ {
		time.Sleep(time.Second)
		wg.Add(1)
		go func() {
			defer wg.Done()
			loc.Lock()
			defer loc.Unlock()
			var s = new(student)
			s.Id = uuid.NewV4().String()
			s.Name = strconv.FormatInt(int64(i), 10)
			s.Sex = "boy"
			s.Age = 18
			body, _ := json.Marshal(s)

			msg.Value = sarama.ByteEncoder(body)
			//producer.Input() <- msg
			//<-producer.Successes()
			paritition, offset, err := producer.SendMessage(msg)

			if err != nil {
				logrus.Error(err)
				fmt.Println("Send Message Fail")
			}

			fmt.Printf("Partion = %d, offset = %d\n", paritition, offset)
		}()

	}

	wg.Wait()
}
