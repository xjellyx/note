package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"golang.org/x/net/websocket"

	"net/http"
)

// 初始化函数
func init() {
	fmt.Println("Entry init")
	// 连接池的  初始化； 数据结构的初始化；

	flag.Set("alsologtostderr", "true") // 日志写入文件的同时，输出到stderr
	flag.Set("log_dir", "./tmp")        // 日志文件保存目录
	flag.Set("v", "3")                  // 配置V输出的等级。
	flag.Parse()

	return
}

type OnlineUser struct {
	Connection *websocket.Conn
}

// 主函数
func main() {
	glog.Info("main")
	fmt.Println("Entry main")
	http.Handle("/Golangweb", websocket.Handler(BuildConnection))
	if err := http.ListenAndServe(":7821", nil); err != nil {
		fmt.Println("main err:", err.Error())
		return
	}
	return
}

func BuildConnection(ws *websocket.Conn) {
	data := ws.Request().URL.Query().Get("data")
	if data == "" {
		fmt.Println("data is nil!!!")
		//return
	}
	// 接受到的 客户端的数据
	fmt.Println("data:", data)
	// 处理-- 用户的登陆  用户的基础行为（走路 ， 开枪 ，死亡 等等）
	onlineUser := &OnlineUser{
		Connection: ws,
	}
	onlineUser.PullFromClient()
}

func (this *OnlineUser) PullFromClient() {

	fmt.Println("PullFromClient")
	for {
		var content string
		if err := websocket.Message.Receive(this.Connection, &content); err != nil {
			break
		}
		if len(content) == 0 {
			break
		}
		glog.Info("data:", content)
	}
}
