// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "192.168.3.85:8820", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/api/v1/security-msg/ws"}
	//u.RawQuery = ""
	log.Printf("connecting to %s", u.String())
	h := http.Header{}
	h.Set("Authorization", "Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJ1c2VJZCI6ImI4NWJkZmM3LWE0MTEtNGFkMy04YjQ0LWY5NDJjZGU4N2UxYSIsImV4cCI6MTY1OTQyOTgyNywiaWF0IjoxNjU5NDIyNjI3LCJuYmYiOjE2NTk0MjI2MjcsInN1YiI6ImI4NWJkZmM3LWE0MTEtNGFkMy04YjQ0LWY5NDJjZGU4N2UxYSJ9.MGkKslRUUWOOJLBcvIdOzUQTubJ6RXSkufMb6Rl9kqTpnUNeSAkT7H8T7YOYrLZNHum4Wg4Do4bw5Ga_1dxOmg")

	c, _, err := websocket.DefaultDialer.Dial(u.String(), h)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			err := c.WriteControl(websocket.PingMessage, []byte(t.String()), time.Time{})
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
