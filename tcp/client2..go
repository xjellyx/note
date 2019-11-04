package main

import (
	"fmt"
	"github.com/srlemon/note/log"
	"net"
	"os"
)

func main() {
	startClient2("127.0.0.1:1253")
}

func startClient2(t string) (err error) {
	var (
		tcpAddress *net.TCPAddr
		conn       *net.TCPConn
	)
	if tcpAddress, err = net.ResolveTCPAddr("tcp4", t); err != nil {
		log.Errorln(err)
		return
	}

	// 向服务器拨号
	if conn, err = net.DialTCP("tcp", nil, tcpAddress); err != nil {
		return
	}

	// 发送消息
	go sendMessage2(conn)

	// 接收来自服务器端的广播消息
	buf := make([]byte, 1024)
	for {
		length, err := conn.Read(buf)
		if err != nil {
			log.Printf("recv server msg failed: %v\n", err)
			conn.Close()
			os.Exit(0)
			break
		}

		log.Infoln(string(buf[0:length]))
	}
	return
}

// 向服务器端发消息
func sendMessage2(conn net.Conn) {
	username := conn.LocalAddr().String()
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
			msg := username + " say:" + input
			_, err := conn.Write([]byte(msg))
			if err != nil {
				conn.Close()
				break
			}
		}
	}
}
