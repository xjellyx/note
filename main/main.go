package main

import "fmt"

func main() {
	fmt.Println(d1(10))
}

// f(n)=f(n-1)+f(n-2)
func f(n int) int {
	if n == 0 {
		return 0
	}
	if n == 2 || n == 1 {
		return 1
	}
	return f(n-1) + f(n-2)
}

func d(n int) int {

	var (
		arr = make([]int, n+1)
	)
	if n == 0 {
		return 0
	}
	arr[1] = 1
	arr[2] = 1
	return f1(n, arr)
}

func f1(n int, arr []int) int {
	if n == 0 {
		return 0
	}
	if arr[n] != 0 {
		return arr[n]
	}
	arr[n] = f1(n-1, arr) + f1(n-2, arr)
	return arr[n]
}

func d1(n int) int {
	var (
		arr = make([]int, n+1)
	)
	if n == 0 {
		return 0
	}
	arr[1] = 1
	arr[2] = 1
	for i := 3; i <= n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr[n]
}
