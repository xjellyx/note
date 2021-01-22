package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	n     string
	start int
	Input string
)

func main() {
	f := bufio.NewReader(os.Stdin) //读取输入的内容
	fmt.Scanln(&start)

	for i := 0; i <= start; i++ {
		Input, _ = f.ReadString('\n') //定义一行输入的内容分隔符。
		if len(Input) == 1 {
			continue //如果用户输入的是一个空行就让用户继续输入。
		}
		fmt.Sscan(Input, &n) //将Input
		if n == "0" || i == start {
			break
		}

		for _i, v := range strings.Split(n, "\\") {
			space := ""
			c := _i
			for c > 1 {
				space += " "
				c--
			}
			if _i == 0 {
				fmt.Println(v)
			} else {
				fmt.Println(space, v)
			}
		}
	}
}
