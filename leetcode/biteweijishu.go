package main

import "fmt"

func countBits(num int) []int {
	var (
		res = make([]int, num+1)
	)
	for i := 0; i <= num; i++ {
		res[i] = res[i>>1] + (i & 1)
	}
	return res
}

func main() {
	fmt.Println(2 ^ 2)
	fmt.Println(countBits(5))
}
