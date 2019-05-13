package main

import (
	"fmt"
	"git.yichui.net/tudy/go-bytebank/contrib/bitcoin"
)

var client, _ = Init()
var allen = "myCmCQ5GH5z92GtY4c9vPTRQh49Gmmeig5"

func main() {
	//fmt.Println(client.OmniSend(allen, "moneyqMan7uh8FqdCA2BV5yZ8qVrc9ikLP", 1, 0.0001,
	//	nil, nil))
	//fmt.Println(client.GetBalance("", 6))
	//fmt.Println(client.OmniGetBalance(allen, 1))

}

// 获取区块高度
func getBlockCount() {
	fmt.Println(client.GetBlockCount())
}

// 返回节点和网络信息
func getInfo() {
	fmt.Println(client.GetInfo())
}

// 新区块
func newAddr() {
	fmt.Println(client.GetNewAddress("allen"))
}
func Init() (ret *bitcoin.Client, err error) {
	testHost := "127.0.0.1"
	testPort := 18332
	testUser := "user"
	testPassword := "password"
	if ret, err = bitcoin.NewClient(
		&bitcoin.ArgRpc{
			Connect:  testHost,
			Port:     testPort,
			User:     testUser,
			Password: testPassword,
		}); err != nil {
		panic(err)
	}
	return
}
