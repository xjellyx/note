package main

import (
	"github.com/srlemon/note/wx/serve"
	"os"
)

func main()  {
	t,_:=serve.PubGetToken()
	b,err:=serve.PubGetAudioByText("百度你好",&t)
	if err!=nil{
		panic(err)
	}
	f,_:=os.Create("/home/allen/t.mp3")
	f.Write(b)
}
