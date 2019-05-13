package main

import (
	"fmt"
	btc "git.yichui.net/tudy/go-bytebank/contrib/bitcoin"
)

func main() {
	c := client()
	fmt.Println(c.GetBlockCount())
	//fmt.Println(c.WalletPassphrase("p1", 3600))
	//fmt.Println(c.GetBestBlockhash())
	//fmt.Println(c.GetBlock("00000000000007e9fa0cd5008dba356971ca285ebdc4860f35a47550be08d961"))
	//fmt.Println(c.GetInfo())
}

func client() *btc.Client {
	var args = &btc.ArgRpc{
		Connect:  "127.0.0.1",
		Port:     8432,
		User:     "btcrpc",
		Password: "123456",
	}
	client, err := btc.NewClient(args)
	if err != nil {
		panic(err)
	}
	return client
}
