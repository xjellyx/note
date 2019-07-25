package main

import (
	"fmt"
)

func task() {
	fmt.Println("I am runnning task.")
}

func taskWithParams(a int, b string) {
	fmt.Println(a, b)
}

func main() {
	var test []string
	test = append(test, "")
	fmt.Println(len(test))
}
