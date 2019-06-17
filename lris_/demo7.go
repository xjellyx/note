package main

import (
	"fmt"
	"github.com/kataras/iris"
)

type user struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	ID   string `json:"id"`
}

func userHandler(ctx iris.Context) {
	u := &user{}
	if err := ctx.ReadJSON(u); err != nil {
		panic(err)
	} else {
		fmt.Println("sssssssssss", u)
		ctx.Writef("user: %#v", u)
	}
}
func main() {
	app := iris.New()
	// curl -d '{"name":"TOM","age":16,"id":"123456789"}' -H "Content-Type:
	// application/json" -X POST http://localhost:8080/user
	app.Post("/user", userHandler)
	app.Run(iris.Addr(":8080"))

}
