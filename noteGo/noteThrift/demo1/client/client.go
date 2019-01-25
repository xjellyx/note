package main

import (
	"context"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/noteGo/noteGo/noteThrift/demo1/gen-go/echo"
	"log"
	"net"
	"os"
)

func main()  {
var ctx context.Context
client(ctx)

}

func client(ctx context.Context) (err error) {
	var(
		transport *thrift.TSocket
		transportFactory *thrift.TBufferedTransportFactory
		protocoFactory *thrift.TCompactProtocolFactory
		useTransport thrift.TTransport
		client  *echo.EchoClient
	)
	transportFactory=thrift.NewTBufferedTransportFactory(8192)
	protocoFactory=thrift.NewTCompactProtocolFactory()
	if transport,err=thrift.NewTSocket(net.JoinHostPort("127.0.0.1","9898"));err!=nil{
		fmt.Fprintln(os.Stderr,"error resolving address:",err)
		os.Exit(1)
		return
	}
	if useTransport,err=transportFactory.GetTransport(transport);err!=nil{
		return
	}
	client =echo.NewEchoClientFactory(useTransport,protocoFactory)
	if err=transport.Open();err!=nil{
		fmt.Fprintln(os.Stderr, "Error opening socket to 127.0.0.1:9898", " ", err)
		os.Exit(1)
		return
	}
	defer transport.Close()
	req:=&echo.EchoRequest{Message:"I Learn thrift now ."}
	res,_err:=client.Echo(ctx,req)
	if _err!=nil{
		err=_err
		log.Println("dddddddddd",_err.Error())
		return
	}
	log.Println(res.Message)
	fmt.Println("well done")
	return
}