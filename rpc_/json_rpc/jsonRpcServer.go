package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
)

type ArithJson struct {
}
type ArithJsonRequest struct {
	Num1 int
	Num2 int
}
type ArithJsonResponse struct {
	Pro int
	Quo int
	Rem int
}

//
func (a *ArithJson) Multiply(req ArithJsonRequest, res *ArithJsonResponse) (err error) {
	res.Pro = req.Num2 * req.Num1
	return
}

//
func (a *ArithJson) Divide(req ArithJsonRequest, res *ArithJsonResponse) (err error) {
	if req.Num2 == 0 {
		err = errors.New("参数非法")
		return
	}
	res.Quo = req.Num1 / req.Num2
	res.Rem = req.Num1 % req.Num2
	return

}

func main() {
	var (
		lis net.Listener
		err error
	)
	rpc.Register(new(ArithJson))
	lis, err = net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal("1234", err)
	}
	fmt.Fprintf(os.Stdout, "%s", "start connection")

	for {
		conn, err := lis.Accept()
		if err != nil {
			continue
		} // 接收请求
		go func(conn net.Conn) {
			fmt.Fprintf(os.Stdout, "%s", "new client is coming\n")
			jsonrpc.ServeConn(conn)
		}(conn)
	}
}
