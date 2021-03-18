package main

import (
	"fmt"
	"unicode"
)

func main() {

	fmt.Println(IsChinese("fgdasdf"))
}
func IsChinese(str string) bool {
	var count int
	for _, v := range str {
		if unicode.Is(unicode.Han, v) {
			count++
			break
		}
	}
	return count > 0
}
