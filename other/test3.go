package main

import (
	"git.yichui.net/tudy/wechat-go/wxweb"
	"log"
	"os"
)

var (
	sess *wxweb.Session
)

func main() {
	f, err := os.Create("/home/allen/test.jpg")
	if err != nil {
		log.Fatal(err)
	}
	if _, err = f.Write([]byte("sdsd")); err != nil {
		log.Println(err)
	}
}

func register(sess *wxweb.Session) (err error) {

	return
}
