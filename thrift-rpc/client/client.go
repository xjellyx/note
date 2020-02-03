package main

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/olefen/note/thrift-rpc/gen-go/demo"
	"log"
	"net"
)

const (
	HOST = "0.0.0.0"
	PORT = "8898"
)

type Client struct {
	BaseClient *demo.BaseServiceClient
}

var (
	client      = new(Client)
	transport   thrift.TTransport
	xiaomingUId = "fb189006-9a8b-4b50-a343-f221be1cce7b"
	xiaohongUid = "a98ed979-881b-49a1-b52d-5747eedd3fe8"
)

func setClient() {
	var (
		tSocket          *thrift.TSocket
		transportFactory thrift.TTransportFactory

		protocolFactory *thrift.TBinaryProtocolFactory
		err             error
	)
	// 获得rpc服务
	if tSocket, err = thrift.NewTSocket(net.JoinHostPort(HOST, PORT)); err != nil {
		panic(err)
	}
	// 开启工厂运输模式
	transportFactory = thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())

	// 获取运输载体
	if transport, err = transportFactory.GetTransport(tSocket); err != nil {
		panic(err)
	}
	// 协议
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	// 打开客户端
	client.BaseClient = demo.NewBaseServiceClientFactory(transport, protocolFactory)
}

var DefaultXCtx = context.Background()

func main() {
	setClient()
	// 打开开运输模式
	var (
		err error
	)
	// 打开连接服务端
	if err = transport.Open(); err != nil {
		log.Fatalln("Error opening:", HOST+":"+PORT)
	}
	defer transport.Close()

	fmt.Println(client.BaseClient.ModifyStudent(DefaultXCtx, xiaomingUId, &demo.FormStudent{ClassName: "sasdsad"}))
	fmt.Println(client.BaseClient.GetStudentByUID(DefaultXCtx, xiaohongUid))

}

func NewClient(c *Client) (s *Client) {
	if c != nil {
		s = c
	}
	s = new(Client)
	return
}
