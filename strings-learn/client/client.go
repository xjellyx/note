package main

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/srlemon/note/thrift-rpc/gen-go/example"
	"log"
	"net"
)

const (
	HOST = "0.0.0.0"
	PORT = "8898"
)

type Client struct {
	BaseClient *example.BaseServiceClient
}

var (
	client      *Client
	transport   thrift.TTransport
	xiaomingUId = "fb189006-9a8b-4b50-a343-f221be1cce7b"
	xiaohongUid = "a98ed979-881b-49a1-b52d-5747eedd3fe8"
	c           *example.BaseServiceClient
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
	c = example.NewBaseServiceClientFactory(transport, protocolFactory)
	fmt.Println(c.GetStudentByUID(DefaultXCtx, xiaomingUId))
}

func (c *Client) GetStudentByUID(ctx context.Context, uid string) (ret *example.Student, err error) {
	ret = new(example.Student)
	if ret, err = c.GetStudentByUID(ctx, uid); err != nil {
		return
	}
	return
}
func GetStudentByUID(ctx context.Context, uid string) (ret *example.Student, err error) {
	if ret, err = client.GetStudentByUID(ctx, uid); err != nil {
		return
	}
	return
}

func (c *Client) ModifyStudent(ctx context.Context, uid string, form *example.FormStudent) (ret *example.Student, err error) {
	ret = new(example.Student)
	if ret, err = c.ModifyStudent(ctx, uid, form); err != nil {
		return
	}
	return
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

	//var (
	//	ctx   = DefaultXCtx
	//	data  = new(example.Student)
	//	data2 = new(example.Student)
	//)
	//
	//if data, err = GetStudentByUID(ctx, xiaomingUId); err != nil {
	//	panic(err)
	//} else {
	//	fmt.Println("小明的信息", data.Age, data.Sex, data.ClassName, data.Name)
	//}
	//
	////
	//var (
	//	form = new(example.FormStudent)
	//)
	//form.ClassName = "高三三班"
	//if data2, err = client.ModifyStudent(ctx, xiaohongUid, form); err != nil {
	//	panic(err)
	//} else {
	//	fmt.Println("小红的信息", data2.ClassName, data2.Age, data2.Name, data2.Sex)
	//}

}

func NewClient(c *Client) (s *Client) {
	if c != nil {
		s = c
	}
	s = new(Client)
	return
}
