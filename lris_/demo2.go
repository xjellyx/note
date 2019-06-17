package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	recover2 "github.com/kataras/iris/middleware/recover"
)

func main() {
	var (
		app *iris.Application
		err error
	)

	app = iris.New()
	app.Use(recover2.New())
	app.Use(logger.New())

	// 创建路由
	app.Handle("GET", "/test", func(ctx iris.Context) {
		if _, err = ctx.HTML("<h1>Test !!!</h1>"); err != nil {
			app.Logger().Fatal(err)
		}
	})

	// 输出字符
	app.Get("/hello", func(ctx iris.Context) {
		if _, err = ctx.WriteString("world"); err != nil {
			app.Logger().Fatal(err)
		}
	})

	// 输出JSON
	app.Get("/json", func(ctx iris.Context) {
		if _, err = ctx.JSON(iris.Map{"Name": "TOM"}); err != nil {
			app.Logger().Fatal(err)
		}
	})

	//
	if err = app.Run(iris.Addr(":8800")); err != nil {
		panic(err)
	}
}
