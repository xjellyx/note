package main

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/olongfen/contrib/log"
	"net/http"
	"sync"
)

type Call struct {
}

func (c *Call) OnMessage(cli *ClientWS, msg []byte) {
	log.Warnln(string(msg), ":aaaaaaaaaaaaaaaa")
	cli.AsyncWritePacket(msg)

}

func main() {
	c := &Call{}
	s := NewServerWs(c.OnMessage)
	s.Wait()
	http.HandleFunc("/", s.Start)
	log.Fatal(http.ListenAndServe("0.0.0.0:8196", nil))
}

var (
	CloseChan = make(chan int)
)

// Callback
type Callback interface {
	OnMessage(conn *ClientWS, msg []byte)
}

// ServerWS
type ServerWS struct {
	connNum   int
	callback  func(conn *ClientWS, msg []byte)
	lock      *sync.RWMutex
	CacheConn map[int]*ClientWS
	waitGroup *sync.WaitGroup
}

// NewServerWs
func NewServerWs(c func(conn *ClientWS, msg []byte)) *ServerWS {
	s := &ServerWS{
		waitGroup: &sync.WaitGroup{},
		callback:  c,
		lock:      &sync.RWMutex{},
		CacheConn: map[int]*ClientWS{},
	}
	return s
}

// DeleteCloseConn 删除已经关闭的客户端
func (s *ServerWS) DeleteCloseConn() {
	go func() {
		for {
			select {
			case id := <-CloseChan:
				s.lock.Lock()
				delete(s.CacheConn, id)
				s.lock.Unlock()
			}
		}
	}()
}

func (s *ServerWS) Start(w http.ResponseWriter, r *http.Request) {
	var (
		upgrader = websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}
		err  error
		conn *websocket.Conn
	)
	if conn, err = upgrader.Upgrade(w, r, nil); err != nil {
		log.Fatal("[websocket upgrade failed] err: ", err)
	}

	// 发送连接成功消息
	msg := make(map[string]interface{})
	msg["id"] = s.connNum + 1
	msg["message"] = "success"
	b, _ := json.Marshal(msg)
	conn.WriteMessage(websocket.BinaryMessage, b)
	s.waitGroup.Add(1)
	s.connNum++
	go func() {
		defer s.waitGroup.Done()
		c := NewClientWS(s.connNum, s.callback, conn)
		s.lock.Lock()
		s.CacheConn[c.ID] = c
		s.lock.Unlock()
		c.Do()
	}()
	s.DeleteCloseConn()

}

// Wait wait service
func (s *ServerWS) Wait() {
	s.waitGroup.Wait()
}

type ClientWS struct {
	ID       int
	Conn     *websocket.Conn
	callback func(conn *ClientWS, msg []byte)
	Send     chan []byte
	Msg      chan []byte
}

func NewClientWS(id int, cl func(conn *ClientWS, msg []byte), conn *websocket.Conn) *ClientWS {
	c := &ClientWS{
		Conn:     conn,
		ID:       id,
		callback: cl,
		Send:     make(chan []byte),
		Msg:      make(chan []byte),
	}
	c.Conn.SetPongHandler(func(appData string) error {
		log.Warnln("asaaaaaaaaaaaaaaaaaaaaa")
		return c.Conn.WriteMessage(websocket.PongMessage, []byte(appData))
	})
	return c
}

// AsyncWritePacket async writes a packet, this method will never block
func (c *ClientWS) AsyncWritePacket(data []byte) {
	c.Send <- data
}

// Close
func (c *ClientWS) Close() {
	c.Conn.Close()
	go func() {
		CloseChan <- c.ID
	}()
}

// Write
func (c *ClientWS) Write() {
	defer func() {
		c.Close()
	}()

	for {
		select {
		case msg, ok := <-c.Send:
			if ok {
				c.Conn.WriteMessage(websocket.BinaryMessage, msg)
			}
		}
	}
}

// Read
func (c *ClientWS) Read() {
	defer func() {
		c.Close()
	}()

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			break
		}
		c.Msg <- message
	}
}

// HandleMsg
func (c *ClientWS) HandleMsg() {
	defer func() {
		c.Close()
	}()

	for {
		select {
		case p := <-c.Msg:
			c.callback(c, p)
		}
	}
}

// Do
func (c *ClientWS) Do() {
	var (
		wg = sync.WaitGroup{}
	)
	defer wg.Done()
	wg.Add(3)

	go func() {
		defer wg.Done()
		c.Write()
	}()

	go func() {
		defer wg.Done()
		c.Read()
	}()

	go func() {
		defer wg.Done()
		c.HandleMsg()
	}()

	wg.Wait()
}
