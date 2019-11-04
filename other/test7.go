package main

import (
	"fmt"
	"time"
)

// 生产者: 生成 factor 整数倍的序列
func Producer(factor int, out chan<- int) {
	switch {
	case factor <= 0:

		for i := 0; i < -factor; i++ {
			out <- i * factor
		}
	case factor > 0 && factor < 5:
		for i := 0; i < factor; i++ {
			out <- i * factor
		}
	case factor > 5 && factor < 10:
		for i := 0; i < factor; i++ {
			out <- i * factor
		}
	default:
		for i := 0; i < factor; i++ {
			out <- i * factor
		}

	}
}

// 消费者
func Consumer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}
func main() {
	ch := make(chan int, 64) // 成果队列
	go Producer(30000, ch)   // 生成 3 的倍数的序列
	go Producer(500000, ch)
	go Producer(100000, ch)
	go Producer(-100000, ch)
	go Consumer(ch) // 消费 生成的队列

	// 运行一定时间后退出
	time.Sleep(5 * time.Second)
}
