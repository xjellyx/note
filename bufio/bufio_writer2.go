package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

// 示例：带检查扫描
func main() {
	var input = "1234 5678 1234567901234567890 90"
	scanner := bufio.NewScanner(strings.NewReader(input))
	// 自定义匹配函数
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		// 获取一个单词 ScanWords 是一个“匹配函数”，用来找出 data 中以空白字符分隔的单词。
		// 空白字符由 unicode.IsSpace 定义。
		advance, token, err = bufio.ScanWords(data, atEOF)
		// 判断其能否转换为整数，如果不能则返回错误
		if err == nil && token != nil {
			_, err = strconv.ParseInt(string(token), 10, 64)
		}
		// 这里包含了 return 0, nil, nil 的情况
		return
	}
	// 设置匹配函数
	scanner.Split(split)
	// 开始扫描
	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s", err)
	}
}
