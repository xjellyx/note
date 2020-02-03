package main

import (
	"net"
)

func main() {
	t, _err := net.Dial("tcp", "192.168.31.180:9091")
	if _err != nil {
		panic(_err)
	}
	defer t.Close()

	for i := 159; i < 1158; i++ {
		t.Write()
	}

}
