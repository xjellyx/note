package main

import (
	"fmt"
	"regexp"
	"strings"
)

func Between(str, starting, ending string) string {
	s := strings.Index(str, starting)
	if s < 0 {
		return ""
	}
	s += len(starting)
	e := strings.Index(str[s:], ending)
	if e < 0 {
		return ""
	}
	return str[s : s+e]
}

func main() {
	b := "https://www.codedesigner.net/exec/1079/1902"
	re, _ := regexp.Compile(`\d+`)

	//查找符合正则的第一个
	all := re.FindAll([]byte(b), -1)
	for _, item := range all {
		fmt.Println(string(item))
	}
}
