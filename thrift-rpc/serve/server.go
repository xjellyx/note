package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/LnFen/note/thrift-rpc/gen-go/demo"
	"github.com/apache/thrift/lib/go/thrift"
	"strings"
)

type Sever struct {
}

const (
	HOST = "0.0.0.0"
	PORT = "8898"
)

func (s *Sever) DoFormat(ctx context.Context, in *demo.Data) (ret *demo.Data, err error) {
	if in == nil {
		err = errors.New("data is nil")
		return
	} else {
		ret = new(demo.Data)
		// 把text参数改为大写
		ret.Text = strings.ToUpper(in.Text)
		ret.StringArr = in.StringArr
	}
	return
}

func main() {
	var (
		hander           = &Sever{}
		processor        *demo.BaseServiceProcessor
		serveTransport   *thrift.TServerSocket
		transportFactory thrift.TTransportFactory
		protocolFactory  *thrift.TBinaryProtocolFactory
		err              error
	)
	processor = demo.NewBaseServiceProcessor(hander)
	// 创建服务
	if serveTransport, err = thrift.NewTServerSocket(HOST + ":" + PORT); err != nil {
		panic(err)
	}

	transportFactory = thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	serve := thrift.NewTSimpleServer4(processor, serveTransport, transportFactory, protocolFactory)
	fmt.Printf("Running at:%s", HOST+":"+PORT)
	if err = serve.Serve(); err != nil {
		panic(err)
	}

}
