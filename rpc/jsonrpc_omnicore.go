package main

import (
	"fmt"
	"net"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.DialTimeout("tcp", "127.0.0.1:8432", 1000*1000*1000*30)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := jsonrpc.NewClient(conn)

	// args := make(map[string]interface{})
	/*args["address"] = "064a3e5c01e37cf41952bc26e2bfa415d1785ad808991cccc102c9ab35494380"
	args["propertyid"] = 1*/
	err = client.Call("/home/allen/bin/omnicore-0.4.0/bin/omnicore-cli omni_getinfo", nil, nil)
	fmt.Println(err)
}
