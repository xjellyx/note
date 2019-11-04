package main

import (
	"encoding/json"
	"fmt"
	"github.com/gansidui/gotcp/examples/echo"
	"log"
	"net"
	"os"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:8896")
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	//
	//echoProtocol := &echo.EchoProtocol{}
	//
	//// ping <--> pong
	//for i := 0; i < 3; i++ {
	//	// write
	//	conn.Write(echo.NewEchoPacket([]byte("hello"), false).Serialize())
	//
	//	// read
	//	p, err := echoProtocol.ReadPacket(conn)
	//	if err == nil {
	//		echoPacket := p.(*echo.EchoPacket)
	//		fmt.Printf("Server reply:[%v] [%v]\n", echoPacket.GetLength(), string(echoPacket.GetBody()))
	//	}
	//
	//	time.Sleep(2 * time.Second)
	//}

	go sendMessage(conn)
	readMessage(conn)

	conn.Close()
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// 发送消息体
type messageData struct {
	ClientIP string `json:"client_ip"`
	Message  string `json:"message"`
}

// sendMessage 向服务器端发消息
func sendMessage(conn *net.TCPConn) {
	var m = new(messageData)
	m.ClientIP = conn.LocalAddr().String()
	for {
		var input string

		// 接收输入消息，放到input变量中
		fmt.Scanln(&input)

		if input == "/q" || input == "/quit" {
			fmt.Println("Byebye ...")
			conn.Close()
			os.Exit(0)
		}

		// 只处理有内容的消息
		if len(input) > 0 {
			m.Message = input
			_d, _ := json.Marshal(m)
			_, err := conn.Write(echo.NewEchoPacket(_d, false).Serialize())
			if err != nil {
				conn.Close()
				break
			}
		}
	}
}

// readMessage 读取消息
func readMessage(conn *net.TCPConn) {
	// 接收来自服务器端的广播消息
	echoProtocol := &echo.EchoProtocol{}
	for {
		p, err := echoProtocol.ReadPacket(conn)
		if err == nil {
			echoPacket := p.(*echo.EchoPacket)
			fmt.Printf("Server reply:[%v] [%v]\n", echoPacket.GetLength(), string(echoPacket.GetBody()))
		}
	}
}
