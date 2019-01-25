package main

// 服务端
import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
)

// 算数运算结构体
type Arith struct {
}

// 算数运算请求结构体
type ArithRequest struct {
	A int
	B int
}

// 算数运算响应
type ArithResponse struct {
	Pro int // 乘积
	Quo int // 商
	Rem int // 余数
}

// 乘法运算
func (a *Arith) Multiply(req ArithRequest, res *ArithResponse) (err error) {
	res.Pro = req.A * req.B
	return
}

// 除法运算
func (a *Arith) Divide(req ArithRequest, res *ArithResponse) (err error) {
	if req.B == 0 {
		err = errors.New("非法参数")
		return
	}
	res.Quo = req.A / req.B
	res.Rem = req.A % req.B
	return
}

func main() {
	rpc.Register(new(Arith))                        // 注册服务
	rpc.HandleHTTP()                                // 采用http协议作为rpc载体
	lis, err := net.Listen("tcp", "127.0.0.1:8096") // 设置服务端地址
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(os.Stdout, "%s", "start connection")
	http.Serve(lis, nil) // 生成服务端

}

// 服务端函数运行时间内客户端可以请求使用
