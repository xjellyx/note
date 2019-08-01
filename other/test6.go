package main

import (
	"log"
	"net/http"
)

func main() {
	//注册一个函数，响应某一个路由
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello this is version 1!!"))
	})
	//这里可以单独写一个函数传递给当前的路由
	log.Println("Start version v1")
	log.Fatal(http.ListenAndServe(":4000", nil))

}
