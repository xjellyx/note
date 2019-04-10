package main

import (
	"context"
	"fmt"
	"github.com/LnFen/note/thrift-rpc/gen-go/demo"
	"github.com/apache/thrift/lib/go/thrift"
	"log"
	"net"
)

const (
	HOST = "0.0.0.0"
	PORT = "8898"
)

var (
	Client    *demo.BaseServiceClient
	transport thrift.TTransport
)

func setClient() {
	var (
		tSocker          *thrift.TSocket
		transportFactory thrift.TTransportFactory

		protocolFactory *thrift.TBinaryProtocolFactory
		err             error
	)
	// 获得rpc服务
	if tSocker, err = thrift.NewTSocket(net.JoinHostPort(HOST, PORT)); err != nil {
		panic(err)
	}
	// 工厂运输模式
	transportFactory = thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())

	// 获取运输载体
	if transport, err = transportFactory.GetTransport(tSocker); err != nil {
		panic(err)
	}
	// 协议
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	// 打开客户端
	Client = demo.NewBaseServiceClientFactory(transport, protocolFactory)
}

func doFormat(ctx context.Context, in *demo.Data) (ret *demo.Data, err error) {
	if ret, err = Client.DoFormat(ctx, in); err != nil {
		return
	}
	return
}
func main() {
	setClient()
	// 打开开运输模式
	var (
		err error
	)
	if err = transport.Open(); err != nil {
		log.Fatalln("Error opening:", HOST+":"+PORT)
	}
	defer transport.Close()
	data := &demo.Data{Text: "hello,thrift rpc !", StringArr: []string{"rpc", "thrift", "go"}}
	var (
		ctx context.Context
		ret = new(demo.Data)
	)
	if ret, err = doFormat(ctx, data); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(ret.StringArr, ret.Text)
	}

}
