package main

import (
	"fmt"
)

func main() {
	var (
		// 无缓冲
		// chan1 = make(chan int)
		// 有缓冲
		chan2 = make(chan int, 11)
	)
	chan2 <- 98
	if v, ok := <-chan2; ok {
		// 读取之后缓存的数据被清除
		fmt.Println(v)
	}

	// 写如10个数据到缓冲区
	for i := 0; i < 10; i++ {
		chan2 <- i
	}
	// 再写入数据
	chan2 <- 9898
	for v := range chan2 {
		fmt.Println(v)
		// 这里输入chan里面的任何一个value
		if v == 9898 {
			// 关闭管道
			close(chan2)
		}
	}
}
