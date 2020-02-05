package main

import (
	"github.com/olongfen/note/wx/serve"
	"github.com/songtianyi/wechat-go/wxweb"
	"github.com/suboat/sorm/log"
)

var (
	sess *wxweb.Session
)

func main() {
	var (
		err error
	)
	if sess, err = wxweb.CreateSession(nil, nil, wxweb.TERMINAL_MODE); err != nil {
		log.Panic(err)
	}

	if err = serve.Register(sess); err != nil {
		log.Panic(err)
	}
	// 登录并接收消息
	if err := sess.LoginAndServe(false); err != nil {
		log.Panic(err)
	}
}
