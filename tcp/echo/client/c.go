package main

import (
	"github.com/gorilla/websocket"
	"github.com/srlemon/note/log"
	"github.com/srlemon/note/tcp/echo"
	"time"
)

func main() {
	c, _, err := websocket.DefaultDialer.Dial("ws://192.168.31.180:8896", nil)
	println("aaaaaaaaaaaa")
	if err != nil {
		log.Errorln(err)
		return
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
			println(t.String())
			err := c.WriteMessage(websocket.TextMessage, echo.NewEchoPacket([]byte("请求连接"), false, 1).Serialize())
			if err != nil {
				log.Println("write:", err)
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
