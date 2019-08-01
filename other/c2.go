package main

import (
	"fmt"
	"net/http"
)

func main() {
	url := "127.0.0.1:8080/getFile"
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
