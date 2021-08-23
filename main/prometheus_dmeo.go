package main

import (
	"context"
	"fmt"
)

type a struct {
}

func main() {

	ctx := context.WithValue(context.Background(), &a{}, 10)
	var a = &a{}
	fmt.Println(ctx.Value(a))
}
