package main

import (
	"fmt"
	"log"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type ClientJsonRequest struct {
	Num1 int
	Num2 int
}
type ClientJsonResponse struct {
	Pro int
	Quo int
	Rem int
}

func main() {
	var (
		conn *rpc.Client
		err  error
		req  ClientJsonRequest
		res  ClientJsonResponse
	)
	if conn, err = jsonrpc.Dial("tcp", "127.0.0.1:8080"); err != nil { // 请求服务器
		log.Fatal(err)
	}
	req = ClientJsonRequest{88, 15}
	if err = conn.Call("ArithJson.Multiply", req, &res); err != nil {
		log.Fatal(err)
	}
	fmt.Println(req.Num1, req.Num2, res.Pro)
	if err = conn.Call("ArithJson.Divide", req, &res); err != nil {
		log.Fatal(err)
	}
	fmt.Println(req.Num1, req.Num2, res.Quo, res.Rem)

}
