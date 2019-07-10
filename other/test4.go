package main

import (
	"github.com/srlemon/note/wx/serve"
	"os"
)

func main() {

	var (
		r = new(serve.RequestAudioTenxun)
	)
	r.Text = "你好阿,哈哈哈"
	r.Region = serve.Guangzhou
	r.ModelType = 1
	r.SecretId = serve.SecretIdTenxun
	r.SecretKey = serve.SecretKeyTenxun

	d, err := serve.PubGetAudioByTextTenxun(r)
	if err != nil {
		panic(err)
	}
	f, _ := os.Create("/home/allen/test.wav")
	f.Write(d)
}
