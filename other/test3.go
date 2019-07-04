package main

import (
	"fmt"
	"git.yichui.net/tudy/wechat-go/wxweb"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	sess *wxweb.Session
)

func main() {
	path := "http://img3.imgtn.bdimg.com/it/u=2585830458,3269303407&fm=26&gp=0.jpg"
	ss := strings.Split(path, "/")
	fmt.Println(ss[len(ss)-1])
	if data, _err := http.Get("https://ss3.bdstatic.com/70cFv8Sh_Q1YnxGkpoWK1HF6hhy/it/u=1815312391,2410878200&fm=26&gp=0.jpg"); _err != nil {
		panic(_err)
	} else {
		defer data.Body.Close()
		body, _err := ioutil.ReadAll(data.Body)
		if _err != nil {
			panic(_err)
		}
		fmt.Println(string((body)))
	}
}

func register(sess *wxweb.Session) (err error) {

	return
}
