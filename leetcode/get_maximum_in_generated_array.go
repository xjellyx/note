package main

import (
	"fmt"
	"sort"
)

func getMaximumGenerated(n int) int {
	var (
		fn = make([]int, n+1)
	)
	if n == 0 {
		return 0
	}
	fn[0] = 0
	fn[1] = 1
	if n == 1 {
		return 1
	}
	for i := 0; i <= n; i++ {
		o := 2 * i
		j := o + 1
		if o >= 2 && o <= n {
			fn[o] = fn[i]
		}
		if j >= 2 && j <= n {
			fn[j] = fn[i] + fn[i+1]
		}
	}

	sort.Ints(fn)
	return fn[n]
}

func main() {
	fmt.Println(getMaximumGenerated(11))
}
