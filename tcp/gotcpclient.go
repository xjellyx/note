package main

import (
	"fmt"
	"github.com/srlemon/note/log"
	"net"
	"os"
)

func main() {
	startGoTcp()
}

func startGoTcp() {
	var (
		tcpAddress *net.TCPAddr
		conn       *net.TCPConn
		err        error
	)
	if tcpAddress, err = net.ResolveTCPAddr("tcp4", "127.0.0.1:9188"); err != nil {
		panic(err)
	}
	if conn, err = net.DialTCP("tcp", nil, tcpAddress); err != nil {
		panic(err)
	}
	go sendMessageTcp(conn)
	// 接收来自服务器端的广播消息
	buf := make([]byte, 1024)
	for {
		length, err := conn.Read(buf)
		if err != nil {
			log.Errorf("recv server msg failed: %v\n", err)
			conn.Close()
			os.Exit(0)
			break
		}

		log.Infoln(string(buf[0:length]))
	}
}

// 向服务器端发消息
func sendMessageTcp(conn net.Conn) {
	username := conn.LocalAddr().String()
	log.Println(username, "aaaaaaaaaaaaaaaaa")
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
				log.Error(err.Error())
				conn.Close()
				break
			} else {
				println("success")
			}
		}
	}
}
