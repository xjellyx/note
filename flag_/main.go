package main

import (
	"flag"
	"fmt"
)

/*
	Flag 包实现了命令行标志解析。
*/

func main() {
	var n *int
	n = flag.Int("num", 199, "test-num")
	var str string // 这样的函数第一个参数换成了变量地址，后面的参数和flag.String是一样的
	flag.StringVar(&str, "address", "NanNin", "where is your address")
	flag.Parse() //解析输入的参数

	// 是否已经解析
	fmt.Println(flag.Parsed())
	fmt.Println(*n)
	fmt.Println(str)

}
