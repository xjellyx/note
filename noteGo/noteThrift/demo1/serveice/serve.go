package main

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/noteGo/noteGo/noteThrift/demo1/gen-go/echo"

	"log"
)

type EchoServeImp struct {

}

// 初始化
func (e *EchoServeImp)Echo(req *echo.EchoRequest)( *echo.EchoResponse,error)  {
	log.Printf("message from client:%v\n",req.GetMessage())
	res:=&echo.EchoResponse{
		Message:req.GetMessage(),
	}
	return res,nil
}

func main()  {
	var(
		transport *thrift.TServerSocket
		err error
		processor *echo.EchoProcessor
		serve *thrift.TSimpleServer
	)
	if transport,err=thrift.NewTServerSocket(":9898");err!=nil{
		panic(err)
	}
	processor=echo.NewEchoProcessor(&EchoServeImp{})
	serve=thrift.NewTSimpleServer4(
		processor,
		transport,
		thrift.NewTBufferedTransportFactory(8192),
		thrift.NewTCompactProtocolFactory(),
		)
	if err=serve.Serve();err!=nil{
		panic(err)
	}
}
