package main

import (
	"fmt"
	"log"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Response struct {
	Country   string
	Province  string
	City      string
	Latitude  float64
	Longitude float64
	TimeZone  string
	Data      string
}
type Client struct {
	*rpc.Client
}

type agrs struct {
	IpString string
	Data     string
}

func main() {
	var (
		c   = new(Client)
		err error
	)
	if c.Client, err = jsonrpc.Dial("tcp", "127.0.0.1:8080"); err != nil {
		log.Fatal("dialing:", err)
	}
	// Synchronous call
	var res Response
	var a agrs
	a.IpString = "219.140.227.235"
	a.Data = "is my data"
	err = c.Call("Serve.GetData", &a, &res)
	if err != nil {
		log.Fatal("ip2addr error:", err)
	}
	fmt.Println(res)

}
