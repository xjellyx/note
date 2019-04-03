package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	{
		timeout := 3 * time.Second
		ctx, _ := context.WithTimeout(context.Background(), timeout)
		fmt.Println(add(ctx))
	}
}
func cdd(ctx context.Context) (ret int) {
	select {
	case <-ctx.Done():
		return -3
	default:
		fmt.Println(ctx.Value("NLJB"))
		return 3

	}
	return
}

func bdd(ctx context.Context) (ret int) {

	ctx = context.WithValue(ctx, "NLJB", "NULIJIABEI")
	go fmt.Println(cdd(ctx))
	select {
	case <-ctx.Done():
		return -1
	default:
		fmt.Println(ctx.Value("HELLO"))
		fmt.Println(ctx.Value("WORLD"))
		return 1

	}
	return
}

func add(ctx context.Context) (ret int) {
	ctx = context.WithValue(ctx, "HELLO", "WORLD")
	ctx = context.WithValue(ctx, "WORLD", "HELLO")
	go fmt.Println(bdd(ctx))
	select {
	case <-ctx.Done():
		return -2
	default:
		return 2
	}
	return
}
