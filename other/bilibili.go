package main

import (
	"fmt"
	"time"
)

const (
	workOneDone   = "oneDone"
	workTwoDone   = "twoDone"
	workThreeDone = "threeDone"
)

func main() {
	fmt.Println("开始施工")
	ch := make(chan string, 1)

	go func() {
		workTwo(ch)
	}()
	go func() {
		workThree(ch)
	}()
	go func() {
		workOne(ch)
	}()

	go func() {
		test(ch)
	}()
	// 关闭chan
	defer close(ch)
	// 等待总工程完成
	time.Sleep(time.Second * 5)

}

func test(ch chan string) {
	for {
		select {
		case v := <-ch:
			fmt.Printf("验收该工程队完成,%s \n", v)
		default:

		}
	}
}

// workOne 工程队1修建一部分
func workOne(ch chan string) {
	workTime := time.Millisecond * 3000 // 完成工程时间
	time.Sleep(workTime)
	fmt.Println("修建完成: workOne")
	ch <- workOneDone
}

// workTwo 工程队2修建一部分
func workTwo(ch chan string) {
	workTime := time.Millisecond * 2500 // 完成工程时间
	time.Sleep(workTime)
	fmt.Println("修建完成: workTwo")
	ch <- workTwoDone
}

// workThree 工程队3修建一部分
func workThree(ch chan string) {
	workTime := time.Millisecond * 2000 // 完成工程时间
	time.Sleep(workTime)
	fmt.Println("修建完成: workThree")
	ch <- workThreeDone
	return
}
