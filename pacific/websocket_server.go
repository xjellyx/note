// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"flag"
	"go.uber.org/zap"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{} // use default options

type gorillaWebsocket struct {
	lock *sync.Mutex
	conn map[*connClient]*websocket.Conn
	log  *zap.Logger
}

type connClient struct {
	conn *websocket.Conn
	done chan struct{}
	log  *zap.Logger
}

func (g *gorillaWebsocket) SocketHandler(w http.ResponseWriter, r *http.Request) {
	g.lock.Lock()
	defer g.lock.Unlock()
	conn, err := upgrader.Upgrade(w, r, nil)
	g.log.Info("conn", zap.String("client", conn.RemoteAddr().String()))
	if err != nil {
		g.log.Fatal("SocketHandler", zap.Error(err))
		return
	}
	cli := &connClient{done: make(chan struct{}), conn: conn, log: g.log}
	g.conn[cli] = conn
	go cli.Read()
	go cli.Close(g)

}

func (c *connClient) Close(g *gorillaWebsocket) {
	for {
		select {
		case <-c.done:
			c.log.Info("conn close", zap.String("client", c.conn.RemoteAddr().String()))
			c.conn.Close()
			g.lock.Lock()
			delete(g.conn, c)
			g.lock.Unlock()
			c.log.Info("conn length", zap.Int("conn", len(g.conn)))
			return
		}
	}
}

func (c *connClient) Read() {
	defer close(c.done)
	for {
		_, _, err := c.conn.ReadMessage()
		if err != nil {
			c.log.Info("read:", zap.Error(err))
			break
		}
	}
}

func (g *gorillaWebsocket) Broadcast(ctx context.Context, data interface{}) {
	for k, v := range g.conn {
		if err := v.WriteJSON(data); err != nil {
			g.log.Info("send:", zap.Error(err))
			close(k.done)
			continue
		}
	}
	return
}

func NewGorillaWebsocket(log *zap.Logger) *gorillaWebsocket {
	return &gorillaWebsocket{
		log:  log,
		lock: &sync.Mutex{},
		conn: map[*connClient]*websocket.Conn{},
	}
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	l, _ := zap.NewProduction()
	g := NewGorillaWebsocket(l)
	http.HandleFunc("/echo", g.SocketHandler)
	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				g.Broadcast(context.Background(), map[string]string{"aaa": "bbbb"})
			}
		}
	}()
	log.Fatal(http.ListenAndServe(*addr, nil))
}
