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
	app.Get("/", func(ctx iris.Context) {
		if _, err = ctx.HTML("<b>Hello!</b>"); err != nil {
			app.Logger().Fatal(err)
		}
	})

	if err = app.Run(
		iris.Addr(":8080"), iris.WithConfiguration(iris.TOML("./main.tml"))); err != nil {
		app.Logger().Fatal(err)
	}
}
