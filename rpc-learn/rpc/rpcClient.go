package main

// 客户端
import (
	"fmt"
	"log"
	"net/rpc"
)

// 声明两个结构体类型要与服务端一样,结构体里面的参数名必须要与服务端一样

// 算数运算请求结构体
type clientRequest struct {
	A int
	B int
}

// 算数运算响应
type clientResponse struct {
	Pro int // 乘积
	Quo int // 商
	Rem int // 余数
}

func main() {
	conn, err := rpc.DialHTTP("tcp", "127.0.0.1:8096")
	if err != nil {
		log.Fatal("6666\n", err)
	}
	req := clientRequest{10, 3}
	var res clientResponse
	err = conn.Call("Arith.Multiply", req, &res) // 呼叫服务端执行乘法函数
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(req.A, req.B, res.Pro)
	if err = conn.Call("Arith.Divide", req, &res); err != nil { // 呼叫服务端执行除法函数
		log.Fatal(err)
	}
	fmt.Println(req.A, req.B, res.Quo, res.Rem)

}
