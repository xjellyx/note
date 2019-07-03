package main

import (
	"context"
	"fmt"
	"time"
)

var key = "的发挥反对宦官科隆会议"

func watch(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Value(key), "stop！！！")
			return
		default:
			fmt.Println(ctx.Value(key), "正在进行!!!")
			time.Sleep(time.Second)

		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	valueCtx := context.WithValue(ctx, key, "萨的和哦法法")
	go watch(valueCtx)
	time.Sleep(10 * time.Second)
	fmt.Println("执行完毕")
}
