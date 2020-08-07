package main

import (
	"bufio"
	"fmt"
	"github.com/olongfen/note/log"
	"os"
	"time"
)

func main() {
	readEachLineScanner("/data/gocode/src/starwiz-customer-micro/log/starwiz-customer-micro/signal_task_33/33.2020-08.log")
}

func readEachLineScanner(filePath string) {
	start1 := time.Now()
	FileHandle, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
		return
	}
	defer FileHandle.Close()
	lineScanner := bufio.NewScanner(FileHandle)
	for lineScanner.Scan() {
		// 相同使用场景下可以使用如下方法
		// func (s *Scanner) Bytes() []byte
		// func (s *Scanner) Text() string
		// 实际逻辑 : 对读取的内容进行某些业务操作
		// 如下代码打印每次读取的文件行内容
		fmt.Println(lineScanner.Text())
	}
	fmt.Println("readEachLineScanner spend : ", time.Now().Sub(start1))
}
