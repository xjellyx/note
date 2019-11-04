package main

import (
	"fmt"
	"github.com/gansidui/gotcp"
	"github.com/gansidui/gotcp/examples/echo"
	"github.com/srlemon/note/log"
	"net"
)

func main() {
	var (
		tcpAddr  *net.TCPAddr
		listener *net.TCPListener
		err      error
	)
	// 获取tcp地址
	if tcpAddr, err = net.ResolveTCPAddr("tcp4", ":"+"9188"); err != nil {
		panic(err)
	}

	// 开启监听
	if listener, err = net.ListenTCP("tcp", tcpAddr); err != nil {
		panic(err)
	}
	s := &server{}
	s0 := gotcp.NewServer(&gotcp.Config{
		PacketSendChanLimit:    500,
		PacketReceiveChanLimit: 500,
	}, s, &echo.EchoProtocol{})
	s0.Start(listener, 1)
	s0.Stop()
}

type server struct {
	gotcp.ConnCallback
}

func (s *server) OnMessage(conn *gotcp.Conn, packet gotcp.Packet) bool {
	var (
		err error
	)
	fmt.Println("sqqqqqqqqqqqqqqqqqqqqq")
	fmt.Println(packet, "wwwwwwwwwwww")
	if err = conn.AsyncWritePacket(echo.NewEchoPacket([]byte("sdsadasdasdqwer"), false), 1); err != nil {
		fmt.Println(err.Error())
	}

	return true
}

func (s *server) OnConnect(conn *gotcp.Conn) bool {
	addr := conn.GetRawConn().RemoteAddr()
	conn.PutExtraData(addr)
	log.Println("OnConnect:", addr)
	return true
}

func (s *server) Onclose(conn *gotcp.Conn) {
	conn.Close()
}
