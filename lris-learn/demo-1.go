package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	recover2 "github.com/kataras/iris/middleware/recover"
	"github.com/valyala/tcplisten"
	"net"
)

func main() {
	var (
		listen net.Listener
		err    error
		app    *iris.Application
	)
	// 创建新的app
	app = iris.New()
	app.Use(recover2.New())
	app.Use(logger.New())

	//输出html
	// 请求方式: GET
	// 直接请求
	app.Get("/", func(ctx iris.Context) {
		if _, err = ctx.HTML("<h1>Hello World!</h1>"); err != nil {
			panic(err)
		}
	})
	listenerCfg := tcplisten.Config{
		ReusePort:   true,
		DeferAccept: true,
		FastOpen:    true,
	}
	if listen, err = listenerCfg.NewListener("tcp4", ":8080"); err != nil {
		app.Logger().Fatal(err)
	}
	if err = app.Run(iris.Listener(listen)); err != nil {
		panic(err)
	}

}
