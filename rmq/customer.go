package main

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
)

type rmqCCfg struct {
	uri          string
	exchange     string
	exchangeType string
	queue        string
	bindingKey   string
	tag          string
}

type consumer struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	tag     string
	done    chan error
}

func main() {
	var (
		rmq = rmqCCfg{
			uri:          "amqp://admin:123456@192.168.3.42:5682/",
			exchange:     "test-exchange1",
			exchangeType: "fanout",
			queue:        "test-queue",
			bindingKey:   "test-key",
			tag:          "simple-consumer",
		}
		err error
		c   = &consumer{
			conn:    nil,
			channel: nil,
			tag:     rmq.tag,
			done:    make(chan error),
		}
		queue      amqp.Queue
		deliveries <-chan amqp.Delivery
	)
	cfg := amqp.Config{Properties: amqp.NewConnectionProperties()}
	cfg.Properties.SetClientConnectionName(rmq.tag)
	c.conn, err = amqp.DialConfig(rmq.uri, cfg)
	if err != nil {
		panic(err)
	}
	if c.channel, err = c.conn.Channel(); err != nil {
		panic(err)
	}
	if err = c.channel.ExchangeDeclare(rmq.exchange, rmq.exchangeType, true, false,
		false, false, nil); err != nil {
		panic(err)
	}
	if queue, err = c.channel.QueueDeclare(rmq.queue, true, false, false, false, nil); err != nil {
		panic(err)
	}
	if deliveries, err = c.channel.Consume(queue.Name, c.tag, true, false, false, false, nil); err != nil {
		panic(err)
	}

	go func() {
		for d := range deliveries {
			fmt.Println(string(d.Body), d.DeliveryTag)
		}
	}()
	<-c.done
}
