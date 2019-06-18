package main

import "github.com/kataras/iris"

func main() {
	app := iris.New()
	//请在参数化路径部分
	users := app.Party("/users", myAuthMiddlewareHandler)
	// http://localhost:8080/users/42/profile
	users.Get("/{id:int}/profile", userProfileHandler)
	// http://localhost:8080/users/inbox/1
	users.Get("/inbox/{id:int}", userMessageHandler)
	if err := app.Run(
		iris.Addr(":8080"), iris.WithConfiguration(iris.TOML("./main.tml"))); err != nil {
		app.Logger().Fatal(err)
	}
}
func myAuthMiddlewareHandler(ctx iris.Context) {
	ctx.WriteString("Authentication failed")
	ctx.Next() //继续执行后续的handler
}
func userProfileHandler(ctx iris.Context) { //
	id := ctx.Params().Get("id")
	ctx.WriteString(id)
}
func userMessageHandler(ctx iris.Context) {
	id := ctx.Params().Get("id")
	ctx.WriteString(id)
}
