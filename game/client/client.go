package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/topfreegames/pitaya/conn/packet"
	"time"
)

func main() {
	c, _, err := websocket.DefaultDialer.Dial("ws://127.0.0.1:3250", nil)
	if err != nil {
		panic(err)
	}

	a := map[string]interface{}{
		"name":    "dddd",
		"content": "csadfadsff",
		//"name1":    "dddd",
		//"content1": "csadfadsff",
		//"name2":    "dddd",
		//"content2": "csadfadsff",
		//"name3":    "dddd",
		//"content3": "csadfadsff",
		//"name4":    "dddd",
		//"content4": "csadfadsff",
		//"name5":    "dddd",
		//"content5": "csadfadsff",
		//"name6":    "dddd",
		//"content6": "csadfadsff",
		//"name7":    "dddd",
		//"content7": "csadfadsff",
		//"name8":    "dddd",
		//"content8": "csadfadsff",
	}
	d, _ := json.Marshal(a)
	var b = make([]byte, 4+len(d))
	b[0] = packet.Data
	binary.BigEndian.PutUint16(b[1:4], uint16(len(d)))
	fmt.Println(b[1:4])
	fmt.Println(uint16(len(d)))
	copy(b[4:], d)
	fmt.Println("ss", b)
	for i := 0; i < 1; i++ {
		if err = c.WriteMessage(2, b); err != nil {
			panic(err)
		}
	}

	time.Sleep(time.Second * 10)
	defer c.Close()
}

func IntToBytes(n int) []byte {
	bytebuf := bytes.NewBuffer([]byte{})
	binary.Write(bytebuf, binary.BigEndian, int64(n))
	return bytebuf.Bytes()
}
