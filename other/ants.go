package main

import (
	"crypto/tls"
	"github.com/gorilla/websocket"
	"github.com/olefen/horse/examples/echo"
	"github.com/olefen/note/log"
	"time"
)

func main() {
	var (
		i int64
		//j int64
	)

	for i = 0; i <= 10000000; i++ {
		go do()
		println(i)
	}

	time.Sleep(time.Second * 300)
}

func do() {
	config := tls.Config{InsecureSkipVerify: true}
	dialer := websocket.Dialer{TLSClientConfig: &config}
	// h5001.zb16999.com
	conn, _, err := dialer.Dial("ws://0.0.0.0:8196", nil)
	if err != nil {
		checkError(err)
	}
	echoProtocol := &echo.EchoProtocol{}

	// ping <--> pong
	var (
		i = 0
	)
	go func() {
		for i < 2 {
			i++
			// write
			conn.WriteMessage(2, ([]byte("hello")))
		}
	}()
	for j := 0; j < 2; j++ {
		// read
		_, err := echoProtocol.ReadPacket(conn)
		if err == nil {
			//echoPacket := p.(*echo.EchoPacket)
			//fmt.Printf("Server reply:[%v] [%v]\n", echoPacket.GetLength(), string(echoPacket.GetBody()))
		} else {
			log.Infoln(err)
		}
	}
	time.Sleep(time.Microsecond)
	conn.Close()
}

func checkError(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
