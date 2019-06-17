package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
)

// Resp 返回值结构体
type Resp struct {
	Country   string
	Province  string
	City      string
	Latitude  float64
	Longitude float64
	TimeZone  string
	Data      string
}

// Serve
type Serve struct {
}

// 参数结构体
type Agrs struct {
	IpString string
	Data     string
}

//json rpc 处理请求
// GetData 获取数据
func (s *Serve) GetData(agr Agrs, res *Resp) error {
	res.City = "南宁"
	res.Province = "广西"
	res.Country = "中国"
	res.Latitude = 888.888
	res.Longitude = 111.111
	res.TimeZone = "ssss"
	res.Data = agr.Data // 返回你传的数据
	return nil
}

func main() {
	// 初始化jsonRPC
	serve := &Serve{}
	// 注册服务
	rpc.Register(serve)
	//绑定端口
	address := "127.0.0.1:8080"
	tcpAddr, err := net.ResolveTCPAddr("tcp", address)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	log.Println("json rpc is listening", tcpAddr)
	// 一直连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go jsonrpc.ServeConn(conn)
	}

}

// 验证错误
func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
