// This example declares a durable Exchange, and publishes a single message to
// that Exchange with a given routing key.
package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	uri          = flag.String("uri", "amqp://admin:123456@192.168.3.42:5682/", "AMQP URI")
	exchangeName = flag.String("exchange", "data-management-exchange", "Durable AMQP exchange name")
	exchangeType = flag.String("exchange-type", "fanout", "Exchange type - direct|fanout|topic|x-custom")
	routingKey   = flag.String("key", "system-manage", "AMQP routing key")
	body         = flag.String("body", "foobar", "Body of message")
	reliable     = flag.Bool("reliable", true, "Wait for the publisher confirmation before exiting")
	continuous   = flag.Bool("continuous", true, "Keep publishing messages at a 1msg/sec rate")
	ErrLog       = log.New(os.Stderr, "[ERROR] ", log.LstdFlags|log.Lmsgprefix)
	Log          = log.New(os.Stdout, "[INFO] ", log.LstdFlags|log.Lmsgprefix)
)

func init() {
	flag.Parse()
}

func main() {
	done := make(chan bool)

	SetupCloseHandler(done)

	if err := publish(done, *uri, *exchangeName, *exchangeType, *routingKey, *body, *reliable); err != nil {
		ErrLog.Fatalf("%s", err)
	}
}

type project struct {
	Name         string     `json:"name"`
	CustomerUni  string     `json:"customerUni"`
	PlatformUuid string     `json:"platformUuid"`
	Locations    []location `json:"locations"`
}

type location struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Address   string `json:"address"`
	Continent string `json:"continent"`
	Country   string `json:"country"`
	City      string `json:"city"`
}

func SetupCloseHandler(done chan bool) {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		done <- true
		Log.Printf("Ctrl+C pressed in Terminal")
	}()
}

func publish(done chan bool, amqpURI, exchange, exchangeType, routingKey, body string, reliable bool) error {
	// This function dials, connects, declares, publishes, and tears down,
	// all in one go. In a real service, you probably want to maintain a
	// long-lived connection as state, and publish against that.
	config := amqp.Config{Properties: amqp.NewConnectionProperties()}
	config.Properties.SetClientConnectionName("sample-producer")
	Log.Printf("dialing %q", amqpURI)
	connection, err := amqp.DialConfig(amqpURI, config)
	if err != nil {
		return fmt.Errorf("dial: %s", err)
	}
	defer connection.Close()

	Log.Printf("got Connection, getting Channel")
	channel, err := connection.Channel()
	if err != nil {
		return fmt.Errorf("channel: %s", err)
	}

	Log.Printf("got Channel, declaring %q Exchange (%q)", exchangeType, exchange)
	if err := channel.ExchangeDeclare(
		exchange,     // name
		exchangeType, // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // noWait
		nil,          // arguments
	); err != nil {
		return fmt.Errorf("exchange Declare: %s", err)
	}

	var publishes chan uint64 = nil
	var confirms chan amqp.Confirmation = nil

	// Reliable publisher confirms require confirm.select support from the
	// connection.
	if reliable {
		Log.Printf("enabling publisher confirms.")
		if err := channel.Confirm(false); err != nil {
			return fmt.Errorf("channel could not be put into confirm mode: %s", err)
		}
		// We'll allow for a few outstanding publisher confirms
		publishes = make(chan uint64, 8)
		confirms = channel.NotifyPublish(make(chan amqp.Confirmation, 1))

		go confirmHandler(done, publishes, confirms)
	}

	Log.Println("declared Exchange, publishing messages")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for {
		seqNo := channel.GetNextPublishSeqNo()
		Log.Printf("publishing %dB body (%q)", len(body), body)
		bodys, _ := json.Marshal(project{
			Name:        "demo",
			CustomerUni: "1",
			Locations: []location{
				{Latitude: "11.34",
					Longitude: "111.34",
					Address:   "dfasdf"},
			},
		})
		if err := channel.PublishWithContext(ctx,
			exchange,   // publish to an exchange
			routingKey, // routing to 0 or more queues
			false,      // mandatory
			false,      // immediate
			amqp.Publishing{
				Headers:         amqp.Table{},
				Type:            "project-add",
				ContentType:     "application/json",
				ContentEncoding: "",
				Body:            []byte(bodys),
				DeliveryMode:    amqp.Transient, // 1=non-persistent, 2=persistent
				Priority:        0,              // 0-9
				// a bunch of application/implementation-specific fields
			},
		); err != nil {
			return fmt.Errorf("Exchange Publish: %s", err)
		}

		Log.Printf("published %dB OK", len(body))
		if reliable {
			publishes <- seqNo
		}

		if *continuous {
			select {
			case <-done:
				Log.Println("producer is stopping")
				return nil
			case <-time.After(time.Second):
				continue
			}
		} else {
			break
		}
	}

	return nil
}

func confirmHandler(done chan bool, publishes chan uint64, confirms chan amqp.Confirmation) {
	m := make(map[uint64]bool)
	for {
		select {
		case <-done:
			Log.Println("confirmHandler is stopping")
			return
		case publishSeqNo := <-publishes:
			Log.Printf("waiting for confirmation of %d", publishSeqNo)
			m[publishSeqNo] = false
		case confirmed := <-confirms:
			if confirmed.DeliveryTag > 0 {
				if confirmed.Ack {
					Log.Printf("confirmed delivery with delivery tag: %d", confirmed.DeliveryTag)
				} else {
					ErrLog.Printf("failed delivery of delivery tag: %d", confirmed.DeliveryTag)
				}
				delete(m, confirmed.DeliveryTag)
			}
		}
		if len(m) > 1 {
			Log.Printf("outstanding confirmations: %d", len(m))
		}
	}
}
