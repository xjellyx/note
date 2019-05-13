package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:5000")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go process2(conn)
	}
}

func process2(conn net.Conn) {
	defer conn.Close()

	for {
		buf := make([]byte, 512)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read err:", err)
			return
		}
		fmt.Printf(string(buf[0:n]))
	}

}
