package main

import (
	jsoniter "github.com/json-iterator/go"
	"gopkg.in/gomail.v2"
)

func main() {
	m := gomail.NewMessage()
	m.SetHeader("From", "cpic@starwiz.cn")
	m.SetHeader("To", "824426699@qq.com")
	m.SetHeader("Subject", "Hello!")
	data := map[string]string{"aa": "bb"}
	dd, _ := jsoniter.Marshal(data)
	m.SetBody("text/plain", string(dd))

	d := gomail.NewDialer("smtp.exmail.qq.com", 465, "cpic@starwiz.cn", "ksJrYt6vkaQ8VJGa")
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

}
